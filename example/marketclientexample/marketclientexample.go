package marketclientexample

import (
	"encoding/json"
	"github.com/hopex-hk/go_sdk/client"
	"github.com/hopex-hk/go_sdk/core/logging"
	"github.com/hopex-hk/go_sdk/core/logging/zaplogger"
	"github.com/hopex-hk/go_sdk/example/config"
	"time"
)

var logger logging.Logger = &zaplogger.ZapLogger{}

func init() {
	logger.SetLevel(logging.INFO)
}

func RunAllExample() {
	GetKLine()
	GetMarketTicker()
	GetMarkets()
	GetTrades()
	GetMarketDepth()
}

func logRes(res interface{}) {
	bs, err := json.Marshal(res)
	if err != nil {
		logger.Error("has error: %+v", err)
		return
	}
	logger.Info("res: %s", string(bs))
}

func GetKLine() {
	client := new(client.MarketClient).InitByDefault(config.DemoConfig)

	now := time.Now()
	duration, _ := time.ParseDuration("-30m")
	startTime := now.Add(duration).Unix()
	endTime := now.Unix()

	res, err := client.GetKline("BTCUSDT", "5m", startTime, endTime)
	if err != nil {
		logger.Error("has error: %+v", err)
		return
	}

	logRes(res)
}

func GetMarketTicker() {
	client := new(client.MarketClient).InitByDefault(config.DemoConfig)

	res, err := client.GetMarketTicker("BTCUSDT")

	if err != nil {
		logger.Error("has error: %+v", err)
		return
	}

	logRes(res)
}

func GetMarkets() {
	client := new(client.MarketClient).InitByDefault(config.DemoConfig)

	res, err := client.GetMarkets()

	if err != nil {
		logger.Error("has error: %+v", err)
		return
	}

	logRes(res)
}

func GetTrades() {
	client := new(client.MarketClient).InitByDefault(config.DemoConfig)

	res, err := client.GetTrades("BTCUSDT", 20)

	if err != nil {
		logger.Error("has error: %+v", err)
		return
	}

	logRes(res)
}

func GetMarketDepth() {
	client := new(client.MarketClient).InitByDefault(config.DemoConfig)

	res, err := client.GetMarketDepth("BTCUSDT", 20)

	if err != nil {
		logger.Error("has error: %+v", err)
		return
	}

	logRes(res)
}
