package dal

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/check_in_sys/biz/common"
	"github.com/lutasam/check_in_sys/biz/model"
	"github.com/lutasam/check_in_sys/biz/repository"
	"github.com/lutasam/check_in_sys/biz/vo"
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

func (ins *UserDal) FindUsers(c *gin.Context, currentPage, pageSize, healthCodeStatus int, name string, recordStatus, needRecordStatus bool, departmentID uint64) ([]*model.User, int64, error) {
	var users []*model.User
	var count int64
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
	err := sql.Where("identity != ?", common.SUPER_ADMIN).Count(&count).Limit(pageSize).Offset((currentPage - 1) * pageSize).Find(&users).Error
	if err != nil {
		return nil, 0, common.DATABASEERROR
	}
	return users, count, nil
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

func (ins *UserDal) FindAllDepartmentAdmins(c *gin.Context) ([]*model.User, error) {
	var users []*model.User
	err := repository.GetDB().WithContext(c).Table(model.User{}.TableName()).
		Where("identity = ?", common.DEPARTMENT_ADMIN.Ints()).
		Find(&users).Error
	if err != nil {
		return nil, common.DATABASEERROR
	}
	return users, nil
}

func (ins *UserDal) FindAllUsersInDepartmentGroup(c *gin.Context, departmentIDs []uint64) ([]*model.User, error) {
	var users []*model.User
	sql := repository.GetDB().WithContext(c).Table(model.User{}.TableName())
	if len(departmentIDs) > 0 && departmentIDs[0] != common.ALLDEPARTMENTS {
		sql = sql.Where("department_id in ?", departmentIDs)
	}
	err := sql.Where("identity = ?", common.USER.Ints()).Find(&users).Error
	if err != nil {
		return nil, common.DATABASEERROR
	}
	return users, nil
}

type SummaryResult struct {
	HealthCode int
	PeopleNum  int
}

// SummaryHealthCodeStatusInSpecificUserGroup summary and analyze user's health code status in specific group
func (ins *UserDal) SummaryHealthCodeStatusInSpecificUserGroup(c *gin.Context, userIDs []uint64) (int, int, []*vo.HealthCodePartVO, error) {
	var recordNum int64
	err := repository.GetDB().WithContext(c).Table(model.User{}.TableName()).
		Where("id in ? and today_record_status = ?", userIDs, true).Count(&recordNum).Error
	if err != nil {
		return 0, 0, nil, common.DATABASEERROR
	}

	var parts []*SummaryResult
	err = repository.GetDB().WithContext(c).Table(model.User{}.TableName()).
		Select("today_health_code_status as health_code, count(*) as people_num").Group("health_code").
		Where("id in ?", userIDs).Scan(&parts).Error
	if err != nil {
		return 0, 0, nil, common.DATABASEERROR
	}

	abnormalNum := 0
	var healthCodeParts []*vo.HealthCodePartVO
	for _, part := range parts {
		if part.HealthCode != common.GREEN.Ints() {
			abnormalNum += part.PeopleNum
		}
		healthCodeParts = append(healthCodeParts, &vo.HealthCodePartVO{
			HealthCode: common.ParseHealthCodeStatus(part.HealthCode).String(),
			PeopleNum:  part.PeopleNum,
		})
	}

	return int(recordNum), abnormalNum, healthCodeParts, nil
}
