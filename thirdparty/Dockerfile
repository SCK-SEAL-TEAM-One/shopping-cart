FROM golfapipol/stubby4j-openjdk8-alpine:6.0.1
ADD ./bank-gateway/visa.yml /data/visa.yml
CMD java -jar /app/stubby4j-6.0.1.jar -l 0.0.0.0 -d /data/visa.yml
