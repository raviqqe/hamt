# hamt.go

[![Circle CI](https://img.shields.io/circleci/project/github/raviqqe/hamt.go.svg?style=flat-square)](https://circleci.com/gh/raviqqe/hamt.go)
[![Go Report Card](https://goreportcard.com/badge/github.com/raviqqe/hamt.go?style=flat-square)](https://goreportcard.com/report/github.com/raviqqe/hamt.go)
[![License](https://img.shields.io/github/license/raviqqe/hamt.go.svg?style=flat-square)][unlicense]

Immutable HAMT implementation in Go.

[Hash-Array Mapped Trie (HAMT)](https://en.wikipedia.org/wiki/Hash_array_mapped_trie)
is a data structure popular as associative arrays (a.k.a. maps or dictionaries)
or sets.
Its immutable variant is adopted widely by functional programming languages
like Scala and Clojure to implement immutable but memory-efficient associative
arrays.

## Installation

```
go get github.com/raviqqe/hamt.go
```

## Documentation

[GoDoc](https://godoc.org/github.com/raviqqe/hamt.go)

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
