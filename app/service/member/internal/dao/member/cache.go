package model

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"go-zero-study/app/service/member/internal/model"
	"go-zero-study/core/ecode"
	"go-zero-study/core/stores/sqlc"
	"go-zero-study/core/stores/sqlx"
	"go-zero-study/core/stringx"
	"go-zero-study/tools/goctl/model/sql/builderx"
	"strconv"
	"strings"
)

var (
	memberRows = strings.Join(builderx.FieldNames(&model.Member{}), ",")

	cacheMemberPhonePrefix = "member:phone:"
	cacheMemberIdPrefix    = "member:id:"
)

const (
	_shard = 100
)

// 分表命名:表名+hit
func (dao *Dao) memberHit(id int64) string {
	return fmt.Sprintf("%s_%d", dao.table, id%_shard)
}

// 批量创建
func (dao *Dao) BatchAddMember(args []model.Member) (affectRow int64, err error) {
	sql := `INSERT INTO member (phone,name,age,address) VALUES `
	var valueString []string
	var valueArgs []interface{}
	for _, arg := range args {
		valueString = append(valueString, "(?,?,?,?)")
		valueArgs = append(valueArgs, arg.Phone, arg.Name, arg.Age, arg.Address)
	}
	result, err := dao.ExecNoCache(sql+strings.Join(valueString, ","), valueArgs...)
	if err != nil {
		return 0, errors.Wrapf(err, "BatchAddMember arg(%v)", args)
	}
	affectRow, _ = result.RowsAffected()
	return
}

func (dao *Dao) Init(data model.Member) (err error) {
	query := `insert ignore into ` + dao.table + ` (phone,name,age,address) values (?, ?, ?, ?)`
	_, err = dao.ExecNoCache(query, data.Phone, data.Name, data.Age, data.Address)
	return
}

func (dao *Dao) Insert(data model.Member) (id int64, err error) {
	query := `insert into ` + dao.table + ` (phone,name,age,address) values (?, ?, ?, ?)`
	result, err := dao.ExecNoCache(query, data.Phone, data.Name, data.Age, data.Address)
	if err != nil {
		return 0, err
	}
	id, _ = result.LastInsertId()
	return id, nil
}

func (dao *Dao) FindOne(id int64) (member *model.Member, err error) {
	memberIdKey := fmt.Sprintf("%s%v", cacheMemberIdPrefix, id)
	member = &model.Member{}
	err = dao.QueryRow(member, memberIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := `select ` + memberRows + ` from ` + dao.table + ` where id = ? limit 1`
		return conn.QueryRow(v, query, id)
	})

	switch err {
	case nil:
		return member, nil
	case sqlc.ErrNotFound:
		return nil, ecode.ErrNotFound
	default:
		return nil, err
	}
}

func (dao *Dao) FindOneByPhone(phone string) (member *model.Member, err error) {
	memberPhoneKey := fmt.Sprintf("%s%v", cacheMemberPhonePrefix, phone)
	member = &model.Member{}
	err = dao.QueryRowIndex(member, memberPhoneKey, func(primary interface{}) string {
		return fmt.Sprintf("%s%v", cacheMemberIdPrefix, primary)
	}, func(conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := `select ` + memberRows + ` from ` + dao.table + ` where phone = ? limit 1`
		if err := conn.QueryRow(member, query, phone); err != nil {
			return nil, err
		}
		return member.Id, nil
	}, func(conn sqlx.SqlConn, v, primary interface{}) error {
		query := `select ` + memberRows + ` from ` + dao.table + ` where id = ? limit 1`
		return conn.QueryRow(v, query, primary)
	})
	switch err {
	case nil:
		return member, nil
	case sqlc.ErrNotFound:
		return nil, ecode.ErrNotFound
	default:
		return nil, err
	}
}

func (dao *Dao) GetMemberMaxAge() (age int64, err error) {
	sql := `SELECT IFNULL(MAX(age),0) FROM member WHERE deleted = 0 `
	if err = dao.QueryRowNoCache(&age, sql); err != nil {
		return 0, errors.Wrapf(err, "GetMemberMaxAge")
	}
	return
}

func (dao *Dao) GetMemberSumAge() (age int64, err error) {
	sql := `SELECT IFNULL(SUM(age),0) FROM member WHERE deleted = 0`
	if err = dao.QueryRowNoCache(&age, sql); err != nil {
		return 0, errors.Wrapf(err, "GetMemberMaxAge")
	}
	return
}

func (dao *Dao) CountMember() (count int64, err error) {
	sql := `SELECT COUNT(*) FROM member `
	if err = dao.QueryRowNoCache(&count, sql); err != nil {
		return 0, errors.Wrapf(err, "CountMember")
	}
	return
}

func (dao *Dao) HasMemberByID(id int64) (has bool, err error) {
	var count int
	s := `SELECT COUNT(*) FROM member WHERE id = ? `
	err = dao.QueryRowNoCache(&count, s, id)
	fmt.Println(err)
	if err != nil && err != sqlc.ErrNotFound {
		return false, errors.Wrapf(err, "HasMemberByID id(%d)", id)
	}
	return count > 0, nil
}

// 根据其他属性查询列表
func (dao *Dao) QueryByName(name string) (res []*model.Member, err error) {
	res = make([]*model.Member, 0, 0) // 返回nil还是空切片会影响json里的结构
	sql := `SELECT ` + memberRows + ` FROM member WHERE name like ?  `
	if err = dao.QueryRowsNoCache(&res, sql, name+"%"); err != nil {
		return res, errors.Wrapf(err, "QueryMemberByName name(%s)", name)
	}
	return
}

// 根据ids查询列表
func (dao *Dao) QueryMemberByIDs(ids []int64) (res map[int64]*model.Member, err error) {
	var t []*model.Member
	res = make(map[int64]*model.Member)
	sql := `SELECT ` + memberRows + `  FROM member WHERE id IN (` + stringx.JoinInts(ids) + ` ) `
	if err = dao.QueryRowsNoCache(&t, sql); err != nil {
		return res, errors.Wrapf(err, "QueryMemberByIDs ids(%v)", ids)
	}
	for _, r := range t {
		res[r.Id] = r
	}
	return
}

// 更新单个
func (dao *Dao) UpdateSome(data model.Member) (err error) {
	memberIdKey := fmt.Sprintf("%s%v", cacheMemberIdPrefix, data.Id)

	sqlStr := "UPDATE member SET  "
	sqlSli := []string{}
	var updateMap []interface{}
	if data.Phone != "-1" {
		sqlSli = append(sqlSli,"phone =?")
		updateMap = append(updateMap, data.Phone)
	}
	if data.Name != "-1" {
		sqlSli = append(sqlSli,"name =?")
		updateMap = append(updateMap, data.Name)
	}
	if data.Address != "-1" {
		sqlSli = append(sqlSli,"address =?")
		updateMap = append(updateMap, data.Address)
	}
	if data.Age != -1 {
		sqlSli = append(sqlSli,"age =?")
		updateMap = append(updateMap, data.Age)
	}
	sqlStr += strings.Join(sqlSli,",")+" WHERE id =?"
	updateMap = append(updateMap, data.Id)

	_, err = dao.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		return conn.Exec(sqlStr, updateMap...)
	}, memberIdKey)
	return err
}

func (dao *Dao) Update(data model.Member) (err error) {
	memberIdKey := fmt.Sprintf("%s%v", cacheMemberIdPrefix, data.Id)
	_, err = dao.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := `update ` + dao.table + ` set phone=?,name=?,age=?,address=? where id = ?`
		return conn.Exec(query, data.Phone, data.Name, data.Age, data.Address, data.Id)
	}, memberIdKey)
	return err
}

func (dao *Dao) Set(data model.Member) (err error) {
	memberIdKey := fmt.Sprintf("%s%v", cacheMemberIdPrefix, data.Id)
	_, err = dao.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := `insert into ` + dao.table + ` (id,phone,name,age,address) values (?, ?, ?, ?, ?)` +
			"ON DUPLICATE KEY UPDATE phone=?,name=?,age=?,address=?"
		return conn.Exec(query, data.Id, data.Phone, data.Name, data.Age, data.Address, data.Phone, data.Name, data.Age, data.Address)
	}, memberIdKey)
	return err
}

// 批量更改顺序
func (dao *Dao) SortMember(args model.ArgMemberSort) (err error) {
	var (
		buf bytes.Buffer
		ids []int64
	)
	buf.WriteString("UPDATE member SET order_num = CASE id")
	for _, arg := range args {
		buf.WriteString(" WHEN ")
		buf.WriteString(strconv.FormatInt(arg.Id, 10))
		buf.WriteString(" THEN ")
		buf.WriteString(strconv.FormatInt(arg.OrderNum, 10))
		ids = append(ids, arg.Id)
	}
	buf.WriteString(" END  WHERE id IN (")
	buf.WriteString(stringx.JoinInts(ids))
	buf.WriteString(")")
	if _, err = dao.ExecNoCache(buf.String()); err != nil {
		return errors.Wrapf(err, "BatchUpdateMemberSort args(%v)", args)
	}
	return
}

func (dao *Dao) Delete(id int64) (err error) {
	data, err := dao.FindOne(id)
	if err != nil {
		return err
	}

	memberIdKey := fmt.Sprintf("%s%v", cacheMemberIdPrefix, id)
	memberPhoneKey := fmt.Sprintf("%s%v", cacheMemberPhonePrefix, data.Phone)
	_, err = dao.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := `delete from ` + dao.table + ` where id = ?`
		return conn.Exec(query, id)
	}, memberPhoneKey, memberIdKey)
	return err
}
