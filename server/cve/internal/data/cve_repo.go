package data

import (
	"context"
	"github.com/hyperits/subscription-sets/cve/internal/biz"
	"github.com/hyperits/subscription-sets/cve/internal/pkg"
)

// CveRepo repo
type CveRepo struct {
	db *pkg.Sql
}

// NewCveRepo init
func NewCveRepo() *CveRepo {
	return &CveRepo{
		db: pkg.NewSql(),
	}
}

// InsertCve insert
func (c *CveRepo) InsertCve(b *biz.Cve) error {
	if err := c.db.MDB.Table("").Create(b).Error; err != nil {
		c.db.MDB.Logger.Error(context.Background(), "InsertCve err:%s", err.Error())
		return err
	}
	return nil
}
