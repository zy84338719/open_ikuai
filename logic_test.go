package open_ikuai

import (
	"context"
	"github.com/zy84338719/open_ikuai/model"
	"reflect"
	"testing"
)

var l = NewLogic("-----BEGIN PUBLIC KEY-----\n"+
	""+
	"\n-----END PUBLIC KEY-----", "")

func TestLogic_AccessToken(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    *model.TokenResponse
		wantErr bool
	}{
		{name: "test", args: args{ctx: context.Background()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := l.AccessToken(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("AccessToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccessToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogic_AccountCreation(t *testing.T) {
	type fields struct {
		client   *client
		pubKey   string
		clientId string
		token    *model.TokenResponse
	}
	type args struct {
		ctx  context.Context
		gwid string
		req  model.AccountData
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logic{
				client:   tt.fields.client,
				pubKey:   tt.fields.pubKey,
				clientId: tt.fields.clientId,
				token:    tt.fields.token,
			}
			if err := l.AccountCreation(tt.args.ctx, tt.args.gwid, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("AccountCreation() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLogic_AccountDelete(t *testing.T) {

	type args struct {
		ctx  context.Context
		gwid string
		id   int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "test", args: args{ctx: context.Background(), gwid: "", id: 40}},
	}
	l.InitRam(context.Background())
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := l.AccountDelete(tt.args.ctx, tt.args.gwid, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("AccountDelete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLogic_AccountEdit(t *testing.T) {
	type args struct {
		ctx  context.Context
		gwid string
		req  model.AccountData
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if err := l.AccountEdit(tt.args.ctx, tt.args.gwid, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("AccountEdit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLogic_AccountList(t *testing.T) {

	type args struct {
		ctx  context.Context
		gwid string
	}
	tests := []struct {
		name    string
		args    args
		want    []model.AccountData
		wantErr bool
	}{
		{name: "test", args: args{gwid: "", ctx: context.Background()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l.InitRam(context.Background())
			got, err := l.AccountList(tt.args.ctx, tt.args.gwid)
			if (err != nil) != tt.wantErr {
				t.Errorf("AccountList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccountList() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogic_AccountStatus(t *testing.T) {

	type args struct {
		ctx    context.Context
		gwid   string
		id     int64
		status model.AccountStatus
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "test", args: args{ctx: context.Background(),
			gwid:   "",
			id:     1,
			status: model.Up,
		}},
	}
	l.InitRam(context.Background())
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if err := l.AccountStatus(tt.args.ctx, tt.args.gwid, tt.args.id, tt.args.status); (err != nil) != tt.wantErr {
				t.Errorf("AccountStatus() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLogic_CouponAdding(t *testing.T) {
	type fields struct {
		client   *client
		pubKey   string
		clientId string
		token    *model.TokenResponse
	}
	type args struct {
		ctx  context.Context
		gwid string
		req  model.CouponParam
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logic{
				client:   tt.fields.client,
				pubKey:   tt.fields.pubKey,
				clientId: tt.fields.clientId,
				token:    tt.fields.token,
			}
			if err := l.CouponAdding(tt.args.ctx, tt.args.gwid, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("CouponAdding() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLogic_CouponDelete(t *testing.T) {
	type fields struct {
		client   *client
		pubKey   string
		clientId string
		token    *model.TokenResponse
	}
	type args struct {
		ctx  context.Context
		gwid string
		id   int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logic{
				client:   tt.fields.client,
				pubKey:   tt.fields.pubKey,
				clientId: tt.fields.clientId,
				token:    tt.fields.token,
			}
			if err := l.CouponDelete(tt.args.ctx, tt.args.gwid, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("CouponDelete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLogic_CouponEdit(t *testing.T) {
	type fields struct {
		client   *client
		pubKey   string
		clientId string
		token    *model.TokenResponse
	}
	type args struct {
		ctx  context.Context
		gwid string
		req  model.CouponParam
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logic{
				client:   tt.fields.client,
				pubKey:   tt.fields.pubKey,
				clientId: tt.fields.clientId,
				token:    tt.fields.token,
			}
			if err := l.CouponEdit(tt.args.ctx, tt.args.gwid, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("CouponEdit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLogic_CouponList(t *testing.T) {
	type fields struct {
		client   *client
		pubKey   string
		clientId string
		token    *model.TokenResponse
	}
	type args struct {
		ctx   context.Context
		gwid  string
		skip  int64
		limit int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []model.Coupon
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logic{
				client:   tt.fields.client,
				pubKey:   tt.fields.pubKey,
				clientId: tt.fields.clientId,
				token:    tt.fields.token,
			}
			got, err := l.CouponList(tt.args.ctx, tt.args.gwid, tt.args.skip, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("CouponList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CouponList() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogic_CustomAuthUserDelete(t *testing.T) {
	type fields struct {
		client   *client
		pubKey   string
		clientId string
		token    *model.TokenResponse
	}
	type args struct {
		ctx      context.Context
		gwid     string
		username string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logic{
				client:   tt.fields.client,
				pubKey:   tt.fields.pubKey,
				clientId: tt.fields.clientId,
				token:    tt.fields.token,
			}
			if err := l.CustomAuthUserDelete(tt.args.ctx, tt.args.gwid, tt.args.username); (err != nil) != tt.wantErr {
				t.Errorf("CustomAuthUserDelete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLogic_CustomAuthUserUpdate(t *testing.T) {
	type fields struct {
		client   *client
		pubKey   string
		clientId string
		token    *model.TokenResponse
	}
	type args struct {
		ctx  context.Context
		gwid string
		req  model.Draft
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logic{
				client:   tt.fields.client,
				pubKey:   tt.fields.pubKey,
				clientId: tt.fields.clientId,
				token:    tt.fields.token,
			}
			if err := l.CustomAuthUserUpdate(tt.args.ctx, tt.args.gwid, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("CustomAuthUserUpdate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLogic_FindAccountByUsername(t *testing.T) {

	type args struct {
		ctx      context.Context
		gwid     string
		username string
	}
	tests := []struct {
		name    string
		args    args
		want    []model.AccountData
		wantErr bool
	}{
		{"test", args{ctx: context.Background(), gwid: "", username: "18600321498"}, nil, false},
	}
	l.InitRam(context.Background())
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := l.FindAccountByUsername(tt.args.ctx, tt.args.gwid, tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindAccountByUsername() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindAccountByUsername() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogic_GetCustomAuthUserList(t *testing.T) {
	type fields struct {
		client   *client
		pubKey   string
		clientId string
		token    *model.TokenResponse
	}
	type args struct {
		ctx  context.Context
		gwid string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []model.Draft
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logic{
				client:   tt.fields.client,
				pubKey:   tt.fields.pubKey,
				clientId: tt.fields.clientId,
				token:    tt.fields.token,
			}
			got, err := l.GetCustomAuthUserList(tt.args.ctx, tt.args.gwid)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCustomAuthUserList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCustomAuthUserList() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogic_GetDeviceList(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    []model.Device
		wantErr bool
	}{
		// TODO: Add test cases.
	}

	l.InitRam(context.Background())
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := l.GetDeviceList(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDeviceList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDeviceList() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogic_GetRouterInfo(t *testing.T) {

	type args struct {
		ctx  context.Context
		gwid string
	}
	tests := []struct {
		name    string
		args    args
		want    *model.RouterBasic
		wantErr bool
	}{
		{"test", args{ctx: context.Background(), gwid: ""}, nil, false},
	}
	l.InitRam(context.Background())
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := l.GetRouterInfo(tt.args.ctx, tt.args.gwid)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRouterInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRouterInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogic_OnlineUserKickByMac(t *testing.T) {
	type fields struct {
		client   *client
		pubKey   string
		clientId string
		token    *model.TokenResponse
	}
	type args struct {
		ctx  context.Context
		gwid string
		mac  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logic{
				client:   tt.fields.client,
				pubKey:   tt.fields.pubKey,
				clientId: tt.fields.clientId,
				token:    tt.fields.token,
			}
			if err := l.OnlineUserKickByMac(tt.args.ctx, tt.args.gwid, tt.args.mac); (err != nil) != tt.wantErr {
				t.Errorf("OnlineUserKickByMac() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLogic_OnlineUserKickByUserName(t *testing.T) {
	type fields struct {
		client   *client
		pubKey   string
		clientId string
		token    *model.TokenResponse
	}
	type args struct {
		ctx      context.Context
		gwid     string
		username string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logic{
				client:   tt.fields.client,
				pubKey:   tt.fields.pubKey,
				clientId: tt.fields.clientId,
				token:    tt.fields.token,
			}
			if err := l.OnlineUserKickByUserName(tt.args.ctx, tt.args.gwid, tt.args.username); (err != nil) != tt.wantErr {
				t.Errorf("OnlineUserKickByUserName() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLogic_OnlineUserList(t *testing.T) {

	type args struct {
		ctx   context.Context
		gwid  string
		skip  int64
		limit int64
	}
	tests := []struct {
		name    string
		args    args
		want    []model.OnlineUser
		wantErr bool
	}{
		{name: "test", args: args{ctx: context.Background(), gwid: "", skip: 0, limit: 10}},
	}
	l.InitRam(context.Background())
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := l.OnlineUserList(tt.args.ctx, tt.args.gwid, tt.args.skip, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("OnlineUserList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OnlineUserList() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogic_PushCustomAuthInfo(t *testing.T) {

	type args struct {
		ctx  context.Context
		gwid string
		req  model.Draft
	}
	tests := []struct {
		name    string
		args    args
		want    []model.Draft
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := l.PushCustomAuthInfo(tt.args.ctx, tt.args.gwid, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("PushCustomAuthInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PushCustomAuthInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogic_RefreshToken(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    *model.TokenResponse
		wantErr bool
	}{
		{name: "test"},
	}
	l.InitRam(context.Background())
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := l.RefreshToken(tt.args.ctx, l.token.AccessToken, l.clientId)
			if (err != nil) != tt.wantErr {
				t.Errorf("RefreshToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RefreshToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogic_SetRouterWebWhiteList(t *testing.T) {

	type args struct {
		ctx     context.Context
		gwid    string
		webList model.WhiteList
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "test", args: args{ctx: context.Background(), gwid: "", webList: model.WhiteList{Whiteip: "8.8.8.8", WhitelistHttps: "baidu.com", Whitelist: "baidu.com"}}, wantErr: false},
	}
	l.InitRam(context.Background())
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if err := l.SetRouterWebWhiteList(tt.args.ctx, tt.args.gwid, tt.args.webList); (err != nil) != tt.wantErr {
				t.Errorf("SetRouterWebWhiteList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLogic_refreshToken(t *testing.T) {

	tests := []struct {
		name string
	}{
		{"test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l.InitRam(context.Background())
			l.refreshToken()
		})
	}
}
