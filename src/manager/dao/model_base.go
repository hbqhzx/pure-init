package dao

import (
	"github.com/jinzhu/gorm"
	. "lib/db"
	"time"
)

type BaseScope struct {
	DB *MyDB
}

type ModelBase struct {
	CreateTime int    `json:"create_time" comment:"创建时间" key:""`
	UpdateTime int    `json:"update_time" comment:"更新时间" key:""`
	CreatedBy  string `json:"created_by" len:"31" comment:"创建人"`
	UpdatedBy  string `json:"updated_by" len:"31" comment:"最后修改人"`

	scope BaseScope `json:"-" gorm:"-"`
}

func (mb *ModelBase) BeforeCreate(scope *gorm.Scope) {
	scope.SetColumn("CreateTime", Now())
	scope.SetColumn("UpdateTime", Now())
	if mb.UpdatedBy != "" {
		scope.SetColumn("CreatedBy", mb.UpdatedBy)
	}
}

func (*ModelBase) BeforeUpdate(scope *gorm.Scope) {
	scope.SetColumn("UpdateTime", Now())
}

func (m *ModelBase) DB() *MyDB {
	if m.scope.DB == nil {
		m.scope.DB = &DB
	}
	return &MyDB{m.scope.DB.New()}
}

func (m *ModelBase) SetDB(db *MyDB) {
	m.scope.DB = db
}

func Now() int {
	return int(time.Now().Unix())
}
