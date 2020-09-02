# -*- coding: utf-8 -*-

import datetime
import json
import time
from confluent_kafka import Consumer, KafkaError


c = Consumer({
    'bootstrap.servers': '127.0.0.1:9092',
    'group.id': 'demo',
    'auto.offset.reset': 'earliest'
})

c.subscribe(['demo'])

while True:
    msg = c.poll(1.0)

    if msg is None:
        continue
    if msg.error():
        print("Consumer error: {}".format(msg.error()))
        continue

    print('Received message: {}'.format(msg.value().decode('utf-8')))

c.close()
