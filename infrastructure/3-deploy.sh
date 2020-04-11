export KUBE_MASTER=ec2-13-229-235-183.ap-southeast-1.compute.amazonaws.com
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