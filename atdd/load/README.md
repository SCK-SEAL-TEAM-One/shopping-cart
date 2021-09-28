# ISSUES

## Add HEAP

```cmd
================================================================================
Don't use GUI mode for load testing !, only for Test creation and Test debugging.
For load testing, use CLI Mode (was NON GUI):
   jmeter -n -t [jmx file] -l [results file] -e -o [Path to web report folder]
& increase Java Heap to meet your test requirements:
   Modify current env variable HEAP="-Xms1g -Xmx1g -XX:MaxMetaspaceSize=256m" in the jmeter batch file
Check : https://jmeter.apache.org/usermanual/best-practices.html
================================================================================
```

## MYSQL MAX CONNECTION

[](https://stackoverflow.com/questions/14331032/mysql-error-1040-too-many-connection)

echo "set global max_connections = 500;" | kubectl exec -it $(kubectl get pods | grep store-database-deployment| awk '{print $1}') -- /usr/bin/mysql -u root --password=root --default-character-set=utf8  toy

echo "show variables like 'max_connections';" | kubectl exec -it $(kubectl get pods | grep store-database-deployment| awk '{print $1}') -- /usr/bin/mysql -u sealteam --password=sckshuhari --default-character-set=utf8  toy

## How do I troubleshoot DNS failures with Amazon EKS?

 kubectl edit deployment coredns -n kube-system

 set replica to 4

## Run with Parameter

```cmd
jmeter -n -t flow-ramp-1min-to-100con.jmx -l flow-ramp-1min-to-100con.jtl -e -o flow-ramp-1min-to-100con -Dhost=host.name.com
```

## Non HTTP response message: The target server failed to respond: Is my server failing to handle load

[Non HTTP response message: The target server failed to respond: Is my server failing to handle load](https://stackoverflow.com/questions/27942583/non-http-response-message-the-target-server-failed-to-respond-is-my-server-fai/27943565)

## Apache JMeter - Basics, Common Errors, Issues, Best Practices And Troubleshooting

[Apache JMeter - Basics, Common Errors, Issues, Best Practices And Troubleshooting](https://www.linkedin.com/pulse/apache-jmeter-basics-common-errors-issues-best-practices-prasad/)
