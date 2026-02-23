package event

import (
	base "github.com/bamgoo/base"
)

type (
	Driver interface {
		Connect(*Instance) (Connection, error)
	}

	Connection interface {
		Open() error
		Close() error
		Start() error
		Stop() error

		Register(name, group string) error
		Publish(name string, data []byte) error
	}

	Instance struct {
		conn    Connection
		Name    string
		Config  Config
		Setting base.Map
	}
)
