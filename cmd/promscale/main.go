// This file and its contents are licensed under the Apache License 2.0.
// Please see the included NOTICE for copyright information and
// LICENSE for a copy of the license.

package main

import (
	"fmt"
	"os"

	"github.com/timescale/promscale/pkg/log"
	"github.com/timescale/promscale/pkg/runner"
	"github.com/timescale/promscale/pkg/version"
	_ "go.uber.org/automaxprocs"
)

func main() {
	args := os.Args[1:]
	if shouldProceed := runner.ParseArgs(args); !shouldProceed {
		os.Exit(0)
	}
	cfg := &runner.Config{}
	cfg, err := runner.ParseFlags(cfg, args)
	if err != nil {
		fmt.Println("Version: ", version.Promscale, "Commit Hash: ", version.CommitHash)
		fmt.Println("Fatal error: cannot parse flags: ", err)
		os.Exit(1)
	}
	err = log.Init(cfg.LogCfg)
	if err != nil {
		fmt.Println("Version: ", version.Promscale, "Commit Hash: ", version.CommitHash)
		fmt.Println("Fatal error: cannot start logger: ", err)
		os.Exit(1)
	}
	runner.WarnOnDeprecatedConfig(cfg)
	err = runner.Run(cfg)
	if err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
