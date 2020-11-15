provider "aws" {
  access_key = "AKIA2LXJORFWO6GVEETJ"
  secret_key = "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"
  region     = "ap-southeast-1"
}

resource "aws_vpc" "performance_vpc" {
  cidr_block           = "10.0.0.0/16"
  instance_tenancy     = "default"
  enable_dns_support   = true
  enable_dns_hostnames = true
  tags = {
      Name = "Performance VPC"
  }
}
resource "aws_subnet" "performance_subnet" {
  vpc_id                  = aws_vpc.performance_vpc.id
  cidr_block              = "10.0.1.0/24"
  map_public_ip_on_launch = true
  availability_zone       = "ap-southeast-1a"
  tags = {
    Name = "Performance Subnet"
  }
}

module "security_group" {
  source = "terraform-aws-modules/security-group/aws"

  name        = "shoppingcart"
  description = "shoppingcart Security group"
  vpc_id      = "${aws_vpc.performance_vpc.id}"
  
  ingress_with_self = [{
    rule = "all-all"
  }]
  ingress_cidr_blocks = ["0.0.0.0/0"]
  ingress_rules       = ["all-tcp", "all-icmp", "mysql-tcp", "ssh-tcp"]
  egress_rules        = ["all-all"]
}

resource "aws_internet_gateway" "performance_vpc_gw" {
  vpc_id = aws_vpc.performance_vpc.id
  tags = {
    Name = "Performance VPC Internet Gateway"
  }
} 

resource "aws_route_table" "performance_route_table" {
  vpc_id = aws_vpc.performance_vpc.id
  tags = {
    Name = "Performance VPC Route Table"
  }
} 

resource "aws_route" "performance_vpc_internet_access" {
  route_table_id         = aws_route_table.performance_route_table.id
  destination_cidr_block = "0.0.0.0/0"
  gateway_id             = aws_internet_gateway.performance_vpc_gw.id
}

resource "aws_route_table_association" "performance_vpc_association" {
  subnet_id      = aws_subnet.performance_subnet.id
  route_table_id = aws_route_table.performance_route_table.id
}


resource "random_string" "token_id" {
  length  = 6
  special = false
  upper   = false
}

resource "random_string" "token_secret" {
  length  = 16
  special = false
  upper   = false
}

locals {
  token = "${random_string.token_id.result}.${random_string.token_secret.result}"
}

# kubectl-ready
module "kube_master" {
  source  = "terraform-aws-modules/ec2-instance/aws"
  version = "~> 2.0"

  name = "Kube Master"
  ami  = "ami-0e45ef99a98566a26"

  instance_type = "t3.medium" # slower one is t3.medium faster one is t3.large

  instance_count              = 1
  vpc_security_group_ids      = ["${module.security_group.this_security_group_id}"]
  subnet_id                   = "${aws_subnet.performance_subnet.id}"
  associate_public_ip_address = true
  monitoring                  = true

  root_block_device = [
    {
      volume_type = "gp2"
      volume_size = 10
    },
  ]

  key_name = "sck_default"
  tags = {
    "Type" = "kubernetes"
  }

  user_data = <<-EOF
  #!/bin/bash
  # Install kubeadm and Docker
  sudo ufw disable
  sudo systemctl disable ufw
  # Run kubeadm
  sudo kubeadm init \
    --token "${local.token}" \
    --token-ttl 15m 
    --kubernetes-version v1.13.0 
    --ignore-preflight-errors=all
  # Prepare kubeconfig file for download to local machine
  mkdir -p $HOME/.kube
  sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
  sudo chown $(id -u):$(id -g) $HOME/.kube/config
  # Indicate completion of bootstrapping on this node
  touch /home/ubuntu/done
  EOF
}

# kubectl-ready
module "kube_slave" {
  source  = "terraform-aws-modules/ec2-instance/aws"
  version = "~> 2.0"

  name = "RSS App"
  ami  = "ami-0e45ef99a98566a26"

  instance_type = "t3.medium" # slower one is t3.medium faster one is t3.large 

  instance_count              = 2
  vpc_security_group_ids      = ["${module.security_group.this_security_group_id}"]
  subnet_id                   = "${aws_subnet.performance_subnet.id}"
  associate_public_ip_address = true
  monitoring                  = true

  root_block_device = [
    {
      volume_type = "gp2"
      volume_size = 10
    },
  ]

  key_name = "sck_default"
  tags = {
    Type = "kubernetes"
  }

  user_data = <<-EOF
  #!/bin/bash
  # Install kubeadm and Docker
  sudo ufw disable
  sudo systemctl disable ufw
  sudo kubeadm join ${module.kube_master.private_ip[0]}:6443 \
    --token ${local.token} \
    --discovery-token-unsafe-skip-ca-verification
  touch /home/ubuntu/done
  EOF
}

module "j_meter" {
  source  = "terraform-aws-modules/ec2-instance/aws"
  version = "~> 2.0"

  name = "j meter"
  ami  = "ami-0907d9e140ac01676"

  instance_type = "t3.medium"

  instance_count              = 1
  vpc_security_group_ids      = ["${module.security_group.this_security_group_id}"]
  subnet_id                   = "${aws_subnet.performance_subnet.id}"
  associate_public_ip_address = true
  monitoring                  = false

  root_block_device = [
    {
      volume_type = "gp2"
      volume_size = 10
    },
  ]

  key_name = "sck_default"
  tags = {
    Type = "jmeter"
  }
}

resource "null_resource" "wait_for_bootstrap_to_finish" {
  provisioner "local-exec" {
    command = <<-EOF
    alias ssh='ssh -q -i ${abspath(path.module)}/sck_default.pem -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null'
    while true; do
      sleep 2
      ! ssh ubuntu@${module.kube_master.public_ip[0]} [[ -f /home/ubuntu/done ]] >/dev/null && continue
      %{for worker_public_ip in module.kube_slave.public_ip[*]~}
      ! ssh ubuntu@${worker_public_ip} [[ -f /home/ubuntu/done ]] >/dev/null && continue
      %{endfor~}
      break
    done
    EOF
  }
  triggers = {
    instance_ids = join(",", concat(module.kube_master.id[*], module.kube_slave.id[*]))
  }
}

resource "null_resource" "install_coredns" {
  provisioner "local-exec" {
    command = <<-EOF
    alias ssh='ssh -q -i ${abspath(path.module)}/sck_default.pem -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null'
    ssh ubuntu@${module.kube_master.public_ip[0]} "source /home/ubuntu/.bashrc; mkdir -p ~/.kube; sudo cp -i /etc/kubernetes/admin.conf ~/.kube/config; sudo chown ubuntu:ubuntu ~/.kube/config; kubectl apply -f https://docs.projectcalico.org/v3.4/getting-started/kubernetes/installation/hosted/etcd.yaml; kubectl apply -f https://docs.projectcalico.org/v3.4/getting-started/kubernetes/installation/hosted/calico.yaml"
    EOF
  }
  triggers = {
    wait_for_bootstrap_to_finish = null_resource.wait_for_bootstrap_to_finish.id
  }
}

output "kube_master_private_ip" {
  value = module.kube_master.private_ip[0]
}

output "kube_master_dns" {
  value = module.kube_master.public_dns[*]
}

output "kube_slave_dns" {
  value = module.kube_slave.public_dns[*]
}

output "jmeter_dns" {
  value = module.j_meter.public_dns[*]
}