package client

import (
	"errors"
	"strconv"

	"github.com/hopex-hk/go_sdk/core"
	"github.com/hopex-hk/go_sdk/core/logging/zaplogger"
	"github.com/hopex-hk/go_sdk/model"
)

type AccountClient struct {
	requester core.HttpRequester
}

func (c *AccountClient) InitByDefault(cfg *core.Config) *AccountClient {
	c.requester = new(core.DefaultHttpRequester).Init(cfg, new(zaplogger.ZapLogger))

	return c
}

func (c *AccountClient) Init(requester core.HttpRequester) *AccountClient {
	c.requester = requester

	return c
}

//get account information
//https://hopex-hk.github.io/docs/contract/v1/cn/#hopex-7
func (c *AccountClient) GetUserInfo() (*model.UserInfo, error) {
	res := model.NewGetUserInfoResponse()
	err := c.requester.Get("/api/v1/userinfo", nil, true, res)

	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

//get account wallet information
//doc: https://hopex-hk.github.io/docs/contract/v1/cn/#224168edc1
func (c *AccountClient) GetWalletInfo() (*model.WalletInfo, error) {
	res := model.NewGetWalletResponse()
	err := c.requester.Get("/api/v1/wallet", nil, true, res)

	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

//get account wallet information v2
func (c *AccountClient) GetWalletInfoV2() (*model.WalletInfo, error) {
	res := model.NewGetWalletResponse()
	err := c.requester.Get("/api/v2/wallet", nil, true, res)

	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

//get account history
//doc: https://hopex-hk.github.io/docs/contract/v1/cn/#8e62101ca9
func (c *AccountClient) GetAccountHistory(page, limit int) (*model.AccountHistoryList, error) {
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
	res := model.NewGetAccountHistoryResponse()
	err := c.requester.Get("/api/v1/account_records", queries, true, res)

	if err != nil {
		return nil, err
	}

	return res.Data, nil
}
