// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go or https://gitee.com/gitee_kun/RuoYi-Go
// Email: hot_kun@hotmail.com or 867917691@qq.com

package usecase

import (
	"RuoYi-Go/internal/common"
	"RuoYi-Go/internal/domain/model"
	"RuoYi-Go/internal/ports/input"
	"RuoYi-Go/internal/ports/output"
	"RuoYi-Go/pkg/cache"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
)

type SysUserService struct {
	repo     output.SysUserRepository
	roleRepo output.SysRoleRepository
	deptRepo output.SysDeptRepository
	cache    *cache.FreeCacheClient
	logger   *zap.Logger
}

func NewPageSysUserService(repo output.SysUserRepository, roleRepo output.SysRoleRepository, deptRepo output.SysDeptRepository, cache *cache.FreeCacheClient, logger *zap.Logger) input.SysUserService {
	return &SysUserService{repo: repo, roleRepo: roleRepo, deptRepo: deptRepo, cache: cache, logger: logger}
}

func NewSysUserService(repo output.SysUserRepository, cache *cache.FreeCacheClient, logger *zap.Logger) input.SysUserService {
	return &SysUserService{repo: repo, cache: cache, logger: logger}
}

func (this *SysUserService) QueryUserByUserName(username string) (*model.SysUser, error) {
	structEntity := &model.SysUser{}
	// 尝试从缓存中获取
	userBytes, err := this.cache.Get([]byte(fmt.Sprintf("UserName:%s", username)))
	if err == nil {
		// 缓存命中
		err = json.Unmarshal(userBytes, &structEntity)
		if err == nil && structEntity.UserID != 0 {
			// 缓存命中
			return structEntity, nil
		}
	}

	structEntity, err = this.repo.QueryUserByUserName(username)
	if err != nil {
		this.logger.Error("查询用户信息失败", zap.Error(err))
		return nil, err
	} else if structEntity.UserID != 0 {
		// 序列化用户对象并存入缓存
		userBytes, err = json.Marshal(structEntity)
		if err == nil {
			this.cache.Set([]byte(fmt.Sprintf("UserName:%s", username)), userBytes, common.EXPIRESECONDS)          // 第三个参数是过期时间，0表示永不过期
			this.cache.Set([]byte(fmt.Sprintf("UserID:%d", structEntity.UserID)), userBytes, common.EXPIRESECONDS) // 第三个参数是过期时间，0表示永不过期
		}
		return structEntity, nil
	}

	this.logger.Debug("查询用户信息失败", zap.Error(err))
	return nil, fmt.Errorf("查询用户信息失败", zap.Error(err))
}

func (this *SysUserService) QueryUserLikeUserName(username string) ([]*model.SysUser, error) {
	structEntity, err := this.repo.QueryUserLikeUserName(username)
	if err != nil {
		this.logger.Error("查询用户信息失败", zap.Error(err))
		return nil, err
	}
	return structEntity, nil
}

func (this *SysUserService) QueryUserByUserId(userId int64) (*model.SysUser, error) {
	structEntity := &model.SysUser{}
	if userId == 0 {
		return structEntity, nil
	}

	// 尝试从缓存中获取
	userBytes, err := this.cache.Get([]byte(fmt.Sprintf("UserID:%d", userId)))
	if err == nil {
		// 缓存命中
		err = json.Unmarshal(userBytes, &structEntity)
		if err == nil && structEntity.UserID != 0 {
			// 缓存命中
			return structEntity, nil
		}
	}

	structEntity, err = this.repo.QueryUserByUserId(userId)
	if err != nil {
		this.logger.Error("查询用户信息失败", zap.Error(err))
		return nil, err
	} else if structEntity.UserID != 0 {
		// 序列化用户对象并存入缓存
		userBytes, err = json.Marshal(structEntity)
		if err == nil {
			this.cache.Set([]byte(fmt.Sprintf("UserID:%d", structEntity.UserID)), userBytes, common.EXPIRESECONDS) // 第三个参数是过期时间，0表示永不过期
		}
		return structEntity, nil
	}

	this.logger.Debug("查询用户信息失败", zap.Error(err))
	return nil, fmt.Errorf("查询用户信息失败", zap.Error(err))
}

func (this *SysUserService) QueryUserPage(pageReq common.PageRequest, u *model.SysUserRequest) ([]*model.UserInfoStruct, int64, error) {
	data, total, err := this.repo.QueryUserPage(pageReq, u)
	if err != nil {
		this.logger.Error("查询用户分页信息失败", zap.Error(err))
		return nil, 0, err
	}

	userList := make([]*model.UserInfoStruct, 0)
	for _, user := range data {
		userInfo := &model.UserInfoStruct{}
		userInfo.SysUser = user
		if user.UserID == common.ADMINID {
			userInfo.Admin = true
		}

		roles, err := this.roleRepo.QueryRolesByUserId(user.UserID)
		if err != nil {
			this.logger.Error("QueryRolesByUserId error,", zap.Error(err))
			return nil, 0, fmt.Errorf("getInfo error", zap.Error(err))
		}
		userInfo.Roles = roles

		dept, err := this.deptRepo.QueryDeptById(user.DeptID)
		if err != nil {
			this.logger.Error("QueryDeptById error,", zap.Error(err))
			return nil, 0, fmt.Errorf("getInfo error", zap.Error(err))
		}
		userInfo.Dept = dept
		userList = append(userList, userInfo)
	}

	return userList, total, nil
}

func (this *SysUserService) QueryUserList(u *model.SysUserRequest) ([]*model.SysUser, error) {
	data, err := this.repo.QueryUserList(u)
	if err != nil {
		this.logger.Error("查询用户列表信息失败", zap.Error(err))
		return nil, err
	}
	return data, nil
}

func (this *SysUserService) DeleteUserByUserId(userId int64) (int64, error) {
	result, err := this.repo.DeleteUserByUserId(userId)
	if err != nil {
		this.logger.Error("删除用户信息失败", zap.Error(err))
		return 0, err
	}
	if result > 0 {
		this.cache.Del(fmt.Sprintf("UserID:%d", userId))
	}
	return result, nil
}

func (this *SysUserService) ChangeUserStatus(user *model.ChangeUserStatusRequest) (int64, error) {
	result, err := this.repo.ChangeUserStatus(user)
	if err != nil {
		this.logger.Error("修改用户状态失败", zap.Error(err))
		return 0, err
	}
	if result > 0 {
		structEntity, err := this.repo.QueryUserByUserId(user.UserID)
		if err != nil {
			this.logger.Error("查询用户信息失败", zap.Error(err))
			return 0, err
		} else if structEntity.UserID != 0 {
			// 序列化用户对象并存入缓存
			userBytes, err := json.Marshal(structEntity)
			if err == nil {
				this.cache.Set([]byte(fmt.Sprintf("UserID:%d", structEntity.UserID)), userBytes, common.EXPIRESECONDS) // 第三个参数是过期时间，0表示永不过期
			}
			return result, nil
		}
	}
	return result, nil
}

func (this *SysUserService) ResetUserPwd(user *model.ResetUserPwdRequest) (int64, error) {
	result, err := this.repo.ResetUserPwd(user)
	if err != nil {
		this.logger.Error("修改用户状态失败", zap.Error(err))
		return 0, err
	}
	return result, nil
}

func (this *SysUserService) AddUser(post *model.SysUser) (*model.SysUser, error) {
	data, err := this.repo.AddUser(post)
	if err != nil {
		this.logger.Error("AddUser", zap.Error(err))
		return nil, err
	}
	if data != nil && data.UserID != 0 {
		// 序列化用户对象并存入缓存
		userBytes, err := json.Marshal(data)
		if err == nil {
			this.cache.Set([]byte(fmt.Sprintf("UserID:%d", data.UserID)), userBytes, common.EXPIRESECONDS) // 第三个参数是过期时间，0表示永不过期
		}
	}
	return data, nil
}

func (this *SysUserService) EditUser(post *model.SysUser) (*model.SysUser, int64, error) {
	data, result, err := this.repo.EditUser(post)
	if err != nil {
		this.logger.Error("EditUser", zap.Error(err))
		return nil, 0, err
	}
	if data != nil && data.UserID != 0 && result == 1 {
		// 序列化用户对象并存入缓存
		userBytes, err := json.Marshal(data)
		if err == nil {
			this.cache.Set([]byte(fmt.Sprintf("UserID:%d", data.UserID)), userBytes, common.EXPIRESECONDS) // 第三个参数是过期时间，0表示永不过期
		}
	}
	return data, result, nil
}

func (this *SysUserService) CheckUserNameUnique(id int64, name string) (int64, error) {
	result, err := this.repo.CheckUserNameUnique(id, name)
	if err != nil {
		this.logger.Error("CheckUserNameUnique", zap.Error(err))
		return -1, err
	}
	return result, nil
}

func (this *SysUserService) UserLogin(user *model.SysUser) (int64, error) {
	result, err := this.repo.UserLogin(user)
	if err != nil {
		this.logger.Error("UserLogin", zap.Error(err))
		return -1, err
	}
	return result, nil
}
