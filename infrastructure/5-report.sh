export JMETER=ec2-13-229-236-162.ap-southeast-1.compute.amazonaws.com

ssh -i sck_default.pem ubuntu@$JMETER cd apache-jmeter-5.1.1/bin && nohup java -jar ApacheJMeter.jar -n -t api-get-all-product-ramp-1min-to-100con.jmx -l api-get-all-product-ramp-1min-to-100con.log -e -o api-get-all-product-ramp-1min-to-100con &
scp -i sck_default.pem ubuntu@$JMETER:~/apache-jmeter-5.1.1/bin/api-get-all-product-ramp-1min-to-100con.log .
scp -i sck_default.pem -r ubuntu@$JMETER:~/apache-jmeter-5.1.1/bin/api-get-all-product-ramp-1min-to-100con .

ssh -i sck_default.pem ubuntu@$JMETER cd apache-jmeter-5.1.1/bin && nohup java -jar ApacheJMeter.jar -n -t api-get-by-id-product-ramp-1min-to-100con.jmx -l api-get-by-id-product-ramp-1min-to-100con.log -e -o api-get-by-id-product-ramp-1min-to-100con &
scp -i sck_default.pem ubuntu@$JMETER:~/apache-jmeter-5.1.1/bin/api-get-by-id-product-ramp-1min-to-100con.log .
scp -i sck_default.pem -r ubuntu@$JMETER:~/apache-jmeter-5.1.1/bin/api-get-by-id-product-ramp-1min-to-100con .

ssh -i sck_default.pem ubuntu@$JMETER cd apache-jmeter-5.1.1/bin && nohup java -jar ApacheJMeter.jar -n -t order-placement-ramp-1min-to-200con.jmx -l order-placement-ramp-1min-to-200con.log -e -o order-placement-ramp-1min-to-200con &
scp -i sck_default.pem ubuntu@$JMETER:~/apache-jmeter-5.1.1/bin/order-placement-ramp-1min-to-200con.log .
scp -i sck_default.pem -r ubuntu@$JMETER:~/apache-jmeter-5.1.1/bin/order-placement-ramp-1min-to-200con .
