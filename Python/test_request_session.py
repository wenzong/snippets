import logging
import sys
import time

import requests

logger = logging.getLogger()
logger.setLevel(logging.INFO)

stdout = logging.StreamHandler(sys.stdout)
stdout.setLevel(logging.INFO)
logger.addHandler(stdout)

sess = requests.Session()
# adapter = requests.adapters.HTTPAdapter(max_retries=1)
# sess.mount('http://', adapter)

while True:
    try:
        resp = sess.get('https://www.baidu.com', timeout=1)
    except Exception as e:
        logger.exception(e)
    finally:
        logger.info(resp)
        time.sleep(1)
