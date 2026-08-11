package com.javax0.babbler;

import java.util.ArrayList;
import java.util.List;
import java.util.Set;

/** Shared constants: CV syllable map, blacklist, and derived good-pair table. */
public final class Constants {

    private Constants() {
    }

    public static final String[] CV_MAP = {
        "ba", "bo", "bu", "da", "do", "du", "ga", "go", "gu", "ka",
        "ko", "ku", "la", "lo", "lu", "ma", "mo", "mu", "na", "no",
        "nu", "pa", "po", "pu", "ra", "ro", "ru", "sa", "so", "su",
        "ta", "to", "tu", "va", "zo", "zu"
    };

    public static final Set<String> BLACKLIST = Set.of(
        "bobo", "dodo", "dudu", "gaga", "kaka", "kuku",
        "lala", "mumu", "popo", "soso", "toto", "tutu",
        "nuna", "suna"
    );

    /**
     * All 1,296 two-syllable combinations, filtered down to the ones that
     * don't spell a blacklisted word. A name is built from two entries of
     * this list, so it is structurally impossible for {@link IdEncoder} to
     * produce a blacklisted word — no retry loop needed.
     *
     * <p>Computed once, in this static initializer, from {@link #CV_MAP}/
     * {@link #BLACKLIST} (not hardcoded) so those two stay the single
     * source of truth and this can never drift out of sync with them. The
     * cost (under a millisecond) is paid once per JVM, not per call.
     */
    public static final int[] GOOD_PAIRS = buildGoodPairs();

    /** Total number of names {@link IdEncoder} can produce: GOOD_PAIRS.length ^ 2. */
    public static final long GOOD_COUNT = (long) GOOD_PAIRS.length * GOOD_PAIRS.length;

    /** Highest valid input to {@link IdEncoder#numberToId}: GOOD_COUNT - 1. */
    public static final long MAX_NUM = GOOD_COUNT - 1;

    private static int[] buildGoodPairs() {
        List<Integer> pairs = new ArrayList<>();
        int n = CV_MAP.length;
        for (int i = 0; i < n * n; i++) {
            int d0 = i / n;
            int d1 = i % n;
            if (!BLACKLIST.contains(CV_MAP[d0] + CV_MAP[d1])) {
                pairs.add(i);
            }
        }
        int[] result = new int[pairs.size()];
        for (int i = 0; i < result.length; i++) {
            result[i] = pairs.get(i);
        }
        return result;
    }
}
