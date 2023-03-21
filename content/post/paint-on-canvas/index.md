---
author: "Serhii Polishchuk"
title: "Малюю на <canvas>"
date: 2018-09-04
tags: ["Javascript"]
draft: false
---
<!--more-->
У зв'язку з останніми подіями в моєму житті, а саме вступом до ВНЗ на 3 курс, маю змогу більше часу приділяти навчанню чого небудь.
Наприклад малювати графіки у canvas. 

Тут будуть невеличкі приклади малювання у canvs 

    let canv = document.getElementById('myCanvas'),
        ctx = canv.getContext('2d')
    ;

    ctx.fillStyle = "rgb(200,0,0)";
    ctx.fillRect(10,10,100,100);

    ctx.fillStyle = "rgba(0,0,200,0.5)";
    ctx.fillRect(60,60,100,100);

    ctx.fillStyle = "rgb(139,69,19)";
    ctx.fillRect(180, 10, 100, 100);
    ctx.clearRect(200, 30, 60, 60);
    ctx.strokeRect(205, 35, 50, 50);

    ctx.beginPath();
    ctx.moveTo(120, 10);
    ctx.lineTo(170, 10);
    ctx.lineTo(145, 35);
    ctx.closePath();
    ctx.stroke();
    ctx.closePath();

    for (let j = 0; j < 2; j++) {
        for (let i = 1; i <= 6; i++) {
            ctx.beginPath();
            let degrees = 360/6*i;
            ctx.arc(25 + (50 * (i-1)), 200 + (60 * j), 20, 0, (Math.PI/180)*degrees, 0 === i%2);
            j === 0 ? ctx.stroke() : ctx.fill();
            ctx.closePath();
        }
    }

<canvas id="myCanvas" width="300" height="300">
</canvas>
<script>
let canv = document.getElementById('myCanvas'),
ctx = canv.getContext('2d')
;

ctx.fillStyle = "rgb(200,0,0)";
ctx.fillRect(10,10,100,100);

ctx.fillStyle = "rgba(0,0,200,0.5)";
ctx.fillRect(60,60,100,100);

ctx.fillStyle = "rgb(139,69,19)";
ctx.fillRect(180, 10, 100, 100);
ctx.clearRect(200, 30, 60, 60);
ctx.strokeRect(205, 35, 50, 50);

ctx.beginPath();
ctx.moveTo(120, 10);
ctx.lineTo(170, 10);
ctx.lineTo(145, 35);
ctx.closePath();
ctx.stroke();
ctx.closePath();

for (let j = 0; j < 2; j++) {
for (let i = 1; i <= 6; i++) {
ctx.beginPath();
let degrees = 360/6*i;
ctx.arc(25 + (50 * (i-1)), 200 + (60 * j), 20, 0, (Math.PI/180)*degrees, 0 === i%2);
j === 0 ? ctx.stroke() : ctx.fill();
ctx.closePath();
}
}
</script>
