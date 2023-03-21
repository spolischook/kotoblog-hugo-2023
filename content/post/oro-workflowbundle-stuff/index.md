---
author: "Serhii Polishchuk"
title: "Oro WorkflowBundle stuff"
date: 2015-12-20
tags: ["PHP", "oro/platform", "oro/crm"]
draft: true
---
<!--more-->
You can use some kind of Expressions in Workflow.
See implementations of **Oro\Component\ConfigExpression\ExpressionInterface** for see all of them
This properties of Transition definitions are mostly used.
Most used Conditions:
    @and
    @or
    @blank
    @not_blank
    @not_empty
    @equal
    @greater_or_equal
    @greater
    @less
    @channel_entity_availiable

All actions implements **Oro\Bundle\WorkflowBundle\Model\Action\ActionInterface**.
See implementations.
Most used Actions:
    @create_entity
    @find_entity
    @assign_value
    @unset_value
    @call_method
    @send_email
