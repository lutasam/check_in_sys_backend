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
	userInfo, err := utils.GetCtxUserInfoJWT(c)
	if err != nil {
		return nil, err
	}
	user, err := dal.GetUserDal().TakeUserByEmail(c, userInfo.Email)
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
	user.TodayRecordStatus = req.TodayRecordStatus
	user.TodayHealthCodeStatus = req.TodayHealthCodeStatus
	user.Identity = req.Identity
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
	users, err := dal.GetUserDal().FindUsers(c, req.CurrentPage, req.PageSize, req.TodayHealthCodeStatus, req.Name, req.TodayRecordStatus, true, departmentID)
	if err != nil {
		return nil, err
	}
	userStatuses, err := convertToUserStatusVOs(c, users)
	if err != nil {
		return nil, err
	}
	return &bo.FindAllUserStatusResponse{
		Total:        len(users),
		UserStatuses: userStatuses,
	}, nil
}

func (ins *UserService) FindAllUsers(c *gin.Context, req *bo.FindAllUsersRequest) (*bo.FindAllUsersResponse, error) {
	if req.CurrentPage < 0 || req.PageSize < 0 || req.PageSize > 100 {
		return nil, common.USERINPUTERROR
	}
	department, err := dal.GetDepartmentDal().TakeDepartmentByName(c, req.Department)
	if err != nil {
		return nil, err
	}
	users, err := dal.GetUserDal().FindUsers(c, req.CurrentPage, req.PageSize, common.ALLHEALTHCODE.Ints(), req.Name, false, false, department.ID)
	if err != nil {
		return nil, err
	}
	userVOs, err := convertToUserVOs(c, users)
	if err != nil {
		return nil, err
	}
	return &bo.FindAllUsersResponse{
		Total: len(users),
		Users: userVOs,
	}, nil
}

func (ins *UserService) DeleteUser(c *gin.Context, req *bo.DeleteUserRequest) (*bo.DeleteUserResponse, error) {
	id, err := utils.StringToUint64(req.UserID)
	if err != nil {
		return nil, err
	}
	err = dal.GetUserDal().DeleteUser(c, id)
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
			CreatedAt:  utils.TimeToString(user.CreatedAt),
			UpdatedAt:  utils.TimeToString(user.UpdatedAt),
		})
	}
	return vos, nil
}
