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