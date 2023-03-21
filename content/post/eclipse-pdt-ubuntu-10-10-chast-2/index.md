---
author: "Serhii Polishchuk"
title: "Eclipse PDT + Ubuntu 10.10. Часть 2"
date: 2011-01-23
tags: ["Ubuntu"]
draft: true
---
<!--more-->
<div class="alert alert-warning">Если Вы уже установили <strong>Eclipse</strong>, и подготовили к&nbsp;созданию нового проекта то пора приступить к следующей части, настройка <strong>Eclipse</strong> под наши нужды и создание тестового проекта.<br />
Обязательно <a href="/articles/eclipse-pdt-ubuntu-10-10-chast-1">прочтите часть 1</a></div>
<!--more-->

<p>Запускаем <strong>Eclipse PDT</strong>, идем в Window-&gt;Preferens там находим <strong>PHP-&gt;PHP Servers</strong> <em><strong> Инфо:</strong> Если у вас в настройках нет пункта PHP, значит не установлен пакет PDT. Подробнее об установке Eclipse PDT читайте <a href="/articles/eclipse-pdt-ubuntu-10-10-chast-1">в предыдущей статье</a>.</em></p>

<p><img alt="" class="img-responsive" src="/uploads/2011/01/eclipse-deb-pref3.jpg" style="width: 500px; height: 299px;" /></p>

<p>Выбираем место где у нас будут храниться файлы нашего проэкта, в моем случае это: <code>http://localhost/joomla_development</code> Выглядит это так:</p>

<p><img alt="" class="img-responsive" src="/uploads/2011/01/eclipse-deb-pref4.jpg" style="width: 451px; height: 333px;" /></p>

<p><span style="line-height: 1.2em;">Тест XDebug</span></p>

<p>Теперь у нас будет небольшое практическое задание. Мы напишем небольшой PHP скрипт и изучим на его примере процесс отладки.</p>

<h2>Терминология <strong>Eclipse</strong></h2>

<p>Сперва разберем некоторую терминологию. В пользовательском интерфейсе <strong>Eclipse</strong> есть несколько основных понятий: рабочий стол, вид и перспектива. <strong>Рабочий стол (workbench)</strong> - это впринципе все что помещается на экран. Рабочий стол <strong>Eclipse</strong> состоит из области редактирования, где мы будем писать и править наш <strong>PHP</strong> код. И ряд видов вокруг. Вид - это область где отображается информация о файле, или любой другой вывод. Перспектива - это просто распределенные настройки вида. Нам будут интересны две перспективы&nbsp;<strong>PHP perspective</strong> и <strong>PHP Debug perspective</strong>. Чтобы открыть нужную нам перспективу идем в Window-&gt;Open perspective-&gt;PHP - &nbsp;если <strong>PHP</strong> отсутствует в списке жмем Other, и выбираем в новом окне. Теперь на рабочем столе у нас отобразилось некоторый набор видов.</p>

<p><img alt="" class="img-responsive" src="/uploads/2011/01/eclipse-deb-pref5.jpg" style="width: 411px; height: 237px;" /></p>

<p><img alt="" class="img-responsive" src="/uploads/2011/01/eclipse-deb-pref6.jpg" style="width: 486px; height: 202px;" /></p>

<h2>Создаем проект</h2>

<p>Для каждого отдельного скрипта, мы создаем отдельный проект. Идем в File-&gt;NewPHP Project называем наш проект <strong>Test Debug</strong> и жмем Finish</p>

<p><img alt="" class="img-responsive" src="/uploads/2011/01/php-project-_0031.png" style="width: 500px; height: 514px;" /></p>

<p>Теперь мы можем наблюдать наш проект в <strong>PHP Explorer view</strong>. Для создания первого <strong>PHP</strong> файла нажимаем правой кнопкой на папке с нашим проектом и выбираем New-&gt;PHP file. К нам выйдет мастер создания нового PHP файла где мы укажем имя файла, напр. test.php</p>

<p><img alt="" class="img-responsive" src="/uploads/2011/01/new-php-file-_0042.png" style="width: 584px; height: 500px;" /></p>

<p>В открывшимся редакторе напишем следующий код:</p>

<pre>
<code class="php">&lt;?php echo &quot;Это мой тестовый компонент&quot;;
$myvar = &quot;Это моя переменная&quot;;
phpinfo();
?&gt;</code></pre>

<p><img alt="" class="img-responsive" src="/uploads/2011/01/php-project-_005.png" style="width: 396px; height: 132px;" /></p>

<p>Сохраняем, в <strong>PHP Explorer</strong> (далее проводник) выбираем наш файл, кликаем правой кнопкой мыши, выбираем Run As-&gt;PHP Web page Примечание. У меня на данном этапе вылезла ошибка из-за невозможности определить дефолтный браузер. Решается эта проблема следующим образом. Идем в Window-&gt;Preferens-&gt;General-&gt;Web Browser <img alt="" class="img-responsive" src="/uploads/2011/01/php-project-_006.png" style="width: 500px; height: 406px;" /> Я выбрал дефолтным веб браузером <strong>Chromium</strong> чего и вам желаю :) Для этого нажимаем New и вводим следующие параметры:<br />
name: На ваше усмотрение<br />
Location: /usr/bin/chromium-browser<br />
<img alt="" class="img-responsive" src="/uploads/2011/01/php-project-_007.png" style="width: 500px; height: 242px;" /> Путь к вашему браузеру, да и вобщем то сам браузер может отличаться от моего, для этого найдите (создайте) ярлык для своего приложения, кликните правой кнопкой-&gt;свойства. Копируем поле &quot;команда&quot; без параметров, т.е. до знака % <img alt="" class="img-responsive" src="/uploads/2011/01/php-project-_008.png" style="width: 500px; height: 197px;" /> Не забудьте предварительно запустить Вэб сервер, иначе вам выдаст сообщение о том что &quot;Невозможно отобразить страницу&quot;. После всего проделаного мой скрипт все таки запустился но вместо русского отображались крокозябры. Решил это так. Добавил в начало нашего скрипта код:</p>

<pre>

&nbsp;</pre>
<meta content="text/html; charset=utf-8" http-equiv="content-type" />
<p>Вот что получилось</p>

<p><img alt="" class="img-responsive" src="/uploads/2011/01/php-project-_009.png" style="width: 500px; height: 363px;" /></p>
