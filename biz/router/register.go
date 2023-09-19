// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	audit "jobor/biz/router/audit"
	dashboard "jobor/biz/router/dashboard"
	role "jobor/biz/router/role"
	sys_api "jobor/biz/router/sys_api"
	task "jobor/biz/router/task"
	task_log "jobor/biz/router/task_log"
	user "jobor/biz/router/user"
	worker "jobor/biz/router/worker"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	dashboard.Register(r)

	role.Register(r)

	sys_api.Register(r)

	task_log.Register(r)

	worker.Register(r)

	task.Register(r)

	user.Register(r)

	audit.Register(r)
}
