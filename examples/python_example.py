"""Run from the repo root: python examples/python_example.py"""

import sys
from pathlib import Path

sys.path.insert(0, str(Path(__file__).resolve().parent.parent / "python"))

from babbler import number_to_id

# Typical usage: give a friendly display name to each row in an
# auto-increment sequence (e.g. a database primary key).
for user_id in range(5):
    print(f"user #{user_id} -> {number_to_id(user_id)}")
