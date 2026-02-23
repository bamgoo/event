package event

import . "github.com/bamgoo/base"

func Publish(name string, values ...Map) error {
	return module.publish("", name, values...)
}

func PublishTo(conn, name string, values ...Map) error {
	return module.publish(conn, name, values...)
}

func Broadcast(name string, values ...Map) error {
	return module.broadcast("", name, values...)
}

func BroadcastTo(conn, name string, values ...Map) error {
	return module.broadcast(conn, name, values...)
}

func RegisterDriver(name string, driver Driver) {
	module.RegisterDriver(name, driver)
}

func RegisterConfig(name string, cfg Config) {
	module.RegisterConfig(name, cfg)
}

func RegisterConfigs(cfgs Configs) {
	module.RegisterConfigs(cfgs)
}
