import java.util.Map;

public class Main {
    public static void main(String[] args) {
        Map<String, String> env = System.getenv();

        // for (Map.Entry<String, String> entry: env.entrySet()) {
        //     System.out.printf("%s = %s\n", entry.getKey(), entry.getValue());
        // }

        env.forEach((k, v) -> {
            System.out.printf("%s=%s\n", k, v);
        });
    }
}
