go-ansicon
==========

A Go library that converts ANSI escape sequences to Windows API calls.
Based on [ANSICON](https://github.com/adoxa/ansicon) by Jason Hood, [go-colortext](https://github.com/daviddengcn/go-colortext) by David Deng and [tcsh](http://www.tcsh.org/Welcome) 

This library follows the definition of the [XTerm control sequences](http://invisible-island.net/xterm/ctlseqs/ctlseqs.html).

Usage
-----

This library provides a `Convert` function, that reads data written to an [io Writer](http://golang.org/pkg/io/#Writer), looks for escape sequences and executes the required Windows API calls.
The rest of the data is written using the returned io Writer.

It can be used to make portable command line applications that require ANSI escape sequences and need to run on both Posix and Windows systems.

### Examples

- [Simple SSH client](https://github.com/Bitbored/go-ssh-client)

Current status
--------------
Only VT100 Mode control sequences are supported.

Currently implemented:
- Selecting graphic rendition
- Display reset and blank filling
- Setting cursor position

This makes e.g. a SSH session pretty usable.
