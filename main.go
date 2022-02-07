package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"github.com/urfave/cli/v2"
	"gopkg.in/mcuadros/go-syslog.v2"
)

func accessLogSyslogToTCP(c *cli.Context) error {
	channel := make(syslog.LogPartsChannel)
	handler := syslog.NewChannelHandler(channel)

	tcpServerAddress := c.String("tcp-server")

retryConnect:

	tcpAddr, err := net.ResolveTCPAddr("tcp", tcpServerAddress)
	if err != nil {
		log.Printf("Cannot resolve tcp-server: %s, %s\n", tcpServerAddress, err.Error())
		time.Sleep(time.Second * 3)
		goto retryConnect
	}

	connection, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Printf("Cannot dial tcp-server: %s, %s\n", tcpServerAddress, err.Error())
		time.Sleep(time.Second * 3)
		goto retryConnect
	}
	defer connection.Close()
	log.Printf("connection successfully established: %s\n", tcpServerAddress)

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
					_, connectionErr := connection.Write([]byte(strings.TrimSpace(json) + "\n"))
					if connectionErr != nil {
						log.Printf("Connection lost: %s, %s\n", tcpServerAddress, connectionErr.Error())
						os.Exit(1)
					}
				}
			}
		}
	}(channel)

	server.Wait()

	return nil
}

func errorLogSyslogToTCP(c *cli.Context) error {
	tcpServerAddress := c.String("tcp-server")

retryConnect:

	tcpAddr, err := net.ResolveTCPAddr("tcp", tcpServerAddress)
	if err != nil {
		log.Printf("Cannot resolve tcp-server: %s, %s\n", tcpServerAddress, err.Error())
		time.Sleep(time.Second * 3)
		goto retryConnect
	}

	connection, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Printf("Cannot dial tcp-server: %s, %s\n", tcpServerAddress, err.Error())
		time.Sleep(time.Second * 3)
		goto retryConnect
	}
	defer connection.Close()
	log.Printf("connection successfully established: %s\n", tcpServerAddress)

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
					log.Printf("Parse failed on: %s\n", content)
				} else {
					json, err := parserJSON(ngxParser)
					if err == nil {
						_, connectionErr := connection.Write(append(json, '\n'))
						if connectionErr != nil {
							log.Printf("Connection lost: %s, %s\n", tcpServerAddress, connectionErr.Error())
							os.Exit(1)
						}
					}
				}
			}
		}
	}(channel)

	server.Wait()

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
			Name:  "access-log",
			Usage: "Receive nginx access log(escape=json) from Syslog(UDP) and send to TCP Server",
			Flags: []cli.Flag{
				&cli.StringFlag{Value: "127.0.0.1:5141", Name: "udp-listen", Aliases: []string{"listen"}, Usage: "Listen syslog RFC3164"},
				&cli.StringFlag{Value: "127.0.0.1:6141", Name: "tcp-server", Aliases: []string{"tcp"}, Usage: "Target TCP server"},
			},
			Action: accessLogSyslogToTCP,
		},
		{
			Name:  "error-log",
			Usage: "Receive nginx error log from Syslog(UDP) and send to TCP Server",
			Flags: []cli.Flag{
				&cli.StringFlag{Value: "127.0.0.1:5140", Name: "udp-listen", Aliases: []string{"listen"}, Usage: "Listen syslog RFC3164"},
				&cli.StringFlag{Value: "127.0.0.1:6140", Name: "tcp-server", Aliases: []string{"tcp"}, Usage: "Target TCP server"},
			},
			Action: errorLogSyslogToTCP,
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
