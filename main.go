package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hpcloud/tail"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Usage = "parse error log to ndjson"
	app.EnableBashCompletion = true
	app.Commands = []*cli.Command{
		{
			Name:  "tail",
			Usage: "tail nginx log files and genererate ndjson output",
			Flags: []cli.Flag{
				&cli.StringFlag{Name: "error-log", Required: true, Aliases: []string{"error"}, Usage: "Input nginx error log"},
				&cli.StringFlag{Name: "ndjson-log", Required: true, Aliases: []string{"ndjson"}, Usage: "Output ndjson"},
				&cli.BoolFlag{Name: "follow", Aliases: []string{"f"}, Value: false, Usage: "Follow the error log"},
			},
			Action: func(c *cli.Context) error {

				// tail
				t, _ := tail.TailFile(c.String("error-log"), tail.Config{Follow: c.Bool("follow")})

				// ndjson
				f, err := os.OpenFile(c.String("ndjson-log"), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
				if err != nil {
					log.Fatal(err)
				}
				defer f.Close()

				for entry := range t.Lines {
					ngxParser, e := Parser(entry.Text)

					if e != nil {
						log.Println("Parse failed on: ", entry.Text)
					}

					json, _ := ParserJSON(ngxParser)
					if _, err = f.Write(append(json, "\n"...)); err != nil {
						log.Fatal(err)
					}
				}

				return nil
			},
		},
		{
			Name:  "test",
			Usage: "Test entry from cli and output the json",
			Flags: []cli.Flag{
				&cli.StringFlag{Name: "log", Required: true, Aliases: []string{"error"}, Usage: "Input nginx error log"},
			},
			Action: func(c *cli.Context) error {
				ngxParser, e := Parser(c.String("log"))
				if e != nil {
					log.Fatal(e)
				}
				json, e := ParserJSON(ngxParser)
				if e == nil {
					fmt.Println(string(json))
				} else {
					log.Fatal(e)
				}
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
