package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/savageking-io/savagedog/proto"
)

type Discord struct {
	session    *discordgo.Session
	guildId    string
	channelIds map[string]string
}

func (d *Discord) Init(config *DiscordConfig) error {
	if config == nil {
		return fmt.Errorf("config is nil")
	}

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

	d.guildId = d.session.State.Guilds[0].ID

	return nil
}

func (d *Discord) RegisterChannelForService(serviceName string, channelName string) error {
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

func (d *Discord) verifyChannel(channelName string) (bool, string) {
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
	if d.session == nil {
		return nil
	}
	return d.session.Close()
}

func (d *Discord) SendMessage(service string, msg *proto.NotificationMessage) error {
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

	if msg.Header != "" {
		m.Embed.Title = msg.Header
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
