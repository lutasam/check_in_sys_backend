package bo

import "github.com/lutasam/check_in_sys/biz/vo"

type UpdateUserInfoRequest struct {
	Email                 string `json:"email" binding:"required"`
	Name                  string `json:"name" binding:"-"`
	Avatar                string `json:"avatar" binding:"-"`
	Department            string `json:"department" binding:"-"`
	TodayRecordStatus     *bool  `json:"today_record_status" binding:"-"`
	TodayHealthCodeStatus *int   `json:"today_health_code_status"  binding:"-"`
	Identity              *int   `json:"identity" binding:"-"`
}

type UpdateUserInfoResponse struct{}

type FindAllUserStatusRequest struct {
	CurrentPage           int    `json:"current_page" binding:"required"`
	PageSize              int    `json:"page_size" binding:"required"`
	Name                  string `json:"name" binding:"-"`
	NeedRecordStatus      *bool  `json:"need_record_status" binding:"-"`
	TodayRecordStatus     *bool  `json:"today_record_status" binding:"-"`
	TodayHealthCodeStatus *int   `json:"today_health_code_status"  binding:"-"`
}

type FindAllUserStatusResponse struct {
	Total        int                `json:"total"`
	UserStatuses []*vo.UserStatusVO `json:"user_statuses"`
}

type FindAllUsersRequest struct {
	CurrentPage int    `json:"current_page" binding:"required"`
	PageSize    int    `json:"page_size" binding:"required"`
	Name        string `json:"name" binding:"-"`
	Department  string `json:"department" binding:"-"`
}

type FindAllUsersResponse struct {
	Total int          `json:"total"`
	Users []*vo.UserVO `json:"users"`
}

type DeleteUserRequest struct {
	UserID string `json:"user_id"`
}

type DeleteUserResponse struct{}

type AddUserRequest struct {
	Email      string `json:"email" binding:"required"`
	Password   string `json:"password" binding:"required"`
	Name       string `json:"name" binding:"required"`
	Department string `json:"department" binding:"required"`
}

type AddUserResponse struct{}

type ApplyChangeUserEmailRequest struct {
	Email string `json:"email" binding:"required"`
}

type ApplyChangeUserEmailResponse struct{}

type ActiveChangeUserEmailRequest struct {
	Email      string `json:"email" binding:"required"`
	ActiveCode string `json:"active_code" binding:"required"`
}

type ActiveChangeUserEmailResponse struct{}

type TakeUserInfoRequest struct {
	UserID string `json:"user_id" binding:"-"`
}

type TakeUserInfoResponse struct {
	User *vo.UserVO `json:"user"`
}

type FindAllAdminsRequest struct{}

type FindAllAdminsResponse struct {
	Admins []*vo.AdminVO `json:"admins"`
}
