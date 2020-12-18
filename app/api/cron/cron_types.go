package cron

type AddCron struct {
    CronName string `v:"required#请输入定时任务名"`
    Num      string `v:"required#请输入数字"`
    Unit     string `v:"required#请输入数字单位"`
    Content  string `v:"required#请输入内容"`
}

type DeleteOrSearchCron struct {
    CronName string `v:"required#请输入定时任务名"`
}
