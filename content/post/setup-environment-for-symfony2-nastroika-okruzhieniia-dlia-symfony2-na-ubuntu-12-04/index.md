---
author: "Serhii Polishchuk"
title: "Setup environment for Symfony2 - Настройка окружения для Symfony2 на Ubuntu 12.04"
date: 2012-11-12
tags: ["Symfony2", "Ubuntu"]
draft: false
---
<!--more-->
Вкратце описание настройки среды для разработки на фреймворке **Symfony2**

    sudo su
    apt-get install mysql-server mysql-client
    apt-get install apache2
    apt-get install php5 libapache2-mod-php5
    /etc/init.d/apache2 restart
    apt-get install php5-mysql php5-curl php5-gd php5-intl php-pear php5-imagick php5-imap php5-mcrypt php5-memcache php5-ming php5-ps php5-pspell php5-recode php5-snmp php5-sqlite php5-tidy php5-xmlrpc php5-xsl php5-xdebug
    /etc/init.d/apache2 restart

Устанвливаем MongoDB если нужно:

    pecl install mongo

Туториал по установке [здесь](http://docs.mongodb.org/manual/tutorial/install-mongodb-on-ubuntu/ "Mongodb instalation manual on Ubuntu")

    apt-get install php-apc
Включем же mod rewrite
    gedit /etc/apache2/sites-available/default

Заменяем все AllowOverride None на AllowOverride All

    a2enmod rewrite
    /etc/init.d/apache2 restart
    pecl install xdebug
    apt-get install phpmyadmin
    apt-get install git-core

Не забудьте после этого прописать некоторые настройки в php.ini Для начала найдем куда установился xdebug:
    find / -name 'xdebug.so' 2> /dev/null

Теперь с имеющимися данными идем в php.ini
    gedit /etc/php5/apache2/php.ini
Вставим в конец файла следующее:
    extension=mongo.so

    ;xDebug Configuration starts
       zend_extension=/usr/lib/php5/20100525/xdebug.so
       xdebug.max_nesting_level=250
       xdebug.var_display_max_depth=10
       xdebug.remote_enable=true
       xdebug.remote_handler=dbgp
       xdebug.remote_mode=req
       xdebug.remote_port=9000
       xdebug.remote_host=127.0.0.1
       xdebug.idekey=phpstorm-xdebug
       xdebug.remote_autostart=1
       xdebug.remote_log=/var/log/apache2/xdebug_remote.log
    ;xDebug Configuration ends

В том же **php.ini** необходимо изменить:
    ;date.timezone =
меняем на
    date.timezone = Europe/Kiev
Выключаем короткие теги <? и ?>
    short_open_tag = On
меняем на
    short_open_tag = Off
Включаем отображение ошибок
    display_errors = Off
меняем на
    display_errors = On
Включаем отображение html ошибок:
    html_errors = Off
меняем на
    html_errors = On

На свое усмотрение можно изменить параметры: max_execution_time memory_limit post_max_size upload_max_filesize
    /etc/init.d/apache2 restart

Теперь если все гуд вот этот файл - [check_gh.php](http://kotoblog.pp.ua/wp-content/uploads/2012/11/check_gh.php_.zip "Check your configuration for Symfony2") будет зелёненьким.
Устанавливаем Composer если еще этого не сделали:
    sudo apt-get install curl
    curl -sS https://getcomposer.org/installer | php
    mv composer.phar /usr/local/bin/composer

Также для того чтобы компосер знал что у нас установлена монга нужно это указать в php.ini для cli. В файл:
    gedit /etc/php5/cli/php.ini

нужно добавить
    extension=mongo.so
