package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/savageking-io/savagedog/proto"
	"testing"
)

func TestDiscord_Close(t *testing.T) {
	type fields struct {
		session      *discordgo.Session
		guildId      string
		channelIds   map[string]string
		authors      map[string]string
		authorUrls   map[string]string
		authorImages map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Discord{
				session:      tt.fields.session,
				guildId:      tt.fields.guildId,
				channelIds:   tt.fields.channelIds,
				authors:      tt.fields.authors,
				authorUrls:   tt.fields.authorUrls,
				authorImages: tt.fields.authorImages,
			}
			if err := d.Close(); (err != nil) != tt.wantErr {
				t.Errorf("Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDiscord_Init(t *testing.T) {
	type fields struct {
		session      *discordgo.Session
		guildId      string
		channelIds   map[string]string
		authors      map[string]string
		authorUrls   map[string]string
		authorImages map[string]string
	}
	type args struct {
		config *DiscordConfig
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Discord{
				session:      tt.fields.session,
				guildId:      tt.fields.guildId,
				channelIds:   tt.fields.channelIds,
				authors:      tt.fields.authors,
				authorUrls:   tt.fields.authorUrls,
				authorImages: tt.fields.authorImages,
			}
			if err := d.Init(tt.args.config); (err != nil) != tt.wantErr {
				t.Errorf("Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDiscord_RegisterChannelForService(t *testing.T) {
	type fields struct {
		session      *discordgo.Session
		guildId      string
		channelIds   map[string]string
		authors      map[string]string
		authorUrls   map[string]string
		authorImages map[string]string
	}
	type args struct {
		serviceName string
		channelName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Discord{
				session:      tt.fields.session,
				guildId:      tt.fields.guildId,
				channelIds:   tt.fields.channelIds,
				authors:      tt.fields.authors,
				authorUrls:   tt.fields.authorUrls,
				authorImages: tt.fields.authorImages,
			}
			if err := d.RegisterChannelForService(tt.args.serviceName, tt.args.channelName); (err != nil) != tt.wantErr {
				t.Errorf("RegisterChannelForService() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDiscord_SendMessage(t *testing.T) {
	type fields struct {
		session      *discordgo.Session
		guildId      string
		channelIds   map[string]string
		authors      map[string]string
		authorUrls   map[string]string
		authorImages map[string]string
	}
	type args struct {
		service string
		msg     *proto.NotificationMessage
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Discord{
				session:      tt.fields.session,
				guildId:      tt.fields.guildId,
				channelIds:   tt.fields.channelIds,
				authors:      tt.fields.authors,
				authorUrls:   tt.fields.authorUrls,
				authorImages: tt.fields.authorImages,
			}
			if err := d.SendMessage(tt.args.service, tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("SendMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDiscord_SetServiceAuthor(t *testing.T) {
	type fields struct {
		session      *discordgo.Session
		guildId      string
		channelIds   map[string]string
		authors      map[string]string
		authorUrls   map[string]string
		authorImages map[string]string
	}
	type args struct {
		serviceName string
		author      string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Discord{
				session:      tt.fields.session,
				guildId:      tt.fields.guildId,
				channelIds:   tt.fields.channelIds,
				authors:      tt.fields.authors,
				authorUrls:   tt.fields.authorUrls,
				authorImages: tt.fields.authorImages,
			}
			if err := d.SetServiceAuthor(tt.args.serviceName, tt.args.author); (err != nil) != tt.wantErr {
				t.Errorf("SetServiceAuthor() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDiscord_SetServiceAuthorImage(t *testing.T) {
	type fields struct {
		session      *discordgo.Session
		guildId      string
		channelIds   map[string]string
		authors      map[string]string
		authorUrls   map[string]string
		authorImages map[string]string
	}
	type args struct {
		serviceName string
		image       string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Discord{
				session:      tt.fields.session,
				guildId:      tt.fields.guildId,
				channelIds:   tt.fields.channelIds,
				authors:      tt.fields.authors,
				authorUrls:   tt.fields.authorUrls,
				authorImages: tt.fields.authorImages,
			}
			if err := d.SetServiceAuthorImage(tt.args.serviceName, tt.args.image); (err != nil) != tt.wantErr {
				t.Errorf("SetServiceAuthorImage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDiscord_SetServiceAuthorUrl(t *testing.T) {
	type fields struct {
		session      *discordgo.Session
		guildId      string
		channelIds   map[string]string
		authors      map[string]string
		authorUrls   map[string]string
		authorImages map[string]string
	}
	type args struct {
		serviceName string
		url         string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Discord{
				session:      tt.fields.session,
				guildId:      tt.fields.guildId,
				channelIds:   tt.fields.channelIds,
				authors:      tt.fields.authors,
				authorUrls:   tt.fields.authorUrls,
				authorImages: tt.fields.authorImages,
			}
			if err := d.SetServiceAuthorUrl(tt.args.serviceName, tt.args.url); (err != nil) != tt.wantErr {
				t.Errorf("SetServiceAuthorUrl() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDiscord_verifyChannel(t *testing.T) {
	type fields struct {
		session      *discordgo.Session
		guildId      string
		channelIds   map[string]string
		authors      map[string]string
		authorUrls   map[string]string
		authorImages map[string]string
	}
	type args struct {
		channelName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
		want1  string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Discord{
				session:      tt.fields.session,
				guildId:      tt.fields.guildId,
				channelIds:   tt.fields.channelIds,
				authors:      tt.fields.authors,
				authorUrls:   tt.fields.authorUrls,
				authorImages: tt.fields.authorImages,
			}
			got, got1 := d.verifyChannel(tt.args.channelName)
			if got != tt.want {
				t.Errorf("verifyChannel() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("verifyChannel() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_hexColorToInt(t *testing.T) {
	type args struct {
		color string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := hexColorToInt(tt.args.color)
			if (err != nil) != tt.wantErr {
				t.Errorf("hexColorToInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("hexColorToInt() got = %v, want %v", got, tt.want)
			}
		})
	}
}
