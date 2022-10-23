package util

type Resp struct {
	Code int16 `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}

func Success(data interface{}) *Resp  {
	return &Resp{
		Code: 200,
		Msg: "OK",
		Data: data,
	}
}

func Error(msg string) *Resp  {
	return &Resp{
		Code: 500,
		Msg: msg,
	}
}