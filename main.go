package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "utils",
		Usage: "simple",
		Commands: []*cli.Command{
			{
				Name:    "split-group",
				Usage:   "splits log lines but a series of regular expressions, maintaining a count of each as well as non-matches",
        Flags: []cli.Flag{
          &cli.BoolFlag{
            Name: "multi-match",
            Usage: "allow multiple matches per line",
            Aliases: []string{"m"},
          },
          &cli.BoolFlag{
            Name: "log-non-matches",
            Aliases: []string{"x"},
            Usage: "log lines that don't match any pattern",
          },
        },
				Aliases: []string{"sg"},
				Action: func(cCtx *cli.Context) error {
					_, err := splitGroup(cCtx.Args().Slice(), cCtx.Bool("multi-match"), cCtx.Bool("log-non-matches"))
					return err
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
