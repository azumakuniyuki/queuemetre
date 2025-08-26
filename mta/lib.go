// Copyright (C) 2024-2025 azumakuniyuki, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//   ___                                       _               ____  __ _____  _    
//  / _ \ _   _  ___ _   _  ___ _ __ ___   ___| |_ _ __ ___   / /  \/  |_   _|/ \   
// | | | | | | |/ _ \ | | |/ _ \ '_ ` _ \ / _ \ __| '__/ _ \ / /| |\/| | | | / _ \  
// | |_| | |_| |  __/ |_| |  __/ | | | | |  __/ |_| | |  __// / | |  | | | |/ ___ \ 
//  \__\_\\__,_|\___|\__,_|\___|_| |_| |_|\___|\__|_|  \___/_/  |_|  |_| |_/_/   \_\

// Package "mta" provide MTA-specific functions
package mta
import "os"

var Daemon = []string{"Postfix", "Sendmail", "OpenSMTPD", "Dragonfly"}
var IsThis = make(map[string]func() bool, 4)
var Counts = make(map[string]func() bool, 4)

// canRead returns true if the argument is a directory, is readable, and is executable.
//   Arguments:
//     - argv1 (string): A path to directory
//   Returns:
//     - (bool):         false if the argument is not a directory, is not readable, or is not executable.
func canRead(argv1 string) bool {
	if argv1 == "" { return false }

	v, e := os.Stat(argv1)
	if e != nil || v.IsDir() == false               { return false }
	if v.Mode() & 0444 == 0 || v.Mode() & 0111 == 0 { return false }
	return true
}

