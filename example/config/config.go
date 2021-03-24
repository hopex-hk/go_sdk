package config

import "github.com/hopex-hk/go_sdk/core"

var DemoConfig *core.Config

func init() {
	DemoConfig = core.NewConfig(
		"https://devapi2.hopex.com",
		"a54bf64e-e0b5-4f91-8516-cc6dfc473946",
		"6yJoD76GzpcJUME5SDnCqxUG6oxyJHht",
		"Demo/0.0.1",
	)
}
