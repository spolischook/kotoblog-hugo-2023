---
author: "Serhii Polishchuk"
title: "Установка Composer в Ubuntu 12.04"
date: 2012-11-12
tags: ["Composer"]
draft: true
---
<!--more-->
Правильно устанавливать так:
    curl -s https://getcomposer.org/installer | php
    sudo mv composer.phar /usr/local/bin/composer

Теперь вместо
    php composer.phar install

Мы будем писать:
    composer install

Ну и в дальнейшем также. Согласитесь так же намного проще :)
