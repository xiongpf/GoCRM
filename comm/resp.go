package comm

/*
"statusCode":"200",
"message":"操作成功", "navTabId":"",
"rel":"", "callbackType":"closeCurrent", "forwardUrl":""
*/


//删除后刷新
func RespJson_Delete() map[string]string {
	return map[string]string{"statusCode": "200", "message": "删除成功", "navTabId": "", "rel": "", "callbackType": "", "forwardUrl": "", "confirmMsg": ""}
}

// 适用于 add update 表单
//关闭Dialog或NavTab并且刷新
func RespJsonRefresh_Close(navTableId string) map[string]string {
	return map[string]string{"statusCode": "200", "message": "操作成功", "navTabId": navTableId, "rel": "", "callbackType": "closeCurrent", "forwardUrl": "", "confirmMsg": ""}
}
