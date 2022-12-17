package cronjob

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lutasam/check_in_sys/biz/dal"
	"time"
)

func TestCron() {
	fmt.Println(time.Now())
}

func CleanUserRecordStatus() {
	c := &gin.Context{}
	err := dal.GetUserDal().UpdateAllUserStatus(c)
	if err != nil {
		panic(err)
	}
}
