import re

import pytest

from babbler import BLACKLIST, MAX_NUM, number_to_id

ID_PATTERN = re.compile(r"^[A-Z][a-z]{3} [A-Z][a-z]{3}$")


def test_number_to_id_zero_returns_valid_id():
    assert ID_PATTERN.match(number_to_id(0))


def test_number_to_id_one_differs_from_zero():
    assert number_to_id(1) != number_to_id(0)


def test_number_to_id_10000_matches_reference():
    assert number_to_id(10000) == "Talo Buno"


def test_number_to_id_10001_matches_reference():
    assert number_to_id(10001) == "Patu Luta"


def test_number_to_id_max_num_matches_reference():
    assert number_to_id(MAX_NUM) == "Dobu Zusa"


def test_number_to_id_out_of_range_raises():
    with pytest.raises(ValueError):
        number_to_id(MAX_NUM + 1)


def test_number_to_id_negative_raises():
    with pytest.raises(ValueError):
        number_to_id(-1)


def test_number_to_id_non_integer_raises():
    with pytest.raises(ValueError):
        number_to_id(1.5)


def test_capitalization_first_letter_only():
    word1, word2 = number_to_id(42).split(" ")
    assert word1 == word1[0].upper() + word1[1:].lower()
    assert word2 == word2[0].upper() + word2[1:].lower()


def test_determinism():
    assert number_to_id(987654) == number_to_id(987654)


def test_uniqueness_and_blacklist_avoidance_across_entire_valid_range():
    seen = set()
    for n in range(MAX_NUM + 1):
        id_ = number_to_id(n)
        assert id_ not in seen, f"{id_} (from {n}) was already issued for a different number"
        seen.add(id_)

        word1, word2 = id_.lower().split(" ")
        assert word1 not in BLACKLIST, f"{word1} from {n} is blacklisted"
        assert word2 not in BLACKLIST, f"{word2} from {n} is blacklisted"
    assert len(seen) == MAX_NUM + 1
