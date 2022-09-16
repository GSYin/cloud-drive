# cloud-drive
This is a mini cloud disk, which can be used to store files and share files with others.

### ✨ Features
- **User manager(register user, user info etc.)**
- **Store files and share files with others**
- **Files associated storage and Folder level management**

### 📦 How to use
```shell
# install go-zero
go get -u github.com/zeromicro/go-zero
go install github.com/zeromicro/go-zero/tools/goctl@latest
goctl api new core
go mod tidy
cd core
go run core.go -f etc/core-api.yaml 
goctl api go -api core.api -dir . -style go_zero 
curl -i -X GET http://localhost:8888/user/login

```


