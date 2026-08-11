'use strict';

const { CV_MAP, GOOD_PAIRS, MAX_NUM } = require('./constants');

const PAIR_COUNT = GOOD_PAIRS.length;
const GOOD_COUNT = MAX_NUM + 1;

// Fixed affine-permutation constants. A modular multiplication x -> A*x + C
// (mod GOOD_COUNT) is a bijection on Z_GOOD_COUNT whenever gcd(A, GOOD_COUNT)
// is 1 — that's the only requirement. These two were picked arbitrarily once
// and are now fixed forever so output is stable across runs and identical
// across every language implementation.
const A = 1256797;
const C = 1443960;

function capitalize(word) {
  return word.charAt(0).toUpperCase() + word.slice(1);
}

function pairToWord(pairIndex) {
  const d0 = Math.floor(pairIndex / CV_MAP.length);
  const d1 = pairIndex % CV_MAP.length;
  return CV_MAP[d0] + CV_MAP[d1];
}

/**
 * Converts an integer counter into a pronounceable, human-friendly name.
 * Guaranteed unique per num and never a blacklisted word, both by
 * construction (see constants.js), not by retrying.
 *
 * @param {number} num Integer in [0, MAX_NUM].
 * @returns {string} e.g. "Talo Buno"
 */
function numberToId(num) {
  if (!Number.isInteger(num) || num < 0 || num > MAX_NUM) {
    throw new RangeError(`num must be an integer in [0, ${MAX_NUM}], got ${num}`);
  }

  const scrambled = (A * num + C) % GOOD_COUNT;
  const outer = Math.floor(scrambled / PAIR_COUNT);
  const inner = scrambled % PAIR_COUNT;

  const word1 = pairToWord(GOOD_PAIRS[outer]);
  const word2 = pairToWord(GOOD_PAIRS[inner]);

  return `${capitalize(word1)} ${capitalize(word2)}`;
}

module.exports = { numberToId };
