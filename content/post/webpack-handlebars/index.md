---
author: "Serhii Polishchuk"
title: "Webpack + Handlebars"
date: 2018-01-31
tags: ["Javascript", "Webpack"]
draft: false
---
<!--more-->
Вовремя очередного набега на **javascript** заметил как разрастается **jQuery** лапша если ее размешивать html кодом.
Вспомнились те времена когда писал на EmberJs и отличнейший шаблонизатор **Handlebars**. Так как я уже использую Webpack для своих проектов, я попробовал тот же способ require и compile - все как в доке. Но вебпак в своей привычной манере посоветовал воспользоватся лоадером.

Чтож сказано - сделано. Итак лоадер - https://github.com/pcardune/handlebars-loader 
PhpStorm plugin - https://github.com/dmarcotte/idea-handlebars

Далее все как обычно, добавляем новый лоадер в webpack.config.js:
```js
let webpackConfig = {
    module: {
        rules: [
            {
                test: /\.hbs$/,
                use: "handlebars-loader"
            },
        ]
    }
};
```
Создаем шаблон 
```handlebars
<div class="alert alert-{{#if alertType}}{{alertType}}{{else}}info{{/if}} alert-dismissible fade show" role="alert">
    {{ message }}
    <button type="button" class="close" data-dismiss="alert" aria-label="Close">
        <span aria-hidden="true">&times;</span>
    </button>
</div>
```
Генерим html:
```js
let alertHtml = AlertTemplate({
    alertType: 'success',
    message: 'You are successfully using handlebars template!'
})
```
Больше примеров использования здесь https://github.com/pcardune/handlebars-loader/tree/master/examples
