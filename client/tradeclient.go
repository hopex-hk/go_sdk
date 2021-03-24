package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/shopspring/decimal"
	"hopex-hk/go_sdk/core"
	"hopex-hk/go_sdk/core/logging/zaplogger"
	coremodel "hopex-hk/go_sdk/core/model"
	"hopex-hk/go_sdk/model"
)

type TradeClient struct {
	requester core.HttpRequester
}

func (c *TradeClient) InitByDefault(cfg *core.Config) *TradeClient {
	c.requester = new(core.DefaultHttpRequester).Init(cfg, new(zaplogger.ZapLogger))

	return c
}
func (c *TradeClient) Init(requester core.HttpRequester) *TradeClient {
	c.requester = requester

	return c
}

//get open orders
func (c *TradeClient) GetOpenOrders(contractCode string) ([]*model.OpenOrder, error) {
	if len(contractCode) == 0 {
		return nil, errors.New("contractCode must not empty")
	}

	queries := map[string]string{
		"contractCode": contractCode,
	}

	res := model.NewGetOpenOrdersResponse()
	err := c.requester.Get("/api/v1/order_info", queries, true, res)

	if err != nil {
		return nil, err
	}
	return res.Data, nil
}

//get positions
func (c *TradeClient) GetPositions(contractCode string) ([]*model.Position, error) {
	if len(contractCode) == 0 {
		return nil, errors.New("contractCode must not empty")
	}

	queries := map[string]string{
		"contractCode": contractCode,
	}

	res := model.NewGetPositionsResponse()
	err := c.requester.Get("/api/v1/position", queries, true, res)

	if err != nil {
		return nil, err
	}
	return res.Data, nil
}

//cancel condition order
//doc: https://hopex-hk.github.io/docs/contract/v1/cn/#6a6d683558
func (c *TradeClient) CancelConditionOrder(contractCode string, taskId int64) error {
	if len(contractCode) == 0 {
		return errors.New("contractCode must not empty")
	}
	if taskId <= 0 {
		return errors.New("taskId must > 0")
	}

	param := &coremodel.ApiRequestModel{
		Param: &struct {
			ContractCode string `json:"contractCode"`
			TaskId       int64  `json:"taskId"`
		}{contractCode, taskId},
	}
	body, err := json.Marshal(param)
	if err != nil {
		return err
	}

	return c.requester.Post("/api/v1/cancel_condition_order", body, nil, true, coremodel.NewBoolResponse())
}

//cancel open order
//doc: https://hopex-hk.github.io/docs/contract/v1/cn/#11d9e25be8
func (c *TradeClient) CancelOrder(contractCode string, orderId int64) error {
	if len(contractCode) == 0 {
		return errors.New("contractCode must not empty")
	}
	if orderId <= 0 {
		return errors.New("orderId must > 0")
	}

	queries := map[string]string{
		"contractCode": contractCode,
		"orderId":      strconv.FormatInt(orderId, 10),
	}

	return c.requester.Get("/api/v1/cancel_order", queries, true, coremodel.NewBoolResponse())
}

//put condition order
//doc: https://hopex-hk.github.io/docs/contract/v1/cn/#f9e87fb381
func (c *TradeClient) CreateConditionOrder(contractCode string, side int, _type string, trigPrice decimal.Decimal, expectedQuantity int, expectedPrice decimal.NullDecimal) error {
	if len(contractCode) == 0 {
		return errors.New("contractCode must not empty")
	}

	param := &coremodel.ApiRequestModel{
		Param: &struct {
			ContractCode     string              `json:"contractCode"`
			Side             int                 `json:"side"`
			Type             string              `json:"type"`
			TrigPrice        decimal.Decimal     `json:"trigPrice"`
			ExpectedQuantity int                 `json:"expectedQuantity"`
			ExpectedPrice    decimal.NullDecimal `json:"expectedPrice"`
		}{contractCode, side, _type, trigPrice, expectedQuantity, expectedPrice},
	}
	body, err := json.Marshal(param)
	if err != nil {
		return err
	}

	return c.requester.Post("/api/v1/condition_order", body, nil, true, coremodel.NewBoolResponse())
}

func isValidOrderSide(side int) bool {
	return side >= 1 && side <= 4
}

//put order
//doc: https://hopex-hk.github.io/docs/contract/v1/cn/#73292b0f57
func (c *TradeClient) CreateOrder(contractCode string, side int, orderPrice decimal.NullDecimal, orderQuantity int) (int64, error) {
	if len(contractCode) == 0 {
		return 0, errors.New("contractCode must not empty")
	}
	if orderQuantity <= 0 {
		return 0, errors.New("orderQuantity must >0")
	}
	if orderPrice.Valid && orderPrice.Decimal.LessThanOrEqual(decimal.Zero) {
		return 0, errors.New("orderPrice must >0")
	}
	if !isValidOrderSide(side) {
		return 0, errors.New("invalid side.")
	}

	param := &coremodel.ApiRequestModel{
		Param: &struct {
			ContractCode  string              `json:"contractCode"`
			Side          int                 `json:"side"`
			OrderPrice    decimal.NullDecimal `json:"orderPrice"`
			OrderQuantity int                 `json:"orderQuantity"`
		}{contractCode, side, orderPrice, orderQuantity},
	}
	body, err := json.Marshal(param)
	if err != nil {
		return 0, err
	}
	res := coremodel.NewInt64Response()
	err = c.requester.Post("/api/v1/order", body, nil, true, res)
	if err != nil {
		return 0, err
	}
	return res.Data, nil
}

//get condition orders
//doc: https://hopex-hk.github.io/docs/contract/v1/cn/#4d2cea97f7
func (c *TradeClient) GetConditionOrders(param *model.GetConditionOrdersParam, page, limit int) (*model.ConditionOrderList, error) {
	if page < 1 {
		return nil, errors.New("page must >=1")
	}
	if limit <= 0 {
		return nil, errors.New("limit must >0")
	}

	queries := map[string]string{
		"page":  strconv.Itoa(page),
		"limit": strconv.Itoa(limit),
	}
	req := &coremodel.ApiRequestModel{
		Param: param,
	}
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	res := model.NewGetConditionOrdersResponse()
	err = c.requester.Post("/api/v1/condition_order_info", body, queries, true, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

//get condition orders
//doc: https://hopex-hk.github.io/docs/contract/v1/cn/#4212d667d1
func (c *TradeClient) GetHistoryOrders(param *model.GetHistoryOrdersParam, page, limit int) (*model.HistoryOrderList, error) {
	if page < 1 {
		return nil, errors.New("page must >=1")
	}
	if limit <= 0 {
		return nil, errors.New("limit must >0")
	}

	queries := map[string]string{
		"page":  strconv.Itoa(page),
		"limit": strconv.Itoa(limit),
	}
	req := &coremodel.ApiRequestModel{
		Param: param,
	}
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	res := model.NewGetHistoryOrdersResponse()
	err = c.requester.Post("/api/v1/order_history", body, queries, true, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

//get liquidation histories
//doc: https://hopex-hk.github.io/docs/contract/v1/cn/#17307ac747
func (c *TradeClient) GetLiquidationHistories(param *model.GetLiquidationHistoryParam, page, limit int) (*model.LiquidationHistoryList, error) {
	if page < 1 {
		return nil, errors.New("page must >=1")
	}
	if limit <= 0 {
		return nil, errors.New("limit must >0")
	}

	queries := map[string]string{
		"page":  strconv.Itoa(page),
		"limit": strconv.Itoa(limit),
	}
	req := &coremodel.ApiRequestModel{
		Param: param,
	}
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	res := model.NewGetLiquidationHistoryResponse()
	err = c.requester.Post("/api/v1/liquidation_history", body, queries, true, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

//get liquidation histories
//doc: https://hopex-hk.github.io/docs/contract/v1/cn/#ed645b50df
func (c *TradeClient) GetOrderParameter(contractCode string) (map[string]interface{}, error) {
	if len(contractCode) == 0 {
		return nil, errors.New("contractCode must not empty")
	}

	req := &coremodel.ApiRequestModel{
		Param: &struct {
			ContractCode string `json:"contractCode"`
		}{contractCode},
	}
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	res := model.NewGetOrderParameterResponse()
	err = c.requester.Post("/api/v1/get_orderParas", body, nil, true, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

//set leverage
//doc: https://hopex-hk.github.io/docs/contract/v1/cn/#24da12c248
func (c *TradeClient) SetLeverage(contractCode string, direct int, leverage float32) (decimal.Decimal, error) {
	if len(contractCode) == 0 {
		return decimal.Zero, errors.New("contractCode must not empty")
	}
	queries := map[string]string{
		"contractCode": contractCode,
		"direct":       strconv.Itoa(direct),
		"leverage":     fmt.Sprintf("%.2f", leverage),
	}

	res := coremodel.NewDecimalResponse()
	err := c.requester.Get("/api/v1/set_leverage", queries, true, res)
	if err != nil {
		return decimal.Zero, err
	}
	return res.Data, nil
}
