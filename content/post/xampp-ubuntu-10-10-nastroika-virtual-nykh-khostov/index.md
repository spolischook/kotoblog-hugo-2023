---
author: "Serhii Polishchuk"
title: "XAMPP + Ubuntu 10.10. Настройка виртуальных хостов"
date: 2011-02-16
tags: ["Ubuntu"]
draft: true
---
<!--more-->
<p>Мало установить <strong>XAMPP</strong> нужно еще на нем работать. Обычно одновременно на разработке находяться несколько сайтов, поэтому для полноценной работы нам необходимо настроить <strong>виртуальные хосты</strong>.<!--more--> Для начала раскомментируем <strong>виртуальные хосты</strong> в файле</p>

<pre>
<code class="bash">sudo gedit /opt/lampp/etc/httpd.conf</code></pre>

<p><span style="line-height: 1.6em;">убрав знак # в начале строки:</span></p>

<pre>
<code>Include etc/extra/httpd-vhosts.conf</code></pre>

<p><img alt="" class="img-responsive" src="/uploads/2011/02/vhost.png" style="width: 500px; height: 409px;" /></p>

<p>Перезапускаем сервер кнопкой на панеле, а лучше командой в командной строке:</p>

<pre>
<code class="bash">sudo /opt/lampp/lampp restart</code></pre>

<p>Так мы увидем все ошибки, если таковые будут. У меня к примеру выдало следующее:</p>

<pre>
<code class="bash">Warning: DocumentRoot [/www/docs/dummy-host.example.com] does not exist
Warning: DocumentRoot [/www/docs/dummy-host2.example.com] does not exist</code></pre>

<p>Как <strong>Apache</strong> разбирает файл конфигурации, до него мы еще дойдем, можно увидеть при выводе следующей комманды:</p>

<pre>
<code class="bash">/opt/lampp/bin/httpd -S</code></pre>

<p>У меня вывалились те же ошибки + конфигурация. Давайте начнем настраивать наш сервер для работы с <strong>виртуальными хостами</strong>. Для начала откроем <strong>файл конфигурации</strong>:</p>

<pre>
<code class="bash">sudo gedit /opt/lampp/etc/extra/httpd-vhosts.conf</code></pre>

<p>Проверяем чтобы строка:</p>

<pre>
<code>NameVirtualHost *:80</code></pre>

<p>была раскомментирована. В <strong>NameVirtualHost</strong> мы должны указать IP адрес и (по-возможности) порт на нашем сервере, который будет принимать запросы для наших хостов. Обычно, а наш случай совсем обычный :), когда любой IP на нашем сервере может быть использован, мы указываем *. Порт будем использовать 80 (<a href="http://ru.wikipedia.org/wiki/%D0%A1%D0%BF%D0%B8%D1%81%D0%BE%D0%BA_%D0%BF%D0%BE%D1%80%D1%82%D0%BE%D0%B2_TCP_%D0%B8_UDP" target="_blank">HTTP</a>) Следующим шагом у нас идет создание двух тестовых (для примера) виртуальных хоста заключенных в тэги:</p>

<pre>
<code>&lt;VirtualHost&gt;&lt;/VirtualHost&gt;</code></pre>

<p><span style="line-height: 1.6em;">С указанием порта в первом (в нашем случае это опять *:80) Минимально блок должен содержать дерективы </span><strong style="line-height: 1.6em;">ServerName</strong><span style="line-height: 1.6em;"> для определения имени </span><strong style="line-height: 1.6em;">виртуального хоста</strong><span style="line-height: 1.6em;">, и </span><strong style="line-height: 1.6em;">DocumentRoot</strong><span style="line-height: 1.6em;"> для определения папки размещения файлов проэкта. Первый блок считается блоком по умолчанию, именно поэтому при наборе http://localhost, там где раньше у нас был </span><strong style="line-height: 1.6em;">XAMPP</strong><span style="line-height: 1.6em;"> теперь странная ошибка:</span></p>

<p><img alt="" class="img-responsive" src="/uploads/2011/02/error-vhost.png" style="width: 533px; height: 281px;" /></p>

<p><span style="line-height: 1.6em;">Итак первая запись будет выглядеть следующим образом:</span></p>

<pre>
<code>&lt;VirtualHost *:80&gt;
    ServerAdmin admin@xampp.ua
    DocumentRoot /opt/lampp/htdocs/xampp
    ServerName xampp.ua
&lt;/VirtualHost&gt;</code></pre>

<p><span style="line-height: 1.6em;">Теперь наш виртуальный хост пропишем здесь:</span></p>

<pre>
<code class="bash">sudo gedit /etc/hosts</code></pre>

<p>В этом файле впишем строку:</p>

<pre>
<code>127.0.0.1	       xampp.ua</code></pre>

<p>Перезапускаем <strong>XAMPP</strong> из командной строки:</p>

<pre>
<code class="bash">sudo /opt/lampp/lampp restart</code></pre>

<p>Ошибок теперь должно вылазить на одну меньше, и по адресам http://localhost и http://xampp.ua у нас будет выдавать одинаковые данные:</p>

<p><img alt="" class="img-responsive" src="/uploads/2011/02/xamppua.png" style="width: 530px; height: 256px;" /></p>

<p>По аналогии остальные <strong>хосты</strong>. Можно сделать это с помощью автоматического скрипта, правда я не знаю насколько это необходимо. Жду Ваших комментариев.</p>
