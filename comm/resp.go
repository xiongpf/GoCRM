package comm

/*
"statusCode":"200",
"message":"操作成功", "navTabId":"",
"rel":"", "callbackType":"closeCurrent", "forwardUrl":""
*/

type DwzResp struct {
	StatusCode   string `json:"statusCode"`
	Message      string `json:"message"`
	NavTabId     string `json:"navTabId"`
	Rel          string `json:"rel"`
	CallBackType string `json:"callbackType"`
	ForwardUrl   string `json:"forwardUrl"`
	ConfirmMsg   string `json:"confirmMsg"`
}
