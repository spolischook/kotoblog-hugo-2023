---
author: "Serhii Polishchuk"
title: "Валюта не по умолчанию для Tienda 0.70"
date: 2012-02-20
tags: ["Joomla"]
draft: true
---
<!--more-->
<p>Я так подозреваю что возможность отображать цены во фронтенде в валюте отличной от той что в базе будет в последующих версиях, но все равно опыт интересен, и я по просьбе lmd и по своему хотению решил немного покопаться в коде. Вот что я нарыл... <!--more--> Посмотрев шаблоны вывода цены, обнаружил что за это отвечает метод <strong>dispayPriceWithTax</strong> класса <strong>TiendaHelperProduct</strong> его можно найти по адресу <code>administrator/components/co_tienda/helpers/product.php</code> Выглядит он примерно так:</p>

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

<p>Как видим в любом случае (с налогами или без них) цену для нас добывает метод <strong>TiendaHelperBase::currency</strong>. Да прибудет с нами священный поиск... :) Находим его в: <code>administrator/components/co_tienda/helpers/_base.php</code> И вот так он выглядит:</p>

<pre>
<code class="php">function currency($amount, $currency=&#39;&#39;, $options=&#39;&#39;)
    {
        $currency_helper =&amp; TiendaHelperBase::getInstance( &#39;Currency&#39; );
        $amount = $currency_helper-&gt;_($amount, $currency, $options);
        return $amount;
    }</code></pre>

<p>Второй параметр <strong>$currency</strong> у нас пуст. По умолчанию дадим ему id валюты в которой мы хотим отображать цены на сайте. Это значение мы возьмем из таблицы валют <strong>tienda_currencies</strong> поле <strong>currency_id</strong>. У меня валютой по умолчанию будет UAH и ее id = 5. Для меня код будет следующим:</p>

<pre>
<code class="php">function currency($amount, $currency=&#39;5&#39;, $options=&#39;&#39;)
    {
        $currency_helper =&amp; TiendaHelperBase::getInstance( &#39;Currency&#39; );
        $amount = $currency_helper-&gt;_($amount, $currency, $options);
        return $amount;
    }</code></pre>

<p>&nbsp;</p>
