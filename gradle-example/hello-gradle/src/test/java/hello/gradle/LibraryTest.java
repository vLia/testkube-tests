package hello.gradle;

import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.*;

class LibraryTest {
    @Test void someLibraryMethodReturnsTrue() {
        String users = System.getProperty("users","/tmp/log");
        assertTrue(users == "set_users", users);
    }
}
