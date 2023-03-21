---
author: "Serhii Polishchuk"
title: "www-data php mail и ssmtp - неправильный FROM в письмах"
date: 2012-07-13
tags: []
draft: true
---
<!--more-->
<p>Хоть я и указывал в заголовках phpmail от кого письмо все равно приходило от www-data &lt;comemail@mail.ru&gt;. Оказалось что по-умолчанию ssmtp настроен таким образом чтобы пользователь не мог изменить этот самый FROM. Посему заходим в конфиг:</p>

<pre>
<code class="bash">sudo vim /etc/ssmtp/ssmtp.conf</code></pre>

<p>и раскоментируем строку FromLineOverride=YES убрав в начале строки знак диеза (для тех кто еще не в курсе) После чего перезапускаем сервис:</p>

<pre>
<code class="bash">sudo /etc/init.d/sendmail restart</code></pre>
