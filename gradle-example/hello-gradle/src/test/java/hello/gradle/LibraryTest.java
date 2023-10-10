package hello.gradle;

import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.*;

class LibraryTest {
    @Test
    void someLibraryMethodReturnsTrue() {
        String users = System.getProperty("users");
        assertEquals("set_users", users, "unexpected value for users");
    }
}
