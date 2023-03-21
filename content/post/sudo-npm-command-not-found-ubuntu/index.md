---
author: "Serhii Polishchuk"
title: "sudo npm command not found ubuntu"
date: 2013-11-22
tags: ["Ubuntu", "NodeJs"]
draft: true
---
<!--more-->
I had trouble when try to install **less module** for **nodejs**. First I try this:

    npm install -g less

But got an error "**sudo npm command not found**". So I find a solution:

    sudo $(which npm) install -g less
