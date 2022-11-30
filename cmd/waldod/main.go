package main

import (
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	"github.com/waldolove-ventures/waldo/app"
	"github.com/waldolove-ventures/waldo/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd(
		"waldo",
		"cosmos",
		app.DefaultNodeHome,
		"waldo-1",
		app.ModuleBasics,
		app.New,
		// this line is used by starport scaffolding # root/arguments
	)

	rootCmd.AddCommand(cmd.AddConsumerSectionCmd(app.DefaultNodeHome))

	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
