<?php
namespace App;

use \ReflectionMethod;
use guymers\proxy\MethodHook;

class TestInvoke implements MethodHook {

    public function supports(ReflectionMethod $method) {
        return true;
    }

    public function invoke($proxy, ReflectionMethod $method, array $args) {
        var_dump('Before invoke');

        $returnValue = $method->invokeArgs($proxy, $args);

        return $returnValue;
    }
}
