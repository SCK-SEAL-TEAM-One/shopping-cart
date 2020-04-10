export KUBE_MASTER=<KUBE_MASTER>
scp -i shoppingcart_key.pem -r ./deploy ubuntu@$KUBE_MASTER:~/
ssh -i shoppingcart_key.pem ubuntu@$KUBE_MASTER kubectl apply -f deploy/mysql-service.yml
ssh -i shoppingcart_key.pem ubuntu@$KUBE_MASTER kubectl apply -f deploy/mysql-deployment.yml
ssh -i shoppingcart_key.pem ubuntu@$KUBE_MASTER kubectl rollout status deployments/store-database-deployment
ssh -i shoppingcart_key.pem ubuntu@$KUBE_MASTER kubectl apply -f deploy/bank-gateway.yml
ssh -i shoppingcart_key.pem ubuntu@$KUBE_MASTER kubectl rollout status deployments/bank-gateway
ssh -i shoppingcart_key.pem ubuntu@$KUBE_MASTER kubectl apply -f deploy/shipping-gateway.yml
ssh -i shoppingcart_key.pem ubuntu@$KUBE_MASTER kubectl rollout status deployments/shipping-gateway
ssh -i shoppingcart_key.pem ubuntu@$KUBE_MASTER kubectl apply -f deploy/store-service-service.yml
ssh -i shoppingcart_key.pem ubuntu@$KUBE_MASTER kubectl apply -f deploy/store-service-deployment.yml
ssh -i shoppingcart_key.pem ubuntu@$KUBE_MASTER kubectl rollout status deployments/store-service
ssh -i shoppingcart_key.pem ubuntu@$KUBE_MASTER kubectl apply -f deploy/store-web-service.yml
ssh -i shoppingcart_key.pem ubuntu@$KUBE_MASTER kubectl apply -f deploy/store-web-deployment.yml
ssh -i shoppingcart_key.pem ubuntu@$KUBE_MASTER kubectl rollout status deployments/store-web
ssh -i shoppingcart_key.pem ubuntu@$KUBE_MASTER kubectl apply -f deploy/store-loadbalancer-service.yml
ssh -i shoppingcart_key.pem ubuntu@$KUBE_MASTER kubectl apply -f deploy/store-loadbalancer-deployment.yml
ssh -i shoppingcart_key.pem ubuntu@$KUBE_MASTER kubectl rollout status deployments/store-loadbalancer