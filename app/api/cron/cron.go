package cron

import (
    "github.com/gogf/gf/net/ghttp"
    "github.com/gogf/gf/os/gcron"
    "github.com/gogf/gf/text/gstr"
    "kingProject/library/response"
    "kingProject/utils"
)

// 定时任务对象
type C struct{}

func (c *C) AddCron(r *ghttp.Request) {
    var (
        data *AddCron
    )
    if err := r.Parse(&data); err != nil {
        response.JsonExit(r, 1, err.Error())
    }
    if cron := gcron.Search(data.CronName); cron != nil {
        response.JsonExit(r, 1, "任务名已存在！")
    }
    // 每隔Num的Unit运行
    pattern := gstr.Join([]string{"@every ", data.Num, data.Unit}, "")
    gcron.AddTimes(pattern, 1, func() {
        utils.SendWechat(data.Content)
    }, data.CronName)
    response.JsonExit(r, 0, "ok")
}

func (c *C) DeleteCron(r *ghttp.Request) {
    var (
        data *DeleteOrSearchCron
    )
    if err := r.Parse(&data); err != nil {
        response.JsonExit(r, 1, err.Error())
    }
    if cron := gcron.Search(data.CronName); cron == nil {
        response.JsonExit(r, 1, "任务不存在！")
    } else {
        gcron.Remove(data.CronName)
        response.JsonExit(r, 0, "ok")
    }
}

func (c *C) GetOneCron(r *ghttp.Request) {
    var (
        data *DeleteOrSearchCron
    )
    if err := r.Parse(&data); err != nil {
        response.JsonExit(r, 1, err.Error())
    }
    if cron := gcron.Search(data.CronName); cron == nil {
        response.JsonExit(r, 1, "任务不存在！")
    } else {
        response.JsonExit(r, 0, "ok", gcron.Search(data.CronName))
    }
}

func (c *C) GetAllCron(r *ghttp.Request) {
    response.JsonExit(r, 0, "ok", gcron.Entries())
}
