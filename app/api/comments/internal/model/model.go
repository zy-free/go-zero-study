package model

type Comment struct {
	Id         int    `json:"id"`
	UserId     int    `json:"user_id"`
	RoomId     int    `json:"room_id"`
	Nickname   string `json:"nickname"`
	Avatar     string `json:"avatar"`
	Content    string `json:"content"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

type CreateCommentReq struct {
	Content string `json:"content"`
}

type CreateCommentResp struct {
}

type BatchQueryCommentsReq struct {
	RoomId int `path:"room_id"`
	Page   int `from:"page,range=[1:9999999]"`
	Size   int `form:"size,range=[1:30]"`
}

type BathQueryCommentsResp struct {
	TotalPage int `json:"total_page"`
	TotalSize int `json:"total_size"`
	Comments  []*Comment
}
