package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/etcinit/phabulous/app"
	"github.com/facebookgo/inject"
	"github.com/jacobstr/confer"
)

// VERSION is the application version string.
var VERSION string

func main() {
	// Seed rand.
	rand.Seed(time.Now().UTC().UnixNano())

	// Create the configuration
	// In this case, we will be using the environment and some safe defaults
	config := confer.NewConfig()
	config.ReadPaths("config/main.yml", "config/main.production.yml")
	config.AutomaticEnv()

	// Create the logger.
	logger := logrus.New()

	// Next, we setup the dependency graph
	var g inject.Graph
	var phabulous app.Phabulous
	g.Provide(
		&inject.Object{Value: config},
		&inject.Object{Value: &phabulous},
		&inject.Object{Value: logger},
	)
	if err := g.Populate(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Setup the command line application
	app := cli.NewApp()
	app.Name = "phabulous"
	app.Usage = "A Phabricator bot in Go"

	// Set version and authorship info
	app.Version = VERSION
	app.Author = "Eduardo Trujillo <ed@chromabits.com>"

	// Setup the default action. This action will be triggered when no
	// subcommand is provided as an argument
	app.Action = func(c *cli.Context) error {
		fmt.Println(
			"Usage: phabulous [global options] command [command options] " +
				"[arguments...]",
		)
		return nil
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config",
			Usage: "Provide an alternative configuration file",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "serve",
			Aliases: []string{"s", "server", "listen"},
			Usage:   "Start the API server",
			Action:  phabulous.Serve.Run,
		},
	}

	// Begin
	app.Run(os.Args)
}
