---
author: "Serhii Polishchuk"
title: "MongoDB Symfony2 Travis-ci"
date: 2013-11-21
tags: []
draft: true
---
<!--more-->
Задался целью подключить свой любимый OpenSource **Symfony2** проект, к самому популярному сервису тестирования **Travis-ci**. Изначально планировалось проект делать на **MongoDB** в качестве хранилища, поэтому задача выдалась нетривиальной. Первое что нужно знать, это то что **Travis-ci** уже имеет поддержку множества популярных pecl расширений. Их список можно посмотреть здесь - http://about.travis-ci.org/docs/user/languages/php/#Preinstalled-PHP-extensions. Но они все отключены, из-за чего нужно бы включить те что нам нужны. Для этого есть два способа. Я выбрал самый простой, т.к. мне нужно было всего **mongodb** подключить.
<!--more-->

    before_script:
    - echo "extension = mongo.so" >> ~/.phpenv/versions/$(phpenv version-name)/etc/php.ini

Также я явно указал на то что буду использовать **Монгу**:

    services: mongodb

Это все. Остальную часть **.travis.yml** я оставил без изменений. Итого, я получил следующее:

    language: php
 
    php:
    - 5.3.3
    - 5.3
    - 5.4
    
    before_script:
    - echo "extension = mongo.so" >> ~/.phpenv/versions/$(phpenv version-name)/etc/php.ini
    - composer install -n  --prefer-source
    
    services: mongodb
    
    script: phpunit -c app
