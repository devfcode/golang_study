# golang_study
 之前golang项目用go.mod管理包

在Windows开发环境下将goland项目编译成Linux下可执行的二进制文件
1、windows下执行
set GOARCH=amd64
set GOOS=linux
go build main.go
编译成二进制文件 main 后，将main传送到linux下
2、在linux下 首先给限，然后执行
chmod 777 main
./main

在Mac开发环境下将goland项目编译成Linux下可执行的二进制文件
1、Mac下编译
GOOS=linux GOARCH=amd64 go build main.go
或
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
2、在linux下 首先给限，然后执行
chmod 777 main
./main