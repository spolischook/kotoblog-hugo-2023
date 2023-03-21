---
author: "Serhii Polishchuk"
title: "О TypeScript по человечески"
date: 2018-02-26
tags: ["Javascript"]
draft: false
---
<!--more-->
## Введение

Все примеры кода и сама статья актуальна для версии TypeScript 2.7
Актуальную версию можно скчать и установить с оффициального сайта https://www.typescriptlang.org
Там же можно прочитать официальную документацию, вольный пересказ которой на
русском можно читать ниже.

Устанавливаем: 
```bash
➜ npm install -g typescript
```

Пишем Hello World скрипт:
```typescript
// hello-world.ts
let greeting: string;
greeting = 'Hello World!';
console.log(greeting);
```

Пробуем запускать как обычный js, и получаем ошибку:
```bash
➜ node typescript.ts 
./typescript.ts:1
(function (exports, require, module, __filename, __dirname) { let greeting: string;

SyntaxError: Unexpected token :
```

**Компилируем typescript в javascript**:
```bash
➜ tsc hello-world.ts
```

Запускаем скомпилированный код:
```bash
➜ node hello-world.js
Hello World!
```

## Базовые типы данных

Три базовых самых простых типа:

```typescript
let isAwesome: boolean; // Boolean
let cowCount: number; // Number
let cowColor: string; // String
```

С ними все просто - если объявили переменную как **String**
нельзя присвоить ей тип **Boolean** and so on. Ошибка произойдет
еще на этапе компиляции:
```typescript
let cowColor: string; // String
cowColor = true;
```
```shell
Type 'true' is not assignable to type 'string'
```
Это же относится ко всем типам кроме **Any** котороый собственно
для того чтобы можно было присвоить значение любого типа

### Array (Массив)

Массивы можно объявлять двумя способами
```typescript
let array1: string[]; // Array
let array2: Array<string>; // Array
```
Кроме скаляров можно указывать имя класа или интерфейса:
```typescript
let array3: Date[];
```
Кроме указаного типа, добавлять в массив элементы нельзя - это
приведет к ошибке компиляции.
Проверяем:
```typescript
array1 = ['hello', 'world'];
array1[3] = 55; // << Type '55' is not assignable to type 'string'
```

Для того чтобы добавить элементы разных типов нужно использовать
юнион оператор (**|**) или Кортеж

### Tuple (Кортеж)

По сути тот же массив строго типизированх значений и длинны:
```typescript
let tuple: [string, number]; // Tuple
tuple = [22, 'age']; // << Type '[number, string]' is not assignable to type '[string, number]'
tuple = ['age', 22, 53]; // << Types of property 'length' are incompatible. Type '3' is not assignable to type '2'
tuple = ['age', 22]; // OK
tuple[3] = 53; // OK
tuple[4] = true; // << Type 'true' is not assignable to type 'string | number'
```

После объявления Кортежа его можно создать строго по объявленному типу (4).
Но уже после создания, можно добавить больше элементов одного из множества
типов которые в нем уже хранятся. Другие типы не допускаются и приводят
к ошибке компиляции.

### Enum

```typescript
enum Animal {Cow, Elephant, Frog}
let animal: Animal;

animal = Animal.Frog;
console.log(animal); // 2
animal = 1;
console.log(Animal[animal]); // Elephant
```

Значения хранятся в виде целых чисел начиная с нуля. Можно начать отсчет 
с любого другого числа или присвоить произвольные индексы:
```typescript
enum Animal {Cow = 1, Elephant, Frog}
enum Animal {Cow = 1, Elephant = 3, Frog = 5}
```

### Any

Тип данных который своидт на нет всю суть статической типизации.
```typescript
let doubt: any;
doubt = 4;
doubt = true;
doubt = 'string';
```

Может быть полезно в некоторых случаях но при возможности стоит избегать.
Следует заметить что тип **Any** это не то же самое что и отсутствие типа.
В TypeScript есть система выведение типа (**Type Inference**). При объявлении переменной со значением, 
переменной будет назначен тип присваемого значения:
```typescript
let doubt = 'string';
doubt = true; // << Error: Type 'true' is not assignable to type 'string'
```
Для того чтобы сэмулировать поведение Any нужно 
(на самом деле нет, просто используйте тип **Any**) объявить переменную
без значения:
```typescript
let doubt;
doubt = 'string';
doubt = true;
```

### Void

Обычно используется как тайпхинт для возвращаемых значений
```typescript
function makeHappy(): void {
    console.log('Weee!');
}
```
Тип ```void``` соответствует двум значениям ```null``` и ```undefined``` 
При этом ```null``` и ```undefined``` имеют свои собственные типы, которые
являются подтипами всех остальных типов. Это значит что данный код будет валиден:
```typescript
let foo: number;
foo = 5;
foo = undefined;
foo = null;
```
До тех пока код компилируется без флага ```--strictNullChecks```. 
Этот влаг желательно (такое желательно как обязательно) использовать всегда.

### Never

Тип который стоит особнячком, и даже **Any** не может принимать тип **Never**
Я просто скопирую сюда примеры из доки:
```typescript
function error(message: string): never {
    throw new Error(message);
}
function fail(): never {
    return error('Something failed');
}
function infiniteLoop(): never {
    while (true) {
    }
}
```

### Type assertions

Можно дополнительно, строго контролировать типы там где компилятор
этого по каким либо причинам делать не может. Это полезно, как вообщем
и с описанием типов, для того чтобы отлавливать ошибки на этапе компиляции.

Есть два способа дополнительного ассерта типов, первый, часто встречается в Angular
и не поддерживается в JSX:
```typescript
let someValue: any = "this is a string";

let strLength: number = (<string>someValue).length;
```
Второй:
```typescript
let someValue: any = "this is a string";

let strLength: number = (someValue as string).length;
```

## Интерфейсы

Интерфейсы все еще не стали частью EcmaScript спецификации (пока что).
Но это не проблема для TypeScript где их можно использовать уже сейчас.

Интерфейсы в TypeScript это нечто что принято называть "Утиной типизацией",
или "Структурной типизацией" - объект который имплементирует интерфейс
не обязан где либо об этом заявлять, но обязатеьно должен имплементировать
заявленные свойства.

Есть два способа объявить интерфейс.
1. Тайпить параметр прямо во время его обьявлении:

```typescript
function haveFun(greeting: {label: string}) {
    console.log(greeting.label);
}

haveFun({label: 'Weeee!'});
```

2. Использовать отдельный интерфейс:

```typescript
interface Greeting {
    label: string;
}

function haveFun(greeting: Greeting) {
    console.log(greeting.label);
}

haveFun({label: 'Weeee!'});
```

### Необязательные параметры

Лично мне необязательные парамтеры в интерфейсе звучит как оксюморон.
Тем не менее они могут быть полезны, хоть и перед использованием нужно
будет проверить что значение там таки есть, или смирится с ```undefined```

```typescript
interface Greeting {
    label: string;
    description?: string;
}

function haveFun(greeting: Greeting) {
    console.log(greeting.label);
    if (greeting.description) {
        console.log(greeting.description);
    }
}
```

### "Лишние" свойства

Интерфейсы не только диктуют каким свойствам быть, но и каким не быть (всем остальным)

```typescript
interface User {
    username: string;
}
let user: User = {username: 'admin', firstName: 'John'}; // << Error: Object literal may only specify known properties, and 'firstName' does not exist in type 'User'
```

Точно так же для параметров:

```typescript
interface User {
    username: string;
}
function echoUser(user: {username: string}) {
    console.log(user.username);
}
echoUser({username: 'admin', firstName: 'John'}); // << Error: Object literal may only specify known properties, and 'firstName' does not exist in type '{ username: string; }'
```

Есть несколько способов избежать этой ошибки, 
кроме очевидного с необязательными параметрами.

1. При помощи Type assertion:

```typescript
interface User {
    username: string;
}
function echoUser(user: User) {
    console.log(user.username);
}
let user: User = {username: 'admin', firstName: 'John'} as User;
echoUser({username: 'admin', firstName: 'John'} as User);
```

2. Самый необъяснимый способ это передать переменную с объектом:

```typescript
interface User {
    username: string;
}

function echoUser(user: {username: string}) {
    console.log(user.username);
}

let user = {username: 'admin', firstName: 'John'};
echoUser(user); // OK
let myUser: User;
myUser = user; // OK
```

3. Или можно разрешить в интерфейсе любый другие свойства любого типа:

```typescript
interface User {
    username: string;
    [property: string]: any;
}
function echoUser(user: User) {
    console.log(user.username);
}
let user: User = {username: 'admin', firstName: 'John'};
echoUser({username: 'admin', firstName: 'John'});
```

### Иммутабельные свойства

Точно также как ```const``` для переменных ```readonly``` для свойств
не позволяет присвоить новое значение после инициализации.

```typescript
interface User {
    readonly firstName: string;
    readonly lastName: string;
}

let user: User = {firstName: 'John', lastName: 'Doe'};
user.firstName = 'Fred'; // << Error: Cannot assign to 'firstName' because it is a constant or a read-only property
```

### Иммутабельный массив

В TypeScript есть дополнительный тип ```ReadonlyArray```.
```typescript
let arr: ReadonlyArray<number> = [1, 2, 3, 4, 5];
arr.push(6); // << Error: Property 'push' does not exist on type 'ReadonlyArray<number>'
arr[0] = 0; // << Error: Index signature in type 'ReadonlyArray<number>' only permits reading
let arr2: Array<number>;
arr2 = arr; // << Error: Type 'ReadonlyArray<number>' is not assignable to type 'number[]'
```

Как видим присвоить иммутабельный массив в обычный нельзя, для этого
можно сконвертировать тип при помощи **Type assertions**

```typescript
let arr: ReadonlyArray<number> = [1, 2, 3, 4, 5];
let arr2: Array<number>;
arr2 = arr as number[]; // OK
```

### Функциональный тип

К функции, у которой есть параметры и возвращаемые значения можно также
типизировать покрыв интерфейсом. Есть два способа описания Function Type
которые идентичны по своей сути:

```typescript
interface HappyNumber {
    (min: number, max: number, nb: number): boolean;
}

type HappyNumber = (min: number, max: number, nb: number) => boolean;

let makeHappy: HappyNumber = function (min, max, number) {
    return Math.floor(Math.random() * (max - min + 1)) + min === number;
};
```

### Интерфейс для методов

Точно как для методов, так и для методов можно добавить интерфейсы:

```typescript
interface DotInterface {
    move(x: number, y: number): boolean;
}

class Dot implements DotInterface {
    move(x, y) {
        return true;
    }
}
```

### Наследование интерфейсов

Итерфейсы могут экстендится друг от друга, и от ... классов. WAT? Да, да
вам не послышалось. Именно от классов. Кроме всего прочего интерфейс 
отнаследованый от класса приобритает декларацию приватных методов и свойств:

```typescript
interface A {
    moveA(): void;
}

interface B extends A {
    moveB(): void;
}

class Dot implements B {
    internalProp: boolean;
    private internalMove() {}
    moveA() {}
    moveB() {}
}

interface C extends Dot {
    moveC(): void;
}

// Class 'DotC' incorrectly implements interface 'C'
// Property 'internalProp' is missing in type 'DotC'
class DotC implements C {
    moveA() {}
    moveB() {}
    moveC() {}
}
```

Добавим проперти:

```typescript
// Class 'DotC' incorrectly implements interface 'C'
// Property 'internalMove' is missing in type 'DotC'
class DotC implements C {
    internalProp: boolean;
    moveA() {}
    moveB() {}
    moveC() {}
}
```

И метод:
```typescript
// Class 'DotC' incorrectly implements interface 'C'
// Types have separate declarations of a private property 'internalMove'
class DotC implements C {
    internalProp: boolean;
    private internalMove() {}
    moveA() {}
    moveB() {}
    moveC() {}
}
```

OOps! А вот так, будет работать:
```typescript
// OK
class DotC extends Dot implements C {
    moveA() {}
    moveB() {}
    moveC() {}
}
```

Итого - наследовать интерфейс от класса будет полезным если нужно убедиться
что класс является наследником некого базового класса

## Классы

Собственно болльшая часть информации уже доступна на русском https://developer.mozilla.org/ru/docs/Web/JavaScript/Reference/Classes

Далее немного особенностей.

### Параметры-свойства

Если к аргументам конструктора прилагаются модификаторы видимости (private, 
public, protected, readonly) то аргументы автоматически становятся свойствами
класса:

```typescript
class Foo {
    constructor(readonly bar, private baz) {}
    public getBaz() {
        return this.baz;
    }
}

let foo = new Foo('5', 'baaaz');
console.log(foo.bar);
console.log(foo.getBaz());
```

### Геттеры, сеттеры

Синтаксис Javascript который позволяет добавить кастомную логику
на запись и чтение свойств

```typescript
class Person {
    private _name: string;
    private nameSet = false;

    set name(name: string) {
        if (this.nameSet) {
            console.error('You can set name only once');
            return;
        }
        this._name = name;
        this.nameSet = true;
    }

    get name(): string {
        console.log('Fetched name of Person');
        return this._name;
    }
}

let person = new Person();
person.name = 'John';
console.log(person.name);
person.name = 'John Doe';
console.log(person.name);
```

Два небольших но важных момента.

1. Гетеры и сеттеры работают только в ES5, поэтому при компиляции прийдется
это явно указать:
```shell
tsc typescript.ts -t ES5
```

2. Если не указан сетер то проперти будет считаться ```readonly``` 

### Модификаторы видимости

Мы уже видили private и public. Если еще static и protected.
Работают они точно так же как и в других ООП языках - protected 
доступен в наследниках, но не может быть вызван из инстанса класса.
Static в свою очередь может быть вызван прямо из класса но не из инстанса

Базовый пример:

```typescript
abstract class Human {}
class Person extends Human {
    private _name: string;
    private nameSet = false;

    static callMe() {
        console.log('static method called');
    }

    protected greeting() {
        console.log(`Hello ${this._name}`);
    }
}

class Employee extends Person {
}

Employee.callMe(); // static method called
```

Теперь попробуем немного поламать наш класс:

```typescript
class Employee extends Person {
    public myMethod() {
        console.log(this._name); // Property '_name' is private and only accessible within class 'Person'
    }
}
let h = new Human(); // Cannot create an instance of an abstract class
let e = new Employee();
e.greeting(); // Property 'greeting' is protected and only accessible within class 'Person' and its subclasses
```
