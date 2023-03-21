---
author: "Serhii Polishchuk"
title: "PrestaShop полезные мелочи"
date: 2011-08-13
tags: ["PrestaShop"]
draft: true
---
<!--more-->
<p>Здесь немного практических полезных примеров при разработке сайта на PrestaShop. <!--more--> Получаем <strong>массив всех категорий</strong> из БД:</p>

<pre>
<code class="php">$categories = Category::getCategories((int)($cookie-&gt;id_lang), false);</code></pre>

<p>Получаем <strong>массив категорий</strong> к которым пренадлежит продукт с ИД в запросе (GET or POST):</p>

<pre>
<code class="php">Product::getProductCategories((int)Tools::getValue(&#39;id_product&#39;));</code></pre>

<p>&nbsp;</p>
