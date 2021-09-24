package main

import "CMDB/api/cmd"



func main(){
cmd.Execute()
}

/* go run main.go -f ../api/etc/demo.toml start
1.127.0.0.1:8050/host
2.快捷启动：
./CMDB-api start -f ../api/etc/demo.toml
3.开始前端 3:04:10
4.前端启动指令 进入ui后 输入 npm run serve
5.28讲看完
*/