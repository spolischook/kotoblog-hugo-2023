---
author: "Serhii Polishchuk"
title: "Установка cUrl Ubuntu 12.04"
date: 2012-07-03
tags: ["Ubuntu"]
draft: true
---
<!--more-->
<p><a href="http://kotoblog.pp.ua/php/curl-php5-ubuntu-pravilnaya-ustanovka.html" title="cURL + PHP5 + Ubuntu правильная установка">В одной из своих статтей</a> я упоминал про установку cUrl на Ubuntu. С приходом 12.04 все немного изменилось. Дабы избежать проблем ставим так:</p>

<pre><code class="bash">apt-get install apache2 php5 php5-json php5-gd php5-sqlite curl libcurl3 libcurl3-dev php5-curl php5-common php-xml-parser</code></pre>
