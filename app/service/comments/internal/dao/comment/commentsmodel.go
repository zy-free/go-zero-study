package comment

import (
	"database/sql"
	"fmt"
	sqlBuild "go-zero-study/core/sql-build"
	"go-zero-study/core/stores/sqlc"
	"go-zero-study/core/stores/sqlx"
	"go-zero-study/core/stringx"
	"go-zero-study/tools/goctl/model/sql/builderx"
	"strings"
	"time"
)

var (
	commentsFieldNames          = builderx.FieldNames(&Comments{})
	commentsRows                = strings.Join(commentsFieldNames, ",")
	commentsRowsExpectAutoSet   = strings.Join(stringx.Remove(commentsFieldNames, "id", "create_time", "update_time"), ",")
	commentsRowsWithPlaceHolder = strings.Join(stringx.Remove(commentsFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"

	cacheCommentsIdPrefix = "cache#Comments#id#"
)

type (
	CommentsModel interface {
		Insert(data Comments) (sql.Result, error)
		FindOne(id int64) (*Comments, error)
		BatchFindByRoomId(roomId int64, start int64, size int64) ([]*Comments, error)
		Count(roomId int64) (int64, error)
		Update(data Comments) error
		Delete(id int64) error
	}

	Comments struct {
		Id         int64     `db:"id"`          // 评论主键
		RoomId     int64     `db:"room_id"`     // 房间id
		UserId     int64     `db:"user_id"`     // 用户id
		Nickname   string    `db:"nickname"`    // 用户昵称
		Avatar     string    `db:"avatar"`      // 用户头衔
		Content    string    `db:"content"`     // 评论内存
		CreateTime time.Time `db:"create_time"` // 创建时间
		UpdateTime time.Time `db:"update_time"` // 更新时间
	}
)

func (m *defaultCommentsModel) Insert(data Comments) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, commentsRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.RoomId, data.UserId, data.Nickname, data.Avatar, data.Content)

	return ret, err
}

func (m *defaultCommentsModel) FindOne(id int64) (*Comments, error) {
	commentsIdKey := fmt.Sprintf("%s%v", cacheCommentsIdPrefix, id)
	var resp Comments
	err := m.QueryRow(&resp, commentsIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", commentsRows, m.table)
		return conn.QueryRow(v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultCommentsModel) Update(data Comments) error {
	commentsIdKey := fmt.Sprintf("%s%v", cacheCommentsIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		//query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, commentsRowsWithPlaceHolder)
		query, err := sqlBuild.Update(m.table).
			Set(data.Nickname, "nickname").
			Set(data.Content, "content").String()
		if err != nil {
			return nil, err
		}
		return conn.Exec(query, data.RoomId, data.UserId, data.Nickname, data.Avatar, data.Content, data.Id)
	}, commentsIdKey)
	return err
}

func (m *defaultCommentsModel) Delete(id int64) error {
	commentsIdKey := fmt.Sprintf("%s%v", cacheCommentsIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, commentsIdKey)
	return err
}

func (m *defaultCommentsModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheCommentsIdPrefix, primary)
}

func (m *defaultCommentsModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", commentsRows, m.table)
	return conn.QueryRow(v, query, primary)
}

func (m *defaultCommentsModel) Count(roomId int64) (int64, error) {
	var count int64
	query, err := sqlBuild.Select(m.table).Column("count(*)").Where_(roomId, "room_id").String()
	if err != nil {
		return count, err
	}

	err = m.QueryRowNoCache(&count, query)
	switch err {
	case nil:
		return count, nil
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	default:
		return 0, err
	}
}

func (m *defaultCommentsModel) BatchFindByRoomId(roomId int64, start int64, size int64) ([]*Comments, error) {
	var comments []*Comments
	query, err := sqlBuild.Select(m.table).Column("id").
		Column("room_id").
		Column("user_id").
		Column("avatar").
		Column("nickname").
		Column("content").
		Column("create_time").
		Column("update_time").
		Where_(roomId, "room_id").
		Offset(int(start)).
		Limit(int(size)).
		String()
	if err != nil {
		return nil, err
	}
	err = m.QueryRowsNoCache(&comments, query)
	switch err {
	case nil:
		return comments, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
