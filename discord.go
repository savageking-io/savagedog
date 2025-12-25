package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/savageking-io/savagedog/proto"
)

type Discord struct {
	session      *discordgo.Session
	guildId      string
	channelIds   map[string]string // Map of service -> discord channel
	authors      map[string]string // Author names [Service]Name
	authorUrls   map[string]string // Author URLs [Service]URL
	authorImages map[string]string // Author Images [Service]Image
}

func (d *Discord) Init(config *DiscordConfig) error {
	log.Traceln("Discord::Init")
	if config == nil {
		return fmt.Errorf("config is nil")
	}

	d.channelIds = make(map[string]string)
	d.authors = make(map[string]string)
	d.authorUrls = make(map[string]string)
	d.authorImages = make(map[string]string)

	var err error
	d.session, err = discordgo.New("Bot " + config.Token)
	if err != nil {
		return fmt.Errorf("error creating Discord session: %v", err)
	}

	err = d.session.Open()
	if err != nil {
		return fmt.Errorf("error opening connection: %v", err)
	}

	if len(d.session.State.Guilds) > 1 {
		return fmt.Errorf("more than one guild found, but only one discord server supported")
	}

	if len(d.session.State.Guilds) == 0 {
		return fmt.Errorf("no guild found, make sure you are in a discord server")
	}

	d.guildId = d.session.State.Guilds[0].ID

	return nil
}

func (d *Discord) RegisterChannelForService(serviceName string, channelName string) error {
	log.Traceln("Discord::RegisterChannelForService", serviceName, channelName)
	if d.session == nil {
		return fmt.Errorf("session is nil")
	}

	if d.guildId == "" {
		return fmt.Errorf("guild id is empty")
	}

	exists, id := d.verifyChannel(channelName)
	if exists {
		d.channelIds[serviceName] = id
		return nil
	}
	return fmt.Errorf("channel %s not found", channelName)
}

func (d *Discord) SetServiceAuthor(serviceName, author string) error {
	log.Traceln("Discord::SetServiceAuthor", serviceName, author)
	if author == "" {
		return fmt.Errorf("author is empty")
	}

	d.authors[serviceName] = author
	return nil
}

func (d *Discord) SetServiceAuthorUrl(serviceName, url string) error {
	log.Traceln("Discord::SetServiceAuthorUrl", serviceName, url)
	if url == "" {
		return fmt.Errorf("url is empty")
	}
	d.authorUrls[serviceName] = url
	return nil
}

func (d *Discord) SetServiceAuthorImage(serviceName, image string) error {
	log.Traceln("Discord::SetServiceAuthorImage", serviceName, image)
	if image == "" {
		return fmt.Errorf("image is empty")
	}
	d.authorImages[serviceName] = image
	return nil
}

func (d *Discord) verifyChannel(channelName string) (bool, string) {
	log.Traceln("Discord::verifyChannel", channelName)
	if d.guildId == "" {
		return false, ""
	}
	channels, err := d.session.GuildChannels(d.guildId)
	if err != nil {
		return false, ""
	}

	for _, ch := range channels {
		if ch.Name == channelName {
			return true, ch.ID
		}
	}
	return false, ""
}

func (d *Discord) Close() error {
	log.Traceln("Discord::Close")
	if d.session == nil {
		return nil
	}
	return d.session.Close()
}

func (d *Discord) SendMessage(service string, msg *proto.NotificationMessage) error {
	log.Traceln("Discord::SendMessage", service)
	if msg == nil {
		return fmt.Errorf("message is nil")
	}

	channelID, ok := d.channelIds[service]
	if !ok {
		return fmt.Errorf("channel for service %s not found", service)
	}

	m := &discordgo.MessageSend{
		Embed: &discordgo.MessageEmbed{},
	}

	authorName, authorExist := d.authors[service]
	authorUrl, authorUrlExist := d.authorUrls[service]
	authorImage, authorImageExist := d.authorImages[service]
	if authorExist || authorUrlExist || authorImageExist {
		m.Embed.Author = &discordgo.MessageEmbedAuthor{}
	}

	if authorExist && authorName != "" {
		m.Embed.Author.Name = authorName
	}

	if authorUrlExist && authorUrl != "" {
		m.Embed.Author.URL = authorUrl
	}

	if authorImageExist && authorImage != "" {
		m.Embed.Author.IconURL = authorImage
	}

	if msg.Color != "" {
		var err error
		m.Embed.Color, err = hexColorToInt(msg.Color)
		if err != nil {
			log.Errorf("Failed to parse color %s: %v", msg.Color, err)
		}
	}

	if msg.Header != "" {
		m.Embed.Title = msg.Header
	}

	if msg.Url != "" {
		m.Embed.URL = msg.Url
	}

	if msg.Content != "" {
		m.Embed.Description = msg.Content
	}

	if msg.Fields != nil && len(msg.Fields) > 0 {
		for _, f := range msg.Fields {
			m.Embed.Fields = append(m.Embed.Fields, &discordgo.MessageEmbedField{
				Name:  f.FieldName,
				Value: f.FieldValue,
			})
		}
	}

	if msg.Sender != "" {
		m.Embed.Footer = &discordgo.MessageEmbedFooter{Text: msg.Sender}
	}

	_, err := d.session.ChannelMessageSendComplex(channelID, m)
	if err != nil {
		fmt.Println("Error:", err)
	}

	return err
}

func hexColorToInt(color string) (int, error) {
	log.Traceln("Discord::hexColorToInt", color)
	color = strings.TrimPrefix(color, "#")

	if len(color) != 6 {
		return 0, fmt.Errorf("invalid hex color: must be 6 digits, got %s", color)
	}

	value, err := strconv.ParseInt(color, 16, 32)
	if err != nil {
		return 0, fmt.Errorf("failed to parse hex color %s: %v", color, err)
	}

	return int(value), nil
}
