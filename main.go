package main

import (
	"github.com/project-ginza/notifications/cmd"
	"github.com/project-ginza/notifications/config"
	_const "github.com/project-ginza/notifications/const"
	flag "github.com/spf13/pflag"
)

var strEnv string

func init() {
	_ = flag.StringP("env", "e", "", "Runtime Environment")
	flag.Parse()

	strEnv = flag.Lookup("env").Value.String()
	_const.Env = &strEnv
}

func main() {
	// load configuration
	config.Load(strEnv)

	// start the service.
	if err := cmd.Root().Execute(); err != nil {
		panic(err)
	}
}
