Laravel
=========

## Prerequisite

+ Composer

## Install Laravel

```
$ php composer.phar create-project --prefer-dist laravel/laravel demo
```

options:

```
-vvv --repository=http://packagist.phpcomposer.com --no-secure-http
```

## Install Packages & Configure composer

+ In Mainland China, Update composer.json for packages mirror

```
{
    ...

    "config": {
        "secure-http": false
    },
    "repositories": [
        {
            "packagist": false
        },
        {
            "type": "composer",
            "url": "http://packagist.phpcomposer.com"
        }
    ],

    ...
}
```
