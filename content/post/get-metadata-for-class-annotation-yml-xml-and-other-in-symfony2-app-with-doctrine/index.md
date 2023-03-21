---
author: "Serhii Polishchuk"
title: "Get metadata for class (annotation, yml, xml and other) in Symfony2 app with Doctrine"
date: 2013-10-25
tags: ["Doctrine2", "Symfony2"]
draft: false
---
<!--more-->
The common case for application that used Doctrine2 is get metadata from class. This very simple, you only need inject manager (document or entity) in your service, and than simply use:

```php
$this->documentManager->getMetadataFactory()->getMetadataForClass($class);
```
