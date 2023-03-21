---
author: "Serhii Polishchuk"
title: "Canon IP1800 в Ubuntu 10.10"
date: 2012-05-05
tags: ["Ubuntu"]
draft: true
---
<!--more-->
<p>Спасибо что есть компания Canon, что бы мы без них делали... наверное лишний час отдохнули бы от дурной работы. На этот раз сюрприз с принтером Canon IP1800 - компания &quot;забыла&quot; выпустить драйвера под Линукс. И вот опять после продолжительного гугления решение было найдено...<!--more--> <a href="http://www.ubuntubuzz.com/2011/06/download-install-canon-printer-driver.html" rel="nofollow" target="_blank">Сайт автора</a></p>

<pre>
<code class="bash">sudo add-apt-repository ppa:michael-gruz/canon
sudo apt-get update</code></pre>

<p>Далее сама установка:</p>

<pre>
<code class="bash">sudo apt-get install cnijfilter-ip100series</code></pre>

<p>Из всевозможных моделей:</p>

<p>Canon Pixma iP Series Ubuntu driver<br />
Canon iP100 Ubuntu Driver &mdash; cnijfilter-ip100series<br />
Canon iP1800 Ubuntu Driver &mdash; cnijfilter-ip1800series<br />
Canon iP1000 Ubuntu Driver &mdash; cnijfilter-pixmaip1000series<br />
Canon iP1500 Ubuntu Driver &mdash; cnijfilter-pixmaip1500series<br />
Canon iP1900 Ubuntu Driver &mdash; cnijfilter-ip1900series<br />
Canon iP 2200 Ubuntu Driver &mdash; cnijfilter-ip2200series<br />
Canon iP2500 Ubuntu Driver &mdash; cnijfilter-ip2500series<br />
Canon iP2600 Ubuntu Driver &mdash; cnijfilter-ip2600series<br />
Canon iP2700 Ubuntu Driver &mdash; cnijfilter-ip2700series<br />
Canon iP3300 Ubuntu Driver &mdash; cnijfilter-ip3300series<br />
Canon iP3500 Ubuntu Driver &mdash; cnijfilter-ip3500series<br />
Canon iP3600 Ubuntu Driver &mdash; cnijfilter-ip3600series<br />
Canon iP4200 Ubuntu Driver &mdash; cnijfilter-ip4200series<br />
Canon iP4500 Ubuntu Driver &mdash; cnijfilter-ip4500series<br />
Canon iP4700 Ubuntu Driver &mdash; cnijfilter-ip4700series<br />
Canon iP4800 Ubuntu Driver &mdash; cnijfilter-ip4800series<br />
Canon iP5200 Ubuntu Driver &mdash; cnijfilter-ip5200series<br />
Canon iP6600 Ubuntu Driver &mdash; cnijfilter-ip6600series<br />
Canon iP7500 Ubuntu Driver &mdash; cnijfilter-ip7500series<br />
Canon Pixma MG Series Ubuntu Driver<br />
Canon MG5100 Ubuntu Driver &mdash; cnijfilter-mg5100series<br />
Canon MG5200 Ubuntu Driver &mdash; cnijfilter-mg5200series<br />
Canon MG6100 Ubuntu Driver &mdash; cnijfilter-mg6100series<br />
Canon MG8100 Ubuntu Driver &mdash; cnijfilter-mg8100series<br />
Canon PIXMA MP Series Ubuntu Driver<br />
Canon MP140 Ubuntu Driver &mdash; cnijfilter-mp140series<br />
Canon MP160 Ubuntu Driver &mdash; cnijfilter-mp160series<br />
Canon MP190 Ubuntu Driver &mdash; cnijfilter-mp190series<br />
Canon MP210 Ubuntu Driver &mdash; cnijfilter-mp210series<br />
Canon MP240 Ubuntu Driver &mdash; cnijfilter-mp240series<br />
Canon MP490 Ubuntu Driver &mdash; cnijfilter-mp490series<br />
Canon MP500 Ubuntu Driver &mdash; cnijfilter-mp500series<br />
Canon MP510 Ubuntu Driver &mdash; cnijfilter-mp510series<br />
Canon MP520 Ubuntu Driver &mdash; cnijfilter-mp520series<br />
Canon MP540 Ubuntu Driver &mdash; cnijfilter-mp540series<br />
Canon MP550 Ubuntu Driver &mdash; cnijfilter-mp550series<br />
Canon MP560 Ubuntu Driver &mdash; cnijfilter-mp560series<br />
Canon MP600 Ubuntu Driver &mdash; cnijfilter-mp600series<br />
Canon MP610 Ubuntu Driver &mdash; cnijfilter-mp610series<br />
Canon MP630 Ubuntu Driver &mdash; cnijfilter-mp630series<br />
Canon MP640 Ubuntu Driver &mdash; cnijfilter-mp640series<br />
Canon MX Series Ubuntu Driver<br />
Canon MX320 Ubuntu Driver &mdash; cnijfilter-mx320series<br />
Canon MX330 Ubuntu Driver &mdash; cnijfilter-mx330series<br />
Canon MX350 Ubuntu Driver &mdash; cnijfilter-mx350series<br />
Canon MX360 Ubuntu Driver &mdash; cnijfilter-mx360series<br />
Canon MX410 Ubuntu Driver &mdash; cnijfilter-mx410series<br />
Canon MX420 Ubuntu Driver &mdash; cnijfilter-mx420series<br />
Canon MX860 Ubuntu Driver &mdash; cnijfilter-mx860series<br />
Canon MX870 Ubuntu Driver &mdash; cnijfilter-mx870series<br />
Canon MX880 Ubuntu Driver &mdash; cnijfilter-mx880series<br />
Canon PIXUS Series Ubuntu Driver<br />
Pixus 550 Ubuntu Driver &mdash; cnijfilter-pixus5510iseries<br />
Pixus 560 Ubuntu Driver &mdash; cnijfilter-pixus560iseries<br />
Pixus 850 Ubuntu Driver &mdash; cnijfilter-pixus850iseries<br />
Pixus 860 Ubuntu Driver &mdash; cnijfilter-pixus860iseries<br />
Pixus 865 Ubuntu Driver &mdash; cnijfilter-pixus865iseries<br />
Pixus 950 Ubuntu Driver &mdash; cnijfilter-pixus950iseries<br />
Pixus 990 Ubuntu Driver &mdash; cnijfilter-pixus990iseries<br />
Canon Pixus ip3100 Ubuntu Driver &mdash; cnijfilter-pixusip3100series<br />
Canon Pixus ip4100 Ubuntu Driver &mdash; cnijfilter-pixusip4100series<br />
Canon Pixus ip8600 Ubuntu Driver &mdash; cnijfilter-pixusip8600series</p>
