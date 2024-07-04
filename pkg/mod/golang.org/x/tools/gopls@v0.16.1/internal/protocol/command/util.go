// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package command

import (
	"encoding/json"
	"fmt"
)

// A Command identifies one of gopls' ad-hoc extension commands
// that may be invoked through LSP's executeCommand.
type Command string

func (c Command) String() string { return string(c) }

// IsAsync reports whether the command is asynchronous:
// clients must wait for the "end" progress notification.
func (c Command) IsAsync() bool {
	switch string(c) {
	// TODO(adonovan): derive this list from interface.go somewhow.
	// Unfortunately we can't even reference the enum from here...
	case "gopls.run_tests", "gopls.run_govulncheck", "gopls.test":
		return true
	}
	return false
}

// MarshalArgs encodes the given arguments to json.RawMessages. This function
// is used to construct arguments to a protocol.Command.
//
// Example usage:
//
//	jsonArgs, err := MarshalArgs(1, "hello", true, StructuredArg{42, 12.6})
func MarshalArgs(args ...interface{}) ([]json.RawMessage, error) {
	var out []json.RawMessage
	for _, arg := range args {
		argJSON, err := json.Marshal(arg)
		if err != nil {
			return nil, err
		}
		out = append(out, argJSON)
	}
	return out, nil
}

// UnmarshalArgs decodes the given json.RawMessages to the variables provided
// by args. Each element of args should be a pointer.
//
// Example usage:
//
//	var (
//	    num int
//	    str string
//	    bul bool
//	    structured StructuredArg
//	)
//	err := UnmarshalArgs(args, &num, &str, &bul, &structured)
func UnmarshalArgs(jsonArgs []json.RawMessage, args ...interface{}) error {
	if len(args) != len(jsonArgs) {
		return fmt.Errorf("DecodeArgs: expected %d input arguments, got %d JSON arguments", len(args), len(jsonArgs))
	}
	for i, arg := range args {
		if err := json.Unmarshal(jsonArgs[i], arg); err != nil {
			return err
		}
	}
	return nil
}
