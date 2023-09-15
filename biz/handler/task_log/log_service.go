// Code generated by hertz generator.

package task_log

import (
	"context"
	"fmt"
	"jobor/biz/model"
	"jobor/biz/pack/dispatcher"
	"jobor/biz/response"
	"jobor/kitex_gen/task_log"
	"jobor/pkg/convert"

	"github.com/cloudwego/hertz/pkg/app"
)

// GetLog .
//
//	@Summary		jobor log get summary
//	@Description	jobor log get
//	@Tags			jobor log
//	@Param			name		query	string	false	"name"
//	@Param			page		query	int		false	"page"
//	@Param			pageSize	query	int		false	"pageSize"
//
// @router /api/v1/jobor/log [GET]
func GetLog(ctx context.Context, c *app.RequestContext) {
	var err error
	userinfo, err := model.GetUserSession(c, false)
	if err != nil {
		response.SendBaseResp(ctx, c, err)
		return
	}
	var req task_log.LogQuery
	err = c.BindAndValidate(&req)
	if err = c.BindAndValidate(&req); err != nil {
		response.ParamFailed(ctx, c, err)
		return
	}

	var objs model.Logs

	resp := response.InitPageData(ctx, c, &objs, false)

	if _, err = objs.List(&req, &userinfo, &resp); err != nil {
		response.SendBaseResp(ctx, c, err)
		return
	}
	response.SendDataResp(ctx, c, response.Succeed, resp)
}

// AbortTask .
//
//	@Summary		jobor task abort summary
//	@Description	jobor task abort
//	@Tags			jobor log
//
// @router /api/v1/jobor/log/:id/abort [POST]
func AbortTask(ctx context.Context, c *app.RequestContext) {
	var err error
	var req task_log.PostLogReq
	err = c.BindAndValidate(&req)
	if err = c.BindAndValidate(&req); err != nil {
		response.ParamFailed(ctx, c, err)
		return
	}
	_id := c.Params.ByName("id")
	_, ok := dispatcher.CacheTask.TaskLogs[convert.ToInt(_id)]
	if !ok {
		response.SendBaseResp(ctx, c, fmt.Errorf("任务[%s]已经完成或不存在", _id))
		return
	}
	dispatcher.CacheTask.TaskLogs[convert.ToInt(_id)].TaskCancel()

	response.SendDataResp(ctx, c, response.Succeed, "")
}

// GetLogById .
//
//	@Summary		jobor task log get summary
//	@Description	jobor task log get
//	@Tags			jobor log
//	@Param			id		path	int				true	"int valid"
//
// @router /api/v1/jobor/log/{id} [GET]
func GetLogById(ctx context.Context, c *app.RequestContext) {
	var err error
	_id := c.Params.ByName("id")

	Resp, err := model.GetLogInfoById(_id, false)
	if err != nil {
		response.SendBaseResp(ctx, c, err)
		return
	}

	response.SendDataResp(ctx, c, response.Succeed, Resp)
}
