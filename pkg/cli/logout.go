// Copyright (C) 2015 Scaleway. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.md file.

package cli

import (
	"github.com/scaleway/scaleway-cli/pkg/commands"
	"github.com/scaleway/scaleway-cli/vendor/github.com/Sirupsen/logrus"
)

var cmdLogout = &Command{
	Exec:        runLogout,
	UsageLine:   "logout [OPTIONS]",
	Description: "Log out from the Scaleway API",
	Help:        "Log out from the Scaleway API.",
}

func init() {
	cmdLogout.Flag.BoolVar(&logoutHelp, []string{"h", "-help"}, false, "Print usage")
}

// FLags
var logoutHelp bool // -h, --help flag

func runLogout(cmd *Command, rawArgs []string) {
	if logoutHelp {
		cmd.PrintUsage()
	}
	if len(rawArgs) != 0 {
		cmd.PrintShortUsage()
	}

	args := commands.LogoutArgs{}
	ctx := cmd.GetContext(rawArgs)
	err := commands.RunLogout(ctx, args)
	if err != nil {
		logrus.Fatalf("Cannot execute 'logout': %v", err)
	}
}