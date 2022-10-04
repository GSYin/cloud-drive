# cloud-drive
This is a mini cloud disk, which can be used to store files and share files with others.

### ✨ Features
- **User manager(register user, user info etc.)**
- **Store files and share files with others**
- **Files associated storage and Folder level management**

### 📦 How to use
```shell
# install go-zero and goctl
go get -u github.com/zeromicro/go-zero
go install github.com/zeromicro/go-zero/tools/goctl@latest

# create project core
goctl api new core
go mod tidy

# generate handler and logic
goctl api go -api core.api -dir . -style go_zero 

# run project
cd core
go run core.go -f etc/core-api.yaml 


# test api
curl -i -X GET http://localhost:8888/user/login

```

### 🌈 Learning resources

- **[go-zero docs](https://go-zero.dev/cn/docs/quick-start)**
- **[video course](https://www.bilibili.com/video/BV1cr4y1s7H4?spm=a2c6h.12873639.article-detail.6.a696fc18NZ02N4)**
- **[api description](https://developer.aliyun.com/article/935464)**
- **[origin GitHub repositorie](https://github.com/GetcharZp/cloud-disk)**