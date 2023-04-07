---
title: "GoLang HTTP Request Validation"
slug: golang-request-validation
date: 2023-04-05T09:31:52+02:00
draft: true
tags: ["golang"]
---
Previously in Symfony I have no such a question "How to validate HTTP Request".
[The documentation](https://symfony.com/doc/current/validation.html) 
has been clearly explain this topic.
In [gin framework](https://gin-gonic.com/docs/examples/binding-and-validation/) 
it's recommended to use go validator with tags.
Honestly I don't think that is a good solution, and also it's not recommended
[by GoLang contributors](https://github.com/golang/go/issues/38641#issuecomment-634834028).
<!--more-->

- [Creating test environment](#creating-test-environment)
- [Tests](#tests)
  - [Unit tests](#unit-tests)
  - [UI tests](#ui-tests)
- [Choosing validation library](#choosing-validation-library)
  - [Ozzo-validation](#ozzo-validation)
  - [Jio](#jio)
- [Conclusion](#conclusion)

So I start to look for validation approach that would be convenient in Go.
I like how validation rules in Symfony described by [YAML](https://g.co/kgs/YRKu47), 
but this approach requires a 
[metaprogramming](https://en.wikipedia.org/wiki/Metaprogramming#:~:text=Metaprogramming%20is%20a%20programming%20technique,even%20modify%20itself%20while%20running.).
This kind of declarative style of
validation rules, that separated from code, helps to keep it simple and demonstrative.
Unfortunately there is no out of the box solution for YAML validation rules in Go.

## Creating test environment

In daily work we faced with issue for validation the request from customer, 3-d party
integration services and other internal services as well.
We deal with **HTTP requests** with **path parameters** in URL, **query params** and **json body** - 
so that is what we should validate.  
Gin Framework used for routing so the main function will look like that:
```go
r := gin.Default()
r.GET("/booking", handler.BookingList)
r.POST("/booking", handler.BookingCreate)
r.POST("/booking/:id", handler.BookingUpdate)
r.Run(Port)
```

First endpoint `bookingList` can have a different query params:
- `from`, `to` - should be in format `YYYY-MM-DD`
- `productId` - must be an integer
- `orderSource` - should be in enum [`b2b`, `retail`]. In case of any other value should be ignored

Second is for creating a new Booking with json body that should be validated like that:
```yaml
BookingCreateRequest:
  properties:
    productId: 
      - type: integer
    orderSource:
      - Choice: [b2b, retail]
    tourDate:
      - type: date
    pax:
      - type: integer
    customer:
      - type: object
      - properties:
        - id:
          - type: integer
        - email: 
          - Email: ~
```

Go structs for body will look like:
```go
type Booking struct {
	Product  Product   `json:"product"`
	Customer Customer  `json:"customer"`
	Pax      int       `json:"pax"`
	TourDate time.Time `json:"tourDate"`
}
type Customer struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
}
type Product struct {
	Id int `json:"id"`
}
```
Field `customer` could be two types:
- if `id` field will be set, we should check that customer with that id is exists, `email` field will be ignored
- if `email` filed will be set, it should be valid email and new customer will be created

The third `BookingUpdate` request very similar to previous one with few differences:
- `orderSource` in request body should be ignored, because we don't want to change it after booking creattion
- `:id` in URL must be an integer - otherwise `400` error must be returned
- Booking with `:id` must exist, otherwise `404` must be returned

This simplified app example covers most validation cases in our services.

## Tests

As I'm usually preach TDD approach, the test cases would be created first.

### Unit tests

There is a validation 

### UI tests

The test data will be a json file with different labeled test cases.
```json
{
  "invalid request": {
    "request": {
      "method": "POST",
      "url": "/booking/123",
      "body": {
        ...
      }
    },
    "response": {
      "status_code": 400,
      "body": {
        ...
      }
    }
  }
}
```


## Choosing validation library

In the [awesome Go]() I found several validation libraries. Few of them are relies on struct tags,
so they automatically were rejected by me. There are few of them that I found quite interesting:
- [ozzo-validation](https://github.com/go-ozzo/ozzo-validation)
- [jio](https://github.com/faceair/jio)

### Ozzo-validation

Ozzo is good enough to make a validation.

### Jio

But **jio** used string for naming the fields,
that could be a good for validation through yaml files.

## Conclusion
