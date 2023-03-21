---
author: "Serhii Polishchuk"
title: "cURL + PHP5 + Ubuntu правильная установка"
date: 2011-10-15
tags: ["PHP"]
draft: true
---
<!--more-->
<div class="alert alert-warning">Внимание! Данная инструкция приминима к 10.10.<br />
Для 12.04 используйте <a href="/articles/ustanovka-curl-ubuntu-12-04">Установка cUrl на Ubuntu 12.04</a></div>

<p>Уже в который раз наступаю на одни и теже грабли. В сети полно мануалов по установке&nbsp;<strong>cURL + PHP5 + Ubuntu</strong> но они все неправильные. Вот как устанавливаю я:</p>

<pre>
<code class="bash">sudo apt-get install curl libcurl3 libcurl3-dev php5-curl php5-mcrypt</code></pre>

<p>После этого в конец файла <strong>php.ini</strong> (у меня по адресу <strong>/etc/php5/apache2/php.ini</strong>) вводим следующее:</p>

<pre>
<code class="ini">;cUrl
extension=/usr/lib/php5/20090626/curl.so</code></pre>

<p>естественно что путь к файлу <strong>curl.so</strong> у Вас может отличатся - я свой нашел через поиск. Теперь перезагружаем<strong> веб сервер</strong>:</p>

<pre>
<code class="bash">/etc/init.d/apache2 restart</code></pre>

<p>Все. Заходите в phpinfo вашего сервера и наблюдайте такую картину:</p>

<p><img alt="" class="img-responsive" src="/uploads/2011/10/curl.jpg" style="width: 544px; height: 860px;" /></p>

<p>Теперь можете использовать &nbsp;&nbsp;<strong>cURL + PHP5 + Ubuntu </strong>в свое удовольствие! :)</p>
