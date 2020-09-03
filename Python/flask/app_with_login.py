# NOTE:
#
# https://github.com/maxcountryman/flask-login/issues/458
# https://blog.miguelgrinberg.com/post/how-secure-is-the-flask-user-session
#
# > Flask uses cookie based sessions by default, but there is support for custom sessions that store data in other places.
# > In particular, the Flask-Session extension is very interesting, as it stores the user session data in the server, giving you a variety of storage options such as plain files, Redis, relational databases, etc.


from flask import Flask, request
from flask_login import LoginManager, login_required, login_user, UserMixin
from flask_cookie_decode import CookieDecode


class User(UserMixin):
    def __init__(self, user_id):
        self.user_id = user_id

    def get_id(self):
        return self.user_id

app = Flask(__name__)
app.config.update(
    {'SECRET_KEY': 'qwerty'}
)
login_manager = LoginManager()
login_manager.init_app(app)

# > FLASK_APP=app_with_login.py flask cookie decode {{ cookie }}
# >>> TrustedCookie(contents={'_fresh': True, '_id': 'xxx', 'user_id': 1}, expiration='2020-10-04T08:43:37')
cookie = CookieDecode()
cookie.init_app(app)

@app.route('/login')
def login():
    user = User(1)
    login_user(user)
    return '200'

@app.route('/')
@login_required
def index():
    return "hello world"

@login_manager.user_loader
def load_user(user_id):
    return User(user_id)
