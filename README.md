# gtab

A guitar tab generator

## Overview

**gtab** is a command line utility that takes a chord name and generates tab (tablature notation) for playing the chord on a six-string guitar in standard tuning.

## How it works

The root of the chord is determined from the chord name, then the chord suffix is used as a map key to lookup the tonal signature of the chord. For example, Cmaj7 is a maj7 chord with a root of C. The chord construction formula for maj7 is composed of a unison, major 3rd, perfect 5th and dominant 7th note or interval. This corresponds to 0, 4, 7 and 11 half steps (frets) from the root. It is then a simple matter of counting guitar frets and strings to generate a tab.

## Tablature notation

Guitar tab denotes fret positions on the guitar fretboard:
- An `x` denotes a string that is not played
- A `0` signifies an open string
- A `number` indicates what fret to finger

The relative position of x's and numbers from left to right indicate which string to play from the low E to the high E string. For example, a Cmaj7 chord tab could be represented as either `x 3 2 0 0 0` or `x 3 5 4 5 3`. In both cases the low E string is not played. Note that the first form is _fingered_ whereas the second form is _barred_ at the third fret.

Tab notation does not tell you which fingers to place, only where to place them: that decision is left to you. In the previous example, two fingers are used to play the fingered form. This chord could be played either with the middle and ring fingers, or the forefinger and middle finger. The second form is barred with the forefinger at the third fret.

A weak attempt is made at determining if a chord is playable by counting the number of non-open fret positions and comparing that to the number of fingers on your hand, but some chords may still prove too difficult to play; feel free to drop a few notes in these cases. A better algorithm may be to count the number of fret transitions. Another consideration is the use of the thumb to finger frets on the low E string.

## Parameters

| parameter | default | description                      |
|:----------|:-------:|:---------------------------------|
| `-chord`  | C       | C major                          |
| `-list`   | false   | list of supported chord suffixes |

## Sample output

Several sample runs are shown below. Note how some tabs are not easily playable (see Futures below).

### Help
    # gtab -help
    Usage of gtab:
      -chord string
            set to chord name, b=flat, #=sharp (default "C")
      -list
            list supported chord formulas

### Cmaj7
    # gtab -chord Cmaj7
    x 3 2 0 0 0
    x 3 5 4 5 3
    x x 10 12 12 12
    x x x 5 5 7

### B7
    # gtab -chord B7
    x 2 1 2 0 2
    x 2 4 2 4 2
    x x 9 11 10 11
    x x x 4 4 5

### Abmaj7
    # gtab -chord Abmaj7
    4 3 1 0 1 3
    4 6 5 5 4 4
    x 11 13 12 13 11
    x x 6 8 8 8
    x x x 1 1 3

### D
    # gtab -chord D
    x 5 7 7 7 5
    x x 0 2 3 2
    x x x 7 7 10
    
### F#sus4
    # gtab -chord F#sus4
    2 2 4 4 2 2
    2 2 4 x 4 1
    x 9 9 11 12 9
    x x 4 4 7 7
    x x x 11 12 14

## Futures

**gtab** does a fair job of generating tabs for common guitar chords played in most forms of Western music, but it does get a few wrong.

Some areas of improvement are:
- generate tabs that aren't based off the root note/fret location
- do a better job of determining whether a chord is playable
- support for additional chord formulas e.g. chord inversions and added notes
- support for alternate guitar tunings
- reverse lookup: find a chord name from tab

## References

- [Chord Construction (Formulas)](https://tedgreene.com/images/lessons/fundamentals/ChordConstructionFormulas_1976-05-26.pdf)
- [Introduction to Intervals](https://musictheory.pugetsound.edu/mt21c/IntervalsIntroduction.html)
- [Table of Intervals in Music Theory](https://www.liveabout.com/table-of-intervals-2455915)
