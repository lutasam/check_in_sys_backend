package service

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/check_in_sys/biz/bo"
	"github.com/lutasam/check_in_sys/biz/common"
	"github.com/lutasam/check_in_sys/biz/dal"
	"github.com/lutasam/check_in_sys/biz/model"
	"github.com/lutasam/check_in_sys/biz/utils"
	"github.com/lutasam/check_in_sys/biz/vo"
	"sync"
)

type StatisticService struct{}

var (
	statisticService     *StatisticService
	statisticServiceOnce sync.Once
)

func GetStatisticService() *StatisticService {
	statisticServiceOnce.Do(func() {
		statisticService = &StatisticService{}
	})
	return statisticService
}

func (ins *StatisticService) TakeUserStatistic(c *gin.Context, req *bo.TakeUserStatisticRequest) (*bo.TakeUserStatisticResponse, error) {
	userInfo, err := utils.GetCtxUserInfoJWT(c)
	if err != nil {
		return nil, err
	}
	user, err := dal.GetUserDal().TakeUserByID(c, userInfo.UserID)
	if err != nil {
		return nil, err
	}
	total, err := dal.GetRecordDal().CountUserRecords(c, user.ID)
	if err != nil {
		return nil, err
	}
	notices, err := dal.GetNoticeDal().FindAllThreeDaysNotice(c)
	if err != nil {
		return nil, err
	}
	return &bo.TakeUserStatisticResponse{
		IsTodayRecordFinished: user.TodayRecordStatus,
		RecordTimes:           int(total),
		Notices:               convertToNoticeVO(notices),
	}, nil
}

func (ins *StatisticService) TakeAdminStatistic(c *gin.Context, req *bo.TakeAdminStatisticRequest) (*bo.TakeAdminStatisticResponse, error) {
	userInfo, err := utils.GetCtxUserInfoJWT(c)
	if err != nil {
		return nil, err
	}
	var departmentIDs []uint64
	if userInfo.Identity == common.SUPER_ADMIN.Ints() {
		departmentIDs = append(departmentIDs, common.ALLDEPARTMENTS)
	} else {
		departments, err := dal.GetDepartmentDal().FindAllDepartmentsByAdminID(c, userInfo.UserID)
		if err != nil {
			return nil, err
		}
		for _, department := range departments {
			departmentIDs = append(departmentIDs, department.ID)
		}
	}
	users, err := dal.GetUserDal().FindAllUsersInDepartmentGroup(c, departmentIDs)
	if err != nil {
		return nil, err
	}
	var userIDs []uint64
	for _, user := range users {
		userIDs = append(userIDs, user.ID)
	}
	peopleNum := len(userIDs)
	recordNum, abnormalNum, healthCodePartVOs, err := dal.GetUserDal().SummaryHealthCodeStatusInSpecificUserGroup(c, userIDs)
	if err != nil {
		return nil, err
	}
	return &bo.TakeAdminStatisticResponse{
		PeopleNum:         peopleNum,
		AbnormalNum:       abnormalNum,
		FinishRecordNum:   recordNum,
		UnFinishRecordNum: peopleNum - recordNum,
		FinishPercentage:  int((float64(recordNum) / float64(peopleNum)) * 100),
		HealthCodeParts:   healthCodePartVOs,
	}, nil
}

func convertToNoticeVO(notices []*model.Notice) []*vo.NoticeVO {
	var vos []*vo.NoticeVO
	for _, notice := range notices {
		vos = append(vos, &vo.NoticeVO{
			ID:        utils.Uint64ToString(notice.ID),
			Date:      utils.TimeToDateString(notice.CreatedAt),
			AdminName: notice.User.Name,
			Content:   notice.Content,
		})
	}
	return vos
}
