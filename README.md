# gtab

A guitar tab generator

## Overview

`gtab` is a command line utility that takes a chord name and generates tablature notation (tab) for playing the chord on a six-string guitar in standard tuning. It determines the root of the chord from the chord name, then performs a map lookup of the tonal signature of the chord suffix. For example, the Cmaj7 chord construction formula is a maj7 (at intervals 1,3,5,7) at the root note C. These intervals correspond to 0,4,7,11 half steps from the root note of the chord. It is then a simple matter of counting guitar frets and strings to generate a chord fingering or tab.

## Parameters

| parameter | default | description  |
|-----------|---------|--------------|
| `-chord`  |    C    | C major      |
| `-debug`  |  false  | debug output |

## Tablature notation

Guitar tab denotes fret positions on the guitar fretboard. An `x` denotes a string that is not played. A `decimal digit` indicates what fret to finger. A `0` signifies an open string. The relative position of digits from left to right indicate which `string` to play from the low to the high `E` string. For example, a Cmaj7 chord tab can be represented as either `x 3 2 0 0 0` or `x 3 5 4 5 3`. In both cases the low E string is not played. Note that the first form is `fingered` whereas the second form is `barred` at the root fret.

Tab notation does not tell you which finger to place, only where to place them. That decision is left up to you. A weak attempt is made at determining if a chord is playable by counting the number of non-open fret positions and comparing that to the number of fingers on your hand, but some chords could still require some amount of contortion to play, which brings us to the next section.

## Futures

This utility does a fair job of generating accurate playable tab for most guitar chords one would reasonably play in most forms of Western music, but it does get a few wrong.

Some areas of improvement are:
- generate chord fingerings further up the neck beyond the first four frets
- support for alternate guitar tunings
- look up a chord name from tab

## References

- [Chord Construction (Formulas)](https://tedgreene.com/images/lessons/fundamentals/ChordConstructionFormulas_1976-05-26.pdf)
- [Introduction to Intervals](https://musictheory.pugetsound.edu/mt21c/IntervalsIntroduction.html)
- [Table of Intervals in Music Theory](https://www.liveabout.com/table-of-intervals-2455915)
