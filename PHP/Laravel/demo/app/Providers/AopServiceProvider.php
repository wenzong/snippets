<?php

namespace App\Providers;

use App\Aspect\MonitorAspect;

use Illuminate\Contracts\Foundation\Application;
use Illuminate\Support\ServiceProvider;

class AopServiceProvider extends ServiceProvider
{
    /**
     * Register the application services.
     *
     * @return void
     */
    public function register()
    {
        $this->app->singleton(MonitorAspect::class, function (Application $app) {
            return new MonitorAspect();
        });

        $this->app->tag([MonitorAspect::class], 'goaop.aspect');
    }
}
