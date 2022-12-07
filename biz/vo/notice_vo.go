package vo

type NoticeVO struct {
	ID        string `json:"id"`
	Date      string `json:"date"`
	AdminName string `json:"admin_name"`
	Content   string `json:"content"`
}
