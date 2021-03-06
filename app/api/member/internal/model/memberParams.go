// Code generated by goctl. DO NOT EDIT.
package model



// options=first|second|third
// optional 可选的，没有此选项则该字段必填
type GetMemberByIDReq struct {
	ID int64 `path:"id,default=3,range=[1:]"`
}

type GetMemberByPhoneReq struct {
	Phone string `form:"phone"`
}

type GetMemberResp struct {
	Member
}

type GetMemberMaxAgeResq struct {
	Age int64 `json:"age"`
}

type QueryMemberByNameReq struct {
	Name string `form:"name"`
}

type QueryMemberByNameResq struct {
	Members []Member `json:"members"`
}

type QueryMemberByIdsReq struct {
	IDs string `form:"ids"`
}

type QueryMemberByIdsResp struct {
	List map[int64]*Member `json:"list"`
}

type LoginMemberResp struct {
	AccessToken  string `json:"access_token"`
	AccessExpire int64  `json:"access_expire"`
	RefreshAfter int64  `json:"refresh_after"`
}

type AddMemberReq struct {
	Phone   string `json:"phone"`
	Name    string `json:"name"`
	Age     int64  `json:"age,default=-1"`
	Address string `json:"address,optional"`
}

type AddMemberResp struct {
	ID int64 `json:"id"`
}

type BatchAddMemberReq struct {
	Args []AddMemberReq `json:"args"`
}

type InitMemberReq struct {
	Phone   string `json:"phone"`
	Name    string `json:"name"`
	Age     int64  `json:"age,default=-1"`
	Address string `json:"address,optional"`
}

type InitMemberResp struct {
	ID int64 `json:"id"`
}

// 参数可选，但是sql目前不支持
//type UpdateMemberReq struct {
//	ID      int64  `path:"id"`
//	Phone   string `json:"phone"`
//	Name    string `json:"name,default=defaut,optional"`
//	Age     int64  `json:"age,default=-1"`
//	Address string `json:"address,default=-1,optional"`
//}

type UpdateMemberReq struct {
	ID      int64  `path:"id"`
	Phone   string `json:"phone"`
	Name    string `json:"name"`
	Age     int64  `json:"age"`
	Address string `json:"address"`
}

type UpdateSomeMemberReq struct {
	ID      int64  `path:"id"`
	Phone   string `json:"phone,default=-1"`
	Name    string `json:"name,default=-1"`
	Age     int64  `json:"age,default=-1"`
	Address string `json:"address,default=-1"`
}

type SetMemberReq struct {
	ID      int64  `path:"id"`
	Phone   string `json:"phone"`
	Name    string `json:"name"`
	Age     int64  `json:"age"`
	Address string `json:"address"`
}

type SortMember struct {
	ID       int64 `json:"id"`
	OrderNum int64 `json:"order_num"`
}

type SortMemberReq struct {
	Args []SortMember `json:"args"`
}

type DelMemberReq struct {
	ID int64 `path:"id"`
}
