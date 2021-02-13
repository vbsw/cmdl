# cmdl

[![GoDoc](https://godoc.org/github.com/vbsw/cmdl?status.svg)](https://godoc.org/github.com/vbsw/cmdl) [![Go Report Card](https://goreportcard.com/badge/github.com/vbsw/cmdl)](https://goreportcard.com/report/github.com/vbsw/cmdl) [![Stability: Experimental](https://masterminds.github.io/stability/experimental.svg)](https://masterminds.github.io/stability/experimental.html)

## About
Package cmdl provides functions to parse command line arguments. It is published on <https://github.com/vbsw/cmdl> and <https://gitlab.com/vbsw/cmdl>.

## Copyright
Copyright 2020, 2021, Vitali Baumtrok (vbsw@mailbox.org).

cmdl is distributed under the Boost Software License, version 1.0. (See accompanying file LICENSE or copy at http://www.boost.org/LICENSE_1_0.txt)

cmdl is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the Boost Software License for more details.

## Usage

### Example A

	package main

	import (
		"fmt"
		"github.com/vbsw/cmdl"
	)

	func main() {
		cl := cmdl.New()

		if cl.NewParam().Parse("--help", "-h").Available() {
			fmt.Println("valid parameters are -h or -v.")

		} else if cl.NewParam().Parse("--version", "-v").Available() {
			fmt.Println("version 1.0.0")

		} else {
			unparsedArgs := cl.UnparsedArgs()

			if len(unparsedArgs) == 1 {
				fmt.Println("error: unknown parameter", unparsedArgs[0])

			} else if len(unparsedArgs) > 1 {
				fmt.Println("error: too many arguments")
			}
		}
	}

### Example B

	package main

	import (
		"fmt"
		"github.com/vbsw/cmdl"
	)

	func main() {
		start := "0"
		end := "0"
		cl := cmdl.New()
		asgOp := cmdl.NewAsgOp(false, false, "=")

		paramStart := cl.NewParam().ParsePairs(asgOp, "start")
		paramEnd := cl.NewParam().ParsePairs(asgOp, "end")

		if paramStart.Available() {
			start = paramStart.Values()[0]
			end = start
		}
		if paramEnd.Available() {
			end = paramEnd.Values()[0]
		}
		fmt.Println("processing from", start, "to", end)
	}

Command line:

	$ ./test start=1 end=10
	$ processing from 1 to 10

## References
- https://golang.org/doc/install
- https://git-scm.com/book/en/v2/Getting-Started-Installing-Git
