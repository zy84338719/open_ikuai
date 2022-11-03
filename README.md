# Open IKuai SDK

#### https://open.ikuai8.com/
Open IKuai 3.0 SDK，封装了部分爱快的接口。
https://open.ikuai8.com/doc/3.0/

### 依赖

```shell
go get -u https://github.com/zy84338719/open_ikuai
```

### 使用

```go
var l = NewLogic("-----BEGIN PUBLIC KEY-----\n"+
""+
"\n-----END PUBLIC KEY-----", "")

l.InitRam(context.Background())

got, err := l.GetRouterInfo(context.Background(), "gwid")
log.Printf("%+v %+v", got, err)
```

