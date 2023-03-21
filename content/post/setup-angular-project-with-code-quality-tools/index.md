---
author: "Serhii Polishchuk"
title: "Setup Angular project with code quality tools"
date: 2018-02-20
tags: ["Javascript", "Angular"]
draft: false
---
<!--more-->
According to my experience right setup project at the start can boost development speed and help to make right choice about architecture of application. So after few weeks of playing around with **Angular** codebase and features I decide to start a new project from scratch. 
In this article I describe my way to setup **TravisCi** and **Scrutinizer** to run js unit tests with **Karma test runner** and 2e2 tests using **Protractor**. Also I add code test coverage calculation to make sure that significant part of code is tested appropriately.

### Travis-Ci

First of all, after ```ng new projectname``` lets connect to TravisCi service. Continuase integration service add ability to run tests again and again, after each commit, and if something will be wrong - Trevis tell us in PR status on GitHub or by email.

The official way for running unit tests in ng project is ```ng test``` - it will run browser, run tests in browser and watch changes for rerun tests. On Travis we have only console and don't have gui, but there is a common way to solve this issue by running python virtual X server. See [Using xvfb to Run Tests That Require a GUI](https://docs.travis-ci.com/user/gui-and-headless-browsers/#Using-xvfb-to-Run-Tests-That-Require-a-GUI) 

The second, more modern way, to run gui tests is a using headless mode for browser, that is supported in modern browsers. Firefox is a default browser in Ubuntu and TravisCi. However I get many issues while setup tests to run with Firefox, so the simplest and easiest way was just install a chrome browser.

The second kind of tests is a e2e tests. It's run in browser by it's nature. But this can be run in headless mode in two ways, same as karma tests.

So here are my configs:
Travis:
```yml
language: node_js
node_js:
  - "6.9"

addons:
  apt:
    sources:
      - google-chrome
    packages:
      - google-chrome-stable

branches:
  only:
    - master

before_script:
  - google-chrome --version
  - npm install -g @angular/cli
  - npm install webdriver-manager -g
  - webdriver-manager update
  - webdriver-manager start > ./webdriver.log 2>&1 &

script:
  - ng test --watch=false --browser=ChromeHeadlessNoSandbox --code-coverage=true
  - ng e2e

notifications:
  email:
    on_failure: change
    on_success: change

after_failure:
  - cat ./webdriver.log

after_script:
  - npm install ocular.js
  - $(npm bin)/ocular coverage/clover.xml
```
Two line at the end, in **after_script** sections needed to generate code coverage - about it below.

karma.conf.js:
```js
// Karma configuration file, see link for more information
// https://karma-runner.github.io/1.0/config/configuration-file.html

module.exports = function (config) {
  config.set({
    basePath: '',
    frameworks: ['jasmine', '@angular/cli'],
    plugins: [
      require('karma-jasmine'),
      require('karma-coverage'),
      require('karma-chrome-launcher'),
      require('karma-jasmine-html-reporter'),
      require('karma-coverage-istanbul-reporter'),
      require('@angular/cli/plugins/karma')
    ],
    client:{
      clearContext: false // leave Jasmine Spec Runner output visible in browser
    },
    preprocessors: {
        'src/**/*.js': ['coverage']
    },
    coverageReporter: {
        type: 'clover',
        subdir: './',
        file: 'clover.xml'
    },
    coverageIstanbulReporter: {
      reports: [ 'html', 'lcovonly' ],
      fixWebpackSourcePaths: true
    },
    angularCli: {
      environment: 'dev'
    },
    reporters: ['progress', 'kjhtml', 'coverage'],
    port: 9876,
    colors: true,
    logLevel: config.LOG_INFO,
    autoWatch: true,
    browsers: ['Chrome', 'ChromeHeadlessNoSandbox'],
    customLaunchers: {
        ChromeHeadlessNoSandbox: {
            base: 'ChromeHeadless',
            flags: ['--no-sandbox']
        }
    },
    singleRun: false
  });
};
```
protractor.conf.js:
```js
const { SpecReporter } = require('jasmine-spec-reporter');

exports.config = {
  allScriptsTimeout: 11000,
  specs: [
    './e2e/**/*.e2e-spec.ts'
  ],
  capabilities: {
    'browserName': 'chrome',
    chromeOptions: {
        args: [ "--headless", "--disable-gpu", "--window-size=800,600", "--no-sandbox" ]
    }
  },
  directConnect: true,
  baseUrl: 'http://localhost:4200/',
  framework: 'jasmine',
  jasmineNodeOpts: {
    showColors: true,
    defaultTimeoutInterval: 30000,
    print: function() {}
  },
  onPrepare() {
    require('ts-node').register({
      project: 'e2e/tsconfig.e2e.json'
    });
    jasmine.getEnv().addReporter(new SpecReporter({ spec: { displayStacktrace: true } }));
  }
};
```
