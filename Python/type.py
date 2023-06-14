from typing import List, Optional
from pydantic import BaseModel, create_model


external_data = {
    'id': '123',
    'signup_ts': '2019-06-01 12:22',
    'friends': [1, 2, '3'],
}

User = type("User", (BaseModel,), {"id": 0, "signup_ts": '123123123123', "friends": [1,2,3]})
user = User(**external_data)
print(user.schema())

User = create_model('User', id=(int, ...), signup_ts=(str, None), friends=(List[int], None))
user = User(**external_data)
print(user.schema())
