---
author: "Serhii Polishchuk"
title: "DoctrineCacheBundle + APC"
date: 2014-10-06
tags: ["Doctrine2", "Symfony2"]
draft: false
---
<!--more-->
I don't know why, but I still can't connect **DoctrineCacheBundle with Memcache**, so I use **APC** as it work perfect and fast. After [install bundle](https://github.com/doctrine/DoctrineCacheBundle#installation) you need some simple configuration:
    doctrine_cache:
        providers:
            apc_cache:
                type: apc
                namespace: your_unique-namespace
                aliases:
                    - apc_cache</pre>
And now you can simply use your cache service:
    container->get('apc_cache');
    $cache->save('qwerty', 'DataDataData');
    var_dump($cache->fetch('qwerty'));
In service you have some very simple methods (from [documentation](http://docs.doctrine-project.org/en/2.0.x/reference/caching.html#cache-drivers)):
- fetch($id) - Fetches an entry from the cache.
- contains($id) - Test if an entry exists in the cache.
- save($id, $data, $lifeTime = false) - Puts data into the cache.
- delete($id) - Deletes a cache entry.

That all, now you have very powerfull cache system in your **Symfony2** application.
