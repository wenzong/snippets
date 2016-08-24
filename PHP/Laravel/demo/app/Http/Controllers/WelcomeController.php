<?php

namespace App\Http\Controllers;

use App\TestInvoke;
use App\Hello;
use App\World;
use Cache;
use \ReflectionClass;
use guymers\proxy\ProxyFactory;

class WelcomeController extends Controller {

    public function home() {
        // Event / Listener
        // Cache::get('a');

        // GoAOP
        $h = new Hello();
        $h->origin_func();

        // Proxy
        $class = new ReflectionClass('App\World');

        $methodOverrides = [new TestInvoke()];

        $proxy = ProxyFactory::create($class, $methodOverrides);

        $proxy->origin_func();
    }
}
