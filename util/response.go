package util

type Resp struct {
	Code int16 `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}

func Success(data interface{}) *Resp  {
	return &Resp{
		Code: 200,
		Data: data,
	}
}

func Ok() * Resp  {
	return &Resp{
		Code: 200,
	}
}

func Error(msg string) *Resp  {
	return &Resp{
		Code: 500,
		Msg: msg,
	}
}