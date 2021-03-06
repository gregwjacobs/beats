package cmd

import (
	"flag"

	"github.com/spf13/pflag"

	"github.com/elastic/beats/packetbeat/beater"

	cmd "github.com/elastic/beats/libbeat/cmd"
)

// Name of this beat
var Name = "packetbeat"

// RootCmd to handle beats cli
var RootCmd *cmd.BeatsRootCmd

func init() {
	var runFlags = pflag.NewFlagSet(Name, pflag.ExitOnError)
	runFlags.AddGoFlag(flag.CommandLine.Lookup("I"))
	runFlags.AddGoFlag(flag.CommandLine.Lookup("t"))
	runFlags.AddGoFlag(flag.CommandLine.Lookup("O"))
	runFlags.AddGoFlag(flag.CommandLine.Lookup("l"))
	runFlags.AddGoFlag(flag.CommandLine.Lookup("dump"))

	// TODO Deprecate this one in favor of subcommand
	runFlags.AddGoFlag(flag.CommandLine.Lookup("devices"))

	RootCmd = cmd.GenRootCmdWithRunFlags(Name, "", beater.New, runFlags)
}
