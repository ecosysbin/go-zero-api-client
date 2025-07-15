# 根据.api文件生成http go client
go run .\generate-client.go --api .\user.api --output client\user.go --package client

# 根据.api文件生成http go server (goctl安装参照: https://go-zero.dev/docs/tasks/installation/goctl)
goctl api go -api ./user.api -dir .

注意：.api文件格式参照: https://go-zero.dev/docs/tasks/dsl/api