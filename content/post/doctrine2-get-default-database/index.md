---
author: "Serhii Polishchuk"
title: "Doctrine2 get default database"
date: 2014-11-06
tags: ["Doctrine2", "Symfony2"]
draft: false
---
<!--more-->
Simple use:

    $this->getContainer()->get('doctrine.odm.mongodb.document_manager')->getConfiguration()->getDefaultDB();
