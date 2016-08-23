<?php


namespace App\Aspect;

use Go\Aop\Aspect;
use Go\Aop\Intercept\FieldAccess;
use Go\Aop\Intercept\MethodInvocation;
use Go\Lang\Annotation\After;
use Go\Lang\Annotation\Before;
use Go\Lang\Annotation\Around;
use Go\Lang\Annotation\Pointcut;

/**
 * Monitor aspect
 */
class MonitorAspect implements Aspect
{

    /**
     * Method that will be called around real method
     *
     * @param MethodInvocation $invocation Invocation
     * @Around("execution(public App\Hello->origin_*(*))")
     */
    public function aroundMethodExecution(MethodInvocation $invocation)
    {
        var_dump("Calling Before Interceptor for method\n");
        $result = $invocation->proceed();
        var_dump("Calling After Intercept for method\n");

        return $result;
    }
}
