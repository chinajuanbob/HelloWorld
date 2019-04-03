package main

import "github.com/chinajuanbob/helloworld/cmd/hwcli/cmd"
import "github.com/davecgh/go-spew/spew"

func main() {
	spew.Config = *spew.NewDefaultConfig()
	spew.Config.ContinueOnMethod = true
	cmd.Execute()
}
