# gtab

A guitar tab generator

## Overview

**gtab** is a utility that takes a chord name and generates tablature notation (tab) for playing it on a six-string guitar in standard tuning. It can be invoked either as a command line application or with REST API calls.

## How it works

The root note is determined from the chord name, then the suffix is used as a map key to lookup the tonal signature of the chord. For example, Cmaj7 is a maj7 chord with a root of C. The chord construction formula for maj7 is composed of a unison, major 3rd, perfect 5th and dominant 7th note or interval. This corresponds to 0, 4, 7 and 11 half steps (frets) from the root. It is then a simple matter of counting guitar frets and strings to generate a tab.

See [ARCHITECTURE.md](ARCHITECTURE.md) for instructions on how to build and run **gtab**.

## Tablature notation

Guitar tab denotes fret positions on the guitar fretboard:
- An `x` denotes a string that is not played
- A `0` signifies an open string
- A `number` indicates what fret to finger

The relative position of x's and numbers from left to right indicate which string to play from the low E to the high E string. Chords can be played in multiple locations along the fretboard. For example, common tabs for Cmaj7 are `x 3 2 0 0 0` and `x 3 5 4 5 3`. In both cases the low E string is not played. Note that the first form is _fingered_ whereas the second form is _barred_ at the third fret.

Tab notation does not tell you which fingers to place, only where to place them: that decision is left to you. In the Cmaj7 example above, two fingers are used to play the fingered form. The chord can be played either with the middle and ring fingers, or the forefinger and middle finger. The second form is barred with the forefinger at the third fret.

A weak attempt is made at determining if a chord is playable by counting the number of non-open fret positions and comparing that to the number of fingers on your hand, but some chords may still prove too difficult to play; feel free to drop a note or two in these cases. A better algorithm may be to count the number of fret transitions. Another consideration is the use of the thumb to finger frets on the low E string.

## Futures

**gtab** does a fair job of generating tabs for common guitar chords played in most forms of Western music, but it does get a few wrong.

Some areas of improvement are:
- generate tabs that aren't anchored at the root note fret
- do a better job of determining whether a chord is playable
- support for additional chord formulas e.g. chord inversions and added notes
- support for alternate guitar tunings
- reverse lookup: find a chord name from tab
- swagger for REST API

## References

- [Chord Construction (Formulas)](https://tedgreene.com/images/lessons/fundamentals/ChordConstructionFormulas_1976-05-26.pdf)
- [Introduction to Intervals](https://musictheory.pugetsound.edu/mt21c/IntervalsIntroduction.html)
- [Table of Intervals in Music Theory](https://www.liveabout.com/table-of-intervals-2455915)
