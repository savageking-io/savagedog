package main

import (
	"io"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"testing"
)

func TestMain(m *testing.M) {
	origExit := log.StandardLogger().ExitFunc
	log.StandardLogger().ExitFunc = func(code int) {}
	defer func() { log.StandardLogger().ExitFunc = origExit }()
	log.SetOutput(io.Discard)
	m.Run()
}

func TestNotify(t *testing.T) {
	origFrom, origDest, origFields := From, Dest, Fields
	defer func() {
		From, Dest, Fields = origFrom, origDest, origFields
	}()

	type args struct {
		c *cli.Context
	}
	tests := []struct {
		name    string
		prep    func()
		args    args
		wantErr bool
	}{
		{
			name: "missing sender (--from)",
			prep: func() {
				From = ""
				Dest = "localhost:12005"
				Fields = ""
			},
			args:    args{},
			wantErr: true,
		},
		{
			name: "missing destination (--dest)",
			prep: func() {
				From = "svc"
				Dest = ""
				Fields = ""
			},
			args:    args{},
			wantErr: true,
		},
		{
			name: "invalid fields query string",
			prep: func() {
				From = "svc"
				Dest = "localhost:12005"
				Fields = "a%zz" // invalid percent-encoding triggers parse error
			},
			args:    args{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// apply case-specific globals
			if tt.prep != nil {
				tt.prep()
			}
			if err := Notify(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Notify() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServe(t *testing.T) {
	origConfig := ConfigFilepath
	defer func() { ConfigFilepath = origConfig }()

	type args struct {
		c *cli.Context
	}
	tests := []struct {
		name    string
		prep    func()
		args    args
		wantErr bool
	}{
		{
			name: "empty --config path",
			prep: func() { ConfigFilepath = "" },
			args: args{}, wantErr: true,
		},
		{
			name: "nonexistent config file",
			prep: func() { ConfigFilepath = "definitely-does-not-exist.yaml" },
			args: args{}, wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.prep != nil {
				tt.prep()
			}
			if err := Serve(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Serve() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
