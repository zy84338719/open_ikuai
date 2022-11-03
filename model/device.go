package model

const (
	OpenDevList = "/open/dev_list" // 获取授权设备列表
	Open9XXX    = "/open/9/%s"     // 获取路由器基本信息
	Open10XXX   = "/open/10/%s"    // 设置路由器白名单
	Open1XXX    = "/open/1/%s"     // 获取账号列表
	Open3XXX    = "/open/3/%s"     // 在线用户踢下线
	Open4XXX    = "/open/4/%s"     // 获取在线用户列表
	Open5XXX    = "/open/5/%s"     // 添加账号
	Open6XXX    = "/open/6/%s"     // 修改账号
	Open7XXX    = "/open/7/%s"     // 启用/停用账号
	Open8XXX    = "/open/8/%s"     // 删除账户
	Open11XXX   = "/open/11/%s"    // 查询账号
	Open18XXX   = "/open/18/%s"    // 添加优惠券
	Open19XXX   = "/open/19/%s"    // 删除优惠券
	Open20XXX   = "/open/20/%s"    // 编辑优惠券
	Open21XXX   = "/open/21/%s"    // 获取优惠券列表
	Open22XXX   = "/custom/22/%s"  // 推送/更新自定义认证信息
	Open23XXX   = "/custom/23/%s"  // 获取自定义认证用户列表
	Open24XXX   = "/custom/24/%s"  // 删除自定义认证信息
	Open25XXX   = "/custom/25/%s"  // 更新自定义认证信息
)

type DeviceList struct {
	ErrorCode int      `json:"ErrorCode"`
	ErrorMsg  string   `json:"ErrorMsg"`
	Data      []Device `json:"Data"`
	Extra     Extra    `json:"extra"`
}

type Device struct {
	DevId       string `json:"dev_id"`
	DevRemark   string `json:"dev_remark"`
	OwnerMobile string `json:"owner_mobile"`
	OwnerQq     string `json:"owner_qq"`
	OpenApis    []int  `json:"open_apis"`
}

type Extra struct {
	Version     string `json:"version"`
	OtherField1 int    `json:"other_field1"`
	OtherField2 string `json:"other_field2"`
	RowId       int    `json:"RowId"`
}

type Open9XXXResponse struct {
	BaseResponse
	Data RouterBasic `json:"Data"`
}

type RouterBasic struct {
	ApCount          int64  `json:"ap_count"`
	UpSpeed          int64  `json:"up_speed"`
	DownSpeed        int64  `json:"down_speed"`
	TotalUp          int64  `json:"total_up"`
	TotalDown        int64  `json:"total_down"`
	Online           int64  `json:"online"`
	Libver           string `json:"libver"`
	Sysuptime        int64  `json:"sysuptime"`
	RouterVer        string `json:"router_ver"`
	Partner          string `json:"partner"`
	Gwid             string `json:"gwid"`
	Enterprise       string `json:"enterprise"`
	Firmware         string `json:"firmware"`
	Mac              string `json:"mac"`
	ReflexiveAddress string `json:"reflexive_address"`
}

type Open10XXXRequest struct {
	Param WhiteList `json:"param"`
}

type WhiteList struct {
	Whitelist      string `json:"whitelist"`
	WhitelistHttps string `json:"whitelist_https"`
	Whiteip        string `json:"whiteip"`
	NoauthMac      string `json:"noauth_mac"`
}

type Open1XXXRequest struct {
	Param Param `json:"param"`
}

type Open1XXXResponse struct {
	BaseResponse
	Data []AccountData `json:"Data"`
}

type Open3XXXRequest struct {
	Param Param `json:"param"`
}

type Open4XXXRequest struct {
	Param Param `json:"param"`
}

type AccountStatus string

const (
	Up   AccountStatus = "up"
	Down AccountStatus = "down"
)

type Param struct {
	TYPE     string        `json:"TYPE,omitempty"`
	Skip     int64         `json:"skip"`
	Limit    int64         `json:"limit,omitempty"`
	Id       int64         `json:"id,omitempty"`
	Status   AccountStatus `json:"status,omitempty"`
	Username string        `json:"username,omitempty"`
	Mac      string        `json:"mac,omitempty"`
}

type Open4XXXResponse struct {
	BaseResponse
	Data []OnlineUser `json:"Data"`
}

type OnlineUser struct {
	IpAddr       string `json:"ip_addr"`
	CheckVlanRes int    `json:"check_vlan_res"`
	Mac          string `json:"mac"`
	Username     string `json:"username"`
	Ppptype      string `json:"ppptype"`
	Id           int    `json:"id"`
	Webid        int    `json:"webid"`
	AuthTime     int    `json:"auth_time"`
	Session      string `json:"session"`
	Uid          string `json:"uid"`
	IpAddrInt    string `json:"ip_addr_int"`
	Pppdev       string `json:"pppdev"`
	Upload       int    `json:"upload"`
	Download     int    `json:"download"`
	Comment      string `json:"comment"`
	Interface    string `json:"interface"`
	Name         string `json:"name"`
	Expires      int    `json:"expires"`
	Packages     int    `json:"packages"`
	Packname     string `json:"packname"`
	Phone        string `json:"phone"`
}

type Open5XXXRequest struct {
	Param AccountData `json:"param"`
}

type Open6XXXRequest struct {
	Param AccountData `json:"param"`
}

type Open7XXXRequest struct {
	Param Param `json:"param"`
}

type Open8XXXRequest struct {
	Param Param `json:"param"`
}

type Open11XXXRequest struct {
	Param Param `json:"param"`
}

type Open11XXXResponse struct {
	BaseResponse
	Data []AccountData `json:"Data"`
}

type Open18XXXRequest struct {
	Param CouponParam `json:"param"`
}

type CouponParam struct {
	Id       int64  `json:"id"`
	Comment  string `json:"comment"`
	Username string `json:"username"`
	Enabled  string `json:"enabled"`
	Expires  int64  `json:"expires"`
	Timeout  int64  `json:"timeout"`
}

type Open18XXXResponse struct {
	ErrorCode int         `json:"ErrorCode"`
	ErrorMsg  interface{} `json:"ErrorMsg"`
	Data      struct {
	} `json:"Data"`
	Extra Extra `json:"Extra"`
}

type Open19XXXRequest struct {
	Param CouponParam `json:"param"`
}

type BaseResponse struct {
	ErrorCode int         `json:"ErrorCode"`
	ErrorMsg  interface{} `json:"ErrorMsg"`
	Extra     Extra       `json:"Extra"`
}

type Open20XXXRequest struct {
	Param CouponParam `json:"param"`
}

type Open21XXXRequest struct {
	Param Param `json:"param"`
}

type Open21XXXResponse struct {
	BaseResponse
	Data []Coupon `json:"Data"`
}

type Coupon struct {
	Id       int    `json:"id"`
	Comment  string `json:"comment"`
	Username string `json:"username"`
	Enabled  string `json:"enabled"`
	Expires  int    `json:"expires"`
	Timeout  int    `json:"timeout"`
	Used     int    `json:"used"`
}

type Open22XXXRequest struct {
	Param Draft `json:"param"`
}

type Draft struct {
	Username   string `json:"username"`
	Mobile     string `json:"mobile"`
	CardId     string `json:"card_id"`
	AdslId     string `json:"adsl_id"`
	Passport   string `json:"passport"`
	EepId      string `json:"eep_id"`
	RoomNumber string `json:"room_number"`
	Other      string `json:"other"`
	Imsi       string `json:"imsi"`
	Imei       string `json:"imei"`
	Name       string `json:"name"`
	Upload     int    `json:"upload"`
	Download   int    `json:"download"`
	Expires    int    `json:"expires"`
	Comment    string `json:"comment"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type Open22XXXResponse struct {
	BaseResponse
	Data []Draft `json:"Data"`
}

type Open23XXXResponse struct {
	BaseResponse
	Data []Draft `json:"Data"`
}

type Open24XXXRequest struct {
	Param Param `json:"param"`
}

type Open25XXXRequest struct {
	Param Draft `json:"param"`
}
