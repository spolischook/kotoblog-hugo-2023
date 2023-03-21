---
author: "Serhii Polishchuk"
title: "Select pager with javascript reload page"
date: 2015-11-25
tags: ["Javascript"]
draft: false
---
<!--more-->
In Twig:

    {% set limit = app.request.query.get('limit')|default(10) %}
    <select class="form-control" onchange="pager()" id="perPageSelector">
        <option {% if limit == 10 %}selected{% endif %} value="10">10</option>
        <option {% if limit == 25 %}selected{% endif %} value="25">25</option>
        <option {% if limit == 50 %}selected{% endif %} value="50">50</option>
        <option {% if limit == 100 %}selected{% endif %} value="100">100</option>
    </select>

In javascript:

    function pager(val) {
        var pager = document.getElementById("log_search_type_limit"),
            perpage = pager.options[pager.selectedIndex].value,
            url = replaceQueryParam('limit', perpage, window.location.search);
        window.location = window.location.pathname + url;
    }
    
    function replaceQueryParam(param, newval, search) {
        var regex = new RegExp("([?;&])" + param + "[^&;]*[;&]?");
        var query = search.replace(regex, "$1").replace(/&$/, '');
    
        return (query.length > 2 ? query + "&" : "?") + (newval ? param + "=" + newval : '');
    }
