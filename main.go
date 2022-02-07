package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/nxadm/tail"
	"github.com/urfave/cli/v2"
	"gopkg.in/mcuadros/go-syslog.v2"
)

func accessLogSyslogToTCP(c *cli.Context) error {
	channel := make(syslog.LogPartsChannel)
	handler := syslog.NewChannelHandler(channel)

	tcpServerAddress := c.String("tcp-server")

	tcpAddr, err := net.ResolveTCPAddr("tcp", tcpServerAddress)
	if err != nil {
		log.Fatal(err)
	}

	// server connection
	connection, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()

	server := syslog.NewServer()
	server.SetFormat(syslog.RFC3164)
	server.SetHandler(handler)
	server.ListenUDP(c.String("udp-listen"))
	server.Boot()

	go func(channel syslog.LogPartsChannel) {
		for logParts := range channel {
			if content, ok := logParts["content"]; ok {
				json := fmt.Sprintf("%v", content)
				if isJSON(json) {
					connection.Write([]byte(strings.TrimSpace(json) + "\n"))
				}
			}
		}
	}(channel)

	server.Wait()

	return nil
}

func errorLogSyslogToTCP(c *cli.Context) error {
	tcpServerAddress := c.String("tcp-server")

	tcpAddr, err := net.ResolveTCPAddr("tcp", tcpServerAddress)
	if err != nil {
		log.Fatal(err)
	}

	// server connection
	connection, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()

	channel := make(syslog.LogPartsChannel)
	handler := syslog.NewChannelHandler(channel)
	server := syslog.NewServer()
	server.SetFormat(syslog.RFC3164)
	server.SetHandler(handler)
	server.ListenUDP(c.String("udp-listen"))
	server.Boot()

	go func(channel syslog.LogPartsChannel) {
		for logParts := range channel {
			if content, ok := logParts["content"]; ok {
				ngxParser, err := Parser(fmt.Sprintf("%v", content))

				if err != nil {
					log.Println("Parse failed on: ", content)
				} else {
					json, err := parserJSON(ngxParser)
					if err == nil {
						connection.Write(append(json, '\n'))
					}
				}
			}
		}
	}(channel)

	server.Wait()

	return nil
}

func tailToNDJSON(c *cli.Context) error {
	// tail
	t, err := tail.TailFile(c.String("error-log"), tail.Config{Follow: c.Bool("follow")})
	if err != nil {
		log.Fatal(err)
	}

	// ndjson
	f, err := os.OpenFile(c.String("ndjson-log"), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for entry := range t.Lines {
		ngxParser, err := Parser(entry.Text)

		if err != nil {
			log.Println("Parse failed on: ", entry.Text)
		} else {
			json, err := parserJSON(ngxParser)
			if err == nil {
				if _, err = f.Write(append(json, "\n"...)); err != nil {
					log.Fatal(err)
				}
			}
		}
	}

	return nil
}

func tailToTCP(c *cli.Context) error {

	// tail
	t, _ := tail.TailFile(c.String("error-log"), tail.Config{Follow: c.Bool("follow")})

	tcpServerAddress := c.String("tcp-server")

	tcpAddr, err := net.ResolveTCPAddr("tcp", tcpServerAddress)
	if err != nil {
		log.Fatal(err)
	}

	// server connection
	connection, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()

	for entry := range t.Lines {
		ngxParser, err := Parser(entry.Text)

		if err != nil {
			log.Println("Parse failed on: ", entry.Text)
		} else {
			json, err := parserJSON(ngxParser)
			if err == nil {
				connection.Write(append(json, '\n'))
			}
		}
	}

	return nil
}

func testLog(c *cli.Context) error {
	ngxParser, e := Parser(c.String("log"))
	if e != nil {
		log.Fatal(e)
	}
	json, e := parserJSON(ngxParser)
	if e == nil {
		fmt.Println(string(json))
	} else {
		log.Fatal(e)
	}
	return nil
}

func main() {
	app := cli.NewApp()
	app.Usage = "parse nginx error log to structured JSON"
	app.EnableBashCompletion = true
	app.Commands = []*cli.Command{
		{
			Name:  "access-log-syslog",
			Usage: "Receive nginx access log from Syslog(UDP) and send to TCP Server",
			Flags: []cli.Flag{
				&cli.StringFlag{Value: "127.0.0.1:5141", Name: "udp-listen", Aliases: []string{"listen"}, Usage: "Listen syslog RFC3164"},
				&cli.StringFlag{Value: "127.0.0.1:6171", Name: "tcp-server", Aliases: []string{"tcp"}, Usage: "Target TCP server"},
			},
			Action: accessLogSyslogToTCP,
		},
		{
			Name:  "error-log-syslog",
			Usage: "Receive nginx error log from Syslog(UDP) and send to TCP Server",
			Flags: []cli.Flag{
				&cli.StringFlag{Value: "127.0.0.1:5140", Name: "udp-listen", Aliases: []string{"listen"}, Usage: "Listen syslog RFC3164"},
				&cli.StringFlag{Value: "127.0.0.1:6172", Name: "tcp-server", Aliases: []string{"tcp"}, Usage: "Target TCP server"},
			},
			Action: errorLogSyslogToTCP,
		},
		{
			Name:  "tail-to-tcp",
			Usage: "Tail nginx log files and send to TCP server",
			Flags: []cli.Flag{
				&cli.StringFlag{Name: "error-log", Required: true, Aliases: []string{"error"}, Usage: "Input nginx error log"},
				&cli.StringFlag{Name: "tcp-server", Required: true, Aliases: []string{"tcp"}, Usage: "Target TCP server"},
				&cli.BoolFlag{Name: "follow", Aliases: []string{"f"}, Value: false, Usage: "Follow the error log"},
			},
			Action: tailToTCP,
		},
		{
			Name:  "tail-to-ndjson",
			Usage: "Tail nginx error log files and generate NDJSON file",
			Flags: []cli.Flag{
				&cli.StringFlag{Name: "error-log", Required: true, Aliases: []string{"error"}, Usage: "Input nginx error log"},
				&cli.StringFlag{Name: "ndjson-log", Required: true, Aliases: []string{"ndjson"}, Usage: "Output ndjson"},
				&cli.BoolFlag{Name: "follow", Aliases: []string{"f"}, Value: false, Usage: "Follow the error log"},
			},
			Action: tailToNDJSON,
		},
		{
			Name:  "test",
			Usage: "Test entry from cli and output the json",
			Flags: []cli.Flag{
				&cli.StringFlag{Name: "log", Required: true, Aliases: []string{"error"}, Usage: "Input nginx error log"},
			},
			Action: testLog,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
