package open_ikuai

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/zy84338719/open_ikuai/model"
	"strconv"
	"time"
)

type client struct {
	Domain string
	*resty.Client
}

func newClinet(domain string) *client {
	c := &client{domain, resty.New()}
	return c
}

func (c client) OpenAccessToken(ctx context.Context, clientId string, pubKey []byte, tokenResp *model.TokenResponse) error {
	timeNowSecond := time.Now().Unix()
	request := model.AccessTokenRequest{
		ClientType:  "OPEN_PARTNER",
		ClientId:    clientId,
		RequestTime: timeNowSecond,
		Options: struct {
			TokenLifetime int `json:"token_lifetime"`
		}{
			TokenLifetime: 7200,
		},
		Sign: base64.StdEncoding.EncodeToString(rsaEncrypt([]byte(strconv.Itoa(int(timeNowSecond))), pubKey)),
	}
	return c.Post(ctx, model.OpenAccessToken, request, nil, &tokenResp)
}

func (c client) OpenRefreshToken(ctx context.Context, accessToken string, clientId string, tokenResp *model.TokenResponse) error {
	request := model.RefreshTokenRequest{
		AccessToken: accessToken,
		ClientId:    clientId,
		Options: struct {
			TokenLifetime int `json:"token_lifetime"`
		}{
			TokenLifetime: 7200,
		},
	}
	return c.Post(ctx, model.OpenRefreshToken, request, nil, &tokenResp)
}

// GetDevList 获取授权设备列表
func (c client) GetDevList(ctx context.Context, token string, deviceList *model.DeviceList) error {
	return c.Post(ctx, model.OpenDevList, nil, map[string]string{"IK-Access-Token": token}, deviceList)
}

// GetRouterBaseInfo 获取路由器基本信息
func (c client) GetRouterBaseInfo(ctx context.Context, token string, gwid string, resp *model.Open9XXXResponse) error {
	return c.Post(ctx, fmt.Sprintf(model.Open9XXX, gwid), nil, map[string]string{"IK-Access-Token": token}, resp)
}

// SetRouterWhiteWebList 设置路由器web访问白名单
func (c client) SetRouterWhiteWebList(ctx context.Context, token string, gwid string, param model.WhiteList, resp *model.BaseResponse) error {
	return c.Post(ctx, fmt.Sprintf(model.Open10XXX, gwid), model.Open10XXXRequest{Param: param}, map[string]string{"IK-Access-Token": token}, resp)
}

// BillingRouterGetAccountList 获取计费列表
func (c client) BillingRouterGetAccountList(ctx context.Context, token string, gwid string, resp *model.Open1XXXResponse) error {
	return c.Post(ctx, fmt.Sprintf(model.Open1XXX, gwid), model.Open1XXXRequest{
		Param: model.Param{TYPE: "data"},
	}, map[string]string{"IK-Access-Token": token}, &resp)
}

// BillingRouterUserOffline 用户下线
func (c client) BillingRouterUserOffline(ctx context.Context, token string, gwid string, req model.Open3XXXRequest, resp *model.BaseResponse) error {
	return c.Post(ctx, fmt.Sprintf(model.Open3XXX, gwid), req, map[string]string{"IK-Access-Token": token}, &resp)
}

// BillingRouterGetUserOnlineList 获取在线用户列表
func (c client) BillingRouterGetUserOnlineList(ctx context.Context, token string, gwid string, skip, limit int64, resp *model.Open4XXXResponse) error {
	return c.Post(ctx, fmt.Sprintf(model.Open4XXX, gwid),
		model.Open3XXXRequest{Param: model.Param{TYPE: "data", Skip: skip, Limit: limit}},
		map[string]string{"IK-Access-Token": token}, &resp)
}

// BillingRouterAddUser 添加账号
func (c client) BillingRouterAddUser(ctx context.Context, token string, gwid string, req model.AccountData, resp *model.BaseResponse) error {
	return c.Post(ctx, fmt.Sprintf(model.Open5XXX, gwid),
		model.Open5XXXRequest{Param: req},
		map[string]string{"IK-Access-Token": token}, &resp)
}

// BillingRouterUpdateUser 更新账号
func (c client) BillingRouterUpdateUser(ctx context.Context, token string, gwid string, req model.AccountData, resp *model.BaseResponse) error {
	return c.Post(ctx, fmt.Sprintf(model.Open6XXX, gwid),
		model.Open6XXXRequest{Param: req},
		map[string]string{"IK-Access-Token": token}, &resp)
}

// BillingRouterEnDisAbleUser 启停账号
func (c client) BillingRouterEnDisAbleUser(ctx context.Context, token string, gwid string, id int64, status model.AccountStatus, resp *model.BaseResponse) error {
	return c.Post(ctx, fmt.Sprintf(model.Open7XXX, gwid),
		model.Open7XXXRequest{Param: model.Param{Id: id, Status: status}},
		map[string]string{"IK-Access-Token": token}, &resp)
}

// BillingRouterDeleteUser 删除账号
func (c client) BillingRouterDeleteUser(ctx context.Context, token string, gwid string, id int64, resp *model.BaseResponse) error {
	return c.Post(ctx, fmt.Sprintf(model.Open8XXX, gwid),
		model.Open8XXXRequest{Param: model.Param{Id: id}},
		map[string]string{"IK-Access-Token": token}, &resp)
}

// BillingRouterGetUser 查询账号
func (c client) BillingRouterGetUser(ctx context.Context, token string, gwid string, username string, resp *model.Open11XXXResponse) error {
	return c.Post(ctx, fmt.Sprintf(model.Open11XXX, gwid),
		model.Open11XXXRequest{Param: model.Param{Username: username}},
		map[string]string{"IK-Access-Token": token}, &resp)
}

// BillingRouterAddUserCoupon 添加优惠券
func (c client) BillingRouterAddUserCoupon(ctx context.Context, token string, gwid string, coupon model.CouponParam, resp *model.Open18XXXResponse) error {
	return c.Post(ctx, fmt.Sprintf(model.Open18XXX, gwid),
		model.Open18XXXRequest{Param: coupon},
		map[string]string{"IK-Access-Token": token}, &resp)
}

// BillingRouterDeleteUserCoupon 删除优惠券
func (c client) BillingRouterDeleteUserCoupon(ctx context.Context, token string, gwid string, id int64, resp *model.BaseResponse) error {
	return c.Post(ctx, fmt.Sprintf(model.Open19XXX, gwid),
		model.Open19XXXRequest{Param: model.CouponParam{Id: id}},
		map[string]string{"IK-Access-Token": token}, &resp)
}

// BillingRouterUpdateUserCoupon 编辑优惠券
func (c client) BillingRouterUpdateUserCoupon(ctx context.Context, token string, gwid string, req model.CouponParam, resp *model.BaseResponse) error {
	return c.Post(ctx, fmt.Sprintf(model.Open20XXX, gwid),
		model.Open20XXXRequest{Param: req},
		map[string]string{"IK-Access-Token": token},
		&resp)
}

// BillingRouterGetUserCouponList 获取优惠券列表
func (c client) BillingRouterGetUserCouponList(ctx context.Context, token string, gwid string, skip, limit int64, resp *model.Open21XXXResponse) error {
	return c.Post(ctx, fmt.Sprintf(model.Open21XXX, gwid),
		model.Open21XXXRequest{Param: model.Param{TYPE: "data", Skip: skip, Limit: limit}},
		map[string]string{"IK-Access-Token": token}, &resp)
}

// BillingRouterCustomAuthInfo 推送/更新自定义认证信息
func (c client) BillingRouterCustomAuthInfo(ctx context.Context, token string, gwid string, req model.Draft, resp *model.Open22XXXResponse) error {
	return c.Post(ctx, fmt.Sprintf(model.Open22XXX, gwid),
		model.Open22XXXRequest{Param: req},
		map[string]string{"IK-Access-Token": token},
		&resp)
}

// BillingRouterCustomAuthUserList 获取自定义认证用户列表
func (c client) BillingRouterCustomAuthUserList(ctx context.Context, token string, gwid string, resp *model.Open23XXXResponse) error {
	return c.Post(ctx, fmt.Sprintf(model.Open23XXX, gwid),
		nil, map[string]string{"IK-Access-Token": token},
		&resp)
}

// BillingRouterCustomAuthUserDelete 删除自定义认证信息
func (c client) BillingRouterCustomAuthUserDelete(ctx context.Context, token string, gwid string, username string, resp *model.BaseResponse) error {
	return c.Post(ctx,
		fmt.Sprintf(model.Open24XXX, gwid),
		model.Open24XXXRequest{Param: model.Param{Username: username}},
		map[string]string{"IK-Access-Token": token}, &resp)
}

// BillingRouterCustomAuthUserUpdate 更新自定义认证信息
func (c client) BillingRouterCustomAuthUserUpdate(ctx context.Context, token string, gwid string, req model.Draft, resp *model.BaseResponse) error {
	return c.Post(ctx, fmt.Sprintf(model.Open25XXX, gwid),
		model.Open25XXXRequest{Param: req},
		map[string]string{"IK-Access-Token": token},
		&resp)
}

func (c client) Post(ctx context.Context, url string, request interface{}, headers map[string]string, response interface{}) error {
	if headers == nil {
		headers = make(map[string]string)
	}
	headers["Content-Type"] = "application/json"
	headers["IK-Api-Language"] = "en"
	r := c.R()
	resp, err := r.
		SetResult(response).
		SetContext(ctx).
		SetBody(request).
		SetHeaders(headers).
		Post(c.Domain + url)

	if err != nil {
		return err
	}
	if resp != nil && resp.RawResponse != nil && resp.RawResponse.StatusCode != 200 {
		return fmt.Errorf("请求数据返回状态异常 code %+v", resp.RawResponse.StatusCode)
	}

	return nil
}
