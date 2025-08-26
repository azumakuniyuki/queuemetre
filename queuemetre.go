// Copyright (C) 2024-2025 azumakuniyuki, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//   ___                                       _            
//  / _ \ _   _  ___ _   _  ___ _ __ ___   ___| |_ _ __ ___ 
// | | | | | | |/ _ \ | | |/ _ \ '_ ` _ \ / _ \ __| '__/ _ \
// | |_| | |_| |  __/ |_| |  __/ | | | | |  __/ |_| | |  __/
//  \__\_\\__,_|\___|\__,_|\___|_| |_| |_|\___|\__|_|  \___|

// Package "main"
package main

import "os"
import "fmt"
import "github.com/alecthomas/kong"
import "libsisimai.org/queuemetre/mta"

const Version = "1.1.0"
type parameters struct {
	Once bool
	Wait uint
}
type commandArg struct {
	Once    bool `short:"1" help:"Print the current queue status and exit"`
	Verbose uint `short:"v" type:"counter" help:"Verbose mode"`
	Wait    uint `short:"w" help:"Wait <n> seconds before getting the next queue status"`
}

func main() {
	emx := ""
	arg := &commandArg{}
	ctx := kong.Parse(arg,
		kong.Vars{"version": Version},
		kong.Description("queuemetre - A command for displaying the number of email queues like the iostat(8)"),
		kong.UsageOnError(),
	)
	if arg.Wait == 0 { arg.Wait = 60 }
fmt.Printf("MTA = %s\n", mta.Daemon)

	for _, e := range mta.Daemon {
		fmt.Printf("Call %s\n", e)
		if mta.IsThis[e]() == true { emx = e; break }
	}



	fmt.Printf("MTA = %#v\n", emx)
	fmt.Printf("CTX = %#v\n", ctx)
	fmt.Printf("ARG.Wait = %s\n", arg.Wait)

}

// IsContained checks whether any element in argv2 includes argv1 or not. (argv1 < argv2)
//   Arguments:
//     - argv1 (string):   String to be contained as a substring listed in argv2
//     - argv2 ([]string): List of strings
//   Returns:
//     - (bool):           true if one or more string in argv1 was found in argv2
func detectDaemon() string {
	daemonfile := map[string]string{
		"dragonfly": "/var/spool/mda/flush",
		"exim":      "/var/spool/exim/exim-daemon.pid",
		"opensmtpd": "/run/smtpd.pid",
		"postfix":   "/var/spool/postfix/pid/master.pid",
		"sendmail":  "/var/run/sendmail.pid",
	}
	for e := range daemonfile {
		in, oe := os.Stat(daemonfile[e])
		if os.IsNotExist(oe) == false && in.IsDir() == false { return e }
	}
	return ""
}

