package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "savagedog"
	app.Version = AppVersion
	app.Description = "A watchdog and notification service for Discord"
	app.Usage = ""

	app.Authors = []cli.Author{
		{
			Name:  "savageking.io",
			Email: "i@savageking.io",
		},
		{
			Name:  "Mike Savochkin (crioto)",
			Email: "mike@crioto.com",
		},
	}

	app.Copyright = "2025 (c) savageking.io. All Rights Reserved"

	app.Commands = []cli.Command{
		{
			Name:  "serve",
			Usage: "Start daemon",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "config",
					Usage:       "Configuration filepath",
					Value:       ConfigFilepath,
					Destination: &ConfigFilepath,
				},
				cli.StringFlag{
					Name:        "log",
					Usage:       "Specify logging level",
					Value:       "",
					Destination: &LogLevel,
				},
			},
			Action: Serve,
		},
		{
			Name:      "notify",
			ShortName: "n",
			Usage:     "Send a notification to daemon",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "dest",
					Usage:       "Destination to send notification to",
					Value:       Dest,
					Destination: &Dest,
					Required:    true,
				},
				cli.StringFlag{
					Name:        "from",
					Usage:       "Sender identifier, must match one in a backend",
					Value:       From,
					Destination: &From,
					Required:    true,
				},
				cli.StringFlag{
					Name:        "header",
					Usage:       "Header of the notification",
					Value:       Header,
					Destination: &Header,
				},
				cli.StringFlag{
					Name:        "content",
					Usage:       "Content of the message",
					Value:       Content,
					Destination: &Content,
				},
				cli.StringFlag{
					Name:        "sender",
					Usage:       "Specify sender",
					Value:       Sender,
					Destination: &Sender,
				},
				cli.StringFlag{
					Name:        "fields",
					Usage:       "Key=Value pairs separated by & to send as fields",
					Value:       Fields,
					Destination: &Fields,
				},
				cli.StringFlag{
					Name:        "config",
					Usage:       "Configuration filepath",
					Value:       ConfigFilepath,
					Destination: &ConfigFilepath,
				},
				cli.StringFlag{
					Name:        "log",
					Usage:       "Specify logging level",
					Value:       "",
					Destination: &LogLevel,
				},
			},
			Action: Notify,
		},
	}

	_ = app.Run(os.Args)
}

func Serve(c *cli.Context) error {
	log.Infof("Start daemon")
	config := new(Config)
	if err := config.Read(ConfigFilepath); err != nil {
		log.Fatalf("Can't initialize daemon: %s", err)
		return err
	}
	log.Infof("Configuration loaded")

	discord := new(Discord)
	if err := discord.Init(&config.Discord); err != nil {
		log.Fatalf("Can't initialize Discord: %s", err)
		return err
	}
	log.Infof("Discord initialized")

	dog := new(DogService)
	if err := dog.Init(config.Hostname, config.Port, discord); err != nil {
		log.Fatalf("Can't initialize DogService: %s", err)
		return err
	}

	if err := dog.Run(); err != nil {
		log.Fatalf("Service failed: %s", err)
		return err
	}

	return nil
}

func Notify(c *cli.Context) error {
	return nil
}
