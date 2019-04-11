docker-compose exec kafka1 \
    kafka-topics \
    --zookeeper zookeeper1:2181,zookeeper2:2181,zookeeper3:2181/kafka/cluster-x \
    $@
