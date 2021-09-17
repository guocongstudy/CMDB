package main

import "CMDB/api/cmd"



func main(){
cmd.Execute()
}

// go run main.go -f ../api/etc/demo.toml start

//120.0.0.1:8050/host
//快捷启动    ./CMDB-api start -f ../api/etc/demo.toml
//开始前端