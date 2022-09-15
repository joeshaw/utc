# utc

A trivial tool to help me convert times to UTC and back.

## Installation

    go install github.com/joeshaw/utc@latest

## Usage

Times without a time zone are assumed to be in local time.  The time zone abbreviations are not DST-aware.  They assume the location on the current date, meaning that e.g. `EST` and `EDT` are equivalent and map to `America/New_York`.

Convert from local time to UTC:

    $ utc 12:30
    16:30 UTC

Convert from a specific US time zone to UTC:

    $ utc 12:30 pdt
    19:30 UTC

Convert from a UTC time to local time:

    $ utc -l 16:30 utc
    12:30 EDT

Convert from a specific US time zone to local time:

    $ utc -l 12:30 pdt
    15:30 EDT
