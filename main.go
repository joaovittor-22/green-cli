package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

const version = "v0.1.0"

func main() {
	app := &cli.App{
		Name:  "green",
		Usage: "Generate retroactive Git commits",
		Version: version,
		Commands: []*cli.Command{
			{
				Name:  "today",
				Usage: "Create a Git commit for today",
				Action: func(c *cli.Context) error {
					fmt.Println("Committing for today...")
					MakeCommit(time.Now(), 0)
					fmt.Println("Done.")
					return nil
				},
			},
			{
				Name:  "days",
				Usage: "Create backdated commits for past N days",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:    "num",
						Aliases: []string{"n"},
						Value:   1,
						Usage:   "Number of backdated commits",
					},
				},
				Action: func(c *cli.Context) error {
					start := time.Now().AddDate(0, 0, -numDays)
                    for i := 0; i < numDays; i++ {
                        date := start.AddDate(0, 0, i)
                        MakeCommit(date, i)
                    }

					fmt.Printf("Created %d backdated commits.\n", num)
					return nil
				},
			},
			{
				Name:  "version",
				Usage: "Show version",
				Action: func(c *cli.Context) error {
					fmt.Println("green", version)
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}