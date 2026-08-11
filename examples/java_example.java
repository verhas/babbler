// Build & run from the repo root:
//   javac -d /tmp/babbler-example-classes java/src/main/java/com/identifier/encoder/*.java examples/java_example.java
//   java -cp /tmp/babbler-example-classes JavaExample
import com.identifier.encoder.Encoder;
import com.identifier.encoder.IdEncoder;

class JavaExample {
    public static void main(String[] args) {
        IdEncoder encoder = new Encoder();
        // Typical usage: give a friendly display name to each row in an
        // auto-increment sequence (e.g. a database primary key).
        for (int userId = 0; userId < 5; userId++) {
            System.out.println("user #" + userId + " -> " + encoder.numberToId(userId));
        }
    }
}
