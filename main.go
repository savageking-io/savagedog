package main

import (
	"context"
	"fmt"
	"github.com/savageking-io/savagedog/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/url"
	"os"
	"time"

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
					Name:        "color",
					Usage:       "Hex of a message color",
					Value:       Color,
					Destination: &Color,
				},
				cli.StringFlag{
					Name:        "config",
					Usage:       "Configuration filepath",
					Value:       ClientConfigFilePath,
					Destination: &ClientConfigFilePath,
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
	SetLogLevel()
	log.Infof("Start daemon")
	if ConfigFilepath == "" {
		log.Errorf("Config file not specified. Use --config")
		return fmt.Errorf("config file not specified")
	}
	config := new(Config)
	if err := ReadConfig(ConfigFilepath, &config); err != nil {
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

	for _, s := range config.Services {
		log.Debugf("Registering channel for service %s: %s", s.Name, s.Channel)
		if err := discord.RegisterChannelForService(s.Name, s.Channel); err != nil {
			log.Errorf("Can't register channel for service %s: %s", s.Name, err)
			return err
		}
		if s.Author != "" {
			log.Debugf("Setting author for service %s: %s", s.Name, s.Author)
			if err := discord.SetServiceAuthor(s.Name, s.Author); err != nil {
				log.Errorf("Can't set author for service %s: %s", s.Name, err)
				return err
			}
		}
		if s.AuthorURL != "" {
			log.Debugf("Setting author url for service %s: %s", s.Name, s.AuthorURL)
			if err := discord.SetServiceAuthorUrl(s.Name, s.AuthorURL); err != nil {
				log.Errorf("Can't set author url for service %s: %s", s.Name, err)
			}
		}
		if s.AuthorImage != "" {
			log.Debugf("Setting author image for service %s: %s", s.Name, s.AuthorImage)
			if err := discord.SetServiceAuthorImage(s.Name, s.AuthorImage); err != nil {
				log.Errorf("Can't set author image for service %s: %s", s.Name, err)
			}
		}
	}

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
	SetLogLevel()
	log.Infof("Launching notificator")
	config := new(ClientConfig)
	if ClientConfigFilePath != "" {
		log.Infof("Configuration file provided. Loading...")
		if err := ReadConfig(ClientConfigFilePath, &config); err != nil {
			log.Fatalf("Can't initialize daemon: %s", err)
			return err
		}
		log.Infof("Configuration loaded")
	}

	if Dest != "" {
		if config.Dest != "" {
			log.Infof("Overriding destination from config: %s", config.Dest)
		}
		config.Dest = Dest
		log.Debugf("Destination set to %s", Dest)
	}

	if From != "" {
		if config.From != "" {
			log.Infof("Overriding sender from config: %s", config.From)
		}
		config.From = From
		log.Debugf("Sender set to %s", From)
	}

	if Header != "" {
		if config.Header != "" {
			log.Infof("Overriding header from config: %s", config.Header)
		}
		config.Header = Header
		log.Debugf("Header set to %s", Header)
	}

	if Content != "" {
		if config.Content != "" {
			log.Infof("Overriding content from config: %s", config.Content)
		}
		config.Content = Content
		log.Debugf("Content set to %s", Content)
	}

	if Sender != "" {
		if config.Sender != "" {
			log.Infof("Overriding sender from config: %s", config.Sender)
		}
		config.Sender = Sender
		log.Debugf("Sender set to %s", Sender)
	}

	if Fields != "" {
		if config.Fields != "" {
			log.Infof("Overriding fields from config: %s", config.Fields)
		}
		config.Fields = Fields
		log.Debugf("Fields set to %s", Fields)
	}

	if Color != "" {
		if config.Color != "" {
			log.Infof("Overriding color from config: %s", config.Color)
		}
		config.Color = Color
		log.Debugf("Color set to %s", Color)
	}

	if config.From == "" {
		log.Fatalf("Sending service was not specified. Use --from")
		return fmt.Errorf("sender is empty")
	}

	if config.Dest == "" {
		log.Fatalf("Destination was not specified. Use --dest")
		return fmt.Errorf("destination is empty")
	}

	msg := &proto.NotificationMessage{}

	if config.Fields != "" {
		values, err := url.ParseQuery(config.Fields)
		if err != nil {
			log.Fatalf("Can't parse fields: %s", err)
			return err
		}

		for k, v := range values {
			msg.Fields = append(msg.Fields, &proto.MessageField{FieldName: k, FieldValue: v[0]})
		}
	}

	if config.Header != "" {
		msg.Header = config.Header
	}

	if config.Content != "" {
		msg.Content = config.Content
	}

	if config.Sender != "" {
		msg.Sender = config.Sender
	}

	if config.Color != "" {
		msg.Color = config.Color
	}

	msg.FromService = config.From

	log.Infof("Connecting to %s...", config.Dest)
	conn, err := grpc.NewClient(config.Dest, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewDogServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.Notification(ctx, msg)
	if err != nil {
		log.Fatalf("Failed to send notification: %v", err)
	}

	log.Debugf("Response Code: %d, Error: %s", resp.Code, resp.ErrorText)

	if resp.Code != 0 {
		log.Fatalf("Notification failed with code %d: %s", resp.Code, resp.ErrorText)
	}

	log.Infof("Notification sent successfully")

	return nil
}
