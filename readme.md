go-ansicon
==========

A Go library that converts ANSI escape sequences to Windows API calls.
Based on [ANSICON](https://github.com/adoxa/ansicon) by Jason Hood and [go-colortext](https://github.com/daviddengcn/go-colortext) by David Deng

Usage
-----

<<<<<<< HEAD
This library provides a `Convert` function, that reads data written to an [io Writer](http://golang.org/pkg/io/#Writer), looks for escape sequences and executes the required Windows API calls.
The rest of the data is written using the returned io Writer.
=======
This library provides the `Convert(input io.Writer) io.Writer` function that reads data written to an [io Writer](http://golang.org/pkg/io/#Writer), looks for ANSI escape sequences and executes the required Windows API calls.
The rest of the input is ~~written using the returned io Writer~~ currently just printed.
>>>>>>> e75f106294744adfddf154b107f75ddb4b47347e

This library can be used to make portable command line applications that require ANSI escape sequences and need to run on both Posix and Windows systems.

### Examples

- [Simple SSH client](https://github.com/Bitbored/go-ssh-client)

Current status
--------------
Only VT100 Mode control sequences are supported. This library follows the definition of the [XTerm control sequences]()
This is still very much a work in progress, most stuff is not implemented yet. Please don't use it (yet), unless to do some testing.
At the moment only changing colors works.

Github is just an easy way to transfer code to my Windows machine...
