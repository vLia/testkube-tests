package hello.gradle;

import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.*;

class LibraryTest {
    @Test
    void someLibraryMethodReturnsTrue() {
        // System.getProperties().stringPropertyNames().forEach(System.out::println);
        String env = System.getProperty("testkube");
        assertTrue(Boolean.parseBoolean(env), "testkube property should be true");
    }
}