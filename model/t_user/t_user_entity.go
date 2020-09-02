// ==========================================================================
// This is auto-generated by gf cli tool. You may not really want to edit it.
// ==========================================================================

package t_user

import (
	"database/sql"

	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/os/gtime"
)

// Entity is the golang structure for table t_user.
type Entity struct {
	Id         int64       `orm:"id,primary"      json:"id"`          // 流水id
	Nickname   string      `orm:"nickname"        json:"nickname"`    // 昵称
	Phone      string      `orm:"phone,primary"   json:"phone"`       // 手机号
	Salt       string      `orm:"salt"            json:"salt"`        // 动态盐
	Passwd     string      `orm:"pw"              json:"pw"`          // 最终密码串
	Logo       string      `orm:"logo"            json:"logo"`        // 头像
	Source     string      `orm:"source"          json:"source"`      // 来源
	RegisterIp string      `orm:"register_ip"     json:"register_ip"` // 注册ip
	CreateTime *gtime.Time `orm:"create_time"     json:"create_time"` //
	ModifyTime *gtime.Time `orm:"modify_time"     json:"modify_time"` //
}

// OmitEmpty sets OPTION_OMITEMPTY option for the model, which automatically filers
// the data and where attributes for empty values.
func (r *Entity) OmitEmpty() *arModel {
	return Model.Data(r).OmitEmpty()
}

// Inserts does "INSERT...INTO..." statement for inserting current object into table.
func (r *Entity) Insert() (result sql.Result, err error) {
	return Model.Data(r).Insert()
}

// InsertIgnore does "INSERT IGNORE INTO ..." statement for inserting current object into table.
func (r *Entity) InsertIgnore() (result sql.Result, err error) {
	return Model.Data(r).InsertIgnore()
}

// Replace does "REPLACE...INTO..." statement for inserting current object into table.
// If there's already another same record in the table (it checks using primary key or unique index),
// it deletes it and insert this one.
func (r *Entity) Replace() (result sql.Result, err error) {
	return Model.Data(r).Replace()
}

// Save does "INSERT...INTO..." statement for inserting/updating current object into table.
// It updates the record if there's already another same record in the table
// (it checks using primary key or unique index).
func (r *Entity) Save() (result sql.Result, err error) {
	return Model.Data(r).Save()
}

// Update does "UPDATE...WHERE..." statement for updating current object from table.
// It updates the record if there's already another same record in the table
// (it checks using primary key or unique index).
func (r *Entity) Update() (result sql.Result, err error) {
	return Model.Data(r).Where(gdb.GetWhereConditionOfStruct(r)).Update()
}

// Delete does "DELETE FROM...WHERE..." statement for deleting current object from table.
func (r *Entity) Delete() (result sql.Result, err error) {
	return Model.Where(gdb.GetWhereConditionOfStruct(r)).Delete()
}