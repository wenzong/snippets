<?php

namespace App\Listeners;

use Illuminate\Cache\Events\CacheHit;

class EventListener
{
    /**
     * Create the event listener.
     *
     * @return void
     */
    public function __construct()
    {
        //
    }

    /**
     * Handle the event.
     *
     * @param  SomeEvent  $event
     * @return void
     */
    public function handle(CacheHit $event)
    {
        // Send hit metric
    }
}
