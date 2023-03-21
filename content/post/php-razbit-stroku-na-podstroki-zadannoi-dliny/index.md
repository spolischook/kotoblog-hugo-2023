---
author: "Serhii Polishchuk"
title: "PHP: разбить строку на подстроки заданной длины"
date: 2012-08-17
tags: ["PHP"]
draft: false
---
<!--more-->
```php
$result="";
for ($start=0, $length=300; $subtext = substr($text, $start, $length); $start = $start + 300){
	$result .= $subtext;
}
```

А если нам нужно **разбить строку на предложения**? Ответ есть:

```php
$substr = explode(".", $text);
foreach ($substr AS $key => $value){
    if ($value != "") {
        echo $value; //одно предложение
    }
    //fix троеточие
    else {

    }
}
```
