---
author: "Serhii Polishchuk"
title: "Eclipse PDT + Ubuntu 10.10. Часть 1"
date: 2011-01-18
tags: ["Ubuntu"]
draft: true
---
<!--more-->
<p>В принципе нет ничего сложного в том чтобы установить <strong>Eclipse</strong> из репозитория <strong>Ubuntu</strong>. Но в той версии нет самого главного инструмента PHP разработчика - <strong>PHP Development Tools (PDT)</strong>. Об установке <strong>PDT в Eclipse на Ubuntu 10.10 </strong>и пойдет речь в этой статье. <!--more-->Собственно устанавливаем сам <strong>Eclipse</strong> из репозитория (Приложения-&gt;Центр приложения <strong>Ubuntu</strong>), в строку поиска забиваем <strong>Eclipse</strong>, находим устанавливаем. При первом запуске нас попросят указать рабочую папку (Select workspace). Наши&nbsp;проекты&nbsp;будем сохранять в папке: <code>/opt/lampp/htdocs</code> Здесь мы создаем новую папку которую назовем напр.<strong> joomla_development</strong> и ее выбираем в качестве папки проекта. Ставим галочку Use this as a defoult дабы не возвращаться к выбору снова при запуске. Запустили? Закроем окно приветствия</p>

<p><img alt="" class="img-responsive" src="/uploads/2011/01/eclipse-start.png" style="width: 550px; height: 391px;" /></p>

<p>и получим нормальный ;) вид <strong>Eclipse</strong></p>

<p><img alt="" class="img-responsive" src="/uploads/2011/01/eclipse.jpg" style="width: 500px; height: 375px;" /></p>

<p>Мы установили базовый функционал, но т.к. мы будем заниматься <strong>PHP</strong> разработкой нам нужно установить <strong>PDT (PHP Development Tools)</strong>. Сделам мы это так... Открываем <strong>Eclipse</strong> (если мы его закрыли) Help-&gt;Install New Software... Work with: <code>http://download.eclipse.org/releases/galileo/</code> Если дальше <strong>Eclipse</strong> вывалится в ошибку пробуем это: <code>http://download.eclipse.org/releases/helios</code> Находим Web, XML, and Java EE Development в нем подпункт <strong>PHP</strong><strong> Development Tools (PDT) SDK Feature</strong> ставим галочку и нажимаем Next</p>

<p><img alt="" class="img-responsive" src="/uploads/2011/01/eclipse2.jpg" style="width: 500px; height: 375px;" /></p>

<p>Далее нас попросят принять лицензионное соглашение - принимаем без доли сомнения (OpenSource как никак :)</p>

<p><img alt="" class="img-responsive" src="/uploads/2011/01/eclipse3.jpg" style="width: 500px; height: 375px;" /></p>

<p>Далее после длительной установки в меню Windows-Preferens, у нас появиться дополнительный пункт <strong>PHP</strong>.</p>

<h2>Конфигурируем XDebug</h2>

<p><em><strong>Инфо:</strong> Следующие действия не являются обязательными, а лишь дополнительными, т.к. <strong>Eclipse</strong> будет работать и без этого.</em></p>

<p><img alt="" class="img-responsive" src="/uploads/2011/01/eclipse-pref.jpg" style="width: 500px; height: 283px;" /></p>

<p>Как видим на нужно внести изменения в <strong>PHP Debuger</strong>, и проверить галочку Break at First Line/ Жмем configure возле <strong>PHP Debuger</strong> и указываем выбранный порт в <strong>php.ini</strong> (<a href="http://kotoblog.pp.ua/ubuntu/xamppxdebug-na-ubuntu-1010.html">мы устанавливали XAMPP+XDEBUG ранее</a>) Если забыли порт, открываем php.ini следующей командой: <code>sudo gedit /opt/lampp/etc/php.ini</code> Порт указали в xdebug.remote_port, в моем случае это 10000</p>

<p><img alt="" class="img-responsive" src="/uploads/2011/01/eclipse-deb-pref.jpg" style="width: 500px; height: 562px;" /></p>

<p>Также я дал <strong>Zend Debuger </strong>- у 10001 порт. В некоторых случаях мы можем получить ошибку: <code>A Runtime Error has occurred. Do you wish to Debug? Line: 1 Error: Syntax error</code> Для устранения необходимо опцию &quot;Output Capture Settings / Capture stdout&quot; переключить с &quot;copy&quot; на &quot;off&quot;. Дальше переходим на вкладку: <code>PHP-&gt; Debug-&gt;Workbench Options</code> Выбираем установки так как на рисунке ниже:</p>

<p><img alt="" class="img-responsive" src="/uploads/2011/01/eclipse-pref2.jpg" style="width: 500px; height: 274px;" /></p>

<p>В принципе с этими настройками вполне можно экспериментировать для получения наилучших результатов. Keep smiling. :)</p>
