package homeclientexample

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
	GetIndexStatistics()
	GetIndexNotify()
}

func logRes(res interface{}) {
	bs, err := json.Marshal(res)
	if err != nil {
		logger.Error("has error: %+v", err)
		return
	}
	logger.Info("res: %s", string(bs))
}

func GetIndexStatistics() {
	client := new(client.HomeClient).InitByDefault(config.DemoConfig)

	res, err := client.GetIndexStatistics()

	if err != nil {
		logger.Error("has error: %+v", err)
		return
	}

	logRes(res)
}

func GetIndexNotify() {
	client := new(client.HomeClient).InitByDefault(config.DemoConfig)

	page, limit := 1, 20
	res, err := client.GetIndexNotify(page, limit)

	if err != nil {
		logger.Error("has error: %+v", err)
		return
	}

	logRes(res)
}
