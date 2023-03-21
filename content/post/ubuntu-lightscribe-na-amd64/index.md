---
author: "Serhii Polishchuk"
title: "Ubuntu LightScribe на AMD64"
date: 2012-01-02
tags: ["Ubuntu"]
draft: true
---
<!--more-->
<p>Поздравляю всех с наступившим 2012 Новым Годом. Как всегда в этом году я делал подарки ввиде красивых <strong>LightScribe</strong> дисков, и как всегда откладывал это в долгий ящик. Т.к. у меня в 2011 появился ковый комп на базе процессора <strong>AMD64</strong>, которым я к стати очень доволен, то процес установки <strong>LightScvribe</strong> приложения будет немного отличатся нежели для 32 битных систем. Итак приступим...<!--more--> Первое что мы делаем &nbsp;- это открываем консоль и вписываем:</p>

<pre><code class="bash">apt-get install ia32-libs apt-get install alien</code></pre>

<p>Если вдруг мы еще не устанавливали чуждые для <strong>64</strong> битных систем <strong>32</strong>-х битные приложения, эти библиотеки нам очень пригодятся.</p>

<pre><code class="bash">apt-get install libstdc++5 </code></pre>

<p>Устаревшая библиотека <strong>libstdc++5</strong> понадобится в дальнейшем для работы нашего приложения. Теперь скачаем приложения для рисовалки дисков <strong>LightScribe</strong>:</p>

<pre><code class="bash">wget http://download.lightscribe.com/ls/lightscribe-1.18.24.1-linux-2.6-intel.deb 
wget http://www.lacie.com/download/drivers/4L-1.0-r6.i586.rpm</code></pre>

<p>Пожалуйста, обратите внимание что версия первого пакета может отличатся от той что у меня - можете уточнить актуальную версию на <a href="http://lightscribe.com/" target="_blank" title="Официальный сайт LightScribe">официальном сайте</a>, там же можете и скачать. Оба пакета, по большому счету нам не подходят, поэтому установка и отличается:</p>

<pre><code class="bash">dpkg -i --force-all lightscribe-1.18.24.1-linux-2.6-intel.deb</code></pre>

<p>Не забудьте по версию пакета! Теперь преобразовуем rpm в tar/gzip</p>

<pre><code class="bash">alien -t 4L-1.0-r6.i586.rpm</code></pre>

<p>Теперь в .deb</p>

<pre><code class="bash">alien 4L-1.0.tgz</code></pre>

<p>Устанавливаем</p>

<pre><code class="bash">dpkg -i 4l_1.0-2_all.deb</code></pre>

<p>Запускаем</p>

<pre><code class="bash">4L-gui</code></pre>

<p>После всего выше перечисленного, приложение запустилось но работать отказалось, т.к. не определило привод - ошибка <strong>No drives have been selected</strong>. На просторах интернета нашлось простое решение. Приложение ищет нужные ей библиотеки только в папке <strong>lib32</strong>, поэтому нужно создать символьные ссылки, или просто скопировать недостающие библиотеки, например так:</p>

<pre><code class="bash">ln -s /usr/lib/liblightscribe.so.1 /usr/lib32/liblightscribe.so.1</code></pre>

<p>или так, как в моем случае:</p>

<pre><code class="bash">sudo ln -s /usr/lib/libstdc++.so.5 /usr/lib32/libstdc++.so.5</code></pre>

<p>Ошибку которую возникает вы можете подсмотреть в консоли с которой запускали приложение. На этом все :)</p>
