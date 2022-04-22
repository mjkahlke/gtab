# gtab

A guitar tab generator

## Overview

`gtab` is a command line utility that takes a chord name and generates tablature notation (tab) for playing the chord on a six-string guitar in standard tuning. It determines the root of the chord from the chord name, then performs a map lookup of the tonal signature of the chord suffix. For example, the Cmaj7 chord construction formula is a maj7 (at intervals 1,3,5,7) at the root note C. These intervals correspond to 0,4,7,11 half steps from the root note of the chord. It is then a simple matter of counting guitar frets and strings to generate a chord fingering or tab.

## Tablature notation

Guitar tab denotes fret positions on the guitar fretboard:
- An `x` denotes a string that is not played
- A `0` signifies an open string
- A `decimal digit` indicates what fret to finger

The relative position of x's and digits from left to right indicate which `string` to play from the low to the high E string. For example, a `Cmaj7` chord tab can be represented as either `x 3 2 0 0 0` or `x 3 5 4 5 3`. In both cases the low E string is not played. Note that the first form is `fingered` whereas the second form is `barred` at the root fret.

Tab notation does not tell you which fingers to place, only where to place them: that decision is left to you. A weak attempt is made at determining if a chord is playable by counting the number of non-open fret positions and comparing that to the number of fingers on your hand, but some chords may still require some amount of contortion to play; feel free to drop a few notes in these cases.

## Parameters

| parameter | default | description                      |
|-----------|---------|----------------------------------|
| `-chord`  |    C    | C major                          |
| `-list`   |  false  | list of supported chord suffixes |

## Sample output

The examples below were produced from the executable version produced by `go build`. Note how some tabs are repeated and some are not really playable (see Futures below).

### Cmaj7
    # gtab -chord Cmaj7
    Finger Cmaj7:  x 3 2 0 0 0
    Barred Cmaj7:  x 3 5 4 5 3
    Root-E Cmaj7:  8 10 9 9 8 8
    Root-A Cmaj7:  x 3 5 4 5 3
    Root-D Cmaj7:  x x 10 12 12 12
    Root-G Cmaj7:  x x x 5 5 7

### Abmaj7
    # gtab -chord Abmaj7
    Finger Abmaj7:  4 3 1 0 1 3
    Barred Abmaj7:  4 6 5 5 4 4
    Root-E Abmaj7:  4 6 5 5 4 4
    Root-A Abmaj7:  x 11 13 12 13 11
    Root-D Abmaj7:  x x 6 8 8 8
    Root-G Abmaj7:  x x x 1 1 3

### D
    # gtab -chord D
    Finger D:  x x 0 2 3 2
    Barred D:  x x 0 2 3 2
    Root-E D:  10 12 12 11 10 10
    Root-A D:  x 5 7 7 7 5
    Root-D D:  x x 0 2 3 2
    Root-G D:  x x x 7 7 10

## Futures

`gtab` does a fair job of generating tabs for common guitar chords played in most forms of Western music, but it does get a few wrong.

Some areas of improvement are:
- generate additional chord fingerings further up the neck beyond the first four frets
- do a better job of determining whether a chord is playable
- support for additional chord formulas e.g. chord inversions and added notes
- support for alternate guitar tunings
- look up a chord name from tab

## References

- [Chord Construction (Formulas)](https://tedgreene.com/images/lessons/fundamentals/ChordConstructionFormulas_1976-05-26.pdf)
- [Introduction to Intervals](https://musictheory.pugetsound.edu/mt21c/IntervalsIntroduction.html)
- [Table of Intervals in Music Theory](https://www.liveabout.com/table-of-intervals-2455915)
