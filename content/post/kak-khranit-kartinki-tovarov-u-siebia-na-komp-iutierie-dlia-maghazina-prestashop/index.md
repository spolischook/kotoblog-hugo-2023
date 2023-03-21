---
author: "Serhii Polishchuk"
title: "Как хранить картинки товаров у себя на компьютере для магазина Prestashop"
date: 2011-06-05
tags: ["PrestaShop"]
draft: true
---
<!--more-->
<p>У меня пока-что магазин вертится на шаред хостинге, то ли от недостатка уверенности что все получится, то ли от недостатка начального капитала :), к тому же Presta отлично ворочается на виртуальном хостинге, в отличии от Tienda Места у меня всего 500 Мб а картинок на 3Гб, и так как на домашнем сервере я все равно тестурую сайт я решил что пока не заимею свой сервачок, или хотябы не решусь на VDS, буду картинки хранить у себя, благо у мой сервер доступен из сети интернет 24/7. О том как это сделать <a href="http://kotoblog.pp.ua/ubuntu/d-link-domashnij-server-na-apache.html">читайте в моей предыдущей статье.</a> Для этого пришлось сделать не сложную модификацию класса... <!--more--> Файл /classes/Link.php там находим класс getImageLink:</p>

<pre>
<code class="php">public function getImageLink($name, $ids, $type = NULL)
{
    global $protocol_content;
    if ($this-&gt;allow == 1)
        $uri_path = __PS_BASE_URI__.$ids.($type ? &#39;-&#39;.$type : &#39;&#39;).&#39;/&#39;.$name.&#39;.jpg&#39;;
    else
        $uri_path = _THEME_PROD_DIR_.$ids.($type ? &#39;-&#39;.$type : &#39;&#39;).&#39;.jpg&#39;;
    return $protocol_content.Tools::getMediaServer($uri_path).$uri_path;
}</code></pre>

<p>В первом случае мы получаем ссылку с включенным ЧПУ во втором с выключенным. Для того чтобы создать ссылку на картинку на своем сервере я изменил класс следующим образом:</p>

<pre>
<code class="php">public function getImageLink($name, $ids, $type = NULL)
{
    global $protocol_content;
    if ($this-&gt;allow == 1)
        $uri_path = __PS_BASE_URI__.$ids.($type ? &#39;-&#39;.$type : &#39;&#39;).&#39;/&#39;.$name.&#39;.jpg&#39;;
    else
        $uri_path = &#39;http://MYDOMEN.dyndns.org/presta/img/p/&#39;.$ids.($type ? &#39;-&#39;.$type : &#39;&#39;).&#39;.jpg&#39;;
    return $uri_path;
}</code></pre>

<p>Не забудьте выключить ЧПУ, или модифицировать для себя, и заменить http://MYDOMEN.dyndns.org на свой.</p>
