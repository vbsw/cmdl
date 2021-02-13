/*
 *       Copyright 2020, 2021, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

package cmdl

import (
	"testing"
)

func TestParseA(t *testing.T) {
	args := []string{"asdf", "--version"}
	cl := NewFrom(args)
	version := cl.NewParam().Parse("-v", "--version")

	if !version.Available() {
		t.Error()
	} else {
		if version.Keys()[0] != args[1] || version.Values()[0] != "" {
			t.Error(version.Keys()[0])
		}
	}
	if cl.parsed[0] {
		t.Error()
	}
	if !cl.parsed[1] {
		t.Error()
	}
	if cl.parsedCount != 1 {
		t.Error(cl.parsedCount)
	}
}

func TestParseB(t *testing.T) {
	args := []string{"--start", "asdf", "-s", "qwer"}
	cl := NewFrom(args)
	start := cl.NewParam().Parse("-s", "--start")

	if !start.Available() {
		t.Error()
	} else {
		if len(start.Keys()) == 2 {
			if start.Keys()[0] != args[0] || start.Values()[0] != "" {
				t.Error(start.Keys()[0])
			}
			if start.Keys()[1] != args[2] || start.Values()[1] != "" {
				t.Error(start.Keys()[1])
			}
		} else {
			t.Error(len(start.Keys()))
		}
	}
}

func TestParsePairs(t *testing.T) {
	args := []string{"asdf", "--start=123"}
	asgOp := NewAsgOp(false, true, "=")
	cl := NewFrom(args)
	start := cl.NewParam().ParsePairs(asgOp, "-s", "--start")

	if !start.Available() {
		t.Error()
	} else {
		if len(start.Keys()) == 1 {
			if start.Keys()[0] != "--start" || start.Values()[0] != "123" {
				t.Error(start.Keys()[0])
			}
		} else {
			t.Error(len(start.Keys()))
		}
	}
	if cl.parsed[0] {
		t.Error()
	}
	if !cl.parsed[1] {
		t.Error()
	}
	if cl.parsedCount != 1 {
		t.Error(cl.parsedCount)
	}
}

func TestUnparsedArgs(t *testing.T) {
	args := []string{"--start", "asdf", "-s", "qwer"}
	cl := NewFrom(args)
	cl.NewParam().Parse("--start", "-s")
	rest := cl.UnparsedArgs()

	if len(rest) != 2 {
		t.Error(len(rest))

	} else if rest[0] != args[1] {
		t.Error(rest[0])

	} else if rest[1] != args[3] {
		t.Error(rest[1])
	}
}
