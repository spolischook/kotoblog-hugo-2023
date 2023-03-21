---
author: "Serhii Polishchuk"
title: "Права на папки Joomla! автоматически!"
date: 2011-03-14
tags: ["Joomla"]
draft: true
---
<!--more-->
<p>Все очень просто. После установки <strong>Joomla!</strong> на локальный сервер под <strong>Linux</strong>, непременно появляется необходимость изменить права доступа на папки для дальнейшей работы, напр. для установки приложений. Делать это вручную задача долгая, нудная и не благодарная. Нашел простенький скрипт для <strong>Joomla</strong>! 1.0 переделал его под <strong>Joomla</strong>! 1.5.<!--more-->Создаем новый файл через <strong>gedit</strong> или любой другой полюбившийся Вам текстовый редактор. Назовем его <strong>chmod.sh</strong>. В него запишем следующий код:</p>

<pre>
<code class="bash">#!/bin/sh
sudo chmod 777 administrator/backups/
sudo chmod 777 administrator/components/
sudo chmod 777 administrator/language/
sudo chmod 777 administrator/language/en-GB/
sudo chmod 777 administrator/language/ru-RU/
sudo chmod 777 administrator/cache/
sudo chmod 777 language/en-GB/
sudo chmod 777 language/pdf_fonts/
sudo chmod 777 language/ru-RU/
sudo chmod 777 plugins/
sudo chmod 777 plugins/content/
sudo chmod 777 plugins/editors/
sudo chmod 777 plugins/editors-xtd/
sudo chmod 777 plugins/search/
sudo chmod 777 plugins/system/
sudo chmod 777 plugins/user/
sudo chmod 777 plugins/xmlrpc/
sudo chmod 777 administrator/modules/
sudo chmod 777 administrator/templates/
sudo chmod 777 cache/
sudo chmod 777 components/
sudo chmod 777 images/
sudo chmod 777 images/banners/
sudo chmod 777 images/stories/
sudo chmod 777 language/
sudo chmod 777 media/
sudo chmod 777 modules/
sudo chmod 777 templates/
sudo chmod 777 logs/
sudo chmod 777 tmp/
sudo touch configuration.php
sudo chmod 777 configuration.php
sudo mv htaccess.txt .htaccess</code></pre>

<p><span style="line-height: 1.6em;">После этого открываем терминал и переходим в корневую папку где установлен новый </span><strong style="line-height: 1.6em;">Joomla</strong><span style="line-height: 1.6em;"> сайт. Если мы используем </span><strong style="line-height: 1.6em;">XAMPP</strong><span style="line-height: 1.6em;"> наш путь будет выглядеть следующими образом:</span></p>

<pre>
<code class="bash">cd /opt/lampp/htdocs/mysite</code></pre>

<p><span style="line-height: 1.6em;">&#39;mysite&#39; замените на папку с Вашим проэктом. После этого запускаем скрипт указывая полный путь к нему. В моем случае, если скрипт находится в папке с проэктом:</span></p>

<pre>
<code class="bash">sudo bash chmod.sh</code></pre>

<p><span style="line-height: 1.6em;">Все готово! :)</span></p>
