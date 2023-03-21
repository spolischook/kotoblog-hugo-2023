---
author: "Serhii Polishchuk"
title: "Get class name from namespace"
date: 2013-10-20
tags: []
draft: false
---
<!--more-->
It's a good trick from Gedmo doctrine extensions lib:

    $className = substr($namespace, strrpos($namespace, '\\') + 1);
