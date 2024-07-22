// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go
// Email: hot_kun@hotmail.com or 867917691@qq.com

// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSysUser = "sys_user"

// SysUser mapped from table <sys_user>
type SysUser struct {
	UserID      int64     `gorm:"column:user_id;primaryKey;comment:用户ID" json:"userId"`     // 用户ID
	DeptID      int64     `gorm:"column:dept_id;comment:部门ID" json:"deptId"`                // 部门ID
	UserName    string    `gorm:"column:user_name;not null;comment:用户账号" json:"userName"`   // 用户账号
	NickName    string    `gorm:"column:nick_name;not null;comment:用户昵称" json:"nickName"`   // 用户昵称
	UserType    string    `gorm:"column:user_type;comment:用户类型（00系统用户）" json:"userType"`    // 用户类型（00系统用户）
	Email       string    `gorm:"column:email;comment:用户邮箱" json:"email"`                   // 用户邮箱
	Phonenumber string    `gorm:"column:phonenumber;comment:手机号码" json:"phonenumber"`       // 手机号码
	Sex         string    `gorm:"column:sex;comment:用户性别（0男 1女 2未知）" json:"sex"`            // 用户性别（0男 1女 2未知）
	Avatar      string    `gorm:"column:avatar;comment:头像地址" json:"avatar"`                 // 头像地址
	Password    string    `gorm:"column:password;comment:密码" json:"password"`               // 密码
	Status      string    `gorm:"column:status;comment:帐号状态（0正常 1停用）" json:"status"`        // 帐号状态（0正常 1停用）
	DelFlag     string    `gorm:"column:del_flag;comment:删除标志（0代表存在 2代表删除）" json:"delFlag"` // 删除标志（0代表存在 2代表删除）
	LoginIP     string    `gorm:"column:login_ip;comment:最后登录IP" json:"loginIp"`            // 最后登录IP
	LoginDate   time.Time `gorm:"column:login_date;comment:最后登录时间" json:"loginDate"`        // 最后登录时间
	CreateBy    string    `gorm:"column:create_by;comment:创建者" json:"createBy"`             // 创建者
	CreateTime  time.Time `gorm:"column:create_time;comment:创建时间" json:"createTime"`        // 创建时间
	UpdateBy    string    `gorm:"column:update_by;comment:更新者" json:"updateBy"`             // 更新者
	UpdateTime  time.Time `gorm:"column:update_time;comment:更新时间" json:"updateTime"`        // 更新时间
	Remark      string    `gorm:"column:remark;comment:备注" json:"remark"`                   // 备注
}

// TableName SysUser's table name
func (*SysUser) TableName() string {
	return TableNameSysUser
}

// SysUser mapped from table <sys_user>
type SysUserRequest struct {
	Status      string `json:"status"`
	DeptID      int64  `json:"deptId"`
	Phonenumber string `json:"phonenumber"`
	UserName    string `json:"userName"`
	BeginTime   string `json:"beginTime"`
	EndTime     string `json:"endTime"`
}

type ChangeUserStatusRequest struct {
	UserID int64  `json:"userId" validate:"required"`
	Status string `json:"status" validate:"required"`
}

type ResetUserPwdRequest struct {
	UserID   int64  `json:"userId" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type GetUserInfoSuccess struct {
	Code    int             `json:"code"`
	Message string          `json:"msg"`
	User    *UserInfoStruct `json:"data"`
	RoleIds []int64         `gorm:"-" json:"roleIds,omitempty"`
	PostIds []int64         `gorm:"-" json:"postIds,omitempty"`
	Roles   []*SysRole      `gorm:"-" json:"roles,omitempty"`
	Posts   []*SysPost      `gorm:"-" json:"posts,omitempty"`
}
