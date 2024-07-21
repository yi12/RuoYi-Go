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

func newSysJobLog(db *gorm.DB, opts ...gen.DOOption) sysJobLog {
	_sysJobLog := sysJobLog{}

	_sysJobLog.sysJobLogDo.UseDB(db, opts...)
	_sysJobLog.sysJobLogDo.UseModel(&model.SysJobLog{})

	tableName := _sysJobLog.sysJobLogDo.TableName()
	_sysJobLog.ALL = field.NewAsterisk(tableName)
	_sysJobLog.JobLogID = field.NewInt64(tableName, "job_log_id")
	_sysJobLog.JobName = field.NewString(tableName, "job_name")
	_sysJobLog.JobGroup = field.NewString(tableName, "job_group")
	_sysJobLog.InvokeTarget = field.NewString(tableName, "invoke_target")
	_sysJobLog.JobMessage = field.NewString(tableName, "job_message")
	_sysJobLog.Status = field.NewString(tableName, "status")
	_sysJobLog.ExceptionInfo = field.NewString(tableName, "exception_info")
	_sysJobLog.CreateTime = field.NewTime(tableName, "create_time")

	_sysJobLog.fillFieldMap()

	return _sysJobLog
}

type sysJobLog struct {
	sysJobLogDo sysJobLogDo

	ALL           field.Asterisk
	JobLogID      field.Int64  // 任务日志ID
	JobName       field.String // 任务名称
	JobGroup      field.String // 任务组名
	InvokeTarget  field.String // 调用目标字符串
	JobMessage    field.String // 日志信息
	Status        field.String // 执行状态（0正常 1失败）
	ExceptionInfo field.String // 异常信息
	CreateTime    field.Time   // 创建时间

	fieldMap map[string]field.Expr
}

func (s sysJobLog) Table(newTableName string) *sysJobLog {
	s.sysJobLogDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysJobLog) As(alias string) *sysJobLog {
	s.sysJobLogDo.DO = *(s.sysJobLogDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysJobLog) updateTableName(table string) *sysJobLog {
	s.ALL = field.NewAsterisk(table)
	s.JobLogID = field.NewInt64(table, "job_log_id")
	s.JobName = field.NewString(table, "job_name")
	s.JobGroup = field.NewString(table, "job_group")
	s.InvokeTarget = field.NewString(table, "invoke_target")
	s.JobMessage = field.NewString(table, "job_message")
	s.Status = field.NewString(table, "status")
	s.ExceptionInfo = field.NewString(table, "exception_info")
	s.CreateTime = field.NewTime(table, "create_time")

	s.fillFieldMap()

	return s
}

func (s *sysJobLog) WithContext(ctx context.Context) *sysJobLogDo {
	return s.sysJobLogDo.WithContext(ctx)
}

func (s sysJobLog) TableName() string { return s.sysJobLogDo.TableName() }

func (s sysJobLog) Alias() string { return s.sysJobLogDo.Alias() }

func (s sysJobLog) Columns(cols ...field.Expr) gen.Columns { return s.sysJobLogDo.Columns(cols...) }

func (s *sysJobLog) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysJobLog) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 8)
	s.fieldMap["job_log_id"] = s.JobLogID
	s.fieldMap["job_name"] = s.JobName
	s.fieldMap["job_group"] = s.JobGroup
	s.fieldMap["invoke_target"] = s.InvokeTarget
	s.fieldMap["job_message"] = s.JobMessage
	s.fieldMap["status"] = s.Status
	s.fieldMap["exception_info"] = s.ExceptionInfo
	s.fieldMap["create_time"] = s.CreateTime
}

func (s sysJobLog) clone(db *gorm.DB) sysJobLog {
	s.sysJobLogDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysJobLog) replaceDB(db *gorm.DB) sysJobLog {
	s.sysJobLogDo.ReplaceDB(db)
	return s
}

type sysJobLogDo struct{ gen.DO }

func (s sysJobLogDo) Debug() *sysJobLogDo {
	return s.withDO(s.DO.Debug())
}

func (s sysJobLogDo) WithContext(ctx context.Context) *sysJobLogDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysJobLogDo) ReadDB() *sysJobLogDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysJobLogDo) WriteDB() *sysJobLogDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysJobLogDo) Session(config *gorm.Session) *sysJobLogDo {
	return s.withDO(s.DO.Session(config))
}

func (s sysJobLogDo) Clauses(conds ...clause.Expression) *sysJobLogDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysJobLogDo) Returning(value interface{}, columns ...string) *sysJobLogDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysJobLogDo) Not(conds ...gen.Condition) *sysJobLogDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysJobLogDo) Or(conds ...gen.Condition) *sysJobLogDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysJobLogDo) Select(conds ...field.Expr) *sysJobLogDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysJobLogDo) Where(conds ...gen.Condition) *sysJobLogDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysJobLogDo) Order(conds ...field.Expr) *sysJobLogDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysJobLogDo) Distinct(cols ...field.Expr) *sysJobLogDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysJobLogDo) Omit(cols ...field.Expr) *sysJobLogDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysJobLogDo) Join(table schema.Tabler, on ...field.Expr) *sysJobLogDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysJobLogDo) LeftJoin(table schema.Tabler, on ...field.Expr) *sysJobLogDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysJobLogDo) RightJoin(table schema.Tabler, on ...field.Expr) *sysJobLogDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysJobLogDo) Group(cols ...field.Expr) *sysJobLogDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysJobLogDo) Having(conds ...gen.Condition) *sysJobLogDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysJobLogDo) Limit(limit int) *sysJobLogDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysJobLogDo) Offset(offset int) *sysJobLogDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysJobLogDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *sysJobLogDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysJobLogDo) Unscoped() *sysJobLogDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysJobLogDo) Create(values ...*model.SysJobLog) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysJobLogDo) CreateInBatches(values []*model.SysJobLog, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysJobLogDo) Save(values ...*model.SysJobLog) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysJobLogDo) First() (*model.SysJobLog, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysJobLog), nil
	}
}

func (s sysJobLogDo) Take() (*model.SysJobLog, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysJobLog), nil
	}
}

func (s sysJobLogDo) Last() (*model.SysJobLog, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysJobLog), nil
	}
}

func (s sysJobLogDo) Find() ([]*model.SysJobLog, error) {
	result, err := s.DO.Find()
	return result.([]*model.SysJobLog), err
}

func (s sysJobLogDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysJobLog, err error) {
	buf := make([]*model.SysJobLog, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysJobLogDo) FindInBatches(result *[]*model.SysJobLog, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysJobLogDo) Attrs(attrs ...field.AssignExpr) *sysJobLogDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysJobLogDo) Assign(attrs ...field.AssignExpr) *sysJobLogDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysJobLogDo) Joins(fields ...field.RelationField) *sysJobLogDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysJobLogDo) Preload(fields ...field.RelationField) *sysJobLogDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysJobLogDo) FirstOrInit() (*model.SysJobLog, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysJobLog), nil
	}
}

func (s sysJobLogDo) FirstOrCreate() (*model.SysJobLog, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysJobLog), nil
	}
}

func (s sysJobLogDo) FindByPage(offset int, limit int) (result []*model.SysJobLog, count int64, err error) {
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

func (s sysJobLogDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysJobLogDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysJobLogDo) Delete(models ...*model.SysJobLog) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysJobLogDo) withDO(do gen.Dao) *sysJobLogDo {
	s.DO = *do.(*gen.DO)
	return s
}
