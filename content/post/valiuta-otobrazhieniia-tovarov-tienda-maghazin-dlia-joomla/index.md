---
author: "Serhii Polishchuk"
title: "Валюта отображения товаров Tienda - магазин для Joomla!"
date: 2011-04-03
tags: ["Joomla"]
draft: true
---
<!--more-->
<p>Я так подозреваю что возможность отображать цены во фронтенде в <strong>валюте </strong>отличной от той что в базе будет в последующих версиях, но все равно опыт интересен, и я по просьбе lmd и по своему хотению решил немного покопаться в коде. Вот что я нарыл... <!--more--> Посмотрев шаблоны вывода цены, обнаружил что за это отвечает метод <strong>dispayPriceWithTax</strong> класса <strong>TiendaHelperProduct</strong> его можно найти по адресу <code>administrator/components/com_tienda/helpers/product.php</code> Выглядит он примерно так:</p>

<pre>
<code class="php">function dispayPriceWithTax($price=&#39;0&#39;, $tax=&#39;0&#39;, $show=&#39;0&#39;)
{
  $txt = &#39;&#39;;
  if($show &amp;&amp; $tax)
  {
    if ($show == &#39;2&#39;)
    {
      $txt .= TiendaHelperBase::currency($price + $tax);
    }
    else
    {
      $txt .= TiendaHelperBase::currency($price);
      $txt .= sprintf( JText::_(&#39;INCLUDE_TAX&#39;), TiendaHelperBase::currency($tax));
    }
  }
  else
  {
    $txt .= TiendaHelperBase::currency($price);
  }

  return $txt;
}</code></pre>

<p><span style="line-height: 1.6em;">Как видим в любом случае (с налогами или без них) цену для нас добывает метод </span><strong style="line-height: 1.6em;">TiendaHelperBase::currency</strong><span style="line-height: 1.6em;">. Да прибудет с нами священный поиск... :) Находим его в: </span><code style="line-height: 1.6em;">administrator/components/co_tienda/helpers/_base.php</code><span style="line-height: 1.6em;"> И вот так он выглядит:</span></p>

<pre>
<code class="php">function currency($amount, $currency=&#39;&#39;, $options=&#39;&#39;)
{
  $currency_helper =&amp; TiendaHelperBase::getInstance( &#39;Currency&#39; );
  $amount = $currency_helper-&gt;_($amount, $currency, $options);
  return $amount;
}</code></pre>

<p><span style="line-height: 1.6em;">Второй параметр </span><strong style="line-height: 1.6em;">$currency</strong><span style="line-height: 1.6em;"> у нас пуст. По умолчанию дадим ему id валюты в которой мы хотим отображать цены на сайте. Это значение мы возьмем из таблицы валют </span><strong style="line-height: 1.6em;">tienda_currencies</strong><span style="line-height: 1.6em;"> поле </span><strong style="line-height: 1.6em;">currency_id</strong><span style="line-height: 1.6em;">. У меня валютой по умолчанию будет </span><strong style="line-height: 1.6em;">UAH </strong><span style="line-height: 1.6em;">и ее id = 5. Для меня код будет следующим:</span></p>

<pre>
<code class="php">function currency($amount, $currency=&#39;5&#39;, $options=&#39;&#39;)
{
  $currency_helper =&amp; TiendaHelperBase::getInstance( &#39;Currency&#39; );
  $amount = $currency_helper-&gt;_($amount, $currency, $options);
  return $amount;
}</code></pre>

<p><span style="line-height: 1.6em;">Вот с этих пор </span><strong style="line-height: 1.6em;">валюта </strong><span style="line-height: 1.6em;">отображения по умолчанию у нас есть. Но в модуле выбора </span><strong style="line-height: 1.6em;">валюты </strong><span style="line-height: 1.6em;">(если Вы им пользуетесь) при том что отображение в гривнах валюта базы данных по умолчанию т.е. </span><strong style="line-height: 1.6em;">USD</strong><span style="line-height: 1.6em;">, что смотрится не совсем красиво. Я предлагаю, дабы не заморачиваться с кодом, просто присвоить валютам что находятся впереди нашей валюты по умолчанию присвоить более высокие id. В моем случае у валюты </span><strong style="line-height: 1.6em;">USD </strong><span style="line-height: 1.6em;">id=1. Ей, через </span><strong style="line-height: 1.6em;">phpmyadmin</strong><span style="line-height: 1.6em;">, присваеваем id, например 8. Вуаля!</span></p>
