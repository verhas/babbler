'use strict';

/** 36 CV (consonant-vowel) syllables, index 0-35. CV_MAP[0] = 'ba', CV_MAP[35] = 'zu'. */
const CV_MAP = [
  'ba', 'bo', 'bu', 'da', 'do', 'du', 'ga', 'go', 'gu', 'ka',
  'ko', 'ku', 'la', 'lo', 'lu', 'ma', 'mo', 'mu', 'na', 'no',
  'nu', 'pa', 'po', 'pu', 'ra', 'ro', 'ru', 'sa', 'so', 'su',
  'ta', 'to', 'tu', 'va', 'zo', 'zu',
];

/** Two-syllable words excluded from encoder output. Lowercase. */
const BLACKLIST = new Set([
  'bobo', 'dodo', 'dudu', 'gaga', 'kaka', 'kuku',
  'lala', 'mumu', 'popo', 'soso', 'toto', 'tutu',
  'nuna', 'suna',
]);

/**
 * All 1,296 two-syllable combinations, filtered down to the ones that don't
 * spell a blacklisted word. A name is built from two entries of this list,
 * so it is structurally impossible for numberToId to produce a blacklisted
 * word — no retry loop needed.
 *
 * Computed once from CV_MAP/BLACKLIST (not hardcoded) so those two stay the
 * single source of truth and this can never drift out of sync with them.
 * The cost (~0.4ms) is paid once per process, not per call.
 */
function buildGoodPairs() {
  const pairs = [];
  for (let i = 0; i < CV_MAP.length * CV_MAP.length; i += 1) {
    const d0 = Math.floor(i / CV_MAP.length);
    const d1 = i % CV_MAP.length;
    if (!BLACKLIST.has(CV_MAP[d0] + CV_MAP[d1])) {
      pairs.push(i);
    }
  }
  return pairs;
}

const GOOD_PAIRS = buildGoodPairs();

/** Total number of names numberToId can produce: GOOD_PAIRS.length ^ 2. */
const GOOD_COUNT = GOOD_PAIRS.length * GOOD_PAIRS.length;

/** Highest valid input to numberToId: GOOD_COUNT - 1. */
const MAX_NUM = GOOD_COUNT - 1;

module.exports = { CV_MAP, BLACKLIST, GOOD_PAIRS, MAX_NUM };
