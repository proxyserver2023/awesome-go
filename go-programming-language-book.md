# The Go Programming Language

## GO MOD
1. `go mod init`
2. go.mod => describes packages with versions
3. go.sum => tracks whether the versions changed or not with a HASH.
4. `go mod tidy` => tides up all the code
5. `go mod vendor` => create a vendor dir
6. `go mod why -m <package name>` => why we are depending on the lib
7. `go get ./...` => install every deps

## Preface
1. Approach of data abstraction and OOP
2. Native support of concurrency.
3. GCollection.
4. building infrastructure
   1. networked servers
   2. tools and systems for programmers

5. garbage collection, package system, first-class functions, lexical scope, system call interface, immutable strings in which text is generally encoded in UTF-8.
6. no implicit numeric conversions, no constructors or destructors, no operator overloading, no macros, no default param values, no inheritence, no generics, no exceptions, no function annotations and no thread-local storage.
7. API's for I/O, text processing, graphics, cryptography, networking and distributed applications.

## Tutorial
1. go natively handles unicode. process text in all the world's language.
2. package main is special. standalone executable program not a lib.
