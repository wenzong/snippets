# -*- coding: utf-8 -*-

import datetime
import json
import time

from kafka import KafkaConsumer


def run():
    consumer = KafkaConsumer(
            'demo',
            group_id='demo.group',
            bootstrap_servers=[
                '127.0.0.1:9092',
            ])

    i = 0
    for message in consumer:
        i += 1
        print(json.loads(message.value.decode()))
        if i % 1000 == 0:
            print(i)


if __name__ == '__main__':
    run()
