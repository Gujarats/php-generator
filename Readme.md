# Generator for php code

This is cli for generating php code setter and getter.

## Installation 

First Install GO to your environment

```shell
$ git clone git@github.com:Gujarats/php-generator.git
$ cd php-generator
$ go build 
$ go install
```


## How to use

For example we have file called TestClass.php to generate getter and setter simply follow this :  

```php

class TestClass{
    private $myVar;
    private $newVaria;
    private $newVar;
    private $TestVar;

}
```

Use this commadn to generate the code : 

```shell
$ php-generator generator TestClass.php

```

And the result : 

```php

class TestClass{
    private $myVar;
    private $newVaria;
    private $newVar;
    private $TestVar;

    public function __construct($myVar,$newVaria,$newVar,$TestVar){
        $this->myVar=$myVar;
        $this->newVaria=$newVaria;
        $this->newVar=$newVar;
        $this->TestVar=$TestVar;
    }
    public function getMyVar(){
        return $this->myVar;
    }
    public function setMyVar($myVar){
        $this->myVar=$myVar;
    }
    public function getNewVaria(){
        return $this->newVaria;
    }
    public function setNewVaria($newVaria){
        $this->newVaria=$newVaria;
    }
    public function getNewVar(){
        return $this->newVar;
    }
    public function setNewVar($newVar){
        $this->newVar=$newVar;
    }
    public function getTestVar(){
        return $this->TestVar;
    }
    public function setTestVar($TestVar){
        $this->TestVar=$TestVar;
    }
}
```
