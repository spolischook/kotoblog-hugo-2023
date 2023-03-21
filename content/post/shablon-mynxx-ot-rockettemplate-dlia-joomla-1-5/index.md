---
author: "Serhii Polishchuk"
title: "Шаблон Mynxx от Rockettemplate для Joomla 1.5"
date: 2011-02-21
tags: ["Joomla"]
draft: true
---
<!--more-->
<p>Нашел очень хорошую тему для интернет магазина, после недолгого шаманства, шаблон все таки встал, но вот отображаться во фронтэнде напрочь отказался. Выдал ошибку:</p>

<pre>
<code class="html">Fatal error: Class &#39;JFile&#39; not found in /home/content/e/n/f/enfurno23/html/templates/rt_mynxx_j15/rt_utils.php on line 16</code></pre>
<!--more-->

<p>Решение подсмотрел на англоязычном форуме. Такое же простое как и ошибка :) Заменим следующий код:</p>

<pre>
<code class="php">if (JFile::exists($menu_params_file)) :
$menu_params_content = file_get_contents($menu_params_file);
eval(&quot;\$module-&gt;params = \&quot;$menu_params_content\&quot;;&quot;);
endif;
$topnav = $renderer-&gt;render( $module, $options );</code></pre>

<p>на:</p>

<pre>
<code class="php">​if (jimport(&#39;joomla.filesystem.file&#39;)) :
$menu_params_content = file_get_contents($menu_params_file);
eval(&quot;\$module-&gt;params = \&quot;$menu_params_content\&quot;;&quot;);
endif;
$topnav = $renderer-&gt;render( $module, $options );</code></pre>
