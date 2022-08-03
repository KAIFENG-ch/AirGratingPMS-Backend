// Code generated by goctl. DO NOT EDIT!

package production_plan

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
	productionPlanFieldNames          = builder.RawFieldNames(&ProductionPlan{})
	productionPlanRows                = strings.Join(productionPlanFieldNames, ",")
	productionPlanRowsExpectAutoSet   = strings.Join(stringx.Remove(productionPlanFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	productionPlanRowsWithPlaceHolder = strings.Join(stringx.Remove(productionPlanFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheProductionPlanIdPrefix = "cache:productionPlan:id:"
)

type (
	productionPlanModel interface {
		Insert(ctx context.Context, data *ProductionPlan) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*ProductionPlan, error)
		Update(ctx context.Context, data *ProductionPlan) error
		Delete(ctx context.Context, id int64) error
	}

	defaultProductionPlanModel struct {
		sqlc.CachedConn
		table string
	}

	ProductionPlan struct {
		Id             int64     `db:"id"`
		EnterpriseId   int64     `db:"enterprise_id"`
		WorkshopId     int64     `db:"workshop_id"`
		ProductionTime time.Time `db:"production_time"`
		State          int64     `db:"state"`
		CreateTime     time.Time `db:"create_time"`
		UpdateTime     time.Time `db:"update_time"`
		Remark         string    `db:"remark"`
		Version        int64     `db:"version"`
	}
)

func newProductionPlanModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultProductionPlanModel {
	return &defaultProductionPlanModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`production_plan`",
	}
}

func (m *defaultProductionPlanModel) Insert(ctx context.Context, data *ProductionPlan) (sql.Result, error) {
	productionPlanIdKey := fmt.Sprintf("%s%v", cacheProductionPlanIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, productionPlanRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.EnterpriseId, data.WorkshopId, data.ProductionTime, data.State, data.Remark, data.Version)
	}, productionPlanIdKey)
	return ret, err
}

func (m *defaultProductionPlanModel) FindOne(ctx context.Context, id int64) (*ProductionPlan, error) {
	productionPlanIdKey := fmt.Sprintf("%s%v", cacheProductionPlanIdPrefix, id)
	var resp ProductionPlan
	err := m.QueryRowCtx(ctx, &resp, productionPlanIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", productionPlanRows, m.table)
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

func (m *defaultProductionPlanModel) Update(ctx context.Context, data *ProductionPlan) error {
	productionPlanIdKey := fmt.Sprintf("%s%v", cacheProductionPlanIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, productionPlanRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.EnterpriseId, data.WorkshopId, data.ProductionTime, data.State, data.Remark, data.Version, data.Id)
	}, productionPlanIdKey)
	return err
}

func (m *defaultProductionPlanModel) Delete(ctx context.Context, id int64) error {
	productionPlanIdKey := fmt.Sprintf("%s%v", cacheProductionPlanIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, productionPlanIdKey)
	return err
}

func (m *defaultProductionPlanModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheProductionPlanIdPrefix, primary)
}

func (m *defaultProductionPlanModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", productionPlanRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultProductionPlanModel) tableName() string {
	return m.table
}
