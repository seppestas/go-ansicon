go-ansicon
==========

A Go library that converts ANSI escape sequences to Windows API calls.
Based on [ANSICON](https://github.com/adoxa/ansicon) by Jason Hood and [go-colortext](https://github.com/daviddengcn/go-colortext) by David Deng

Usage
-----

This library provides the `Convert(input io.Writer) io.Writer` function that reads data written to an [io Writer](http://golang.org/pkg/io/#Writer), looks for ANSI escape sequences and executes the required Windows API calls.
The rest of the input is ~~written using the returned io Writer~~ currently just printed.

This library can be used to make portable command line applications that require ANSI escape sequences and need to run on both Posix and Windows systems.

### Examples

- Simple SSH client: coming soon

Current status
--------------
This is still very much a work in progress, most stuff is not implemented yet. Please don't use it (yet), unless to do some testing.
At the moment only changing colors works.

Github is just an easy way to transfer code to my Windows machine...
