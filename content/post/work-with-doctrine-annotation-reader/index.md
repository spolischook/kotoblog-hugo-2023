---
author: "Serhii Polishchuk"
title: "Work with doctrine annotation reader"
date: 2013-12-30
tags: ["Doctrine2"]
draft: false
---
<!--more-->
I need for one project to work with class **annotation** - reed and do some work. Because I already usedoctrine dbal, so I decide to use AnnotationReader.
So the first of all I create my own Annotation class:

    https://gist.github.com/e4125658cb1c0c6761aa

@Annotation - this means that this class use as annotation
@Target ("PROPERTY") - use it only for a property. If use on class or method - exception will be throw

We need some class:
    https://gist.github.com/9a109bd47cd335d7115f

Then we need use AnnotationReader of Doctrine with Reflection. So:
    https://gist.github.com/0bd36a2cacbd92b05c31

Annotation reader don't use composer autoload for some reasons. So you need to include annotation class.

Now you can see all mapping information. Enjoy!
