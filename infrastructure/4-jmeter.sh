export JMETER=ec2-13-229-236-162.ap-southeast-1.compute.amazonaws.com

scp -i sck_default.pem ../atdd/load/api-get-all-product-ramp-1min-to-100con.jmx ubuntu@$JMETER:~/apache-jmeter-5.1.1/bin
scp -i sck_default.pem ../atdd/load/api-get-by-id-product-ramp-1min-to-100con.jmx ubuntu@$JMETER:~/apache-jmeter-5.1.1/bin
scp -i sck_default.pem ../atdd/load/order-placement-ramp-1min-to-200con.jmx ubuntu@$JMETER:~/apache-jmeter-5.1.1/bin
