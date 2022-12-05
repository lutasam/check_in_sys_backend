package dal

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/check_in_sys/biz/common"
	"github.com/lutasam/check_in_sys/biz/model"
	"github.com/lutasam/check_in_sys/biz/repository"
	"sync"
)

type UserDal struct{}

var (
	userDal     *UserDal
	userDalOnce sync.Once
)

func GetUserDal() *UserDal {
	userDalOnce.Do(func() {
		userDal = &UserDal{}
	})
	return userDal
}

// TakeUserByEmail if there is no this email in database, it will return error
func (ins *UserDal) TakeUserByEmail(c *gin.Context, email string) (*model.User, error) {
	user := &model.User{}
	err := repository.GetDB().WithContext(c).Table(user.TableName()).Where("email = ?", email).Find(user).Error
	if err != nil {
		return nil, common.DATABASEERROR
	}
	if user.ID == 0 {
		return nil, common.USERDOESNOTEXIST
	}
	return user, nil
}

func (ins *UserDal) CreateUser(c *gin.Context, user *model.User) error {
	err := repository.GetDB().WithContext(c).Table(user.TableName()).Create(user).Error
	if err != nil {
		return common.DATABASEERROR
	}
	return nil
}

func (ins *UserDal) UpdateUser(c *gin.Context, user *model.User) error {
	err := repository.GetDB().WithContext(c).Model(user).Updates(user).Error
	if err != nil {
		return common.DATABASEERROR
	}
	return nil
}

func (ins *UserDal) FindUsers(c *gin.Context, currentPage, pageSize, healthCodeStatus int, name string, recordStatus, needRecordStatus bool, departmentID uint64) ([]*model.User, error) {
	var users []*model.User
	sql := repository.GetDB().WithContext(c).Table(model.User{}.TableName())
	if healthCodeStatus != common.ALLHEALTHCODE.Ints() {
		sql = sql.Where("today_health_code_status = ?", healthCodeStatus)
	}
	if name != "" {
		sql = sql.Where("name like ?", "%"+name+"%")
	}
	if needRecordStatus {
		sql = sql.Where("today_record_status = ?", recordStatus)
	}
	if departmentID != common.ALLDEPARTMENTS {
		sql = sql.Where("department_id = ?", departmentID)
	}
	err := sql.Where("identity != ?", common.SUPER_ADMIN).Limit(pageSize).Offset((currentPage - 1) * pageSize).Find(&users).Error
	if err != nil {
		return nil, common.DATABASEERROR
	}
	return users, nil
}

func (ins *UserDal) DeleteUser(c *gin.Context, userID uint64) error {
	err := repository.GetDB().WithContext(c).Table(model.User{}.TableName()).Where("id = ?", userID).Delete(&model.User{}).Error
	if err != nil {
		return common.DATABASEERROR
	}
	return nil
}

func (ins *UserDal) TakeUserByID(c *gin.Context, userID uint64) (*model.User, error) {
	user := &model.User{}
	err := repository.GetDB().WithContext(c).Table(user.TableName()).Where("id = ?", userID).Find(user).Error
	if err != nil {
		return nil, common.DATABASEERROR
	}
	if user.ID == 0 {
		return nil, common.USERDOESNOTEXIST
	}
	return user, nil
}

func (ins *UserDal) TakeSuperAdmin(c *gin.Context) (*model.User, error) {
	user := &model.User{}
	err := repository.GetDB().WithContext(c).Table(user.TableName()).
		Where("identity = ?", common.SUPER_ADMIN.Ints()).Find(user).Error
	if err != nil {
		return nil, common.DATABASEERROR
	}
	return user, nil
}

func (ins *UserDal) TakeAllDepartmentAdmins(c *gin.Context, currentPage, pageSize int) ([]*model.User, error) {
	var users []*model.User
	err := repository.GetDB().WithContext(c).Table(model.User{}.TableName()).
		Where("identity == ?", common.DEPARTMENT_ADMIN.Ints()).
		Limit(pageSize).Offset((currentPage - 1) * pageSize).Find(&users).Error
	if err != nil {
		return nil, common.DATABASEERROR
	}
	return users, nil
}
