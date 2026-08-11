package com.javax0.babbler;

/**
 * Converts an integer counter into a pronounceable, human-friendly
 * identifier. Guaranteed unique per {@code num} and never a blacklisted
 * word, both by construction (see {@link Constants#GOOD_PAIRS}), not by
 * retrying. One-way: there is no decoder (see the top-level README's
 * "Non-goals").
 */
public interface IdEncoder {

    // Fixed affine-permutation constants. A modular multiplication
    // x -> A*x + C (mod GOOD_COUNT) is a bijection on Z_GOOD_COUNT whenever
    // gcd(A, GOOD_COUNT) is 1 -- that's the only requirement. These two were
    // picked arbitrarily once and are now fixed forever so output is stable
    // across runs and identical across every language implementation.
    long A = 1256797L;
    long C = 1443960L;

    /**
     * @param num integer in {@code [0, Constants.MAX_NUM]}
     * @return a string like {@code "Talo Buno"}
     */
    default String numberToId(int num) {
        if (num < 0 || num > Constants.MAX_NUM) {
            throw new IllegalArgumentException(
                "num must be in [0, " + Constants.MAX_NUM + "], got " + num);
        }

        int pairCount = Constants.GOOD_PAIRS.length;
        long scrambled = (A * num + C) % Constants.GOOD_COUNT;
        int outer = (int) (scrambled / pairCount);
        int inner = (int) (scrambled % pairCount);

        String word1 = pairToWord(Constants.GOOD_PAIRS[outer]);
        String word2 = pairToWord(Constants.GOOD_PAIRS[inner]);

        return capitalize(word1) + " " + capitalize(word2);
    }

    private static String pairToWord(int pairIndex) {
        int n = Constants.CV_MAP.length;
        return Constants.CV_MAP[pairIndex / n] + Constants.CV_MAP[pairIndex % n];
    }

    private static String capitalize(String word) {
        return Character.toUpperCase(word.charAt(0)) + word.substring(1);
    }
}
