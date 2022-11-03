package model

const (
	OpenAccessToken  = "/open/access_token"
	OpenRefreshToken = "/open/refresh_token"
)

//request

type AccessTokenRequest struct {
	ClientType  string `json:"client_type"`
	ClientId    string `json:"client_id"`
	RequestTime int64  `json:"request_time"`
	Options     struct {
		TokenLifetime int `json:"token_lifetime"`
	} `json:"options"`
	Sign string `json:"sign"`
}

//response

type TokenResponse struct {
	AccessToken             string `json:"access_token"`
	TokenCreateTime         int    `json:"token_create_time"`
	TokenExcessTime         int    `json:"token_excess_time"`
	TokenExpireAtServerTime string `json:"token_expire_at_server_time"`
}

type RefreshTokenRequest struct {
	AccessToken string `json:"access_token"`
	ClientId    string `json:"client_id"`
	Options     struct {
		TokenLifetime int `json:"token_lifetime"`
	} `json:"options"`
}
