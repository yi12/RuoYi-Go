// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go
// Email: hot_kun@hotmail.com or 867917691@qq.com

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

func newSysRole(db *gorm.DB, opts ...gen.DOOption) sysRole {
	_sysRole := sysRole{}

	_sysRole.sysRoleDo.UseDB(db, opts...)
	_sysRole.sysRoleDo.UseModel(&model.SysRole{})

	tableName := _sysRole.sysRoleDo.TableName()
	_sysRole.ALL = field.NewAsterisk(tableName)
	_sysRole.RoleID = field.NewInt64(tableName, "role_id")
	_sysRole.RoleName = field.NewString(tableName, "role_name")
	_sysRole.RoleKey = field.NewString(tableName, "role_key")
	_sysRole.RoleSort = field.NewInt32(tableName, "role_sort")
	_sysRole.DataScope = field.NewString(tableName, "data_scope")
	_sysRole.MenuCheckStrictly = field.NewInt16(tableName, "menu_check_strictly")
	_sysRole.DeptCheckStrictly = field.NewInt16(tableName, "dept_check_strictly")
	_sysRole.Status = field.NewString(tableName, "status")
	_sysRole.DelFlag = field.NewString(tableName, "del_flag")
	_sysRole.CreateBy = field.NewString(tableName, "create_by")
	_sysRole.CreateTime = field.NewTime(tableName, "create_time")
	_sysRole.UpdateBy = field.NewString(tableName, "update_by")
	_sysRole.UpdateTime = field.NewTime(tableName, "update_time")
	_sysRole.Remark = field.NewString(tableName, "remark")

	_sysRole.fillFieldMap()

	return _sysRole
}

type sysRole struct {
	sysRoleDo sysRoleDo

	ALL               field.Asterisk
	RoleID            field.Int64  // 角色ID
	RoleName          field.String // 角色名称
	RoleKey           field.String // 角色权限字符串
	RoleSort          field.Int32  // 显示顺序
	DataScope         field.String // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
	MenuCheckStrictly field.Int16  // 菜单树选择项是否关联显示
	DeptCheckStrictly field.Int16  // 部门树选择项是否关联显示
	Status            field.String // 角色状态（0正常 1停用）
	DelFlag           field.String // 删除标志（0代表存在 2代表删除）
	CreateBy          field.String // 创建者
	CreateTime        field.Time   // 创建时间
	UpdateBy          field.String // 更新者
	UpdateTime        field.Time   // 更新时间
	Remark            field.String // 备注

	fieldMap map[string]field.Expr
}

func (s sysRole) Table(newTableName string) *sysRole {
	s.sysRoleDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysRole) As(alias string) *sysRole {
	s.sysRoleDo.DO = *(s.sysRoleDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysRole) updateTableName(table string) *sysRole {
	s.ALL = field.NewAsterisk(table)
	s.RoleID = field.NewInt64(table, "role_id")
	s.RoleName = field.NewString(table, "role_name")
	s.RoleKey = field.NewString(table, "role_key")
	s.RoleSort = field.NewInt32(table, "role_sort")
	s.DataScope = field.NewString(table, "data_scope")
	s.MenuCheckStrictly = field.NewInt16(table, "menu_check_strictly")
	s.DeptCheckStrictly = field.NewInt16(table, "dept_check_strictly")
	s.Status = field.NewString(table, "status")
	s.DelFlag = field.NewString(table, "del_flag")
	s.CreateBy = field.NewString(table, "create_by")
	s.CreateTime = field.NewTime(table, "create_time")
	s.UpdateBy = field.NewString(table, "update_by")
	s.UpdateTime = field.NewTime(table, "update_time")
	s.Remark = field.NewString(table, "remark")

	s.fillFieldMap()

	return s
}

func (s *sysRole) WithContext(ctx context.Context) *sysRoleDo { return s.sysRoleDo.WithContext(ctx) }

func (s sysRole) TableName() string { return s.sysRoleDo.TableName() }

func (s sysRole) Alias() string { return s.sysRoleDo.Alias() }

func (s sysRole) Columns(cols ...field.Expr) gen.Columns { return s.sysRoleDo.Columns(cols...) }

func (s *sysRole) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysRole) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 14)
	s.fieldMap["role_id"] = s.RoleID
	s.fieldMap["role_name"] = s.RoleName
	s.fieldMap["role_key"] = s.RoleKey
	s.fieldMap["role_sort"] = s.RoleSort
	s.fieldMap["data_scope"] = s.DataScope
	s.fieldMap["menu_check_strictly"] = s.MenuCheckStrictly
	s.fieldMap["dept_check_strictly"] = s.DeptCheckStrictly
	s.fieldMap["status"] = s.Status
	s.fieldMap["del_flag"] = s.DelFlag
	s.fieldMap["create_by"] = s.CreateBy
	s.fieldMap["create_time"] = s.CreateTime
	s.fieldMap["update_by"] = s.UpdateBy
	s.fieldMap["update_time"] = s.UpdateTime
	s.fieldMap["remark"] = s.Remark
}

func (s sysRole) clone(db *gorm.DB) sysRole {
	s.sysRoleDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysRole) replaceDB(db *gorm.DB) sysRole {
	s.sysRoleDo.ReplaceDB(db)
	return s
}

type sysRoleDo struct{ gen.DO }

func (s sysRoleDo) Debug() *sysRoleDo {
	return s.withDO(s.DO.Debug())
}

func (s sysRoleDo) WithContext(ctx context.Context) *sysRoleDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysRoleDo) ReadDB() *sysRoleDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysRoleDo) WriteDB() *sysRoleDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysRoleDo) Session(config *gorm.Session) *sysRoleDo {
	return s.withDO(s.DO.Session(config))
}

func (s sysRoleDo) Clauses(conds ...clause.Expression) *sysRoleDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysRoleDo) Returning(value interface{}, columns ...string) *sysRoleDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysRoleDo) Not(conds ...gen.Condition) *sysRoleDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysRoleDo) Or(conds ...gen.Condition) *sysRoleDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysRoleDo) Select(conds ...field.Expr) *sysRoleDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysRoleDo) Where(conds ...gen.Condition) *sysRoleDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysRoleDo) Order(conds ...field.Expr) *sysRoleDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysRoleDo) Distinct(cols ...field.Expr) *sysRoleDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysRoleDo) Omit(cols ...field.Expr) *sysRoleDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysRoleDo) Join(table schema.Tabler, on ...field.Expr) *sysRoleDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysRoleDo) LeftJoin(table schema.Tabler, on ...field.Expr) *sysRoleDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysRoleDo) RightJoin(table schema.Tabler, on ...field.Expr) *sysRoleDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysRoleDo) Group(cols ...field.Expr) *sysRoleDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysRoleDo) Having(conds ...gen.Condition) *sysRoleDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysRoleDo) Limit(limit int) *sysRoleDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysRoleDo) Offset(offset int) *sysRoleDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysRoleDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *sysRoleDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysRoleDo) Unscoped() *sysRoleDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysRoleDo) Create(values ...*model.SysRole) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysRoleDo) CreateInBatches(values []*model.SysRole, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysRoleDo) Save(values ...*model.SysRole) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysRoleDo) First() (*model.SysRole, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysRole), nil
	}
}

func (s sysRoleDo) Take() (*model.SysRole, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysRole), nil
	}
}

func (s sysRoleDo) Last() (*model.SysRole, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysRole), nil
	}
}

func (s sysRoleDo) Find() ([]*model.SysRole, error) {
	result, err := s.DO.Find()
	return result.([]*model.SysRole), err
}

func (s sysRoleDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysRole, err error) {
	buf := make([]*model.SysRole, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysRoleDo) FindInBatches(result *[]*model.SysRole, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysRoleDo) Attrs(attrs ...field.AssignExpr) *sysRoleDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysRoleDo) Assign(attrs ...field.AssignExpr) *sysRoleDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysRoleDo) Joins(fields ...field.RelationField) *sysRoleDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysRoleDo) Preload(fields ...field.RelationField) *sysRoleDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysRoleDo) FirstOrInit() (*model.SysRole, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysRole), nil
	}
}

func (s sysRoleDo) FirstOrCreate() (*model.SysRole, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysRole), nil
	}
}

func (s sysRoleDo) FindByPage(offset int, limit int) (result []*model.SysRole, count int64, err error) {
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

func (s sysRoleDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysRoleDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysRoleDo) Delete(models ...*model.SysRole) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysRoleDo) withDO(do gen.Dao) *sysRoleDo {
	s.DO = *do.(*gen.DO)
	return s
}
