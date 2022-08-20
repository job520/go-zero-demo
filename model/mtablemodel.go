package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ MtableModel = (*customMtableModel)(nil)

type (
	// MtableModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMtableModel.
	MtableModel interface {
		mtableModel
	}

	customMtableModel struct {
		*defaultMtableModel
	}
)

// NewMtableModel returns a model for the database table.
func NewMtableModel(conn sqlx.SqlConn, c cache.CacheConf) MtableModel {
	return &customMtableModel{
		defaultMtableModel: newMtableModel(conn, c),
	}
}
