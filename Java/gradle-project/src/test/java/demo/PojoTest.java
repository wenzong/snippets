package demo;

import junit.framework.TestCase;
import org.junit.Assert.*;

public class PojoTest extends TestCase {

    public void testSetterGetter() {
        String s = "hello world";

        Pojo o = new Pojo();
        o.setStr(s);
        assertEquals(s, o.getStr());
    }
}
