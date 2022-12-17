package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lutasam/check_in_sys/biz/bo"
	"github.com/lutasam/check_in_sys/biz/common"
	"github.com/lutasam/check_in_sys/biz/dal"
	"github.com/lutasam/check_in_sys/biz/model"
	"github.com/lutasam/check_in_sys/biz/utils"
	"github.com/lutasam/check_in_sys/biz/vo"
	"sync"
)

type UserService struct{}

var (
	userService     *UserService
	userServiceOnce sync.Once
)

func GetUserService() *UserService {
	userServiceOnce.Do(func() {
		userService = &UserService{}
	})
	return userService
}

func (ins *UserService) UpdateUserInfo(c *gin.Context, req *bo.UpdateUserInfoRequest) (*bo.UpdateUserInfoResponse, error) {
	if req.Avatar != "" && !utils.IsValidURL(req.Avatar) {
		return nil, common.USERINPUTERROR
	}
	user, err := dal.GetUserDal().TakeUserByEmail(c, req.Email)
	if err != nil {
		return nil, err
	}
	department, err := dal.GetDepartmentDal().TakeDepartmentByName(c, req.Department)
	if err != nil {
		return nil, err
	}
	user.Name = req.Name
	user.Avatar = req.Avatar
	user.DepartmentID = department.ID
	if req.TodayRecordStatus != nil {
		user.TodayRecordStatus = *req.TodayRecordStatus
	}
	if req.TodayHealthCodeStatus != nil {
		user.TodayHealthCodeStatus = *req.TodayHealthCodeStatus
	}
	if req.Identity != nil {
		user.Identity = *req.Identity
	}
	err = dal.GetUserDal().UpdateUser(c, user)
	if err != nil {
		return nil, err
	}
	return &bo.UpdateUserInfoResponse{}, nil
}

func (ins *UserService) FindAllUserStatus(c *gin.Context, req *bo.FindAllUserStatusRequest) (*bo.FindAllUserStatusResponse, error) {
	if req.CurrentPage < 0 || req.PageSize < 0 || req.PageSize > 100 {
		return nil, common.USERINPUTERROR
	}
	userInfo, err := utils.GetCtxUserInfoJWT(c)
	if err != nil {
		return nil, err
	}
	var departmentID uint64
	if userInfo.Identity == common.SUPER_ADMIN.Ints() {
		departmentID = common.ALLDEPARTMENTS
	} else {
		user, err := dal.GetUserDal().TakeUserByEmail(c, userInfo.Email)
		if err != nil {
			return nil, err
		}
		departmentID = user.DepartmentID
	}
	var todayHealthCodeStatus int
	var todayRecordStatus, needRecordStatus bool
	if req.TodayHealthCodeStatus == nil {
		todayHealthCodeStatus = common.ALLHEALTHCODE.Ints()
	} else {
		todayHealthCodeStatus = *req.TodayHealthCodeStatus
	}
	if req.TodayRecordStatus == nil {
		todayRecordStatus = false
	} else {
		todayRecordStatus = *req.TodayRecordStatus
	}
	if req.NeedRecordStatus == nil {
		needRecordStatus = false
	} else {
		needRecordStatus = *req.NeedRecordStatus
	}
	users, total, err := dal.GetUserDal().FindUsers(c, req.CurrentPage, req.PageSize, todayHealthCodeStatus, req.Name, todayRecordStatus, needRecordStatus, departmentID)
	if err != nil {
		return nil, err
	}
	userStatuses, err := convertToUserStatusVOs(c, users)
	if err != nil {
		return nil, err
	}
	return &bo.FindAllUserStatusResponse{
		Total:        int(total),
		UserStatuses: userStatuses,
	}, nil
}

func (ins *UserService) FindAllUsers(c *gin.Context, req *bo.FindAllUsersRequest) (*bo.FindAllUsersResponse, error) {
	if req.CurrentPage < 0 || req.PageSize < 0 || req.PageSize > 100 {
		return nil, common.USERINPUTERROR
	}
	var departmentID uint64
	if req.Department == "" { // 超级管理员可查看所有部门
		departmentID = common.ALLDEPARTMENTS
	} else { // 部门管理员只能查看自己部门
		department, err := dal.GetDepartmentDal().TakeDepartmentByName(c, req.Department)
		if err != nil {
			return nil, err
		}
		departmentID = department.ID
	}
	users, total, err := dal.GetUserDal().FindUsers(c, req.CurrentPage, req.PageSize, common.ALLHEALTHCODE.Ints(), req.Name, false, false, departmentID)
	if err != nil {
		return nil, err
	}
	userVOs, err := convertToUserVOs(c, users)
	if err != nil {
		return nil, err
	}
	return &bo.FindAllUsersResponse{
		Total: int(total),
		Users: userVOs,
	}, nil
}

func (ins *UserService) DeleteUser(c *gin.Context, req *bo.DeleteUserRequest) (*bo.DeleteUserResponse, error) {
	id, err := utils.StringToUint64(req.UserID)
	if err != nil {
		return nil, err
	}
	user, err := dal.GetUserDal().TakeUserByID(c, id)
	if err != nil {
		return nil, err
	}
	if user.DepartmentID == uint64(common.DEPARTMENT_ADMIN.Ints()) || user.DepartmentID == uint64(common.SUPER_ADMIN.Ints()) {
		return nil, common.CANNOTDELETEADMIN
	}
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		err1 := dal.GetUserDal().DeleteUser(c, user)
		if err1 != nil {
			err = err1
		}
	}()
	go func() {
		err1 := dal.GetRecordDal().DeleteUserAllRecords(c, user)
		if err1 != nil {
			err = err1
		}
	}()
	wg.Wait()

	if err != nil {
		return nil, err
	}
	return &bo.DeleteUserResponse{}, nil
}

func (ins *UserService) AddUser(c *gin.Context, req *bo.AddUserRequest) (*bo.AddUserResponse, error) {
	if !utils.IsValidEmail(req.Email) {
		return nil, common.USERINPUTERROR
	}
	_, err := dal.GetUserDal().TakeUserByEmail(c, req.Email)
	if err != nil && errors.Is(err, common.DATABASEERROR) {
		return nil, err
	}
	if err == nil {
		return nil, common.USEREXISTED
	}
	encryptPass, err := utils.EncryptPassword(req.Password)
	if err != nil {
		return nil, err
	}
	department, err := dal.GetDepartmentDal().TakeDepartmentByName(c, req.Department)
	if err != nil {
		return nil, err
	}
	user := &model.User{
		ID:                    utils.GenerateUserID(),
		Email:                 req.Email,
		Password:              encryptPass,
		Name:                  req.Name,
		DepartmentID:          department.ID,
		Avatar:                common.DEFAULTAVATARURL,
		TodayRecordStatus:     false,
		TodayHealthCodeStatus: common.GREEN.Ints(),
		Identity:              common.USER.Ints(),
	}
	err = dal.GetUserDal().CreateUser(c, user)
	if err != nil {
		return nil, err
	}
	return &bo.AddUserResponse{}, nil
}

func (ins *UserService) TakeUserInfo(c *gin.Context, req *bo.TakeUserInfoRequest) (*bo.TakeUserInfoResponse, error) {
	var userID uint64
	if req.UserID == "" {
		userInfo, err := utils.GetCtxUserInfoJWT(c)
		if err != nil {
			return nil, err
		}
		userID = userInfo.UserID
	} else {
		id, err := utils.StringToUint64(req.UserID)
		if err != nil {
			return nil, err
		}
		userID = id
	}
	user, err := dal.GetUserDal().TakeUserByID(c, userID)
	if err != nil {
		return nil, err
	}
	department, err := dal.GetDepartmentDal().TakeDepartmentByID(c, user.DepartmentID)
	if err != nil {
		return nil, err
	}
	return &bo.TakeUserInfoResponse{
		User: &vo.UserVO{
			ID:                    utils.Uint64ToString(user.ID),
			Email:                 user.Email,
			Name:                  user.Name,
			Department:            department.Name,
			Identity:              common.ParseIdentity(user.Identity).String(),
			Avatar:                user.Avatar,
			TodayRecordStatus:     user.TodayRecordStatus,
			TodayHealthCodeStatus: common.ParseHealthCodeStatus(user.TodayHealthCodeStatus).Ints(),
			CreatedAt:             utils.TimeToString(user.CreatedAt),
			UpdatedAt:             utils.TimeToString(user.UpdatedAt),
		},
	}, nil
}

func (ins *UserService) FindAllAdmins(c *gin.Context, req *bo.FindAllAdminsRequest) (*bo.FindAllAdminsResponse, error) {
	admins, err := dal.GetUserDal().FindAllDepartmentAdmins(c)
	if err != nil {
		return nil, err
	}
	return &bo.FindAllAdminsResponse{
		Admins: convertToAdminVOs(admins),
	}, nil
}

func convertToUserStatusVOs(c *gin.Context, users []*model.User) ([]*vo.UserStatusVO, error) {
	var vos []*vo.UserStatusVO
	for _, user := range users {
		department, err := dal.GetDepartmentDal().TakeDepartmentByID(c, user.DepartmentID)
		if err != nil {
			return nil, err
		}
		vos = append(vos, &vo.UserStatusVO{
			ID:                    utils.Uint64ToString(user.ID),
			Name:                  user.Name,
			Department:            department.Name,
			TodayRecordStatus:     user.TodayRecordStatus,
			TodayHealthCodeStatus: common.ParseHealthCodeStatus(user.TodayHealthCodeStatus).String(),
			UpdatedAt:             utils.TimeToString(user.UpdatedAt),
		})
	}
	return vos, nil
}

func convertToUserVOs(c *gin.Context, users []*model.User) ([]*vo.UserVO, error) {
	var vos []*vo.UserVO
	for _, user := range users {
		department, err := dal.GetDepartmentDal().TakeDepartmentByID(c, user.DepartmentID)
		if err != nil {
			return nil, err
		}
		vos = append(vos, &vo.UserVO{
			ID:         utils.Uint64ToString(user.ID),
			Email:      user.Email,
			Name:       user.Name,
			Department: department.Name,
			Identity:   common.ParseIdentity(user.Identity).String(),
			Avatar:     user.Avatar,
			CreatedAt:  utils.TimeToString(user.CreatedAt),
			UpdatedAt:  utils.TimeToString(user.UpdatedAt),
		})
	}
	return vos, nil
}

func convertToAdminVOs(admins []*model.User) []*vo.AdminVO {
	var vos []*vo.AdminVO
	for _, admin := range admins {
		vos = append(vos, &vo.AdminVO{
			ID:   utils.Uint64ToString(admin.ID),
			Name: admin.Name,
		})
	}
	return vos
}
