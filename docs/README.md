# Melrose

[Home](index.html)
[Usage](cli.html)
[Language](dsl.html)
[DAW](daw.html)
[Install](install.html)

The basic musical objects in Melrose are:

- [note](dsl.html#note)
- [sequence](dsl.html#sequence)
- [chord](dsl.html#chord)
- [scale](dsl.html#scale)

Musical objects can be composed using:

- [repeat](dsl.html#repeat)
- [pitch](dsl.html#pitch)
- [reverse](dsl.html#reverse)
- [rotate](dsl.html#rotate)
- [join](dsl.html#join)
- [parallel](dsl.html#parallel)
- [serial](dsl.html#serial)
- [undynamic](dsl.html#undynamci)
- [indexMapper](dsl.html#indexmap)
- [interval](dsl.html#interval)

Musical objects can be played using:

- play
- [loop](dsl.html#loop)
- go

## Notations

### Note

| Notation | Alternative | Description
|----------|-------|-------------
| C4       | ¼C,C,c  | quarter C octave 4
| 2E5      | ½E5,½e5 | Halftone (2 x ¼) E octave 5
| 1C       |        | Full tone C octave 4
| F#       | F♯,f♯  | F sharp
| G_       | G♭    | G flat
| G.       | G.    | duration x 1.5
| =        | =     | quarter rest
| 2=       | ½=    | half rest
| 1=       | 1=    | full rest
| D+       | d+    | quarter D octave 4 MezzoForte
| 16E#.--  | 16e♯.-- | sixteenth E sharp duration x 1.5 Piano

### Note dynamics<a name="note-not"></a>

| Notation    | Description
|-------------|---
| \-\-\- |Pianissimo
| \-\-	|Piano
| \-	  |MezzoPiano
| 0   |Regular
| +	  |MezzoForte
| ++	|Forte
| +++ |Fortissimo

### Sequence<a name="sequence-not"></a>

| Notation    | Description
|-------------|---
| C D E F       | 4 quarter tones
| [C E] [d5 f5] | 2 doublets
| [1C 1E 1G]    | C Chord

### Chord<a name="chord-not"></a>

| Notation    | Description
|-------------|---
| C#5/m/2     | C sharp triad, Octave 5, Minor, 2nd inversion

### Scale<a name="scale-not"></a>

| Notation    | Description
|-------------|---
| C5          | C major scale, Octave 5
| E/m         | E natural minor scale, Octave 4