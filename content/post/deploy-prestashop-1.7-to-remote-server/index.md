---
author: "Serhii Polishchuk"
title: "Deploy PrestaShop 1.7 to the remote server"
date: 2018-08-24
tags: ["PHP", "Ansible"]
draft: false
---
<!--more-->
## Foreword

I have not so much experience either with PrestaShop or other CMS. 
I have started my way with **Joomla 1.5, WordPress and PrestaShop** 10 years ago
but it was another story.
Today I had need to make yet another art shop with prints, 
as simplest solution was made it gain with existing solution.

So I came back and look at **PrestaShop**, that has a few versions through this time.
It has a lot of changes in codebase but nothing happen with its low quality.
**Symfony** got a lot of gain to Presta internals, but **e-commerce is too complex**
and people don't want to change the way that soft was written.

I spent more then two weeks catch bugs one by one, _and the only way that you can use
**PrestaShop** solution is the a way the maintainers suppose to be_. So, I throw away 
all unneeded extensions, got default template, got payment and shipping modules, add content
and deploy all of that stuff to the remote server. 

When you deploy it first time and need to make changes you have two way of doing so, 
because you have two equal instances. 
1. First one is edit your remote application through ftp, ssh etc.
You got inconsistency with your local installation and *make changes on your live website is
a bad practice*.
2. Second is to somehow move your changes to live server from your local app. 
It is more consistent and modern way.

## Ansible

I have find out **Ansible** at *Symfony Camp 2014* 
from *Igor Brovchenko* speech https://www.youtube.com/watch?v=C94jF1PBvEE

Before I have use **Capifony** for *Symfony project deploy*, however Ruby language overhead
and maintainance issues (currently Capifony is unmaintained) move my attention to **Ansible**
  
**Ansible** is wrote on Python but there is nothing to do with code itself. Ansible provide
a flexible way of configuration over code. You can do all you want with tons of modules for
everyone and everything.  

If you don't know about Ansbile yet, I'm strongly recommend the Ryan Weaver's traning course
"Ansible for Automation!" https://knpuniversity.com/screencast/ansible It's a fast and
 most complete tutorial.
 
Another great way is to research it through official website https://www.ansible.com/

## PrestaShop 1.7 Ansible playbook 

Usually I'm prepare the server before deploy by hands, so I simply skip the install php, mysql
and other stuff in my configuration, other parts of deploy stage I comment out.

PrestaShop like many others CMS in their standard edition rely only on the file cache and mysql.
For the database I have create separate mysql role that gathering facts e.g. db name, password
and username, however on both machines I have used **.my.cnf**, 
that I strongly recommend to you.

Mysql role vars:

```yaml
---

database_name: "{{ database_name_std.stdout }}"
database_user: "{{ database_user_std.stdout }}"
database_pass: "{{ database_pass_std.stdout }}"

```

Mysql role tasks:
```yaml
---

- name: Get database name
  shell: php -r 'echo (require "{{current_release}}/app/config/parameters.php")["parameters"]["database_name"].PHP_EOL;'
  register: database_name_std

- name: Get database user
  shell: php -r 'echo (require "{{current_release}}/app/config/parameters.php")["parameters"]["database_name"].PHP_EOL;'
  register: database_user_std

- name: Get database pass
  shell: php -r 'echo (require "{{current_release}}/app/config/parameters.php")["parameters"]["database_name"].PHP_EOL;'
  register: database_pass_std

```

Deploy task

```yaml
---
- hosts: local

  vars:
    project_name: shop_cherevan_art
    project_dir: /home/myuser/www/shop.cherevan
    current_release: "{{project_dir}}" # mysql role has rely on this var

  roles:
    - mysql

  tasks:
    - name: Dump database
      shell: mysqldump {{database_name}} > {{project_dir}}/dump.sql

    - name: Compress whole website
      archive:
        path: "{{project_dir}}/*"
        dest: /tmp/{{project_name}}.tgz

#################################################################
##################### P R O D U C T I O N #######################
#################################################################

- hosts: prod

  vars:
    project_name: shop_cherevan_art
    project_dir: /var/www/shop.cherevan.art
    project_dir_local: /home/myuser/www/shop.cherevan
    host: shop.cherevan.art
    local_host: shop.cherevan.local
    host_regxp: http:\/\/{{host}}
    local_host_regxp: http:\/\/{{local_host}}
    timestamp: "{{ansible_date_time.epoch}}"
    number_of_keeped_releases: 5
    new_release_dir: "{{project_dir}}/releases/{{timestamp}}"
    current_release: "{{project_dir}}/current"
    parameters: app/config/parameters.php

  roles:
    - mysql

  tasks:
    - name: copy archive to remote
      copy:
        src: /tmp/{{project_name}}.tgz
        dest: /tmp

    - name: Dump current mysql dump to current release
      shell: mysqldump {{database_name}} > {{current_release}}/dump.sql

    - name: Recreate current mysql db
      shell: mysql -e "drop database {{database_name}};" && mysql -e "create database {{database_name}};"

    - name: Make directory for new release
      file:
        path: "{{new_release_dir}}"
        state: directory
        mode: 0755

    - name: Extract archive
      unarchive:
        src: /tmp/{{project_name}}.tgz
        dest: "{{new_release_dir}}"
        remote_src: yes

    - name: Copy parameters from current release to new
      copy:
        remote_src: true
        src: "{{current_release}}/{{parameters}}"
        dest: "{{new_release_dir}}/{{parameters}}"

    #############################################################################
    #####    DANGER ZONE! BE CAREFUL WHILE DELETE FOLDERS ON THE SERVER    ######
    #############################################################################

    - name: List of outdated releases
      shell: "find {{project_dir}}/releases/* -maxdepth 0 -type d|sort -r|tail -n +{{number_of_keeped_releases + 1}}"
      register: outdated_releases

    - name: Delete outdated releases
      file:
        state: absent
        path: "{{ item }}"
      with_items: "{{ outdated_releases.stdout_lines }}"

    #############################################################################
    #####    DANGER ZONE! BE CAREFUL WHILE DELETE FOLDERS ON THE SERVER    ######
    #############################################################################

    #############################################################################
    ###################    PREPARE DATABASE AND IMPORT    #######################
    #############################################################################

    - name: Replace url
      replace:
        path: "{{new_release_dir}}/dump.sql"
        regexp: '{{local_host_regxp}}'
        replace: '{{host_regxp}}'

    - name: Replace domain
      replace:
        path: "{{new_release_dir}}/dump.sql"
        regexp: '{{local_host}}'
        replace: '{{host}}'

    - name: Import mysql dump
      shell: mysql {{database_name}} < {{new_release_dir}}/dump.sql

    - name: Enable https
      shell: mysql -e 'UPDATE {{database_name}}.ps_configuration SET value=1 WHERE name="PS_SSL_ENABLED";'

    - name: Enable https everywhere
      shell: mysql -e 'UPDATE {{database_name}}.ps_configuration SET value=1 WHERE name="PS_SSL_ENABLED_EVERYWHERE";'

    #############################################################################
    ###################    PREPARE DATABASE AND IMPORT    #######################
    #############################################################################

    - name: Disable debug mode
      replace:
        path: "{{new_release_dir}}/config/defines.inc.php"
        regexp: "define\\('_PS_MODE_DEV_', true"
        replace: "define('_PS_MODE_DEV_', false"

    - name: Clear cache of new release
      shell: "{{current_release}}/app/console cache:clear -e prod"

    - name: Make symbolic link to new release
      file:
        src: "{{new_release_dir}}"
        dest: "{{current_release}}"
        state: link
```
