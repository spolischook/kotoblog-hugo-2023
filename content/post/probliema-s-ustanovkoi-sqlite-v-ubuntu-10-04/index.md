---
author: "Serhii Polishchuk"
title: "Проблема с установкой sqlite в Ubuntu 10.04"
date: 2012-01-10
tags: ["Ubuntu"]
draft: true
---
<!--more-->
<p>На самом деле задача достаточно простая и я не знаю почему для меня она оказалась такой сложной. Суть проблемы в том что несмотря на то что я установил все возможные пакеты из репозитория касательно <strong>SQLite</strong> и даже явно прописал <strong>php.ini</strong>, поддержка БД не высвечивалась на phpinfo(). Проведя несколько часов поиска в интернете решение было найдено такое же простое как и проблема :) <strong>UPD.</strong> Рецепт от <a href="http://yakimoff.ru/" rel="external nofollow" target="_blank">Павла Якимова</a>, на работоспособность не проверялся, но имеет смысл попробовать:</p>

<pre>
<code class="bash">sudo service apache2 reload</code></pre>

<p>Если ничего не вышло то пробуем следущее - нужно снести все и заново установить:</p>

<pre>
<code class="bash">apt-get --purge remove php5*
sudo apt-get install php5 php5-sqlite php5-mysql
sudo apt-get install php-pear php-apc php5-curl
sudo apt-get autoremove</code></pre>

<p>&nbsp;</p>
