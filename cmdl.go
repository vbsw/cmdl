/*
 *       Copyright 2020, 2021, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

// Package cmdl provides functions to parse command line arguments.
package cmdl

import (
	"os"
	"strings"
)

// CommandLine holds command line arguments.
type CommandLine struct {
	args        []string
	parsed      []bool
	parsedCount int
}

// Parameter represents a parsed argument from command line.
type Parameter struct {
	cl     *CommandLine
	keys   []string
	values []string
}

// AssignmentOperator represents the operator between key and value in argument.
type AssignmentOperator struct {
	ops          []string
	blankAllowed bool
	emptyAllowed bool
}

// New creates and returns a new instance of CommandLine.
func New() *CommandLine {
	cl := NewFrom(osArgsCopy())
	return cl
}

// NewAsgOp creates and returns a new instance of AssignmentOperator.
func NewAsgOp(blankAllowed, emptyAllowed bool, ops ...string) *AssignmentOperator {
	asgOp := new(AssignmentOperator)
	asgOp.ops = ops
	asgOp.blankAllowed = blankAllowed
	asgOp.emptyAllowed = emptyAllowed
	return asgOp
}

// NewFrom creates and returns a new instance of CommandLine.
func NewFrom(args []string) *CommandLine {
	cl := new(CommandLine)
	cl.args = args
	cl.parsed = make([]bool, len(args))
	return cl
}

// NewParam creates and returns a new instance of Parameter.
func (cl *CommandLine) NewParam() *Parameter {
	param := new(Parameter)
	param.cl = cl
	return param
}

// Args returns command line arguments.
func (cl *CommandLine) Args() []string {
	return cl.args
}

// UnparsedArgs returns unparsed arguments.
func (cl *CommandLine) UnparsedArgs() []string {
	unparsedArgsCount := len(cl.args) - cl.parsedCount
	unparsedArgs := make([]string, unparsedArgsCount)
	j := 0

	for i, arg := range cl.args {
		if !cl.parsed[i] {
			unparsedArgs[j] = arg
			j++
		}
	}
	return unparsedArgs
}

func (cl *CommandLine) setParsed(index int) {
	cl.parsed[index] = true
	cl.parsedCount++
}

// Parse searches for keys in command line. Returns itself.
func (param *Parameter) Parse(keys ...string) *Parameter {
	if len(keys) > 0 && len(param.cl.args) > param.cl.parsedCount {
		for i, arg := range param.cl.args {
			if !param.cl.parsed[i] {
				for _, key := range keys {
					if arg == key {
						param.Add(key, "")
						param.cl.setParsed(i)
						break
					}
				}
			}
		}
	}
	return param
}

// ParsePairs searches for keys in command line.
func (param *Parameter) ParsePairs(asgOps *AssignmentOperator, keys ...string) *Parameter {
	if len(asgOps.ops) > 0 || asgOps.blankAllowed || asgOps.emptyAllowed {
		param.parsePairs(asgOps, keys)
	}
	return param
}

// Add appends key and value.
func (param *Parameter) Add(key, value string) {
	param.keys = append(param.keys, key)
	param.values = append(param.values, value)
}

// Available returns true, if parameter has been found in command line.
func (param *Parameter) Available() bool {
	return len(param.keys) > 0
}

// Count returns number of matches in command line.
func (param *Parameter) Count() int {
	return len(param.keys)
}

// Keys returns parsed keys.
func (param *Parameter) Keys() []string {
	return param.keys
}

// Values returns parsed values.
func (param *Parameter) Values() []string {
	return param.values
}

// MatchingOp returns true and matching operator, if str starts with assignment operator.
func (asgOps *AssignmentOperator) MatchingOp(str string) (bool, string) {
	for _, op := range asgOps.ops {
		if strings.HasPrefix(str, op) {
			return true, op
		}
	}
	return asgOps.emptyAllowed, ""
}

func (param *Parameter) parsePairs(asgOps *AssignmentOperator, keys []string) {
	if len(keys) > 0 && len(param.cl.args) > param.cl.parsedCount {
		if asgOps.blankAllowed {
			param.parsePairsBlankAllowed(asgOps, keys)
		} else {
			param.parsePairsWithoutBlank(asgOps, keys)
		}
	}
}

func (param *Parameter) parsePairsBlankAllowed(asgOps *AssignmentOperator, keys []string) {
	for i := 0; i < len(param.cl.args); i++ {
		if !param.cl.parsed[i] {
			arg := param.cl.args[i]

			for _, key := range keys {
				if strings.HasPrefix(arg, key) {
					if len(arg) == len(key) {
						value := ""
						param.cl.setParsed(i)

						if i+1 < len(param.cl.args) {
							value = param.cl.args[i+1]
							param.cl.setParsed(i + 1)
							i++
						}
						param.Add(key, value)
						break

					} else {
						argWithoutKey := arg[len(key):]
						match, op := asgOps.MatchingOp(argWithoutKey)

						if match {
							param.Add(key, argWithoutKey[len(op):])
							param.cl.setParsed(i)
							break
						}
					}
				}
			}
		}
	}
}

func (param *Parameter) parsePairsWithoutBlank(asgOps *AssignmentOperator, keys []string) {
	for i, arg := range param.cl.args {
		if !param.cl.parsed[i] {
			for _, key := range keys {
				if strings.HasPrefix(arg, key) {
					if len(arg) == len(key) {
						param.Add(key, "")
						param.cl.setParsed(i)
						break

					} else {
						argWithoutKey := arg[len(key):]
						match, op := asgOps.MatchingOp(argWithoutKey)

						if match {
							param.Add(key, argWithoutKey[len(op):])
							param.cl.setParsed(i)
							break
						}
					}
				}
			}
		}
	}
}

func osArgsCopy() []string {
	if len(os.Args) > 1 {
		args := make([]string, len(os.Args)-1)
		copy(args, os.Args[1:])
		return args
	}
	return make([]string, 0)
}
