package open_ikuai

import (
	"context"
	"fmt"
	"github.com/zy84338719/open_ikuai/model"
	"time"
)

type Logic struct {
	client   *client
	pubKey   string
	clientId string
	token    *model.TokenResponse
}

func NewLogic(pubKey, clientId string) *Logic {
	return &Logic{client: newClinet("https://open.ikuai8.com/api/v3"), pubKey: pubKey, clientId: clientId}
}

// InitRam 初始化到内存
func (l *Logic) InitRam(ctx context.Context) error {
	response := model.TokenResponse{}
	err := l.client.OpenAccessToken(ctx, l.clientId, []byte(l.pubKey), &response)
	if err != nil {
		return err
	}
	l.token = &response
	go l.refreshToken()
	return nil
}

func (l *Logic) refreshToken() {
	for range time.After(30 * time.Minute) {
		response := model.TokenResponse{}
		err := l.client.OpenRefreshToken(context.Background(), l.token.AccessToken, l.clientId, &response)
		if err != nil {
			return
		}
		l.token = &response
	}
}

// AccessToken token 获取
func (l *Logic) AccessToken(ctx context.Context) (*model.TokenResponse, error) {
	response := model.TokenResponse{}
	err := l.client.OpenAccessToken(ctx, l.clientId, []byte(l.pubKey), &response)
	if err != nil {
		return nil, err
	}
	l.token = &response
	return &response, err
}

// RefreshToken token 刷新
func (l *Logic) RefreshToken(ctx context.Context, accessToken, clientId string) (*model.TokenResponse, error) {
	response := model.TokenResponse{}
	err := l.client.OpenRefreshToken(context.Background(), accessToken, clientId, &response)
	if err != nil {
		return nil, err
	}
	l.token = &response
	return &response, err
}

// GetDeviceList 获取设备列表
func (l *Logic) GetDeviceList(ctx context.Context) ([]model.Device, error) {
	response := model.DeviceList{}
	err := l.client.GetDevList(ctx, l.token.AccessToken, &response)
	if response.ErrorCode%10000 != 0 {
		return nil, fmt.Errorf("%+v", response.ErrorMsg)
	}
	return response.Data, err
}

// GetRouterInfo 获取路由信息
func (l *Logic) GetRouterInfo(ctx context.Context, gwid string) (*model.RouterBasic, error) {
	response := model.Open9XXXResponse{}
	err := l.client.GetRouterBaseInfo(ctx, l.token.AccessToken, gwid, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrorCode%10000 != 0 {
		return nil, fmt.Errorf("%+v", response.ErrorMsg)
	}
	return &response.Data, nil
}

// SetRouterWebWhiteList 获取网络白名单
func (l *Logic) SetRouterWebWhiteList(ctx context.Context, gwid string, webList model.WhiteList) error {
	response := model.BaseResponse{}
	err := l.client.SetRouterWhiteWebList(ctx, l.token.AccessToken, gwid, webList, &response)
	if err != nil {
		return err
	}
	if response.ErrorCode%10000 != 0 {
		return fmt.Errorf("%+v", response.ErrorMsg)
	}
	return nil
}

// AccountList 获取账号列表
func (l *Logic) AccountList(ctx context.Context, gwid string) ([]model.AccountData, error) {
	response := model.Open1XXXResponse{}
	err := l.client.BillingRouterGetAccountList(ctx, l.token.AccessToken, gwid, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrorCode%10000 != 0 {
		return nil, fmt.Errorf("%+v", response.ErrorMsg)
	}
	return response.Data, nil
}

// OnlineUserKickByMac 强制用户下线ByMac
func (l *Logic) OnlineUserKickByMac(ctx context.Context, gwid string, mac string) error {
	response := model.BaseResponse{}
	err := l.client.BillingRouterUserOffline(ctx, l.token.AccessToken, gwid, model.Open3XXXRequest{Param: model.Param{Mac: mac}}, &response)
	if err != nil {
		return err
	}
	if response.ErrorCode%10000 != 0 {
		return fmt.Errorf("%+v", response.ErrorMsg)
	}
	return nil
}

// OnlineUserKickByUserName 强制用户下线ByUserName
func (l *Logic) OnlineUserKickByUserName(ctx context.Context, gwid string, username string) error {
	response := model.BaseResponse{}
	err := l.client.BillingRouterUserOffline(ctx, l.token.AccessToken, gwid, model.Open3XXXRequest{Param: model.Param{Username: username}}, &response)
	if err != nil {
		return err
	}
	if response.ErrorCode%10000 != 0 {
		return fmt.Errorf("%+v", response.ErrorMsg)
	}
	return nil
}

// OnlineUserList 用户在线列表
func (l *Logic) OnlineUserList(ctx context.Context, gwid string, skip, limit int64) ([]model.OnlineUser, error) {
	response := model.Open4XXXResponse{}
	err := l.client.BillingRouterGetUserOnlineList(ctx, l.token.AccessToken, gwid, skip, limit, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrorCode%10000 != 0  {
		return nil, fmt.Errorf("%+v", response.ErrorMsg)
	}
	return response.Data, nil
}

// AccountCreation 创建用户
func (l *Logic) AccountCreation(ctx context.Context, gwid string, req model.AccountData) error {
	response := model.BaseResponse{}
	err := l.client.BillingRouterAddUser(ctx, l.token.AccessToken, gwid, req, &response)
	if err != nil {
		return err
	}
	if response.ErrorCode%10000 != 0  {
		return fmt.Errorf("%+v", response.ErrorMsg)
	}
	return nil
}

// AccountEdit 账户编辑
func (l *Logic) AccountEdit(ctx context.Context, gwid string, req model.AccountData) error {
	response := model.BaseResponse{}
	err := l.client.BillingRouterUpdateUser(ctx, l.token.AccessToken, gwid, req, &response)
	if err != nil {
		return err
	}
	if response.ErrorCode%10000 != 0  {
		return fmt.Errorf("%+v", response.ErrorMsg)
	}
	return nil
}

// AccountStatus 账户状态
func (l *Logic) AccountStatus(ctx context.Context, gwid string, id int64, status model.AccountStatus) error {
	response := model.BaseResponse{}
	err := l.client.BillingRouterEnDisAbleUser(ctx, l.token.AccessToken, gwid, id, status, &response)
	if err != nil {
		return err
	}
	if response.ErrorCode%10000 != 0 {
		return fmt.Errorf("%+v", response.ErrorMsg)
	}
	return nil
}

// AccountDelete 账户删除
func (l *Logic) AccountDelete(ctx context.Context, gwid string, id int64) error {
	response := model.BaseResponse{}
	err := l.client.BillingRouterDeleteUser(ctx, l.token.AccessToken, gwid, id, &response)
	if err != nil {
		return err
	}
	if response.ErrorCode%10000 != 0 {
		return fmt.Errorf("%+v", response.ErrorMsg)
	}
	return nil
}

// FindAccountByUsername 查看用户
func (l *Logic) FindAccountByUsername(ctx context.Context, gwid string, username string) ([]model.AccountData, error) {
	response := model.Open11XXXResponse{}
	err := l.client.BillingRouterGetUser(ctx, l.token.AccessToken, gwid, username, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrorCode%10000 != 0 {
		return nil, fmt.Errorf("%+v", response.ErrorMsg)
	}
	return response.Data, nil
}

// CouponAdding 添加优惠券
func (l *Logic) CouponAdding(ctx context.Context, gwid string, req model.CouponParam) error {
	response := model.Open18XXXResponse{}

	err := l.client.BillingRouterAddUserCoupon(ctx, l.token.AccessToken, gwid, req, &response)
	if err != nil {
		return err
	}
	if response.ErrorCode%10000 != 0  {
		return fmt.Errorf("%+v", response.ErrorMsg)
	}
	return nil
}

// CouponDelete 删除优惠券
func (l *Logic) CouponDelete(ctx context.Context, gwid string, id int64) error {
	response := model.BaseResponse{}

	err := l.client.BillingRouterDeleteUserCoupon(ctx, l.token.AccessToken, gwid, id, &response)
	if err != nil {
		return err
	}
	if response.ErrorCode%10000 != 0 {
		return fmt.Errorf("%+v", response.ErrorMsg)
	}
	return nil
}

// CouponEdit 编辑优惠券
func (l *Logic) CouponEdit(ctx context.Context, gwid string, req model.CouponParam) error {
	response := model.BaseResponse{}

	err := l.client.BillingRouterUpdateUserCoupon(ctx, l.token.AccessToken, gwid, req, &response)
	if err != nil {
		return err
	}
	if response.ErrorCode%10000 != 0  {
		return fmt.Errorf("%+v", response.ErrorMsg)
	}
	return nil
}

// CouponList 优惠券列表
func (l *Logic) CouponList(ctx context.Context, gwid string, skip, limit int64) ([]model.Coupon, error) {
	response := model.Open21XXXResponse{}

	err := l.client.BillingRouterGetUserCouponList(ctx, l.token.AccessToken, gwid, skip, limit, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrorCode%10000 != 0 {
		return nil, fmt.Errorf("%+v", response.ErrorMsg)
	}
	return response.Data, nil
}

// PushCustomAuthInfo 推送更新自定义认证信息
func (l *Logic) PushCustomAuthInfo(ctx context.Context, gwid string, req model.Draft) ([]model.Draft, error) {
	response := model.Open22XXXResponse{}

	err := l.client.BillingRouterCustomAuthInfo(ctx, l.token.AccessToken, gwid, req, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrorCode%10000 != 0 {
		return nil, fmt.Errorf("%+v", response.ErrorMsg)
	}
	return response.Data, nil
}

// GetCustomAuthUserList 获取自定义认证用户列表
func (l *Logic) GetCustomAuthUserList(ctx context.Context, gwid string) ([]model.Draft, error) {
	response := model.Open23XXXResponse{}

	err := l.client.BillingRouterCustomAuthUserList(ctx, l.token.AccessToken, gwid, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrorCode%10000 != 0 {
		return nil, fmt.Errorf("%+v", response.ErrorMsg)
	}
	return response.Data, nil
}

// CustomAuthUserDelete 删除自定义认证信息
func (l *Logic) CustomAuthUserDelete(ctx context.Context, gwid string, username string) error {
	response := model.BaseResponse{}
	err := l.client.BillingRouterCustomAuthUserDelete(ctx, l.token.AccessToken, gwid, username, &response)
	if err != nil {
		return err
	}
	if response.ErrorCode%10000 != 0 {
		return fmt.Errorf("%+v", response.ErrorMsg)
	}
	return nil
}

// CustomAuthUserUpdate 更新自定义认证信息
func (l *Logic) CustomAuthUserUpdate(ctx context.Context, gwid string, req model.Draft) error {
	response := model.BaseResponse{}
	err := l.client.BillingRouterCustomAuthUserUpdate(ctx, l.token.AccessToken, gwid, req, &response)
	if err != nil {
		return err
	}
	if response.ErrorCode%10000 != 0 {
		return fmt.Errorf("%+v", response.ErrorMsg)
	}
	return nil
}
