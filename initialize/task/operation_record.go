package task

import (
	"github.com/bearllflee/go_shop/global"
	"github.com/bearllflee/go_shop/task"
)

func ClearOperationRecord(cronString string) {
	global.Cron.AddFunc(cronString, func() {
		task.ClearOperationRecord()
	})
}
