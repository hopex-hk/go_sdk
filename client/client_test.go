package client

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/hopex-hk/go_sdk/core"
	"github.com/hopex-hk/go_sdk/core/logging"
	"github.com/hopex-hk/go_sdk/core/logging/zaplogger"
	"github.com/hopex-hk/go_sdk/model"
	"github.com/shopspring/decimal"
)

var cfg *core.Config = core.GetUnitTestConfig()

func TestMain(m *testing.M) {
	zaplogger.SetLevel(logging.DEBUG)

	m.Run()
}

func Test_GetUserInfo(t *testing.T) {
	client := new(AccountClient).InitByDefault(cfg)

	res, err := client.GetUserInfo()

	if err != nil {
		t.Errorf("has error: %v", err)
		return
	}

	resbs, err := json.Marshal(res)
	if err != nil {
		t.Errorf("has error: %v", err)
		return
	}
	t.Log(string(resbs))
}

func Test_GetWalletInfo(t *testing.T) {
	client := new(AccountClient).InitByDefault(cfg)

	res, err := client.GetWalletInfo()

	if err != nil {
		t.Errorf("has error: %v", err)
		return
	}

	resbs, err := json.Marshal(res)
	if err != nil {
		t.Errorf("has error: %v", err)
		return
	}
	t.Log(string(resbs))
}

func Test_GetWalletInfoV2(t *testing.T) {
	client := new(AccountClient).InitByDefault(cfg)

	res, err := client.GetWalletInfoV2()

	if err != nil {
		t.Errorf("has error: %v", err)
		return
	}

	resbs, err := json.Marshal(res)
	if err != nil {
		t.Errorf("has error: %v", err)
		return
	}
	t.Log(string(resbs))
}

func Test_GetAccountHistory(t *testing.T) {
	client := new(AccountClient).InitByDefault(cfg)

	res, err := client.GetAccountHistory(1, 2)

	if err != nil {
		t.Errorf("has error: %v", err)
		return
	}

	resbs, err := json.Marshal(res)
	if err != nil {
		t.Errorf("has error: %v", err)
		return
	}
	t.Log(string(resbs))
}

func Test_GetIndexStatistics(t *testing.T) {
	client := new(HomeClient).InitByDefault(cfg)

	res, err := client.GetIndexStatistics()

	if err != nil {
		t.Errorf("has error: %v", err)
		return
	}

	resbs, err := json.Marshal(res)
	if err != nil {
		t.Errorf("has error: %v", err)
		return
	}
	t.Log(string(resbs))
}

func Test_GetIndexNotify(t *testing.T) {
	client := new(HomeClient).InitByDefault(cfg)

	page, limit := 1, 20
	res, err := client.GetIndexNotify(page, limit)

	if err != nil {
		t.Errorf("has error: %v", err)
		return
	}

	if res.Data.Page != page {
		t.Errorf("expect page is %d but is %d", page, res.Data.Page)
	}

	if res.Data.PageSize != limit {
		t.Errorf("expect limit is %d but is %d", limit, res.Data.PageSize)
	}

	resbs, err := json.Marshal(res.Data.Result)
	if err != nil {
		t.Errorf("has error: %v", err)
		return
	}
	t.Log(string(resbs))
}

func Test_GetKLine(t *testing.T) {
	client := new(MarketClient).InitByDefault(cfg)

	now := time.Now()
	duration, _ := time.ParseDuration("-30m")
	startTime := now.Add(duration).Unix()
	endTime := now.Unix()

	res, err := client.GetKline("BTCUSDT", "5m", startTime, endTime)

	if err != nil {
		t.Errorf("has error: %v", err)
		return
	}

	if res.TimeData == nil || len(res.TimeData) == 0 {
		t.Errorf("expect TimeData must not empty but empty")
	}

	resbs, err := json.Marshal(res)
	if err != nil {
		t.Errorf("has error: %v", err)
		return
	}
	t.Log(string(resbs))
}

func Test_GetMarketTicker(t *testing.T) {
	client := new(MarketClient).InitByDefault(cfg)

	res, err := client.GetMarketTicker("BTCUSDT")

	if err != nil {
		t.Errorf("has error: %v", err)
		return
	}

	if len(res.LastPrice) == 0 {
		t.Errorf("expect lastPrice not empty but is %s", res.LastPrice)
	}

	resbs, err := json.Marshal(res)
	if err != nil {
		t.Errorf("has error: %v", err)
		return
	}
	t.Log(string(resbs))
}

func Test_GetMarkets(t *testing.T) {
	client := new(MarketClient).InitByDefault(cfg)

	res, err := client.GetMarkets()

	if err != nil {
		t.Errorf("has error: %v", err)
		return
	}

	if len(res) == 0 {
		t.Error("expect res not empty")
	}

	if !res[0].LastPrice.Valid {
		t.Errorf("expect lastPrice not empty but is %s", res[0].LastPrice.Decimal)
	}

	resbs, err := json.Marshal(res)
	if err != nil {
		t.Errorf("has error: %v", err)
		return
	}
	t.Log(string(resbs))
}

func Test_GetTrades(t *testing.T) {
	client := new(MarketClient).InitByDefault(cfg)

	res, err := client.GetTrades("BTCUSDT", 20)

	if err != nil {
		t.Errorf("has error: %v", err)
		return
	}

	if len(res) == 0 {
		t.Error("expect res not empty")
	}

	if len(res) != 20 {
		t.Errorf("expect trades array length is 20, but %d", len(res))
	}

	if res[0].Id <= 0 || len(res[0].FillPrice) == 0 || len(res[0].FillQuantity) == 0 || res[0].Side != "1" && res[0].Side != "2" || res[0].Timestamp <= 0 {

		trade, _ := json.Marshal(res[0])
		t.Errorf("trade data error: %s", string(trade))
	}

	resbs, err := json.Marshal(res)
	if err != nil {
		t.Errorf("has error: %v", err)
		return
	}
	t.Log(string(resbs))
}

func Test_GetMarketDepth(t *testing.T) {
	client := new(MarketClient).InitByDefault(cfg)

	res, err := client.GetMarketDepth("BTCUSDT", 20)

	if err != nil {
		t.Errorf("has error: %v", err)
		return
	}

	/*if len(res.Asks) != 20 {
		t.Errorf("expect res.Asks length is 20, but %d", len(res.Asks))
	}

	if len(res.Bids) != 20 {
		t.Errorf("expect res.Bids length is 20, but %d", len(res.Bids))
	}*/

	if res.Asks[0].PriceD.LessThanOrEqual(decimal.Zero) || res.Asks[0].Quantity <= 0 {

		ask, _ := json.Marshal(res.Asks[0])
		t.Errorf("ask data error: %s", string(ask))
	}

	if res.Bids[0].PriceD.LessThanOrEqual(decimal.Zero) || res.Bids[0].Quantity <= 0 {

		bid, _ := json.Marshal(res.Bids[0])
		t.Errorf("bid data error: %s", string(bid))
	}
	resbs, err := json.Marshal(res)
	if err != nil {
		t.Errorf("has error: %v", err)
		return
	}
	t.Log(string(resbs))
}

func Test_GetOpenOrders(t *testing.T) {
	client := new(TradeClient).InitByDefault(cfg)

	res, err := client.GetOpenOrders("BTCUSDT")

	if err != nil {
		t.Errorf("has error: %v", err)
		return
	}

	if len(res) == 0 {
		t.Error("expect res not empty")
	}

	resbs, err := json.Marshal(res)
	if err != nil {
		t.Errorf("has error: %v", err)
		return
	}
	t.Log(string(resbs))
}

func Test_GetPositions(t *testing.T) {
	client := new(TradeClient).InitByDefault(cfg)

	res, err := client.GetPositions("BTCUSDT")

	if err != nil {
		t.Errorf("has error: %v", err)
		return
	}

	/*if len(res) == 0 {
		t.Error("expect res not empty")
	}

	if len(res[0].ContractCode) == 0 {
		t.Error("expect res[0].contractCode not empty")
	}*/

	resbs, err := json.Marshal(res)
	if err != nil {
		t.Errorf("has error: %v", err)
		return
	}
	t.Log(string(resbs))
}

func Test_Order(t *testing.T) {
	client := new(TradeClient).InitByDefault(cfg)

	orderId, err := client.CreateOrder("BTCUSDT", 1, decimal.NullDecimal{
		Decimal: decimal.NewFromFloat32(1),
		Valid:   true,
	}, 1)

	if err != nil {
		t.Errorf("create order has error: %v", err)
		return
	}
	if orderId <= 0 {
		t.Errorf("expect orderId >0,but %d", orderId)
	}

	time.Sleep(1000 * time.Millisecond)
	openorders, err := client.GetOpenOrders("BTCUSDT")

	if err != nil {
		t.Errorf("get open orders has error: %v", err)
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
		t.Errorf("open order %d not exists", orderId)
		return
	}

	err = client.CancelOrder("BTCUSDT", orderId)

	if err != nil {
		t.Errorf("cancel order has error: %v", err)
		return
	}
	time.Sleep(1000 * time.Millisecond)
	openorders, err = client.GetOpenOrders("BTCUSDT")

	if err != nil {
		t.Errorf("get open orders has error: %v", err)
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
		t.Errorf("open order %d cancelled but exists", orderId)
		return
	}
}

func Test_ConditionOrder(t *testing.T) {
	client := new(TradeClient).InitByDefault(cfg)
	trigPrice := decimal.NewFromFloat32(1.1)

	err := client.CreateConditionOrder("BTCUSDT", 1, "Market", trigPrice, 1, decimal.NullDecimal{})

	if err != nil {
		t.Errorf("create order has error: %v", err)
		return
	}

	time.Sleep(1000 * time.Millisecond)
	orders, err := client.GetConditionOrders(&model.GetConditionOrdersParam{
		ContractCodeList: []string{"BTCUSDT"},
		Side:             2,
		Direct:           1,
		TaskStatusList:   []int{1},
	}, 1, 20)

	if err != nil {
		t.Errorf("get open orders has error: %v", err)
		return
	}
	if orders.TotalCount == 0 {
		t.Error("expect at least a condition order")
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
		t.Error("condition order not exists")
		return
	}

	err = client.CancelConditionOrder("BTCUSDT", taskId)

	if err != nil {
		t.Errorf("cancel condition order has error: %v", err)
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
		t.Errorf("get condition orders has error: %v", err)
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
		t.Errorf("condition order %d cancelled but exists", taskId)
		return
	}
}
