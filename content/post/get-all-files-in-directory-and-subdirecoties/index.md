---
author: "Serhii Polishchuk"
title: "Get all files in directory and subdirecoties"
date: 2015-12-03
tags: ["PHP"]
draft: false
---
<!--more-->
This code:
```php
<?php
/**
* @param $dir
 * @param array $files
 * @return array
 */
protected function getFiles($dir, array &$files = [])
{
    foreach (scandir($dir) as $file) {
        if (in_array($file, [".", ".."])) {
            continue;
        }
        $fullPath = $dir.'/'.$file;
        if (is_dir($fullPath)) {
            $this->getFiles($fullPath, $files);
        } else {
            $files[] = $fullPath;
        }
    }
     return $files;
}
```
This test:
```php
<?php
public function testGetFiles()
{
    $method = self::getMethod('getFiles');
    $command = new ImportLogsCommand();
    $files = $method->invokeArgs($command, [realpath(__DIR__.'/../..')]);
    $this->assertContains(realpath(__DIR__.'/ImportLogsCommandTest.php'), $files);
    $this->assertContains(realpath(__DIR__.'/../../AppBundle.php'), $files);
}
/**
 * @param string $name Method name
 * @return \ReflectionMethod
 */
protected static function getMethod($name) {
    $class = new \ReflectionClass('AppBundle\Command\ImportLogsCommand');
    $method = $class->getMethod($name);
    $method->setAccessible(true);
    return $method;
}
```
