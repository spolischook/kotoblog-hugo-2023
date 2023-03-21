---
author: "Serhii Polishchuk"
title: "Видео YouTube во всплывающем окне JCE MediaBox+Joomla! 1.5"
date: 2011-01-15
tags: ["Joomla"]
draft: true
---
<!--more-->
<p>Собственно сама <strong>задача:</strong> Улучшить юзабилити сайта, устранив лишние клики мышью при просмотре видео на сайте. <strong>Дано:</strong> сайт на <strong>Joomla! 1.5</strong> с установленным <strong>JCE Editor</strong> <strong>Решение:</strong> Если еще не установлен<strong> JCE Editor</strong> <a href="http://www.joomlacontenteditor.net/downloads/editor">качаем</a> с <a href="http://www.joomlacontenteditor.net/">официального сайта</a> и устанавливаем. Точно также поступаем с <strong><a href="http://www.joomlacontenteditor.net/downloads/mediabox/category/mediabox-2">MediaBox</a>.</strong><!--more-->Далее делаем картинку, я делаю снимок экрана видеоролика на YouTube. Например такую:</p>

<p><img alt="" class="img-responsive" src="/uploads/2011/01/cherry-2010.png" style="width: 300px; height: 183px;" /></p>

<p>Вставляем ее в статью, выделяем картинку и делаем ее ссылкой вот так:</p>

<p><img alt="" class="img-responsive" src="/uploads/2011/01/2011-01-15-1.jpg" style="width: 400px; height: 312px;" /></p>

<p>Далее указываем следующие параметры. <strong> Вкладка Link</strong>. <em>URL</em> - это ссылка на видео <strong>YouTube</strong>. Это просто заходим на видео и копируем адрес из адресной строки браузера. <em>Заголовок:</em> Указываем название и описание через двойное двоеточие, так: <code>&lt;Название_видео&gt;::&lt;Описание_видео&gt;</code> <strong> Вкладка Advanced.</strong> <em>Class List </em>- jcepopup.&nbsp;<em>Relationship page to target </em>- указываем параметры: <code> width[1024], height[768] </code> Ширину и высоту поменять на свои. Сохраняем статью, получаем видео во всплывающем окне. Enjoy and keep smiling</p>
