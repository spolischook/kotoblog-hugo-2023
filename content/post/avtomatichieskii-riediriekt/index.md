---
author: "Serhii Polishchuk"
title: "Автоматический редирект"
date: 2011-01-06
tags: ["PHP"]
draft: true
---
<!--more-->
<p>На днях понадобилось создать <strong>автоматический редирект.</strong> Суть в том что малобюджетные, или некомерческие проэкты по ошибке регистрируются (награждаются :)) платными доменными именами. И когда приходит время платить за имя... Ближе к теме. <strong>Автоматический редирект</strong> - это перенаправление на заданную страницу с заданным временным интервалом. <!--more--> Сразу скажу что пользоваться нужно им осторожно, т.к. поисковые машины относятся к <strong>редиректу</strong> не очень дружелюбно. Некоторые оптимизаторы используют этот метод для привлечения новых посетителей по специально отточенным ключевым запросам используемых на странице с <strong>редиректом</strong>. Для этого нам понадобиться:</p>

<ul>
	<li>Сервер с возможностью выполнения PHP скриптов</li>
	<li>Текстовый редактор (лучше всего подойдет <a href="http://notepad-plus-plus.org/" target="_blank" title="Официальный сайт Notepad++">Notepad++</a>)</li>
	<li>Откроем нашу страницу в текстовом редакторе, и вставим в самом верху:</li>
</ul>

<p><img alt="" class="img-responsive" src="/uploads/2011/01/redir.jpg" style="width: 619px; height: 29px;" /></p>

<ul>
	<li>А внизу: <img alt="" class="img-responsive" src="/uploads/2011/01/redir2.jpg" style="width: 164px; height: 25px;" />&nbsp;Здесь два важных параметра: Если это страница какой нибудь CMS (например WordPress), а не обычная HTML, то&nbsp;придется&nbsp;немного поморочится с условиями if then else, и здесь уже без знаний РНР Вам не обойтись. Но нет ничего не возможного - спрашивайте в комментариях. Это все. Теперь можете обращаться к файлу на сервере и он автоматически Вас перенаправит на выбранный вами сайт.
	<ul>
		<li>Refresh 4 - вместо &quot;4&quot;&nbsp;подставьте&nbsp;свое значение. Это время в секундах, через которое пользователь будет перенаправлен на...</li>
		<li>http://example.com - это сам сайт на который будет перенаправлен пользователь.</li>
	</ul>
	</li>
</ul>
