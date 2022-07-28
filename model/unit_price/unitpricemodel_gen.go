// Code generated by goctl. DO NOT EDIT!

package unit_price

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
	unitPriceFieldNames          = builder.RawFieldNames(&UnitPrice{})
	unitPriceRows                = strings.Join(unitPriceFieldNames, ",")
	unitPriceRowsExpectAutoSet   = strings.Join(stringx.Remove(unitPriceFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	unitPriceRowsWithPlaceHolder = strings.Join(stringx.Remove(unitPriceFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheUnitPriceIdPrefix = "cache:unitPrice:id:"
)

type (
	unitPriceModel interface {
		Insert(ctx context.Context, data *UnitPrice) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*UnitPrice, error)
		Update(ctx context.Context, data *UnitPrice) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUnitPriceModel struct {
		sqlc.CachedConn
		table string
	}

	UnitPrice struct {
		Id           int64          `db:"id"`
		EnterpriseId int64          `db:"enterprise_id"`
		WorkshopId   int64          `db:"workshop_id"`
		TechnologyId int64          `db:"technology_id"`
		ClientId     int64          `db:"client_id"`
		UnitPrice    float64        `db:"unit_price"`
		CreateTime   time.Time      `db:"create_time"`
		UpdateTime   time.Time      `db:"update_time"`
		Remark       sql.NullString `db:"remark"`
		Version      int64          `db:"version"`
	}
)

func newUnitPriceModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultUnitPriceModel {
	return &defaultUnitPriceModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`unit_price`",
	}
}

func (m *defaultUnitPriceModel) Insert(ctx context.Context, data *UnitPrice) (sql.Result, error) {
	unitPriceIdKey := fmt.Sprintf("%s%v", cacheUnitPriceIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, unitPriceRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.EnterpriseId, data.WorkshopId, data.TechnologyId, data.ClientId, data.UnitPrice, data.Remark, data.Version)
	}, unitPriceIdKey)
	return ret, err
}

func (m *defaultUnitPriceModel) FindOne(ctx context.Context, id int64) (*UnitPrice, error) {
	unitPriceIdKey := fmt.Sprintf("%s%v", cacheUnitPriceIdPrefix, id)
	var resp UnitPrice
	err := m.QueryRowCtx(ctx, &resp, unitPriceIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", unitPriceRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
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

func (m *defaultUnitPriceModel) Update(ctx context.Context, data *UnitPrice) error {
	unitPriceIdKey := fmt.Sprintf("%s%v", cacheUnitPriceIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, unitPriceRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.EnterpriseId, data.WorkshopId, data.TechnologyId, data.ClientId, data.UnitPrice, data.Remark, data.Version, data.Id)
	}, unitPriceIdKey)
	return err
}

func (m *defaultUnitPriceModel) Delete(ctx context.Context, id int64) error {
	unitPriceIdKey := fmt.Sprintf("%s%v", cacheUnitPriceIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, unitPriceIdKey)
	return err
}

func (m *defaultUnitPriceModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheUnitPriceIdPrefix, primary)
}

func (m *defaultUnitPriceModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", unitPriceRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUnitPriceModel) tableName() string {
	return m.table
}