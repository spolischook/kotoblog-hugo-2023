---
author: "Serhii Polishchuk"
title: "Denwer и кодировка UTF-8"
date: 2011-03-18
tags: []
draft: false
---
<!--more-->
<p>Недавно в который раз наткнулся на проблему несовместимости Джентельменского набора разработчика <strong>Denwer 3 </strong>и кодировки <strong>UTF-8</strong>. Но в этот я раз всерьез решил заняться решением этой проблемы. И решение нашлось как всегда очень простое. В директорию где разрабатывается Ваш скрипт нужно разместить файл <strong>.htaccess</strong> со следующим содержанием: <code>AddDefaultCharset utf-8</code> Теперь скрипты сохраненные в кодировке <strong>UTF-8</strong> отображаются правильно. Файл можно <a href="http://ubuntuone.com/p/iK7/" target="_blank" title="Скачать файл .htaccess::Denwer и кодировка UTF-8">скачать здесь</a> <strong>UPD:</strong> Пользователь lance подошел глобально:</p>

<blockquote>httpd.conf отвечает за глобальные инструкции для сервера, т.е. внесенные в него изменения повлияют на все сайты, которые у нас будут храниться на сервере. Располагается он в каталоге /usr/local/apache/conf/. Отредактировать нам требуется всего одну строку: ищем &ldquo;AddDefaultCharset windows-1251&Prime; и меняем её на &ldquo;AddDefaultCharset utf-8&Prime;. мне это помогло!!! ))</blockquote>

<p>Спасибо ему за это :)</p>
