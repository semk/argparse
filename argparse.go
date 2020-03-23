// Simple and elegant argument parser for scripts that can be used as a shebang
// (Bash | Python | Perl)
//
// Usage (direct):
// 		argparse --executor=/bin/bash --arg name::No Name Defined::Name printname.sh
//
// Usage (as shebang in the script)
//		#!/usr/bin/argparse --executor=/bin/bash --arg name::No Name Defined::Name
//
// The arguments are then exported as environment variables within the script
// 		$ARG_NAME
//
// Author: Sreejith Kesavan <sreejithemk@gmail.com>

package main

import (
	"flag"
	"fmt"
	"os"
	"syscall"
	"strings"
)

const SEP = "::"

func buildVariableName(name string) string {
	return strings.ToUpper(fmt.Sprintf("ARG_%s", name))
}

type argFlags []string

func (i *argFlags) String() string {
	return ""
}

func (i *argFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func main() {
	exportedVars := make(map[string]*string)
	// --executor option can be used to specify the original shebang
	executor := flag.String(
		"executor",
		"/bin/bash",
		"Path to the script executor (set this to shebang (#!))")

	var argPatterns argFlags
	flag.Var(&argPatterns, "arg", "Define custom argument to the program.")

	// Parse the flags passed to argparse utility
	flag.Parse()

	// Get the script to be executed with the flags passed to the script
	originalScriptWithArgs := flag.Args()

	// Define the script flags
	scriptFlagSet := flag.NewFlagSet(originalScriptWithArgs[0], flag.ExitOnError)
	for _, arg := range argPatterns {
		argOptions := strings.SplitN(arg, SEP, 3)
		if len(argOptions) < 3 {
			fmt.Println("Invaid arguments. Format: argparse <program> [name::val::desc]...")
			os.Exit(1)
		}

		argName := argOptions[0]
		argDefaultVal := argOptions[1]
		argDesc := argOptions[2]

		exportedVars[buildVariableName(argName)] = scriptFlagSet.String(argName, argDefaultVal, argDesc)
	}

	// Validate and parse the flags passed to the script
	scriptFlagSet.Parse(originalScriptWithArgs[1:])

	var scriptEnvVars []string
	for key, val := range exportedVars {
		scriptEnvVars = append(scriptEnvVars, fmt.Sprintf("%s=%s", key, *val))
	}

	scriptArgs := []string{*executor}
	scriptArgs = append(scriptArgs, originalScriptWithArgs...)

	// Exec the program with arguments
	err := syscall.Exec(*executor, scriptArgs, scriptEnvVars)

	if err != nil {
		fmt.Printf("Command execution failed. %s", err.Error())
	}
}
