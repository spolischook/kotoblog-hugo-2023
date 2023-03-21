---
author: "Serhii Polishchuk"
title: "Шаблон вывода списка товаров категории Prestashop"
date: 2011-05-27
tags: ["PrestaShop"]
draft: true
---
<!--more-->
<p>Дело в том что при наличии достаточно обширного асортимента товаров стала необходимо разделить товары по различным группам.</p>

<ol>
	<li>Группа товаров которая отображалась в стандартном виде - Название, картинка, короткое описание, цена, купить</li>
	<li>Группа отображения в виде таблицы - Название, цена, купить.</li>
</ol>

<p>Вторая группа будет использоватся для товаров для которых не важны ни описание ни внешний вид, таких как - расходные материалы, комплектующие и т. п. <!--more-->Для начала ставим принудительную компиляцию шаблона <strong>Smarty</strong> для того чтобы сразу же видеть изменения которые мы будем вносить в код. Для этого есть 2 способа. Первый. Идем в Админ панель сайта -&gt; Настройки -&gt; Производительность -&gt; <strong>Принудительная компиляция</strong> -&gt; Да Второй. Оставляем Принудительную компиляцию выключенной, идем в /config/smarty.config.inc.php ищем smarty-&gt;compile_check и меняем значение на true. Все теперь <strong>Smarty</strong> будет перекомпилировать только измененные шаблоны. Следующим шагом будет добавление свойства категории с названием <strong>flag_template</strong>. Идем в /classes/<strong>Category.php</strong> и в классе <strong>CategoryCore</strong> объявляем новую переменную, после:</p>

<pre>
<code class="php">public $date_upd;</code></pre>

<p>ставляем:</p>

<pre>
<code class="php">/** @var integer category view template */
public	 $flag_template;</code></pre>

<p>Далее инициализируем переменную, находим:</p>

<pre>
<code class="php">$fields[&#39;date_upd&#39;] = pSQL($this-&gt;date_upd);</code></pre>

<p>и вставляем после нее:</p>

<pre>
<code class="php">$fields[&#39;flag_template&#39;] = pSQL($this-&gt;flag_template);</code></pre>

<p>Теперь у нас переменной<strong> flag_template</strong> будет присвоено значение из базы данных, таблицы _category поле&nbsp;<strong><strong>flag_template</strong></strong>. Нужно создать это поле в базе данных вручную. Я для этих целей использовал <strong>PhpMyAdmin</strong> В файле /themes/prestashop/<strong>category.tpl </strong>находим строку включения списка товаров:</p>

<pre>
<code class="php">{include file=&quot;$tpl_dir./product-list.tpl&quot; products=$products}</code></pre>

<p>И заменяем на:</p>

<pre>
<code>{if $category-&gt;flag_template == 1}
{include file=&quot;$tpl_dir./product-list1.tpl&quot; products=$products}
{else}
{include file=&quot;$tpl_dir./product-list.tpl&quot; products=$products}
{/if}</code></pre>

<p>Вобщем то на этом можно закончить. Можно вставлять столько условий, сколько будет шаблонов. Каждый из этих шаблонов необходимо будет перевести через админ панель (Инструменты (Tools) -&gt; Переводы(Translation) -&gt; Изменить перевод -&gt; Фронт офис -&gt; Флаг языка (русский? :))- просто скопировав значения из уже переведенного шаблона! После того как я уменьшил почти в 3-и раза высоту блока отображения (class=&quot;ajax_block_product) товара образовалась пустота внизу страницы. Я подумал что неплохо было бы эти товары выводить в большем количестве. Итак, я нашел класс который отвечает за к-во товаров на странице, а также за нумерацию страниц (pagination) - /classes/FrontController.php. Находим в нем <strong>function pagination</strong> и после</p>

<pre>
<code class="php">if (!self::$initialized)
$this-&gt;init();</code></pre>

<p><span style="line-height: 1.6em;">Вставляем:</span></p>

<pre>
<code class="php">$query = &#39;SELECT `flag_template` FROM `&#39;._DB_PREFIX_.&#39;category` WHERE `id_category` = &#39;.$_GET[&#39;id_category&#39;];
$template = Db::getInstance(_PS_USE_SQL_SLAVE_)-&gt;getValue($query);
if ($template == 1) {$product_per_page = 30;}
else {$product_per_page = Configuration::get(&#39;PS_PRODUCTS_PER_PAGE&#39;);}</code></pre>

<p><span style="line-height: 1.6em;">Далее по тексту ф-ции отслеживаем выражение </span><em style="line-height: 1.6em;">Configuration::get(&#39;PS_PRODUCTS_PER_PAGE&#39;)</em><span style="line-height: 1.6em;"> и наменяем его на </span><em style="line-height: 1.6em;">$product_per_page</em><span style="line-height: 1.6em;"> Огромное спасибо за хорошую подсказку </span><a href="http://www.prestashop.com/forums/member/80483/sors/" style="line-height: 1.6em;"><strong>Sors-у</strong></a><span style="line-height: 1.6em;"> из русского сообщества </span><strong style="line-height: 1.6em;">Prestashop</strong><span style="line-height: 1.6em;">, без него я бы еще б неделю рыл код. З.Ы.: по просьбе Виталия выкладываю скрины как это выглядит на моем сайте. </span><strong style="line-height: 1.6em;">Это стандартный шаблон:</strong></p>

<p><img alt="" class="img-responsive" src="/uploads/2011/08/product_list1.jpg" style="width: 565px; height: 549px;" /></p>

<p><strong>Модифицированный шаблон:</strong></p>

<p><strong><img alt="" class="img-responsive" src="/uploads/2011/08/product_list2.jpg" style="width: 566px; height: 481px;" /></strong></p>
