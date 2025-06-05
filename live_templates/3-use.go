package live_templates

//
//import (
//	"context"
//	"fmt"
//)
//
//func sync3Use(ctx context.Context) (3UseList []share.3Use, err error) {
//
//	var (
//		resp = &3UseResponse{}
//	)
//	// 发起 h1 请求
//	respBody, err := http1.EasyPostUnMarshal(ctx, 50,
//		config.GetIris3UseUrl(), share.ContentTypeJson, map[string]string{"Authorization": config.GetIrisKey()}, nil)
//
//	if err != nil {
//		err = fmt.Errorf("3Use EasyPostUnMarshal err: %s", err.Error())
//		return
//	}
//
//	err = resp.UnmarshalJSON(respBody)
//	if err != nil {
//		err = fmt.Errorf("3Use UnmarshalJSON err: %s", err.Error())
//		return
//	}
//
//	3UseList = resp.Data
//
//	return
//}
//
//func format3Use(ctx context.Context) {
//
//}
//
////easyjson:json
//type 3UseResponse struct {
//	Data []share.3Use `json:"data"`
//}
