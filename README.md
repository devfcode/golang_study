# golang_study
 之前golang项目用go.mod管理包  </br>

## 在Windows开发环境下将goland项目编译成Linux下可执行的二进制文件
1、windows下执行     </br>
set GOARCH=amd64   </br>
set GOOS=linux        </br>
go build main.go      </br>
编译成二进制文件 main 后，将main传送到linux下   </br>
2、在linux下 首先给限，然后执行   </br>
chmod 777 main   </br>
./main   </br>

## 在Mac开发环境下将goland项目编译成Linux下可执行的二进制文件
1、Mac下编译   </br>
GOOS=linux GOARCH=amd64 go build main.go   </br>
或  </br>
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go  </br>
2、在linux下 首先给限，然后执行  </br>
chmod 777 main   </br>
./main  </br>
