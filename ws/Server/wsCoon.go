package ws

//RspBody ws请求响应body体
type RspBody struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

//ReqBody ws请求body体
type ReqBody struct {
	Path string `json:"path"` //路由key
	Msg  string `json:"msg"`  //请求的消息json数据
}

type WsCoon interface {
	SetProperty(key string, value interface{})
	GetProperty(key string) (interface{}, error)
	RemoveProperty(key string)
	Push(data interface{})
}
