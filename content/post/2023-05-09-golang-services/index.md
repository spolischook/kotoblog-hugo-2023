---
title: "Golang Services"
date: 2023-05-09T11:09:38+02:00
draft: true
---

There is no easy to use conceptions of services in GoLang.
I found some implementation of [Dependency Injection from Uber](https://github.com/uber-go/dig) but it's huge and
I don't believe it's well fit into the overall GoLang philosophy.

<!--more-->

Instead of that GoLang has a concept of [singleton](https://refactoring.guru/design-patterns/singleton)
that I found into [GLG logger](https://github.com/kpango/glg).
