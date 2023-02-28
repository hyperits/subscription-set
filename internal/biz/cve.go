package biz

import (
	"database/sql"
)

// Cve 结构体
type Cve struct {
	Id             uint         `gorm:"column:id;type:int(10) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	Uuid           string       `gorm:"column:uuid;type:varchar(128);NOT NULL" json:"uuid"`
	VulnStatus     string       `gorm:"column:vuln_status;type:varchar(64);NOT NULL" json:"vuln_status"`
	Descriptions   string       `gorm:"column:descriptions;type:json" json:"descriptions"`
	Metrics        string       `gorm:"column:metrics;type:json" json:"metrics"`
	Weaknesses     string       `gorm:"column:weaknesses;type:json" json:"weaknesses"`
	Configurations string       `gorm:"column:configurations;type:json" json:"configurations"`
	References     string       `gorm:"column:references;type:json" json:"references"`
	Published      sql.NullTime `gorm:"column:published;type:timestamp" json:"published"`
	LastModified   sql.NullTime `gorm:"column:last_modified;type:timestamp" json:"last_modified"`
}

// TableName table
func (m *Cve) TableName() string {
	return "cve"
}

// CveRepo repo
type CveRepo interface {
	InsertCve(b *Cve) error
}
