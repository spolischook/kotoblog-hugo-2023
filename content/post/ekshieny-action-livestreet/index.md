---
author: "Serhii Polishchuk"
title: "Экшены (Action) LiveStreet"
date: 2012-09-24
tags: []
draft: true
---
<!--more-->
<p>Экшены (Action) предназначены для обработки определенных URL. Например при запросе&nbsp;<strong>http://site.com/action_name/event_name/</strong>&nbsp;будет запущен экшен с именем&nbsp;<strong>action_name</strong>. Соответствие имени экшена и его класса определяется в&nbsp;<a href="http://web.archive.org/web/20100724120349/http://trac.lsdev.ru/livestreet/wiki/DeveloperDoc/Router#Конфигроутинга">конфиге роутинга</a>. Классы экшенов расположены в каталоге&nbsp;<strong>/classes/actions/</strong>&nbsp;и наследуются от класса&nbsp;<strong>Action</strong>(/engine/classes/Action.class.php). Название файла экшена должно формироваться из названия класса, например, для класса&nbsp;<strong>ActionTest</strong>&nbsp;файл будет иметь имя&nbsp;<strong>ActionTest.class.php</strong>. При запуске экшен пытается найти необходимый эвент по&nbsp;<strong>event_name</strong>&nbsp;и если находит, то запускает его. Эвент представляет из себя обычный метод класса, которому определено соответствие с именем эвента, в нашем примере это&nbsp;<strong>event_name</strong>. <!--more--> Пример простого экшена:</p>

<pre>
<code class="php">&lt;?php

class ActionTest extends Action {
        /**
         * Инициализация экшена
         *
         */
        public function Init() {                
                $this-&gt;SetDefaultEvent(&#39;index&#39;);                
        }
        /**
         * Регистрация эвентов
         *
         */
        protected function RegisterEvent() {
                $this-&gt;AddEvent(&#39;index&#39;,&#39;EventIndex&#39;);
                $this-&gt;AddEvent(&#39;edit&#39;,&#39;EventEdit&#39;);            
        }
        /**
         * Метод эвуента index
         *
         */
        protected function EventIndex() {               
                $this-&gt;SetTemplateAction(&#39;index&#39;);
        }
        /**
         * Метод эвента edit
         *
         */
        protected function EventEdit() {
                $this-&gt;Viewer_Assign(&#39;iSomeNumber&#39;,12345);
                $this-&gt;SetTemplateAction(&#39;some_tamplate&#39;);
        }       
        /**
         * Вызывается автоматически после выполнения эвента
         *
         */
        public function EventShutdown() {
                $this-&gt;Viewer_Assign(&#39;sText&#39;,&#39;Event complete&#39;);
        }
}</code></pre>

<p>В данном примере мы создаем экшен&nbsp;<strong>ActionTest</strong>&nbsp;с двумя эвентами:&nbsp;<strong>index</strong>&nbsp;и&nbsp;<strong>edit</strong>. Обязательные методы, которые должны присутствовать в любом экшене, это&nbsp;<strong>Init()</strong>&nbsp;и&nbsp;<strong>RegisterEvent()</strong>, в первом происходит инициализация экшена, например, установка дефолтного эвента, а во втором регистрация всех используемых эвентов в данном экшене. Также предусмотрены два дополнительных метода, которые определять не обязательно, это<strong>EventShutdown()</strong>&nbsp;и&nbsp;<strong>EventNotFound()</strong>, первый автоматически вызывается после выполнения текущего эвента, второй вызывается в том случаи, если не найден необходимый эвент. По умолчанию&nbsp;<strong>EventNotFound()</strong>&nbsp;переадресует управление на экшен&nbsp;<strong>error</strong>, который показывает 404 HTTP ошибку:</p>

<pre>
<code class="php">&lt;?php

protected function EventNotFound() {
        $this-&amp;gt;Message_AddErrorSingle($this-&gt;Lang_Get(&#39;system_error_404&#39;),&#39;404&#39;);
        return Router::Action(&#39;error&#39;);
}</code></pre>

<h2 id="ЭвентыEvent">Эвенты (Event)</h2>

<p>Регистрация евентов происходит в методе&nbsp;<strong>RegisterEvent()</strong>, существует два способа зарегистрировать евент, это&nbsp;<strong>AddEvent()</strong>&nbsp;и&nbsp;<strong>AddEventPreg()</strong>, второй использует для этого регулярные выражения. Регистрация евентов:</p>

<pre>
<code class="php">&lt;?php

// Стандартный способ
$this-&gt;AddEvent(&#39;event_name&#39;,&#39;EventMethod&#39;);

// Через регулярные выражение
/**
* соответствует site.com/action/123.html и site.com/action/123.html/blabla/
*/
$this-&gt;AddEventPreg(&#39;/^(\d+)\.html$/i&#39;,&#39;EventMethod&#39;);

/**
* соответствует site.com/action/123.html, но не соответствует site.com/action/123.html/blabla/
*/
$this-&gt;AddEventPreg(&#39;/^(\d+)\.html$/i&#39;,&#39;/^$/i&#39;,&#39;EventMethod&#39;);

/**
* соответствует site.com/action/123/, site.com/action/123/page1/ ,
* но не соответствует site.com/action/123/blabla/
*/
$this-&gt;AddEventPreg(&#39;/^(\d+)$/i&#39;,&#39;/^(page\d+)?$/i&#39;,&#39;/^$/i&#39;,&#39;EventMethod&#39;);</code></pre>

<p>В примерах выше при запросе правильного URL будет вызван метод&nbsp;<strong>EventMethod()</strong>, в противном случаи метод&nbsp;<strong>EventNotFound()</strong>. В&nbsp;<strong>AddEvent(&#39;event_name&#39;,&#39;EventMethod&#39;)</strong>&nbsp;первый аргумент определяет имя эвента, а второй метод, который будет вызван. В&nbsp;<strong>AddEventPreg(&#39;/^(\d+)\.html$/i&#39;,&#39;/^$/i&#39;,&#39;EventMethod&#39;)</strong>первый параметр определят регулярное выражение для имени эвента, а второй и последующие (кроме последнего) регулярное выражение для параметров. Последний параметр определят название метода для вызова. После выполнения эвента в браузер выводится шаблон, по умолчанию совпадающий с именем эвента. Например, для эвента&nbsp;<strong>AddEvent(&#39;event_name&#39;,&#39;EventMethod&#39;)</strong>&nbsp;будет использоваться шаблон&nbsp;<strong>event_name.tpl</strong>, расположенный в&nbsp;<strong>/templates/skin/skin_name/actions/ActionTest/</strong>, где&nbsp;<strong>skin_name</strong>&nbsp;название скина(общего шаблона), а&nbsp;<strong>ActionTest</strong>&nbsp;название класса экшена в котором зарегистрирован эвент. Шаблон этот можно переопределить используя метод&nbsp;<strong>SetTemplateAction(&#39;some_tamplate&#39;)</strong>, в таком случаи будет использован шаблон&nbsp;<strong>/templates/skin/skin_name/actions/ActionTest/some_tamplate.tpl</strong>. Также есть возможность из одного эвента передать управление на другой (внутренний реврайт), или даже на другой экшен. Делается это с помощью роутера:</p>

<pre>
<code class="php">&lt;?php

protected function EventEdit() {
        // некоторый код
        //.....

        return Router::Action(&#39;test&#39;,&#39;edit&#39;);
}</code></pre>

<p>В данном примере передается управление в экшен&nbsp;<strong>test</strong>&nbsp;с эвентом&nbsp;<strong>edit</strong>, т.е. происходит эмуляция URL&nbsp;<strong>http://site.com/test/edit/</strong>. Для передачи в шаблон каких-либо данный используется метод&nbsp;<strong>Assign($sName,$value)</strong>&nbsp;системного модуля&nbsp;<strong>Viewer</strong>, где&nbsp;<strong>$sName</strong>&nbsp;будущее имя переменной в шаблоне, а&nbsp;<strong>$value</strong>&nbsp;её значение:</p>

<pre>
<code class="php">&lt;?php

$this-&gt;Viewer_Assign(&#39;iSomeNumber&#39;,12345);</code></pre>

<h2 id="ДоступныеметодыAction">Доступные методы Action</h2>

<ul>
	<li><strong>AddEvent($sEventName,$sEventFunction)</strong>&nbsp;- регистрация эвента по имени</li>
	<li><strong>AddEventPreg($sEventPregName,..,$sParamPregN,..,$sEventFunction)</strong>&nbsp;- регистрация евента через регулярные выражения</li>
	<li><strong>SetDefaultEvent($sEvent)</strong>&nbsp;- устанавливает эвент по умолчанию, т.е. тот который будет вызван при URL&nbsp;<strong>http://site.com/action/</strong>. Данные метод должен вызываться из метода<strong>Init()</strong></li>
	<li><strong>GetDefaultEvent()</strong>&nbsp;- возвращает эвент по умолчанию, например,&nbsp;<strong>index</strong></li>
	<li><strong>GetEventMatch($iItem=null)</strong>&nbsp;- возвращает элемент совпадения евента по регулярному выражению. Например, если для евента&nbsp;<strong>AddEventPreg(&#39;/^(page(\d+))$/i&#39;,&#39;EventMethod&#39;)</strong>вызвать метод&nbsp;<strong>GetEventMatch(2)</strong>, то он вернет числовое значение равное номеру страницы</li>
	<li><strong>GetParamEventMatch($iParamNum,$iItem=null)</strong>&nbsp;- возвращает элемент совпадения по регулярному выражению для параметров евента. Например, если для евента<strong>AddEventPreg(&#39;/^blog$/i&#39;,&#39;/^new$/&#39;,&#39;/^page(\d+)$/i&#39;,&#39;EventMethod&#39;)</strong>&nbsp;вызвать метод&nbsp;<strong>GetParamEventMatch(1,1)</strong>, то он вернет числовое значение равное номеру страницы</li>
	<li><strong>GetParam($iOffset,$default=null)</strong>&nbsp;- возвращает параметр по его номеру/смещению, $default определяет дефолтное значение в случаи если параметра не существует. Например, для&nbsp;<strong>http://site.com/user/ort/edit/fast/</strong>&nbsp;GetParam(0) вернет значение&nbsp;<strong>edit</strong></li>
	<li><strong>GetParams()</strong>&nbsp;- возвращает массив текущих параметров, которые были переданные через URL. Например, для&nbsp;<strong>http://site.com/user/ort/edit/fast/</strong>&nbsp;вернет значение&nbsp;<strong>array(&#39;edit&#39;,&#39;fast&#39;)</strong></li>
	<li><strong>SetParam($iOffset,$value)</strong>&nbsp;- устанавливает значение параметра, например,&nbsp;<strong>SetParam(0,&#39;add&#39;)</strong></li>
	<li><strong>SetTemplate($sTemplate)</strong>&nbsp;- устанавливает шаблон для вывода, используя путь относительно общего каталога шаблонов. Например,&nbsp;<strong>SetTemplate(&#39;index.tpl&#39;)</strong>&nbsp;указывает на шаблон&nbsp;<strong>/templates/skin/skin_name/index.tpl</strong></li>
	<li><strong>SetTemplateAction($sTemplate)</strong>&nbsp;- устанавливает шаблон для вывода, используя путь относительно каталога шаблонов экшена. Например,&nbsp;<strong>SetTemplateAction(&#39;index&#39;)</strong>указывает на шаблон&nbsp;<strong>/templates/skin/skin_name/actions/ActionTest/index.tpl</strong>. Внимание! При указании шаблона использовать расширение&nbsp;<strong>.tpl</strong>&nbsp;не нужно!</li>
	<li><strong>GetTemplate()</strong>&nbsp;- возвращает пусть до шаблона относительно общего каталога шаблонов, например,&nbsp;<strong>actions/ActionTest/index.tpl</strong></li>
	<li><strong>GetActionClass()</strong>&nbsp;- возвращает имя класса экшена, например,&nbsp;<strong>ActionTest</strong></li>
</ul>
