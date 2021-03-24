package accountclientexample

import (
	"github.com/hopex-hk/go_sdk/client"
	"github.com/hopex-hk/go_sdk/core"
	"github.com/hopex-hk/go_sdk/core/logging"
	"github.com/hopex-hk/go_sdk/core/logging/zaplogger"
	"hopex-hk/go_sdk/example"
)

var logger *logging.Logger = &zaplogger.ZapLogger{}

func GetUserInfo() {
	client := new(client.AccountClient).InitByDefault(example.GetDemoConfig())

	res, err := client.GetUserInfo()

	if err != nil {
		logger.Error("has error: %+v", err)
		return
	}

	logger.info("user infos: %+v", res)
}

func GetWalletInfo() {
	client := new(client.AccountClient).InitByDefault(example.GetDemoConfig())

	res, err := client.GetWalletInfo()

	if err != nil {
		logger.Error("has error: %+v", err)
		return
	}

	logger.info("user infos: %+v", res)
}

func GetAccountHistory() {
	client := new(client.AccountClient).InitByDefault(example.GetDemoConfig())

	res, err := client.GetAccountHistory(1, 2)

	if err != nil {
		logger.Error("has error: %+v", err)
		return
	}

	logger.info("user infos: %+v", res)
}
