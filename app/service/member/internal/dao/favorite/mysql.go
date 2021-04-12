package favorite

import (
	"go-zero-study/app/service/member/internal/model"
	"go-zero-study/core/ecode"
	"go-zero-study/core/stores/sqlc"
	"go-zero-study/tools/goctl/model/sql/builderx"
	"strings"
)

var (
	memberFavoriteFieldNames = builderx.FieldNames(&model.Favorite{})
	memberFavoriteRows       = strings.Join(memberFavoriteFieldNames, ",")
)

func (m *Dao) Insert(data model.Favorite) (int64, error) {
	query := `insert into ` + m.table + ` (mid,name) values (?, ?)`
	result,err := m.conn.Exec(query, data.Mid, data.Name)
	if err != nil {
		return 0, err
	}
	id, _ := result.LastInsertId()
	return id, nil
}

func (m *Dao) FindOne(id int64) (*model.Favorite, error) {
	query := `select ` + memberFavoriteRows + ` from ` + m.table + ` where id = ? limit 1`
	var resp model.Favorite
	err := m.conn.QueryRow(&resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ecode.ErrNotFound
	default:
		return nil, err
	}
}

func (m *Dao) Update(data model.Favorite) error {
	query := `update ` + m.table + ` set name = ? where id = ?`
	_, err := m.conn.Exec(query, data.Name, data.Id)
	return err
}

func (m *Dao) Delete(id int64) error {
	query := `delete from ` + m.table + ` where id = ?`
	_, err := m.conn.Exec(query, id)
	return err
}
