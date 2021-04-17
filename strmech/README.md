# *strops* Version 3.0.0

[![GoDoc](https://godoc.org/github.com/MikeAustin71/stringopsgo/strops/v3?status.svg)](https://godoc.org/github.com/MikeAustin71/stringopsgo/strops/v3)

*strops* is a software library or in *Go* parlance, a software *package*.

Package *strops*, or **string operations**, is a collection of string
management utilities written in the Go Programming Language. 

This package is written in the *Go* programming language, a.k.a. 'golang'.

Type *StrOps* provides simple string management routines which perform operations
like string centering, justification trimming and character manipulation.

Version 3.0.0 introduces several new important features which come at the price
of breaking changes. 

- Developed with Go Version 1.15.6.
- Thread Safety protocols added for support of parallel processing.
- Enhanced error management capabilities added. 
  - This version now supports documentation of function chains in error messages.

- This version continues support for Go modules

  [Click To View Source Documentation](http://godoc.org/github.com/MikeAustin71/stringopsgo/strops/v3)    

# Table of Contents
+ [Supported Platforms](#supported-platforms)
+ [Getting Started](#getting-started)
    - [Installation](#installation)
    - [Source Code Import](#source-code-import)
+ [Source Code Documentation](#source-code-documentation)
+ [Production File Location](#production-file-location)
+ [Tests](#tests)
+ [Version](#version)
+ [License](#license)
+ [Comments](#comments-and-questions) 

## Supported Platforms
This package was developed and tested on Windows, although the package
was designed to operate on multiple operating systems including 
Mac-OS, Linux and Windows.

While development testing has focused primarily on *Windows*, the unit
tests are now completing successfully on *Linux Mint 19.2* and *Ubuntu 18.04.3*.


## Getting Started

### Installation
Use this command to download and install the *strops* package
locally. Note: Version 3.0.0 supports *Go* modules.

    go get github.com/MikeAustin71/stringopsgo/strops/v3

To update the package run:    
    go get -u github.com/MikeAustin71/stringopsgo/strops/v3


### Source Code Import        
You will need to import and reference this package in your source code
files. Go module support has been available from version 2 onwards.

To import version 3 or later, use the following import statement:

    import "github.com/MikeAustin71/stringopsgo/strops/v3"  

To import version 2, use the following import statement:

    import "github.com/MikeAustin71/stringopsgo/strops/v2"

To import version 1, which does NOT support Go modules, use the following
import statement:

    import "github.com/MikeAustin71/stringopsgo/strops"

## Source Code Documentation

[Source Documentation](http://godoc.org/github.com/MikeAustin71/stringopsgo/strops/v3)    


## Production File Location
All the active production files are located in directory path:

    github.com/MikeAustin71/stringopsgo/strops/v3

## Tests
Currently, the *strops/v3* package has 317 unit tests with a code coverage
of 93%. 

Test coverage and outcomes are documented in:

[github.com/MikeAustin71/stringopsgo/strops/v3/xx_tests.txt](./strops/v3/xx_tests.txt)

[How To Run Tests Documentation](./strops/v3/wt_HowToRunTests.md)

## Version
The latest version is Version 3.0.0. Like Version 2+, Version 3+ also
supports *Go* modules. 

[Version 3.0.0 Release Notes](./strops/v3/releasenotes.md)

## License

Use of this source code is governed by the (open-source)
MIT-style license which can be found in the LICENSE file
located in this directory.

[MIT License](./LICENSE)

## Comments And Questions

Send questions or comments to:

    mike.go@paladinacs.net



