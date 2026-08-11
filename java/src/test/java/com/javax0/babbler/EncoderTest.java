package com.javax0.babbler;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertFalse;
import static org.junit.jupiter.api.Assertions.assertNotEquals;
import static org.junit.jupiter.api.Assertions.assertThrows;
import static org.junit.jupiter.api.Assertions.assertTrue;

import java.util.HashSet;
import java.util.Set;
import java.util.regex.Pattern;
import org.junit.jupiter.api.Test;

class EncoderTest {

    private static final Pattern ID_PATTERN = Pattern.compile("^[A-Z][a-z]{3} [A-Z][a-z]{3}$");
    private static final IdEncoder encoder = new Encoder();

    @Test
    void numberToIdZeroReturnsValidId() {
        assertTrue(ID_PATTERN.matcher(encoder.numberToId(0)).matches());
    }

    @Test
    void numberToIdOneDiffersFromZero() {
        assertNotEquals(encoder.numberToId(0), encoder.numberToId(1));
    }

    @Test
    void numberToId10000MatchesReference() {
        assertEquals("Talo Buno", encoder.numberToId(10000));
    }

    @Test
    void numberToId10001MatchesReference() {
        assertEquals("Patu Luta", encoder.numberToId(10001));
    }

    @Test
    void numberToIdMaxNumMatchesReference() {
        assertEquals("Dobu Zusa", encoder.numberToId((int) Constants.MAX_NUM));
    }

    @Test
    void numberToIdOutOfRangeThrows() {
        assertThrows(IllegalArgumentException.class, () -> encoder.numberToId((int) Constants.MAX_NUM + 1));
    }

    @Test
    void numberToIdNegativeThrows() {
        assertThrows(IllegalArgumentException.class, () -> encoder.numberToId(-1));
    }

    @Test
    void capitalizationFirstLetterOnly() {
        String[] parts = encoder.numberToId(42).split(" ");
        for (String word : parts) {
            assertEquals(Character.toUpperCase(word.charAt(0)) + word.substring(1).toLowerCase(), word);
        }
    }

    @Test
    void determinism() {
        assertEquals(encoder.numberToId(987654), encoder.numberToId(987654));
    }

    @Test
    void uniquenessAndBlacklistAvoidanceAcrossEntireValidRange() {
        Set<String> seen = new HashSet<>();
        for (int n = 0; n <= Constants.MAX_NUM; n++) {
            String id = encoder.numberToId(n);
            assertTrue(seen.add(id), id + " (from " + n + ") was already issued for a different number");

            String[] words = id.toLowerCase().split(" ");
            assertFalse(Constants.BLACKLIST.contains(words[0]), words[0] + " from " + n + " is blacklisted");
            assertFalse(Constants.BLACKLIST.contains(words[1]), words[1] + " from " + n + " is blacklisted");
        }
        assertEquals(Constants.MAX_NUM + 1, seen.size());
    }
}
