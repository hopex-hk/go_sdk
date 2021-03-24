package accountclientexample

import (
	"encoding/json"

	"github.com/hopex-hk/go_sdk/client"
	"github.com/hopex-hk/go_sdk/core/logging"
	"github.com/hopex-hk/go_sdk/core/logging/zaplogger"
	"github.com/hopex-hk/go_sdk/example/config"
)

var logger logging.Logger = &zaplogger.ZapLogger{}

func init() {
	logger.SetLevel(logging.INFO)
}

func RunAllExample() {
	GetUserInfo()
	GetWalletInfo()
	GetAccountHistory()
}

func logRes(res interface{}) {
	bs, err := json.Marshal(res)
	if err != nil {
		logger.Error("has error: %+v", err)
		return
	}
	logger.Info("res: %s", string(bs))
}

func GetUserInfo() {
	client := new(client.AccountClient).InitByDefault(config.DemoConfig)

	res, err := client.GetUserInfo()

	if err != nil {
		logger.Error("has error: %+v", err)
		return
	}

	logRes(res)
}

func GetWalletInfo() {
	client := new(client.AccountClient).InitByDefault(config.DemoConfig)

	res, err := client.GetWalletInfo()

	if err != nil {
		logger.Error("has error: %+v", err)
		return
	}

	logRes(res)
}

func GetAccountHistory() {
	client := new(client.AccountClient).InitByDefault(config.DemoConfig)

	res, err := client.GetAccountHistory(1, 2)

	if err != nil {
		logger.Error("has error: %+v", err)
		return
	}

	logRes(res)
}
