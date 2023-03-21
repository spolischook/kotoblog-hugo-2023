---
author: "Serhii Polishchuk"
title: "Принтер Canon LBP на Ubuntu 10.10"
date: 2011-01-10
tags: ["Ubuntu"]
draft: true
---
<!--more-->
<p>Нашел в интернете полностью автоматическое решение для установки практически всех популярных лазерных принтеров <strong>Canon серии LBP</strong>. Интересно, что с принтерами HP никаких проблем нет, все драйвера под <strong>Linux</strong> идут достаточно качественные, а некоторые даже лучше чем <strong>Windows</strong>. То ли дело <strong>Canon</strong>. Честно говоря, кто хоть раз сталкивался с необходимостью установить &quot;японцев&quot; под <strong>Ubuntu</strong> не раз проклинал и производителя принтеров и саму Канониан, что складывается впечатление что драйвера пишут заядлые ненавистники <strong>Линукса</strong>. И вот совсем недавно добрый программист по имени Раду (<a href="http://radu.cotescu.com/" target="_blank">официальный сайт</a>) выложил самописный скрипт установки драйверов <!--more-->Если в кратце то инструкция такова:</p>

<p>1. Скачать архив <a href="http://ubuntuone.com/p/Xgf/">отсюда</a> или <a href="http://codebin.cotescu.com/canon/lbp_driver/CanonCAPTdriver.tar.gz">отсюда</a></p>

<p>2. Распаковать архив на жесткий диск</p>

<p>3. Через терминал заходим в директорию (открыть папку, жми Ctrl+L, скопируй адрес, в терминале пишем cd &lt;адрес&gt;)</p>

<p>4. Запускаем скрипт следующей коммандой:</p>

<pre>
<code class="bash">sudo bash ./canonLBP_install.sh &lt;Имя_принтера&gt; </code></pre>

<p><span style="line-height: 1.6em;">&lt;Имя_принтера&gt; заменяем на модель из перечисленных ниже: Скрипт поддерживает следующие модели: </span></p>

<p><strong style="line-height: 1.6em;">LBP-1120/810, LBP-1210, LBP2900, LBP3000, LBP3010, LBP3018, LBP3050, LBP3100, LBP3108, LBP3150, LBP3200, LBP3210, LBP3250, LBP3300, LBP3310, LBP3500, LBP5000, LBP5050, LBP5100, LBP5300, LBP7200C, LBP6300dn, LBP9100Cdn</strong><span style="line-height: 1.6em;"> </span></p>

<p><span style="line-height: 1.6em;">Все устанавливается в автоматическом режиме. Красота, не правда ли :) Статья в оригинале, обсуждение и общение с автором </span><a href="http://radu.cotescu.com/how-to-install-canon-lbp-printers-in-ubuntu/" style="line-height: 1.6em;" target="_blank">здесь.</a><span style="line-height: 1.6em;"> </span></p>

<p><strong style="line-height: 1.6em;">UPD: </strong><span style="line-height: 1.6em;">Понадобилось мне его еще расшарить по сети, и это оказалось проще-простого. Проверьте адрес </span><a href="http://localhost:631/" style="line-height: 1.6em;" target="_blank">http://localhost:631/</a><span style="line-height: 1.6em;"> Если все норм, то должна открыться панель управления CUPS. В Windows выбираем установка сетевого принтера и прописываем путь http://ip_линукс_машины:631/printers/имя_принтера У меня к примеру путь выглядит так: http://192.168.0.100:631/printers/LBP-1120 upd. Запуск ccpd</span></p>

<pre>
<code class="bash">sudo /etc/init.d/ccpd start</code></pre>
