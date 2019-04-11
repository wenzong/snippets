# -*- coding: utf-8 -*-

import datetime
import json
import time

from kafka import KafkaConsumer


def run():
    consumer = KafkaConsumer(
            'default', # user_product_events
            group_id='user_observer', # product_observer ...
            bootstrap_servers=[
                'kafka1:9092',
                'kafka2:9092',
                'kafka3:9092',])

    i = 0
    for message in consumer:
        i += 1
        print(message.value)
        if i % 1000 == 0:
            print(i)


if __name__ == '__main__':
    run()
