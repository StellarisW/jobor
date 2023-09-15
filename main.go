// Code generated by hertz generator.

package main

import (
	"fmt"
	Cmd "jobor/cmd"
	"jobor/cmd/srv_http"
	"os"
)

//	@title			Jobor 定时任务 API
//	@version		3.0
//	@description	支持host,server,network等
//	@description	Authorization Bearer token
//	@description	header:  x-enc-data = yes

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
//	@schema						Bearer

//	@contact.name	jobor
//	@contact.url	https://github.com/v-mars/jobor
//
// //@host        localhost:5002
//
//	// @BasePath /
//
// //@schemes     http
func main() {
	//dal.Init()

	srv_http.RootCmd.AddCommand(Cmd.Version())
	if err := srv_http.RootCmd.Execute(); err != nil {
		_ = fmt.Errorf("rootCmd.Execute failed %s", err.Error())
		os.Exit(-1)
	}
}
