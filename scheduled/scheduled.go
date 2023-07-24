package scheduled

import (
	"log"
	"strings"

	"github.com/robfig/cron"
)

type CmdName string
type CmdFun string

type CmdConfig map[CmdName]CmdFun

type Commands interface {
	GetConfig() CmdConfig
	MethodMap(CmdName) func()
}

// ScheduledCommands 定时任务
func ScheduledTask(commands Commands) (c *cron.Cron, err error) {
	c = cron.New()
	for ClassMethod, timing := range commands.GetConfig() {
		ClassMethodSlice := strings.Split(string(ClassMethod), "/")
		if len(ClassMethodSlice) != 1 {
			continue
		}
		function := commands.MethodMap(ClassMethod)
		if function != nil {
			if err := c.AddFunc(string(timing), function); err != nil {
				log.Fatalln(err)
			}
		}
	}
	c.Start()
	return
}
