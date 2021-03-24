package client

import (
	"errors"
	"strconv"

	"hopex-hk/go_sdk/core"
	"hopex-hk/go_sdk/core/logging/zaplogger"
	"hopex-hk/go_sdk/model"
)

type HomeClient struct {
	requester core.HttpRequester
}

func (c *HomeClient) InitByDefault(cfg *core.Config) *HomeClient {
	c.requester = new(core.DefaultHttpRequester).Init(cfg, new(zaplogger.ZapLogger))

	return c
}
func (c *HomeClient) Init(requester core.HttpRequester) *HomeClient {
	c.requester = requester

	return c
}

//获取成交统计数据
func (c *HomeClient) GetIndexStatistics() (*model.IndexStatistics, error) {

	res := model.NewGetIndexStatisticsResponse()
	err := c.requester.Get("/api/v1/indexStat", nil, false, res)

	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

//获取公告
//page: current page, starts from 1
func (c *HomeClient) GetIndexNotify(page, limit int) (*model.GetIndexNotifyResponse, error) {
	if page < 1 {
		return nil, errors.New("page must greater than 0")
	}

	var queries = make(map[string]string)

	queries["page"] = strconv.Itoa(page)
	queries["limit"] = strconv.Itoa(limit)

	res := model.NewGetIndexNotifyResponse()

	err := c.requester.Get("/api/v1/index_notify", queries, false, res)

	if err != nil {
		return nil, err
	}

	return res, nil
}
