## How to start flask app with uWSGI/gunicorn.

+ gunicorn: ve gunicorn -b(backend) :5000 -w(worker) 1 my_app(entry):app(Application object)
+ uwsgi: ve uwsgi --http 127.0.0.1:5000 --wsgi-file my_app.py  --callable app

## What happened under app.route


