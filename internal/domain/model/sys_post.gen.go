// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go or https://gitee.com/gitee_kun/RuoYi-Go
// Email: hot_kun@hotmail.com or 867917691@qq.com

// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSysPost = "sys_post"

// SysPost mapped from table <sys_post>
type SysPost struct {
	PostID     int64     `gorm:"column:post_id;primaryKey;comment:岗位ID" json:"postId"`                       // 岗位ID
	PostCode   string    `gorm:"column:post_code;not null;comment:岗位编码" json:"postCode" validate:"required"` // 岗位编码
	PostName   string    `gorm:"column:post_name;not null;comment:岗位名称" json:"postName" validate:"required"` // 岗位名称
	PostSort   int32     `gorm:"column:post_sort;not null;comment:显示顺序" json:"postSort" validate:"required"` // 显示顺序
	Status     string    `gorm:"column:status;not null;comment:状态（0正常 1停用）" json:"status"`                   // 状态（0正常 1停用）
	CreateBy   string    `gorm:"column:create_by;comment:创建者" json:"createBy"`                               // 创建者
	CreateTime time.Time `gorm:"column:create_time;comment:创建时间" json:"createTime"`                          // 创建时间
	UpdateBy   string    `gorm:"column:update_by;comment:更新者" json:"updateBy"`                               // 更新者
	UpdateTime time.Time `gorm:"column:update_time;comment:更新时间" json:"updateTime"`                          // 更新时间
	Remark     string    `gorm:"column:remark;comment:备注" json:"remark"`                                     // 备注
}

// TableName SysPost's table name
func (*SysPost) TableName() string {
	return TableNameSysPost
}

type SysPostRequest struct {
	Status   string `json:"status"`   // 角色状态（0正常 1停用）
	PostCode string `json:"postCode"` // 岗位编码
	PostName string `json:"postName"` // 岗位名称
}
