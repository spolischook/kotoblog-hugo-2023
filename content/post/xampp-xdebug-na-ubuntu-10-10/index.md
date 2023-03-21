---
author: "Serhii Polishchuk"
title: "XAMPP + XDebug на Ubuntu 10.10"
date: 2011-01-11
tags: ["Ubuntu"]
draft: true
---
<!--more-->
<p>Некоторое время пользовался <strong>LAMP</strong> и вроде как все устраивало, но после переустановки системы решил немного&nbsp;поэкспериментировать&nbsp;и установить <strong>XAMPP</strong> как более совершенный инструмент для разработки вэб приложений.</p>

<h2><!--more-->Что это?</h2>

<p><strong>XAMPP</strong> это крос-платформенный программный пакет для веб программирования.</p>

<p><strong>XAMPP</strong> - это аббревиатура:</p>

<p>X - кросс-платформенность (Windows, <strong>Linux</strong>, MacOs, Solaris)</p>

<p>A - <strong>Apache</strong> (самый распространенный вэб сервер)</p>

<p>M - <strong>MySQL</strong> (реляционная база данных)</p>

<p>P - <strong>PHP</strong> (язык программирования)</p>

<p>P - <strong>Perl </strong>(язык программирования)</p>

<h2>Установка</h2>

<p>На официальном сайте скачать последнюю версию <strong>XAMPP</strong> для <strong>Linux</strong>. <a href="http://www.apachefriends.org/en/xampp-linux.html" target="_blank">Качать здесь.</a> На момент написания статьи - <a href="http://www.apachefriends.org/download.php?xampp-linux-1.7.3a.tar.gz" target="_blank">XAMPP Linux 1.7.3a</a></p>

<p><em><strong>О XAMPP и Skype</strong>: Apache и Skype используют 80-ый порт для входящих соединений. Если Вы желаете использовать их вместе идем в Настройки(Ctrl+O)-&gt;Дополнительные-&gt;Соединение и снимите галочку с &quot;Использовать 80 и 433 порт&quot;. Раньше у меня эта опция была, но сейчас Skype по умолчанию не использует эти порты, может только на Linux? Если у Вас их нет, так тому и быть :)</em></p>

<p>На время установки уполномочим себя правами суперпользователя root:</p>

<pre>
<code class="bash">sudo su</code></pre>

<p><span style="line-height: 1.6em;">Перейдем в каталог куда мы сохранили архив (менеджере файлов нжав Ctrl+L можно получить адресс):</span></p>

<pre>
<code class="bash">cd &lt;Путь_к_архиву&gt; </code></pre>

<p><span style="line-height: 1.6em;">Распакуем архив в папку /opt (незабудьте что Имя_архива у Вас может отличаться):</span></p>

<pre>
<code class="bash">tar xvfz xampp-linux-1.7.3a.tar.gz -C /opt</code></pre>

<p><span style="line-height: 1.6em;">Фактически это все. </span><strong style="line-height: 1.6em;">XAMPP</strong><span style="line-height: 1.6em;"> установлен в директорию /opt/lampp Для запуска:</span></p>

<pre>
<code class="bash">/opt/lampp/lampp start</code></pre>

<p><span style="line-height: 1.6em;">Для проверки перейдите по адресу http://localhost </span></p>

<p><strong style="line-height: 1.6em;">PHPMyAdmin</strong><span style="line-height: 1.6em;"> доступен по адресу http://localhost/phpmyadmin/ </span></p>

<p><span style="line-height: 1.6em;">Далее не обязательный но очень желательный </span><strong style="line-height: 1.6em;">XDebug</strong></p>

<pre>
<code class="bash">sudo apt-get install php-pear php5-dev</code></pre>

<p><span style="line-height: 1.6em;">Соглашаемся со всем и далее:</span></p>

<pre>
<code class="bash">sudo pecl install xdebug</code></pre>

<p><span style="line-height: 1.6em;">В конце должны получить:</span></p>

<pre>
<code>Build process completed successfully
Installing &#39;/usr/lib/php5/20090626+lfs/xdebug.so&#39;
install ok: channel://pecl.php.net/xdebug-2.1.0
configuration option &quot;php_ini&quot; is not set to php.ini location
You should add &quot;extension=xdebug.so&quot; to php.ini</code></pre>

<p><span style="line-height: 1.6em;">Здесь нам сообщают что процесс установки завершился успешно, но нужно еще сообщить об этом серверу через </span><strong style="line-height: 1.6em;">php.ini</strong><span style="line-height: 1.6em;">. </span></p>

<p><span style="line-height: 1.6em;">Но сначала создадим временную папку для </span><strong style="line-height: 1.6em;">XDebuger</strong><span style="line-height: 1.6em;">:</span></p>

<pre>
<code>sudo mkdir /opt/lampp/tmp/xdebug
sudo chmod a+rwx -R /opt/lampp/tmp/xdebug</code></pre>

<p><span style="line-height: 1.6em;">Открываем </span><strong style="line-height: 1.6em;">php.ini</strong><span style="line-height: 1.6em;"> командой:</span></p>

<pre>
<code>sudo gedit /opt/lampp/etc/php.ini</code></pre>

<p><span style="line-height: 1.6em;">А теперь нужно прописать путь к </span><strong style="line-height: 1.6em;">XDebug</strong><span style="line-height: 1.6em;">, у меня это /usr/lib/php5/20090626+lfs но у Вас последняя директория может отличаться от моей, обязательно проверьте. В конец файла добавим следующий код:</span></p>

<pre>
<code class="ini">;xDebug Configuration starts 
  zend_extension=&quot;/usr/lib/php5/20090626+lfs/xdebug.so&quot;
  xdebug.remote_enable=1
  xdebug.profiler_output_dir = &quot;/opt/lampp/tmp/xdebug&quot; ;здесь директория для сохранения результатов профилировщика
  xdebug.profiler_enable = On 
  xdebug.remote_enable=On
  xdebug.remote_host=&quot;localhost&quot;
  xdebug.remote_port=10000
  xdebug.remote_handler=&quot;dbgp&quot;
;xDebug Configuration ends</code></pre>

<p><span style="line-height: 1.6em;">Ок, все установили, все прописали, не забыли сохранить и закрыть </span><strong style="line-height: 1.6em;">php.ini</strong><span style="line-height: 1.6em;">, теперь как нам все проверить? Очень просто. Перезапускаем сервер:</span></p>

<pre>
<code>/opt/lampp/lampp restart</code></pre>

<p><span style="line-height: 1.6em;">Переходим по ссылке http://localhost/xampp/index.php Заходим во вкладку </span><strong style="line-height: 1.6em;">phpinfo()</strong><span style="line-height: 1.6em;"> и листаем вниз где находим информацию по </span><strong style="line-height: 1.6em;">XDebug</strong><span style="line-height: 1.6em;"> примерно вот такую:</span></p>

<p><img alt="" class="img-responsive" src="/uploads/2011/01/xdebug.png" style="width: 884px; height: 506px;" /></p>

<p><span style="line-height: 1.6em;">Обратите внимание на настройки, они должны быть такими же как на рисунке:</span></p>

<p><img alt="" class="img-responsive" src="/uploads/2011/01/xdebug_settings.png" style="width: 524px; height: 134px;" /></p>

<p><span style="line-height: 1.6em;">Дополнение1: Если после перезагрузки компьютера по адресу </span><strong style="line-height: 1.6em;">http://localhost/</strong><span style="line-height: 1.6em;"> выводиться следующее сообщение</span></p>

<p><img alt="" class="img-responsive" src="/uploads/2011/01/itworks.png" style="width: 550px; height: 317px;" /></p>

<p><span style="line-height: 1.6em;">А при запуске из командной строки </span><strong style="line-height: 1.6em;">XAMPP</strong><span style="line-height: 1.6em;"> выдается ошибка, следует убрать из автозагрузки сервер </span><strong style="line-height: 1.6em;">Apache</strong><span style="line-height: 1.6em;">. Делается это следующим образом: Система-&gt;Параметры-&gt;Запускаемые приложения находим Apache и удаляем его. </span></p>

<p><span style="line-height: 1.6em;">Дополнение 2: Для упрощения работы - чтобы каждый раз не запускать сервер с консоли, создадим кнопки на рабочем столе (как в Denwer). 1. В домашнем каталоге (Переход-&gt;Домашний каталог) создадим пустой файл, назовем его например </span><strong style="line-height: 1.6em;">XAMPP</strong><span style="line-height: 1.6em;">, в который добавим следующий код:</span></p>

<pre>
<code class="bash">#!/bin/bash
foo=`gksudo -u root -k -m &quot;введите пароль root&quot; /bin/echo &quot;А вы рут?&quot;`
if [ &quot;$1&quot; = &quot;stop&quot; ]; then
sudo /opt/lampp/./lampp stop
fi
if [ &quot;$1&quot; = &quot;restart&quot; ]; then
sudo /opt/lampp/./lampp restart
fi
if [ &quot;$1&quot; = &quot;start&quot; ]; then
sudo /opt/lampp/./lampp start
fi</code></pre>

<p><span style="line-height: 1.6em;">Теперь перейдем на рабочий стол и создадим кнопку (клик, правой кнопкой -&gt; создать кнопку запуска), дадим ей имя Start, в поле команда впишем /home/&lt;Имя_пользователя&gt;/</span><strong style="line-height: 1.6em;">XAMPP</strong><span style="line-height: 1.6em;"> start жмем Ок. Все готово :) </span></p>

<p><span style="line-height: 1.6em;">То же самое проделываем для Restart и Stop </span></p>

<p><span style="line-height: 1.6em;">Вот так это выглядит у меня:</span></p>

<p><span style="line-height: 1.6em;"><img alt="" class="img-responsive" src="/uploads/2011/01/button.png" style="width: 466px; height: 272px;" /></span></p>
