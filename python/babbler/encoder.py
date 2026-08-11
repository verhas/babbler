"""Encoder: convert an integer counter into a pronounceable identifier."""

from .constants import CV_MAP, GOOD_COUNT, GOOD_PAIRS, MAX_NUM

_PAIR_COUNT = len(GOOD_PAIRS)

# Fixed affine-permutation constants. A modular multiplication x -> A*x + C
# (mod GOOD_COUNT) is a bijection on Z_GOOD_COUNT whenever gcd(A, GOOD_COUNT)
# is 1 -- that's the only requirement. These two were picked arbitrarily once
# and are now fixed forever so output is stable across runs and identical
# across every language implementation.
_A = 1256797
_C = 1443960


def _pair_to_word(pair_index):
    d0, d1 = divmod(pair_index, len(CV_MAP))
    return CV_MAP[d0] + CV_MAP[d1]


def number_to_id(num):
    """Convert an integer counter into a pronounceable two-word identifier.

    Guaranteed unique per ``num`` and never a blacklisted word, both by
    construction (see :mod:`constants`), not by retrying.

    Args:
        num: Integer in ``[0, MAX_NUM]``.

    Returns:
        A string like ``"Talo Buno"``.

    Raises:
        ValueError: If ``num`` is out of range.
    """
    if not isinstance(num, int) or isinstance(num, bool) or num < 0 or num > MAX_NUM:
        raise ValueError(f"num must be an integer in [0, {MAX_NUM}], got {num!r}")

    scrambled = (_A * num + _C) % GOOD_COUNT
    outer, inner = divmod(scrambled, _PAIR_COUNT)

    word1 = _pair_to_word(GOOD_PAIRS[outer])
    word2 = _pair_to_word(GOOD_PAIRS[inner])

    return f"{word1.capitalize()} {word2.capitalize()}"
