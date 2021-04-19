package model

import "strconv"

type Member struct {
	ID        int64  `json:"id"`
	Phone     string `json:"phone"`
	Name      string `json:"name"`
	Age       int64  `json:"age"`
	Address   string `json:"address"`
	OrderNum  int64  `json:"order_num"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

// 贫血模型

// 判断用户是否old
func (v *Member) IsOld() bool {
	return v.Age >= 35
}

func MemberExportFields() []string {
	return []string{"ID", "手机号", "米子", "年龄", "地址"}
}

func (v *Member) ExportStrings() []string {
	return []string{
		strconv.FormatInt(v.ID, 10),
		v.Phone,
		v.Name,
		"年龄：" + strconv.FormatInt(v.Age, 10),
		v.Address,
	}
}
