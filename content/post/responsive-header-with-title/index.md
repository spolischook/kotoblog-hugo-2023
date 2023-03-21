---
author: "Serhii Polishchuk"
title: "Responsive header with title"
date: 2018-08-30
tags: ["CSS"]
draft: false
---
<!--more-->
### Preface

It's very simple and common task to make responsive header with image and title.
Bootstrap already has 
[few breakpoints for mobile devices](https://getbootstrap.com/docs/4.0/layout/overview/#responsive-breakpoints), 
however I'll choose responsive breakpoints based on 
[iOs screen sizes](https://developer.apple.com/library/archive/documentation/DeviceInformation/Reference/iOSDeviceCompatibility/Displays/Displays.html)

### Media Queries

The most dificalt things that I should to understand is ```min-with``` and ```max-width``` queries

#### Max-Width

    @media (max-width: 320px)  {...}
    
This query means “If [device width] is less than or equal to 320px, then do {…}”

#### Min-Width

    @media (min-width: 1024px)  {...}

This query means “If [device width] is greater than or equal to 1024px, then do {…}”

The order of media queries matters, because CSS rules that appear later in the embedded styles 
override earlier rules if both have the same specificity. It is a very common bug that happens with media queries.

### Live example

First we need a picture at the top of the page

    <body>
    <style>
        div.header {
            background-image: url("header-2048.jpeg");
            background-repeat: no-repeat;
            background-position: center;
            background-size: cover;
            min-height: 450px;
        }
    </style>
    <div class="header">
    </div>
    </body>

That was added div with background as image. Next we are going to add title to the header.

    <body>
    <style>
        div.header {
            ...
            position: relative;
        }
        .header-caption span.container {
            display: block;
        }
        div.header h1 {
            margin: 0;
            position: absolute;
            bottom: 0;
            width: 100%;
            padding: 1rem 0 1rem 0;
            color: white;
            font-size: 2.5rem;
            background-color: rgba(0, 0, 0, 0.5);
        }
    </style>
    <div class="header">
        <h1 class="header-caption">
            <span class="container">
                Samui has amazing sunset!
            </span>
        </h1>
    </div>
    </body>

We add relative position for parent div and absolute for h1, with ```bottom: 0``` and ```margin: 0``` 
h1 will be at bottom at parent div, other styles will make it pretty. Span with ```container```
class needed for position title at the center of the page, at the same level as content.

Now we can describe media queries:

    @media (max-width: 1024px)  {
        div.header {
            background-image: url("header-1024.jpeg");
            min-height: 480px;
        }
    }
    @media (max-width: 834px)  {
        div.header {
            background-image: url("header-834.jpeg");
            min-height: 391px;
        }
    }
    @media (max-width: 768px)  {
        div.header {
            background-image: url("header-768.jpeg");
            min-height: 360px;
        }
    }
    @media (max-width: 414px)  {
        div.header {
            background-image: url("header-414.jpeg");
            min-height: 194px;
        }
        div.header h1 {
            font-size: 1.75rem;
            padding: 0.5rem 0 0.5rem 0;
        }
    }
    @media (max-width: 375px)  {
        div.header {
            background-image: url("header-375.jpeg");
            min-height: 176px;
        }
        div.header h1 {
            font-size: 1.5rem;
            /*padding: 0.5rem 0.7rem 0.5rem 0.7rem;*/
        }
    }
    @media (max-width: 320px)  {
        div.header {
            background-image: url("header-320.jpeg");
            min-height: 150px;
        }
        div.header h1 {
            font-size: 1rem;
            /*padding: 0.3rem 0.5rem 0.3rem 0.5rem;*/
        }
    }

```max-width``` media queries should be written from large to small.
e.g. for iPhone SE with width 320px can apply all of the queries but only last will be applied because of order CSS rules.

### Results:

#### Live demo

HTML result available by link - https://kotoblog.pp.ua/responsive-header-with-title.html
You can play with it though google console.

#### iPhone SE/5

<img src="http://kotoblog.s3.eu-central-1.amazonaws.com/Articles/117-responsive-header-with-title/iPhone_SE.png"  alt="iPhone SE/5" class="img-thumbnail img-fluid">
    
#### iPhone 6/7/8

<img src="http://kotoblog.s3.eu-central-1.amazonaws.com/Articles/117-responsive-header-with-title/iPhone_6.png" alt="iPhone 6 7 8" class="img-thumbnail img-fluid">

#### iPhone 6/7/8 Plus

<img src="http://kotoblog.s3.eu-central-1.amazonaws.com/Articles/117-responsive-header-with-title/iPhone_6_plus.png" alt="iPhone 6 7 8 Plus" class="img-thumbnail img-fluid">

#### iPhone X

<img src="http://kotoblog.s3.eu-central-1.amazonaws.com/Articles/117-responsive-header-with-title/iPhone_X.png" alt="iPhone X" class="img-thumbnail img-fluid">

#### iPad

<img src="http://kotoblog.s3.eu-central-1.amazonaws.com/Articles/117-responsive-header-with-title/iPad.png" alt="iPad" class="img-thumbnail img-fluid">

#### iPad Pro

<img src="http://kotoblog.s3.eu-central-1.amazonaws.com/Articles/117-responsive-header-with-title/iPad_Pro.png" alt="iPad Pro" class="img-thumbnail img-fluid">

#### Desktop 1920 width

<img src="http://kotoblog.s3.eu-central-1.amazonaws.com/Articles/117-responsive-header-with-title/desktop-1920.png" alt="Desktop 1920px" class="img-thumbnail img-fluid">
