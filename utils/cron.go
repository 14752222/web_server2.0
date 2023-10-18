package utils

import "github.com/robfig/cron/v3"

// cron 定时任务
var cronTask *cron.Cron

func init() {
	cronTask = cron.New()
	cronTask.Start()
}

// AddCronTask 添加定时任务
func AddCronTask(spec string, cmd func()) (cron.EntryID, error) {
	return cronTask.AddFunc(spec, cmd)
}

// RemoveCronTask 移除定时任务
func RemoveCronTask(id cron.EntryID) {
	cronTask.Remove(id)
}

// StopCronTask 停止定时任务
func StopCronTask() {
	cronTask.Stop()
}

// StartCronTask 启动定时任务
func StartCronTask() {
	cronTask.Start()
}

// AddCronJob 添加重复执行任务
func AddCronJob(spec string, cmd func()) (cron.EntryID, error) {
	return cronTask.AddJob(spec, cron.FuncJob(cmd))
}

// AddCronJobs 添加多个重复执行任务
func AddCronJobs(spec string, cmd []func()) ([]cron.EntryID, error) {
	var entries []cron.EntryID
	for _, v := range cmd {
		entry, err := cronTask.AddJob(spec, cron.FuncJob(v))
		if err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}
	return entries, nil
}
