Hello Kafka
=======

## Prerequisite

+ Docker
+ docker-compose
+ [blockade](http://blockade.worstcase.io/en/latest/index.html)

## Quick Start

```
$ docker-compose up -d
$ docker-compose ps | tail -n +3 | awk '{print $1}' | xargs blockade add
```

```
$ bash bin/reset-topic.sh
```

```
## basic
$ docker build -t kafka-python:0.1 -f client/Dockerfile client/
$ docker-compose run --rm kafka-python python producer.py 100000 0 0 default
$ docker-compose run --rm kafka-python:0.1 python consumer.py

## kafka-console-{producer,consumer}
$ docker-compose exec kafka1 kafka-console-producer --broker-list kafka1:9092,kafka2:9092,kafka3:9092 --topic default
$ docker-compose exec kafka1 kafka-console-consumer --bootstrap-server kafka1:9092,kafka2:9092,kafka3:9092 --topic default --from-beginning
```

```
$ blockade destroy
$ docker-compose down
```

## Hardcode

+ kafka{1,2,3} & zookeeper{1,2,3} hostname resolve depends on docker-compose
+ kafka chroot /kafka/cluster-x

## Concepts

+ Ack Mode
+ Group
+ Configuration
    + Topic-Level
        + Partition
        + Replicate
        + cleanup.policy
        + compression.type
    + Producer
        + key.serializer
        + value.serializer
        + acks
        + compression.type
        + retries
        + connections.max.idle.ms
        + delivery.timeout.ms
    + Consumer
        + key.deserializer
        + value.deserializer
        + group.id
        + heartbeat.interval.ms
        + max.partition.fetch.bytes
        + session.timeout.ms

## Deploy On Production

```
kafka1 $ docker run -d \
    --name kafka-default \
    --restart always \
    --net=host \
    -v /data/kafka/data:/var/lib/kafka/data \
    -v /data/kafka/log:/var/log/kafka \
    -e KAFKA_BROKER_ID=1 \
    -e KAFKA_ZOOKEEPER_CONNECT=zookeeper1,zookeeper2,zookeeper3/kafka/default \
    -e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://kafka1:9092 \
    -e KAFKA_LOG4J_LOGGERS="kafka.controller=INFO,state.change.logger=INFO" \
    -e KAFKA_LOG4J_ROOT_LOGLEVEL=WARN \
    -e KAFKA_TOOLS_LOG4J_LOGLEVEL=ERROR \
    -e KAFKA_HEAP_OPTS="-Xmx8G -Xms8G -server -XX:+UseG1GC -XX:MaxGCPauseMillis=20 -XX:InitiatingHeapOccupancyPercent=35 -XX:G1HeapRegionSize=16M -XX:MinMetaspaceFreeRatio=50 -XX:MaxMetaspaceFreeRatio=80" \
    -e KAFKA_JMX_PORT=9999 \
    confluentinc/cp-kafka:5.2.1
```

## References

+ [Confluent Documentation](https://docs.confluent.io/current/)
    + [On-Premises Deployments](https://docs.confluent.io/current/installation/installing_cp/index.html)
    + [Quick Start using Community Components (Docker)](https://docs.confluent.io/current/quickstart/cos-docker-quickstart.html)
+ [Apache kafka Documentation](https://kafka.apache.org/)
    + [Quick Start](https://kafka.apache.org/quickstart)
+ [Kafka Cluster (Docker)](https://github.com/confluentinc/cp-docker-images/blob/master/examples/kafka-cluster/docker-compose.yml)
+ [How to Lose Messages on a Kafka Cluster](https://github.com/Vanlightly/ChaosTestingCode/tree/master/Kafka)
+ [blockade](https://blockade.worstcase.io)
+ [Monitoring Kafka performance metrics](https://www.datadoghq.com/blog/monitoring-kafka-performance-metrics/)
