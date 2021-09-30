# Load Injector

## Install Jmeter

### Step 1 - Create Atlantic.Net Cloud Server

First, Login to your Atlantic.Net Cloud Server. Create a new server, choosing Ubuntu 20.04 as the operating system, with at least 2 GB RAM. Connect to your Cloud Server via SSH and log in using the credentials highlighted at the top of the page.

Once you are logged into your Ubuntu 20.04 server, run the following command to update your base system with the latest available packages.

```cmd
sudo apt-get update -y
```

### Step 2 - Install Java

Apache JMeter is a Java-base application so Java must be installed in your system. You can install it by running the following command:

```cmd
sudo apt-get install openjdk-8-jdk -y
```

After installing Java, verify the installed version of Java with the following command:

```cmd
java -version

You should get the following output:

openjdk version "1.8.0_272"
OpenJDK Runtime Environment (build 1.8.0_272-8u272-b10-0ubuntu1~20.04-b10)
OpenJDK 64-Bit Server VM (build 25.272-b10, mixed mode)
OpenJDK 64-Bit Server VM (build 25.272-b10, mixed mode)
```

### Step 3 - Install Apache JMeter

By default, Apache JMeter is not available in the Ubuntu 20.04 default repository. So you will need to download it from its official website:

You can download it with the following command:

```cmd
wget https://downloads.apache.org//jmeter/binaries/apache-jmeter-5.3.zip
```

install UnZip

```cmd
sudo apt install unzip
```

Once downloaded, unzip the downloaded file with the following command:

```cmd
unzip apache-jmeter-5.3.zip
```

### Step 4 - Install Apache JMeter Plugin

```cmd
export JMETER=18.138.237.67

scp -i sck_default.pem ../../atdd/load/*.jar ubuntu@$JMETER:~/apache-jmeter-5.3/lib
```

### Step 5 - Upload Performance Script

```cmd
export JMETER=18.138.237.67

scp -i sck_default.pem ../../atdd/load/flow-ramp-1min-to-100con.jmx ubuntu@$JMETER:~/apache-jmeter-5.3/bin
```

### Step 6 - Run Performance Test

```cmd
cd apache-jmeter-5.3/bin/

java -jar ApacheJMeter.jar -n -t flow-ramp-1min-to-100con.jmx -l flow-ramp-1min-to-100con.jtl -e -o flow-ramp-1min-to-100con
java -jar ApacheJMeter.jar -n -t flow-ramp-1min-to-100con.jmx -l flow-ramp-1min-to-100con.jtl -e -o flow-ramp-1min-to-100con -Dhost=a2dbbec4988744ca29d16c438d471528-421367058.ap-southeast-1.elb.amazonaws.com
```

### Step 7 - Get Jmeter Report

```cmd
export JMETER=18.138.237.67
scp -i sck_default.pem -r ubuntu@$JMETER:~/apache-jmeter-5.3/bin/flow-ramp-1min-to-100con ../../atdd/load/
scp -i sck_default.pem ubuntu@$JMETER:~/apache-jmeter-5.3/bin/flow-ramp-1min-to-100con.jtl ../../atdd/load/

open ../../atdd/load/flow-ramp-1min-to-100con/index.html

```

### Step 5.2 - Upload Performance Script

```cmd
export JMETER=18.138.237.67

scp -i sck_default.pem ../../atdd/load/flow-ramp-1min-to-300con.jmx ubuntu@$JMETER:~/apache-jmeter-5.3/bin
```

### Step 6.2 - Run Performance Test

```cmd
cd apache-jmeter-5.3/bin/

java -jar ApacheJMeter.jar -n -t flow-ramp-1min-to-300con.jmx -l flow-ramp-1min-to-300con.jtl -e -o flow-ramp-1min-to-300con -Dhost=a2dbbec4988744ca29d16c438d471528-421367058.ap-southeast-1.elb.amazonaws.com

```

### Step 7.2 - Get Jmeter Report

```cmd
export JMETER=18.138.237.67
scp -i sck_default.pem -r ubuntu@$JMETER:~/apache-jmeter-5.3/bin/flow-ramp-1min-to-300con ../../atdd/load/
scp -i sck_default.pem ubuntu@$JMETER:~/apache-jmeter-5.3/bin/flow-ramp-1min-to-300con.jtl ../../atdd/load/

open ../../atdd/load/flow-ramp-1min-to-300con/index.html
```

```cmd
scp -i sck_default.pem ../../atdd/load/api-get-all-product-ramp-1min-to-100con.jmx ubuntu@$JMETER:~/apache-jmeter-5.3/bin
scp -i sck_default.pem ../../atdd/load/api-get-by-id-product-ramp-1min-to-100con.jmx ubuntu@$JMETER:~/apache-jmeter-5.3/bin
scp -i sck_default.pem ../../atdd/load/order-placement-ramp-1min-to-200con.jmx ubuntu@$JMETER:~/apache-jmeter-5.3/bin
```

reference [How to Install Apache JMeter on Ubuntu 20.04](https://dev.to/hitjethva/how-to-install-apache-jmeter-on-ubuntu-20-04-2di9)
