export SUT=ec2-54-169-205-187.ap-southeast-1.compute.amazonaws.com

ssh -i sck_default.pem ubuntu@$SUT mkdir sut monitoring
ssh -i sck_default.pem ubuntu@$SUT mkdir -p monitoring/data/grafana
scp -i sck_default.pem ../docker-compose.deploy.yml ubuntu@$SUT:~/sut/docker-compose.yml
scp -i sck_default.pem -r ../tearup ubuntu@$SUT:~/sut
scp -i sck_default.pem ./monitoring/docker-compose.monitoring.yml ubuntu@$SUT:~/monitoring/docker-compose.yml
scp -i sck_default.pem ./monitoring/prometheus.yml ubuntu@$SUT:~/monitoring
ssh -i sck_default.pem ubuntu@$SUT "cd ~/monitoring && docker-compose pull && docker-compose up -d "
ssh -i sck_default.pem ubuntu@$SUT "cd ~/sut && docker-compose pull && docker-compose up -d "
