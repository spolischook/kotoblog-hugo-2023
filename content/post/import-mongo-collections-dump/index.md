---
author: "Serhii Polishchuk"
title: "Import mongo collections dump"
date: 2014-09-23
tags: ["MongoDb"]
draft: false
---
<!--more-->
If you have **mongoDb dump** with many **.bson** files, you can import it one by one by this command:
    mongorestore -d database_name -c collection_name /path/to/dump_file.bson
All simple use this simple import script (it work with base script):
    https://gist.github.com/4402cfdf36543e6ccfef
Import script:
    https://gist.github.com/08ebf9462ad9eb2bdeed
