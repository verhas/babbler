"""Shared constants for babbler: CV syllable map, blacklist, and derived
good-pair table."""

CV_MAP = [
    'ba', 'bo', 'bu', 'da', 'do', 'du', 'ga', 'go', 'gu', 'ka',
    'ko', 'ku', 'la', 'lo', 'lu', 'ma', 'mo', 'mu', 'na', 'no',
    'nu', 'pa', 'po', 'pu', 'ra', 'ro', 'ru', 'sa', 'so', 'su',
    'ta', 'to', 'tu', 'va', 'zo', 'zu',
]

BLACKLIST = frozenset({
    'bobo', 'dodo', 'dudu', 'gaga', 'kaka', 'kuku',
    'lala', 'mumu', 'popo', 'soso', 'toto', 'tutu',
    'nuna', 'suna',
})


def _build_good_pairs():
    """All 1,296 two-syllable combinations, filtered down to the ones that
    don't spell a blacklisted word.

    A name is built from two entries of this list, so it is structurally
    impossible for number_to_id to produce a blacklisted word -- no retry
    loop needed.

    Computed once, at import time, from CV_MAP/BLACKLIST (not hardcoded) so
    those two stay the single source of truth and this can never drift out
    of sync with them. The cost (under a millisecond) is paid once per
    process, not per call.
    """
    n = len(CV_MAP)
    pairs = []
    for i in range(n * n):
        d0, d1 = divmod(i, n)
        if CV_MAP[d0] + CV_MAP[d1] not in BLACKLIST:
            pairs.append(i)
    return pairs


GOOD_PAIRS = _build_good_pairs()

#: Total number of names number_to_id can produce: len(GOOD_PAIRS) ** 2.
GOOD_COUNT = len(GOOD_PAIRS) ** 2

#: Highest valid input to number_to_id: GOOD_COUNT - 1.
MAX_NUM = GOOD_COUNT - 1
