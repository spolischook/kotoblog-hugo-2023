---
author: "Serhii Polishchuk"
title: "Аудиоплеер на сайт на javascript"
date: 2012-05-05
tags: []
draft: true
---
<!--more-->
<p>Понадобилось мне как то добавить плеер на свой персональный блог (да-да, есть еще и такой). И я удивился простоте реализации которую я нашел, никаких плагинов, абсолютная независимость от платформы, и пока HTML5 не вступил в оффициальную силу предлагаю следущую реализацию...<!--more--> Вчасть нашего сайта вставляем вызов самого плеера от Yahoo!</p>

<pre>
<code class="javascript">&lt;script type=&quot;text/javascript&quot; src=&quot;http://mediaplayer.yahoo.com/js&quot;&gt;&lt;/script&gt;</code></pre>

<p>Теперь в статьях можем использовать ссылки с произвольными якорями, ссылаясь на аудио файл:</p>

<pre>
<code class="http">&lt;a href=&quot;http://путь/к/МР3/файлу/music.mp3&quot;&gt;Хорошая Песня&lt;/a&gt;</code></pre>

<p><img alt="" class="img-responsive" src="/uploads/2012/05/2012-05-05-audio-player.jpg" style="width: 720px; height: 947px;" /></p>
