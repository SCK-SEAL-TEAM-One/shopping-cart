# Provisioning Infrastructure

## Pre-required

- Access key & Secret key (Programmatic access)

## 1. Initial instance state

```bash
cd infrastructure
terraform init
```

terraform will check dependencies that it needs

```bash
terraform apply
```

after finished will be output dns like this

```
Outputs:

jmeter_dns = [
    "ec2-54-254-218-68.ap-southeast-1.compute.amazonaws.com",
]
kube_master_dns = [
    "ec2-13-212-186-240.ap-southeast-1.compute.amazonaws.com",
]
kube_master_private_ip = "10.0.1.192"
kube_slave_dns = [
    "ec2-54-255-134-154.ap-southeast-1.compute.amazonaws.com",
    "ec2-52-221-252-221.ap-southeast-1.compute.amazonaws.com",
]
```

## 2. use first script to initial kubernetes cluster

```
export KUBE_MASTER=ec2-13-212-186-240.ap-southeast-1.compute.amazonaws.com
ssh -i sck_default.pem ubuntu@$KUBE_MASTER

sudo ufw disable
sudo systemctl disable ufw
sudo kubeadm init --kubernetes-version v1.13.0 --ignore-preflight-errors=all
```

after run kubeadm init it will show some token to join nodes like this

```
kubeadm join 10.0.1.192:6443 --token 3ck0ma.vhzoohejph82dx3l --discovery-token-ca-cert-hash sha256:2019e86de17a93950a677d8d546089213d55c4b4fa2dafaaa5707f7f94af8179
```

run it with sudo

```
sudo kubeadm join 10.0.1.192:6443 --token 3ck0ma.vhzoohejph82dx3l --discovery-token-ca-cert-hash sha256:2019e86de17a93950a677d8d546089213d55c4b4fa2dafaaa5707f7f94af8179
```

## 3. use second script to apply network with calico

```bash
export KUBE_MASTER=ec2-13-212-186-240.ap-southeast-1.compute.amazonaws.com
ssh -i sck_default.pem ubuntu@$KUBE_MASTER

mkdir -p $HOME/.kube
sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config

kubectl apply -f https://docs.projectcalico.org/v3.4/getting-started/kubernetes/installation/hosted/etcd.yaml

kubectl apply -f https://docs.projectcalico.org/v3.4/getting-started/kubernetes/installation/hosted/calico.yaml
```

## 3. deploy all services

```bash
export KUBE_MASTER=ec2-13-212-186-240.ap-southeast-1.compute.amazonaws.com
scp -i sck_default.pem -r ../deploy ubuntu@$KUBE_MASTER:~/
ssh -i sck_default.pem ubuntu@$KUBE_MASTER kubectl apply -f deploy/mysql-service.yml
ssh -i sck_default.pem ubuntu@$KUBE_MASTER kubectl apply -f deploy/mysql-deployment.yml
ssh -i sck_default.pem ubuntu@$KUBE_MASTER kubectl rollout status deployments/store-database-deployment
cat ../tearup/init.sql | docker run -i mysql:5.7 /usr/bin/mysql -h -u sealteam --password=sckshuhari --default-character-set=utf8  toy
ssh -i sck_default.pem ubuntu@$KUBE_MASTER kubectl apply -f deploy/bank-gateway.yml
ssh -i sck_default.pem ubuntu@$KUBE_MASTER kubectl rollout status deployments/bank-gateway
ssh -i sck_default.pem ubuntu@$KUBE_MASTER kubectl apply -f deploy/shipping-gateway.yml
ssh -i sck_default.pem ubuntu@$KUBE_MASTER kubectl rollout status deployments/shipping-gateway
ssh -i sck_default.pem ubuntu@$KUBE_MASTER kubectl apply -f deploy/store-service-service.yml
ssh -i sck_default.pem ubuntu@$KUBE_MASTER kubectl apply -f deploy/store-service-deployment.yml
ssh -i sck_default.pem ubuntu@$KUBE_MASTER kubectl rollout status deployments/store-service
ssh -i sck_default.pem ubuntu@$KUBE_MASTER kubectl apply -f deploy/store-web-service.yml
ssh -i sck_default.pem ubuntu@$KUBE_MASTER kubectl apply -f deploy/store-web-deployment.yml
ssh -i sck_default.pem ubuntu@$KUBE_MASTER kubectl rollout status deployments/store-web
ssh -i sck_default.pem ubuntu@$KUBE_MASTER kubectl apply -f deploy/store-loadbalancer-service.yml
ssh -i sck_default.pem ubuntu@$KUBE_MASTER kubectl apply -f deploy/store-loadbalancer-deployment.yml
ssh -i sck_default.pem ubuntu@$KUBE_MASTER kubectl rollout status deployments/store-loadbalancer
```

## 4. Load Test script to executor

```bash
export JMETER=ec2-13-229-236-162.ap-southeast-1.compute.amazonaws.com

scp -i sck_default.pem ../atdd/load/api-get-all-product-ramp-1min-to-100con.jmx ubuntu@$JMETER:~/apache-jmeter-5.1.1/bin
scp -i sck_default.pem ../atdd/load/api-get-by-id-product-ramp-1min-to-100con.jmx ubuntu@$JMETER:~/apache-jmeter-5.1.1/bin
scp -i sck_default.pem ../atdd/load/order-placement-ramp-1min-to-200con.jmx ubuntu@$JMETER:~/apache-jmeter-5.1.1/bin
```

## 5. Execute load

```bash
export JMETER=ec2-13-229-236-162.ap-southeast-1.compute.amazonaws.com

ssh -i sck_default.pem ubuntu@$JMETER
cd apache-jmeter-5.1.1/bin 

java -jar ApacheJMeter.jar -n -t api-get-all-product-ramp-1min-to-100con.jmx -l api-get-all-product-ramp-1min-to-100con.log -e -o api-get-all-product-ramp-1min-to-100con

java -jar ApacheJMeter.jar -n -t api-get-by-id-product-ramp-1min-to-100con.jmx -l api-get-by-id-product-ramp-1min-to-100con.log -e -o api-get-by-id-product-ramp-1min-to-100con

java -jar ApacheJMeter.jar -n -t order-placement-ramp-1min-to-200con.jmx -l order-placement-ramp-1min-to-200con.log -e -o order-placement-ramp-1min-to-200con
```
 
## Teardown all services 

```bash
terraform destroy
```

destroy all instances and networks. need to confirmation.

## TODO

- script 1,2,3,4,5 ใช้ `user_data` and `null_resource (provisioner [local-exec, file])` เพื่อให้เป็น automated ทั้งหมด
- ลบ access_key กับ secret_key เนื่องจากถ้า key หลุดเป็น public จะโดนเอาไปใช้เปิดเหมือง 
  Solution: ใช้ variable แทน แล้ว pass variable ผ่าน command `terraform apply -var 'access_key=xxx' -var 'secret_key=yyyy'`