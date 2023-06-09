// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userFieldNames          = builder.RawFieldNames(&User{})
	userRows                = strings.Join(userFieldNames, ",")
	userRowsExpectAutoSet   = strings.Join(stringx.Remove(userFieldNames, "`user_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userRowsWithPlaceHolder = strings.Join(stringx.Remove(userFieldNames, "`user_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheUserUserIdPrefix = "cache:user:userId:"
)

type (
	userModel interface {
		Insert(ctx context.Context, data *User) (sql.Result, error)
		FindOne(ctx context.Context, userId int64) (*User, error)
		Update(ctx context.Context, data *User) error
		Delete(ctx context.Context, userId int64) error
	}

	defaultUserModel struct {
		sqlc.CachedConn
		table string
	}

	User struct {
		UserId       int64     `db:"user_id"`       // 用户ID
		PassWord     string    `db:"passWord"`      // 用户密码，MD5加密
		UserNick     string    `db:"user_Nick"`     // 用户昵称
		UserFace     string    `db:"user_Face"`     // 用户头像地址
		UserSex      int64     `db:"User_Sex"`      // 用户性别：0男，1女，2保密
		UserEmail    string    `db:"user_Email"`    // 用户邮箱
		UserPhone    string    `db:"user_Phone"`    // 手机号
		GithubId     string    `db:"github_id"`     // github_id
		QqId         string    `db:"qq_id"`         // qq_id
		LoginAddress string    `db:"login_Address"` // 用户登录IP地址
		CreateTime   time.Time `db:"create_time"`   // 创建时间
		UpdateTime   time.Time `db:"update_time"`   // 更新时间
	}
)

func newUserModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultUserModel {
	return &defaultUserModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`user`",
	}
}

func (m *defaultUserModel) Delete(ctx context.Context, userId int64) error {
	userUserIdKey := fmt.Sprintf("%s%v", cacheUserUserIdPrefix, userId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `user_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, userId)
	}, userUserIdKey)
	return err
}

func (m *defaultUserModel) FindOne(ctx context.Context, userId int64) (*User, error) {
	userUserIdKey := fmt.Sprintf("%s%v", cacheUserUserIdPrefix, userId)
	var resp User
	err := m.QueryRowCtx(ctx, &resp, userUserIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `user_id` = ? limit 1", userRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, userId)
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

func (m *defaultUserModel) Insert(ctx context.Context, data *User) (sql.Result, error) {
	userUserIdKey := fmt.Sprintf("%s%v", cacheUserUserIdPrefix, data.UserId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, userRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.PassWord, data.UserNick, data.UserFace, data.UserSex, data.UserEmail, data.UserPhone, data.GithubId, data.QqId, data.LoginAddress)
	}, userUserIdKey)
	return ret, err
}

func (m *defaultUserModel) Update(ctx context.Context, data *User) error {
	userUserIdKey := fmt.Sprintf("%s%v", cacheUserUserIdPrefix, data.UserId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `user_id` = ?", m.table, userRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.PassWord, data.UserNick, data.UserFace, data.UserSex, data.UserEmail, data.UserPhone, data.GithubId, data.QqId, data.LoginAddress, data.UserId)
	}, userUserIdKey)
	return err
}

func (m *defaultUserModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheUserUserIdPrefix, primary)
}

func (m *defaultUserModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `user_id` = ? limit 1", userRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUserModel) tableName() string {
	return m.table
}
