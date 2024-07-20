// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go
// Email: hot_kun@hotmail.com or BusinessCallKun@gmail.com

// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"RuoYi-Go/internal/domain/model"
)

func newSysMenu(db *gorm.DB, opts ...gen.DOOption) sysMenu {
	_sysMenu := sysMenu{}

	_sysMenu.sysMenuDo.UseDB(db, opts...)
	_sysMenu.sysMenuDo.UseModel(&model.SysMenu{})

	tableName := _sysMenu.sysMenuDo.TableName()
	_sysMenu.ALL = field.NewAsterisk(tableName)
	_sysMenu.MenuID = field.NewInt64(tableName, "menu_id")
	_sysMenu.MenuName = field.NewString(tableName, "menu_name")
	_sysMenu.ParentID = field.NewInt64(tableName, "parent_id")
	_sysMenu.OrderNum = field.NewInt32(tableName, "order_num")
	_sysMenu.Path = field.NewString(tableName, "path")
	_sysMenu.Component = field.NewString(tableName, "component")
	_sysMenu.Query = field.NewString(tableName, "query")
	_sysMenu.IsFrame = field.NewInt32(tableName, "is_frame")
	_sysMenu.IsCache = field.NewInt32(tableName, "is_cache")
	_sysMenu.MenuType = field.NewString(tableName, "menu_type")
	_sysMenu.Visible = field.NewString(tableName, "visible")
	_sysMenu.Status = field.NewString(tableName, "status")
	_sysMenu.Perms = field.NewString(tableName, "perms")
	_sysMenu.Icon = field.NewString(tableName, "icon")
	_sysMenu.CreateBy = field.NewString(tableName, "create_by")
	_sysMenu.CreateTime = field.NewTime(tableName, "create_time")
	_sysMenu.UpdateBy = field.NewString(tableName, "update_by")
	_sysMenu.UpdateTime = field.NewTime(tableName, "update_time")
	_sysMenu.Remark = field.NewString(tableName, "remark")

	_sysMenu.fillFieldMap()

	return _sysMenu
}

type sysMenu struct {
	sysMenuDo sysMenuDo

	ALL        field.Asterisk
	MenuID     field.Int64  // 菜单ID
	MenuName   field.String // 菜单名称
	ParentID   field.Int64  // 父菜单ID
	OrderNum   field.Int32  // 显示顺序
	Path       field.String // 路由地址
	Component  field.String // 组件路径
	Query      field.String // 路由参数
	IsFrame    field.Int32  // 是否为外链（0是 1否）
	IsCache    field.Int32  // 是否缓存（0缓存 1不缓存）
	MenuType   field.String // 菜单类型（M目录 C菜单 F按钮）
	Visible    field.String // 菜单状态（0显示 1隐藏）
	Status     field.String // 菜单状态（0正常 1停用）
	Perms      field.String // 权限标识
	Icon       field.String // 菜单图标
	CreateBy   field.String // 创建者
	CreateTime field.Time   // 创建时间
	UpdateBy   field.String // 更新者
	UpdateTime field.Time   // 更新时间
	Remark     field.String // 备注

	fieldMap map[string]field.Expr
}

func (s sysMenu) Table(newTableName string) *sysMenu {
	s.sysMenuDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysMenu) As(alias string) *sysMenu {
	s.sysMenuDo.DO = *(s.sysMenuDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysMenu) updateTableName(table string) *sysMenu {
	s.ALL = field.NewAsterisk(table)
	s.MenuID = field.NewInt64(table, "menu_id")
	s.MenuName = field.NewString(table, "menu_name")
	s.ParentID = field.NewInt64(table, "parent_id")
	s.OrderNum = field.NewInt32(table, "order_num")
	s.Path = field.NewString(table, "path")
	s.Component = field.NewString(table, "component")
	s.Query = field.NewString(table, "query")
	s.IsFrame = field.NewInt32(table, "is_frame")
	s.IsCache = field.NewInt32(table, "is_cache")
	s.MenuType = field.NewString(table, "menu_type")
	s.Visible = field.NewString(table, "visible")
	s.Status = field.NewString(table, "status")
	s.Perms = field.NewString(table, "perms")
	s.Icon = field.NewString(table, "icon")
	s.CreateBy = field.NewString(table, "create_by")
	s.CreateTime = field.NewTime(table, "create_time")
	s.UpdateBy = field.NewString(table, "update_by")
	s.UpdateTime = field.NewTime(table, "update_time")
	s.Remark = field.NewString(table, "remark")

	s.fillFieldMap()

	return s
}

func (s *sysMenu) WithContext(ctx context.Context) *sysMenuDo { return s.sysMenuDo.WithContext(ctx) }

func (s sysMenu) TableName() string { return s.sysMenuDo.TableName() }

func (s sysMenu) Alias() string { return s.sysMenuDo.Alias() }

func (s sysMenu) Columns(cols ...field.Expr) gen.Columns { return s.sysMenuDo.Columns(cols...) }

func (s *sysMenu) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysMenu) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 19)
	s.fieldMap["menu_id"] = s.MenuID
	s.fieldMap["menu_name"] = s.MenuName
	s.fieldMap["parent_id"] = s.ParentID
	s.fieldMap["order_num"] = s.OrderNum
	s.fieldMap["path"] = s.Path
	s.fieldMap["component"] = s.Component
	s.fieldMap["query"] = s.Query
	s.fieldMap["is_frame"] = s.IsFrame
	s.fieldMap["is_cache"] = s.IsCache
	s.fieldMap["menu_type"] = s.MenuType
	s.fieldMap["visible"] = s.Visible
	s.fieldMap["status"] = s.Status
	s.fieldMap["perms"] = s.Perms
	s.fieldMap["icon"] = s.Icon
	s.fieldMap["create_by"] = s.CreateBy
	s.fieldMap["create_time"] = s.CreateTime
	s.fieldMap["update_by"] = s.UpdateBy
	s.fieldMap["update_time"] = s.UpdateTime
	s.fieldMap["remark"] = s.Remark
}

func (s sysMenu) clone(db *gorm.DB) sysMenu {
	s.sysMenuDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysMenu) replaceDB(db *gorm.DB) sysMenu {
	s.sysMenuDo.ReplaceDB(db)
	return s
}

type sysMenuDo struct{ gen.DO }

func (s sysMenuDo) Debug() *sysMenuDo {
	return s.withDO(s.DO.Debug())
}

func (s sysMenuDo) WithContext(ctx context.Context) *sysMenuDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysMenuDo) ReadDB() *sysMenuDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysMenuDo) WriteDB() *sysMenuDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysMenuDo) Session(config *gorm.Session) *sysMenuDo {
	return s.withDO(s.DO.Session(config))
}

func (s sysMenuDo) Clauses(conds ...clause.Expression) *sysMenuDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysMenuDo) Returning(value interface{}, columns ...string) *sysMenuDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysMenuDo) Not(conds ...gen.Condition) *sysMenuDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysMenuDo) Or(conds ...gen.Condition) *sysMenuDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysMenuDo) Select(conds ...field.Expr) *sysMenuDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysMenuDo) Where(conds ...gen.Condition) *sysMenuDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysMenuDo) Order(conds ...field.Expr) *sysMenuDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysMenuDo) Distinct(cols ...field.Expr) *sysMenuDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysMenuDo) Omit(cols ...field.Expr) *sysMenuDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysMenuDo) Join(table schema.Tabler, on ...field.Expr) *sysMenuDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysMenuDo) LeftJoin(table schema.Tabler, on ...field.Expr) *sysMenuDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysMenuDo) RightJoin(table schema.Tabler, on ...field.Expr) *sysMenuDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysMenuDo) Group(cols ...field.Expr) *sysMenuDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysMenuDo) Having(conds ...gen.Condition) *sysMenuDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysMenuDo) Limit(limit int) *sysMenuDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysMenuDo) Offset(offset int) *sysMenuDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysMenuDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *sysMenuDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysMenuDo) Unscoped() *sysMenuDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysMenuDo) Create(values ...*model.SysMenu) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysMenuDo) CreateInBatches(values []*model.SysMenu, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysMenuDo) Save(values ...*model.SysMenu) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysMenuDo) First() (*model.SysMenu, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysMenu), nil
	}
}

func (s sysMenuDo) Take() (*model.SysMenu, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysMenu), nil
	}
}

func (s sysMenuDo) Last() (*model.SysMenu, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysMenu), nil
	}
}

func (s sysMenuDo) Find() ([]*model.SysMenu, error) {
	result, err := s.DO.Find()
	return result.([]*model.SysMenu), err
}

func (s sysMenuDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysMenu, err error) {
	buf := make([]*model.SysMenu, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysMenuDo) FindInBatches(result *[]*model.SysMenu, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysMenuDo) Attrs(attrs ...field.AssignExpr) *sysMenuDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysMenuDo) Assign(attrs ...field.AssignExpr) *sysMenuDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysMenuDo) Joins(fields ...field.RelationField) *sysMenuDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysMenuDo) Preload(fields ...field.RelationField) *sysMenuDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysMenuDo) FirstOrInit() (*model.SysMenu, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysMenu), nil
	}
}

func (s sysMenuDo) FirstOrCreate() (*model.SysMenu, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysMenu), nil
	}
}

func (s sysMenuDo) FindByPage(offset int, limit int) (result []*model.SysMenu, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s sysMenuDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysMenuDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysMenuDo) Delete(models ...*model.SysMenu) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysMenuDo) withDO(do gen.Dao) *sysMenuDo {
	s.DO = *do.(*gen.DO)
	return s
}