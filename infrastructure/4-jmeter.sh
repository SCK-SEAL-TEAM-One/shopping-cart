export JMETER=<KUBE_MASTER>

scp -i deploy/shoppingcart_key.pem atdd/load ubuntu@$JMETER:~/apache-jmeter-5.1.1/bin
