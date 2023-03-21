---
author: "Serhii Polishchuk"
title: "Eclipse PDT + Ubuntu 10.10. Часть 3. Отладка с помощью XDebug, Установка Subversopn (SVN)"
date: 2011-02-16
tags: ["Ubuntu"]
draft: true
---
<!--more-->
<p>В частях 2 и 1 мы ознакомились с тем как устанавливать<strong> Eclipse XDebug на Ubuntu 10.10</strong> и создать свой первый простейший (так и хочется назвать одноклеточный :)) скрипт. Теперь нам предстоит научиться отлаживать программу с помощью <strong>XDebug</strong>. <!--more-->Жмем правой кнопкой на нашем скрипте и выбираем&nbsp;<strong>Debug As </strong>/ PHP Web Page В этот момент <strong>Eclipse</strong> предложит нам открыть <strong>Debug</strong> перспективу, с чем мы вежливо соглашаемся:</p>

<p><img alt="" class="img-responsive" src="/uploads/2011/01/php-project-_010.png" style="width: 500px; height: 173px;" /></p>

<p>Вообще мы можем выбирать ту перспективу которая нам нужна в данный момент из меню Window-&gt;Open Perspective. После входа в <strong>Debug</strong> песпективу мы увидим что <strong>Debuger</strong> остановился на первой <strong>PHP</strong> строке.</p>

<p><img alt="" class="img-responsive" src="/uploads/2011/01/php-project-_011.png" style="width: 500px; height: 154px;" /></p>

<p>Теперь посмотрим немного выше, где у нас распологаються кнопки управления <strong>XDebug</strong></p>

<p><img alt="" class="img-responsive" src="/uploads/2011/01/eclipse-debug-button.png" style="width: 525px; height: 360px;" /></p>

<p>1. Начать отладку скрипта. Перед тем как запускать отладку, необходимо удостовериться что предыдущая отладка завершена (3) 2. Запустить полностью весь скрипт. На данном этапе, с теми настройками которые у нас есть, мы отлаживаем программу по строчно, но можно &quot;снять с паузы&quot; и продолжить выполнение до конца. 3. Остановить выполнение отладки скрипта. Обязательное условие для повторной отладки. 4. Следующий шаг. Т.к. в настройках ранее мы выбрали что наш скрипт будет отлаживаться по-шагово, то при нажатии на эту кнопку отладчик переходит к следующему событию/оператору. В правом верхнем углу у нас находятся кнопки выбора перспективы:</p>

<p><img alt="" class="img-responsive" src="/uploads/2011/01/eclipse-perspective.png" style="width: 600px; height: 182px;" /></p>

<p>между которыми мы можем свободно переключатся. В правом верхнем углу, чуть по ниже кнопок выбора перспективы находяться вид переменных, в которых мы можем наблюдать какие переменные у нас задействованны, и какие значения они принимают на каждом шаге нашой программы.</p>

<p><img alt="" class="img-responsive" src="/uploads/2011/01/eclipse-variables.png" style="width: 557px; height: 337px;" /></p>

<p>При старте отладки, как Вы уже успели заметить, открываеться выбранный ранее в настройках браузер, в котором также по шагово отображается информация которая приходит от нашей программы находящейся на отладке. Таким образом мы можем контролировать процесс. Если мы дойдем до конца программы, то это вовсе не означает что процесс отладки окончен. Нам нужно нажать на кнопку 3 что бы его завершить, иначе новый сеанс не запуститься. По завершению откроется новое окно браузера в котором нам сообщат об окончании отладки скрипта. А для того чтобы в дальнейшем <strong>Eclipse PDT</strong> не &quot;напрягал&quot; нас вопросами о том &quot;хотим ли мы открыть соответствующую перспективу&quot; при отладке, можно в настройках (Window-&gt;Preference) это явно указать.</p>

<p><img alt="" class="img-responsive" src="/uploads/2011/01/eclipse-preference-perspective-launch.png" style="width: 536px; height: 424px;" /></p>

<p>На этом все. Мы научились отлаживать наши программы в <strong>Eclipse PDT</strong>. В следующей статье я раскрою очень важную тему<strong> Eclipse Subversion</strong>. Keep smiling! :)</p>

<h3>Установка&nbsp;Subversion (SVN)</h3>

<p><strong>Subversion (SVN) </strong>- это сторонний плагин<strong> Eclipse </strong>который необходим для работы с Системой управления версиями&nbsp;исходного кода разрабатываемых (и не только нами, а иначе зачем это все нада :) ) нами программ.</p>

<p><img alt="" class="img-responsive" src="/uploads/2011/01/eclipse-new-plagin.jpg" style="width: 297px; height: 312px;" /></p>

<p>Идем в Help-&gt;Install New Software. В диалоговом окне выбираем Add</p>

<p><img alt="" class="img-responsive" src="/uploads/2011/01/eclipse-add-install.png" style="width: 550px; height: 496px;" /></p>

<p><span style="line-height: 1.6em;">В качестве названия (name) даем любое удобочитаемое имя, напр&nbsp;</span><strong style="line-height: 1.6em;">Subclipse</strong><span style="line-height: 1.6em;">, в поле размещение (Location) пишем &quot;</span><strong style="line-height: 1.6em;">http://subclipse.tigris.org/update_1.6.x</strong><span style="line-height: 1.6em;">&quot;</span></p>

<p><img alt="" class="img-responsive" src="/uploads/2011/01/eclipse-sorce.png" style="width: 640px; height: 221px;" /></p>

<p><span style="line-height: 1.6em;">Выбираем все три доступных пакета</span></p>

<p><img alt="" class="img-responsive" src="/uploads/2011/01/eclipse-add-install2.png" style="width: 550px; height: 430px;" /></p>

<p><span style="line-height: 1.6em;">Жмем Next до упора, пока нам не покажут лицензию, которую мы принимаем и жмем Finish. Возможно нам выдаст некоторые сообщения о возможной несовместимости - соглашаемся со всем. Далее нас попросят перезагрузить </span><strong style="line-height: 1.6em;">Eclipse</strong><span style="line-height: 1.6em;"> с чем мы вежливо соглашаемся. После перезапуска мы можем проверить установку нашего плагина. Заходим в File-&gt;Import...</span></p>

<p><img alt="" class="img-responsive" src="/uploads/2011/01/eclipse-subv.jpg" style="width: 550px; height: 575px;" /></p>

<p><span style="line-height: 1.6em;">В новом окне находим </span><strong style="line-height: 1.6em;">SVN</strong></p>

<p><img alt="" class="img-responsive" src="/uploads/2011/01/eclipse-subv21.png" style="width: 550px; height: 437px;" /></p>

<p><span style="line-height: 1.6em;">Если такой пункт присутствует, то будем считать что мы успешно установили </span><strong style="line-height: 1.6em;">Eclipse Subversion</strong><span style="line-height: 1.6em;">. В следующей статье мы поговорим о том как импортировать </span><strong style="line-height: 1.6em;">Joomla! project</strong><span style="line-height: 1.6em;"> с помощью </span><strong style="line-height: 1.6em;">Eclipse Subversion (SVN)</strong></p>
