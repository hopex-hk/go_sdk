package tradeclientexample

import (
	"encoding/json"
	"github.com/hopex-hk/go_sdk/client"
	"github.com/hopex-hk/go_sdk/model"
	"github.com/hopex-hk/go_sdk/core/logging"
	"github.com/hopex-hk/go_sdk/core/logging/zaplogger"
	"github.com/hopex-hk/go_sdk/example/config"
	"github.com/shopspring/decimal"
	"time"
)

var logger logging.Logger = &zaplogger.ZapLogger{}

func init() {
	logger.SetLevel(logging.INFO)
}

func logRes(res interface{}) {
	bs, err := json.Marshal(res)
	if err != nil {
		logger.Error("has error: %+v", err)
		return
	}
	logger.Info("res: %s", string(bs))
}

func RunAllExample() {
	GetOpenOrders()
	GetPositions()
	Order()
	ConditionOrder()
}

func GetOpenOrders() {
	client := new(client.TradeClient).InitByDefault(config.DemoConfig)

	res, err := client.GetOpenOrders("BTCUSDT")

	if err != nil {
		logger.Error("has error: %+v", err)
		return
	}

	logRes(res)
}

func GetPositions() {
	client := new(client.TradeClient).InitByDefault(config.DemoConfig)

	res, err := client.GetPositions("BTCUSDT")

	if err != nil {
		logger.Error("has error: %+v", err)
		return
	}

	logRes(res)
}

func Order() {
	client := new(client.TradeClient).InitByDefault(config.DemoConfig)

	logger.Info("put a market order")

	orderId, err := client.CreateOrder("BTCUSDT", 1, decimal.NullDecimal{
		Decimal: decimal.NewFromFloat32(1),
		Valid:   true,
	}, 1)

	if err != nil {
		logger.Error("create order has error: %v", err)
		return
	}
	if orderId <= 0 {
		logger.Error("expect orderId >0,but %d", orderId)
	}

	logger.Info("order success. orderid: %d", orderId)

	time.Sleep(1000 * time.Millisecond)

	logger.Info("get open orders")
	openorders, err := client.GetOpenOrders("BTCUSDT")

	if err != nil {
		logger.Error("get open orders has error: %v", err)
		return
	}
	hasOpenOrder := false
	for _, o := range openorders {
		if o.OrderId == orderId {
			hasOpenOrder = true
			break
		}
	}
	if !hasOpenOrder {
		logger.Error("open order %d not exists", orderId)
		return
	}

	logger.Info("cancel order %d", orderId)

	err = client.CancelOrder("BTCUSDT", orderId)

	if err != nil {
		logger.Error("cancel order has error: %v", err)
		return
	}

	logger.Info("%d cancelled success", orderId)

	time.Sleep(1000 * time.Millisecond)

	openorders, err = client.GetOpenOrders("BTCUSDT")

	if err != nil {
		logger.Error("get open orders has error: %v", err)
		return
	}
	hasOpenOrder = false
	for _, o := range openorders {
		if o.OrderId == orderId {
			hasOpenOrder = true
			break
		}
	}
	if hasOpenOrder {
		logger.Error("open order %d cancelled but exists", orderId)
		return
	}
}

func ConditionOrder() {
	client := new(client.TradeClient).InitByDefault(config.DemoConfig)
	trigPrice := decimal.NewFromFloat32(1.1)

	logger.Info("put a condition order")
	err := client.CreateConditionOrder("BTCUSDT", 1, "Market", trigPrice, 1, decimal.NullDecimal{})

	if err != nil {
		logger.Error("create order has error: %v", err)
		return
	}
	logger.Info("order success")

	time.Sleep(1000 * time.Millisecond)

	logger.Info("query condition orders")
	orders, err := client.GetConditionOrders(&model.GetConditionOrdersParam{
		ContractCodeList: []string{"BTCUSDT"},
		Side:             2,
		Direct:           1,
		TaskStatusList:   []int{1},
	}, 1, 20)

	if err != nil {
		logger.Error("get open orders has error: %v", err)
		return
	}
	if orders.TotalCount == 0 {
		logger.Error("expect at least a condition order")
	}
	hasOrder := false
	var taskId int64
	for _, o := range orders.Result {
		if o.TrigPrice.Equal(trigPrice) {
			hasOrder = true
			taskId = o.TaskId
			break
		}
	}
	if !hasOrder {
		logger.Error("condition order not exists")
		return
	}

	err = client.CancelConditionOrder("BTCUSDT", taskId)

	if err != nil {
		logger.Error("cancel condition order has error: %v", err)
		return
	}
	time.Sleep(1000 * time.Millisecond)
	orders, err = client.GetConditionOrders(&model.GetConditionOrdersParam{
		ContractCodeList: []string{"BTCUSDT"},
		Side:             2,
		Direct:           1,
		TaskStatusList:   []int{1},
	}, 1, 20)

	if err != nil {
		logger.Error("get condition orders has error: %v", err)
		return
	}

	hasOrder = false
	for _, o := range orders.Result {
		if o.TrigPrice.Equal(trigPrice) {
			hasOrder = true
			break
		}
	}
	if hasOrder {
		logger.Error("condition order %d cancelled but exists", taskId)
		return
	}
}
