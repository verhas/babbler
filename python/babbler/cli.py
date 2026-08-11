"""Command-line entry point for babbler.

Usage:
    python -m babbler encode 10000
"""

import argparse
import sys

from .encoder import number_to_id


def main(argv=None):
    """Run the babbler CLI.

    Args:
        argv: Argument list to parse, defaults to ``sys.argv[1:]``.

    Returns:
        Process exit code.
    """
    parser = argparse.ArgumentParser(prog="babbler")
    subparsers = parser.add_subparsers(dest="command", required=True)

    encode_parser = subparsers.add_parser("encode", help="Convert a number into an id")
    encode_parser.add_argument("num", type=int)

    args = parser.parse_args(argv)

    if args.command == "encode":
        print(number_to_id(args.num))
    return 0


if __name__ == "__main__":
    sys.exit(main())
