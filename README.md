# gtab

A guitar tab generator from common chord names

## Overview

A command line utility that takes a chord name and generates tablature notation (tab) for playing the chord on a standard tuning guitar. It determines the root of the chord from the chord name, and performs a map lookup up the tonal signature of the chord from the suffix name. For example, the Cmaj7 chord construction formula is a maj7 (at intervals 1,3,5,7) at the root note C. These intervals correspond to 0,4,7,11 half steps from the root note of the chord. It is then a simple matter of counting guitar frets and strings from the root note to generate a chord fingering or tab for the chord.

## Parameters

| parameter | default | description  |
|-----------|---------|--------------|
| `-chord`  |    C    | C major      |
| `-debug`  |  false  | debug output |

## Tablature notation

Guitar tab directly relates to fret positions on the guitar fretboard. An `x` denotes a string that is not played. A `decimal digit` indicates what fret to finger; a `0` signifies an open string. The relative position of digits from left to right indicate which `string` to finger a given fret. For example, a Cmaj7 chord tab can be represented as either `x 3 2 0 0 0` or `x 3 5 4 5 3`. Note that the first form is `fingered` whereas the second form is `barred` meaning the forefinger is held across the entire fret on the fretboard while the remaining fingers finger notes on frets higher up the fretboard.

Note that tab notation does not tell you which finger to place where. That is left up to you. The gtab application makes a weak attempt at determing if a chord is playable by counting the number of non-open fret positions and comparing that to the number of fingers on your hand, but some chords could still require some contortion to play, which brings us to the next section.

## Futures

This utility does a fair job of generating accurate playable tab for most guitar chords one would reasonably play, but it does get a few wrong.

Some areas of improvement are:
- generate chord fingerings further up the neck beyond the first four frets
- support for alternate guitar tunings
- a reverse lookup, given the tab determine what the chord is

## References

- [Chord Construction (Formulas)](https://tedgreene.com/images/lessons/fundamentals/ChordConstructionFormulas_1976-05-26.pdf)
