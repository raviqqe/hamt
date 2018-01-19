# hamt

[![Circle CI](https://img.shields.io/circleci/project/github/raviqqe/hamt.svg?style=flat-square)](https://circleci.com/gh/raviqqe/hamt)
[![Codecov](https://img.shields.io/codecov/c/github/raviqqe/hamt.svg?style=flat-square)](https://codecov.io/gh/raviqqe/hamt)
[![Go Report Card](https://goreportcard.com/badge/github.com/raviqqe/hamt?style=flat-square)](https://goreportcard.com/report/github.com/raviqqe/hamt)
[![License](https://img.shields.io/github/license/raviqqe/hamt.svg?style=flat-square)][unlicense]

Immutable and Memory Efficient Maps and Sets in Go.

This package `hamt` provides immutable collection types of maps (associative arrays)
and sets implemented as Hash-Array Mapped Tries (HAMTs).
All operations of the collections, such as insert and delete, are immutable and
create new ones keeping original ones unmodified.

[Hash-Array Mapped Trie (HAMT)](https://en.wikipedia.org/wiki/Hash_array_mapped_trie)
is a data structure popular as a map (a.k.a. associative array or dictionary)
or set.
Its immutable variant is adopted widely by functional programming languages
like Scala and Clojure to implement immutable and memory-efficient associative
arrays and sets.

## Installation

```
go get github.com/raviqqe/hamt
```

## Documentation

[GoDoc](https://godoc.org/github.com/raviqqe/hamt)

## Technical notes

The implementation canonicalizes tree structures of HAMTs by eliminating
intermediate nodes during delete operations as described
in [the CHAMP paper][champ].

## References

- [Ideal Hash Trees](https://infoscience.epfl.ch/record/64398/files/idealhashtrees.pdf)
- [Optimizing Hash-Array Mapped Tries for Fast and Lean Immutable JVM Collections][champ]

## License

[The Unlicense][unlicense]

[champ]: https://michael.steindorfer.name/publications/oopsla15.pdf

[unlicense]: https://unlicense.org/
