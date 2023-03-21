---
author: "Serhii Polishchuk"
title: "Setup environment for Symfony2 - Настройка окружения для Symfony2 на Ubuntu 14.04"
date: 2014-12-28
tags: ["Symfony2", "Ubuntu", "PHP"]
draft: false
---
<!--more-->
This is simple instruction on how to setup environment for Symfony2 and Ubuntu 14.04:

    https://gist.github.com/31fac5fcfb6eea4f2787

Edit file:

     vim /etc/apache2/sites-available/000-default.conf

Edit DocumentRoot to your value, I prefer "/var/www"

Next step we need to enable mod rewrite:

    a2enmod rewrite
    
To allow use htaccess files in our projects we need to enable this option for our dir (/var/www in this case). 

    vim /etc/apache2/apache2.conf

Replace "AllowOverride None" to "AllowOverride All" for <Directory /var/www/> section.
Restart apache

    service apache2 restart

Install xdebug if it needed:

    pecl install xdebug

Find place where xdebug was installed:

    find / -name 'xdebug.so' 2> /dev/null

Copy this value and edit your php.ini:

    vim /etc/php5/apache2/php.ini

Pase to end of file (Shift + F - go to end of file in vim):

    ;xDebug Configuration starts
        zend_extension=/usr/lib/php5/20121212/xdebug.so
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
Also you need to replace/edit some values in php.ini. 
First of all, set date.timezone. You can choose it [at php.net site](http://php.net/manual/en/timezones.php "List of Supported Timezones"). 
Second turn on display_errors if it needed.
Restart apache

    service apache2 restart

Install composer globally:

    curl -sS https://getcomposer.org/installer | php
    mv composer.phar /usr/local/bin/composer

Install nodejs:

    curl -sL https://deb.nodesource.com/setup | sudo bash -
    apt-get install -y nodejs

Install bower:

    npm install -g bower

That all.
