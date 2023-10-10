package hello.gradle;

import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.*;

class LibraryTest {
    @Test void someLibraryMethodReturnsTrue() {
        String env = System.getenv("TESTKUBE_GRADLE");
        assertTrue(Boolean.parseBoolean(env), "TESTKUBE_GRADLE env should be true");
        String api = System.getenv("API_KEY");
        int users = Integer.getInteger("users", 1); 
        assertTrue(api == "api_key_value", "API_KEY should have the value api_key_value");
        assertTrue(users == 500, "users should be set to 500");
    }
}
