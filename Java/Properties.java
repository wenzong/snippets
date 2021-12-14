import java.io.Reader;
import java.io.StringReader;
import java.io.IOException;
import java.util.Properties;

public class Main {
    public static void main(String[] args) {
        Reader r = new StringReader("hello = world");
        Properties p = new Properties();
        try {
            p.load(r);
        } catch(IOException e) {
        }

        System.out.println(p.getProperty("hello"));
    }
}
