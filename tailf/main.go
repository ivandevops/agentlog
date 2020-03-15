package tailf

//https://blog.csdn.net/wuhao2048/article/details/102681584?depth_1-utm_source=distribute.pc_relevant.none-task&utm_source=distribute.pc_relevant.none-task
import (
	"fmt"
	"github.com/hpcloud/tail"
	"time"
)
func main() {
	filename := "log/agent.log"						//相对路径无论在哪里写都是相对于agentlog的
	tails, err := tail.TailFile(filename, tail.Config{
		ReOpen:    true,							//写满一定程度，挪走文件，打开新的文件写
		Follow:    true,							//跟踪
		//Location:  &tail.SeekInfo{Offset: 0, Whence: 2},  //进程挂掉，重启，记录当前读的位置
		MustExist: false,							//文件不存在，也监控，直到文件存在然后收集
		Poll:      true,							//不断查询文件
	})
	if err != nil {
		fmt.Println("tail file err:", err)
		return
	}

	var msg *tail.Line			//一行日志
	var ok bool
	for true {
		msg, ok = <-tails.Lines //从管道里读一行
		if !ok {				//管道关闭
			fmt.Printf("tail file close reopen, filename:%s\n", tails.Filename)
			time.Sleep(100 * time.Millisecond)
			continue
		}
		fmt.Println("msg:", msg.Text)
	}
}
