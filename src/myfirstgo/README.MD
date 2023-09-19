
프로젝트의 소스코드는 GOPATH\src\{provider}\{account}\{repository} 형태로 저장

GOPATH\src\{코드레포지토리}\{계정}\{프로젝트}
GOPATH\src\github.com\daintree-henry\learngo-master

외부 코드 다운
=> go get go.uber.org/zap

go 모듈 초기화, The go.mod file only appears in the root of the module.
=> $ go mod init example.com/hello
   go mod init github.com/api_server
=> $ go test

- to check go environments
```
go env | grep GOPATH
GOPATH="/Users/simonchoi/workspace_vscode/golang"
GOPATH는 GO 프로젝트의 import 위치
모듈을 다운로드 하는 장소 결정

모듈 폴더에서 수행
go mod init github.com/myfirstgo
go mod edit -replace myfirstgo/modules/mylibs=../modules/mylibs
go mod edit -replace myfirstgo/modules/httpserver=../modules/httpserver
```

- go build
```
go build -o test_main main.go

# go build with windows
$ GOOS=windows GOARCH=amd64 go build -o dist/main-amd64.exe main.go

# 64-bit, macos
$ GOOS=darwin GOARCH=amd64 go build -o dist/main-amd64-darwin main.go

# 64-bit, linux
$ GOOS=linux GOARCH=amd64 go build -o dist/main-amd64-linux main.go

```

- go build setting in vscode


- go build 하면서 os check
```
$ go env
... in mac
GOOS="darwin"
GOARCH="amd64"
```

- go gitignore create
```
gi go
```

- go.work ?
```
go work init ./modules/httpserver ./modules/mylibs
use ./modules/httpserver
use ./modules/mylibs

go work use ./modules/myhttpserver
```

- log 를 파일로 추가하기
```

```

- shell 명령어를 수행하기
- webhook 사용하기

- task.json 에서 build 명령어 추가 완료.. mac만 했다 !!! 끝... 

go env -w GO111MODULE=auto.  ==> 이걸로 변경

src/myfirstgo/main.go:4:2: package myfirstgo/modules/myhttpserver is not in GOROOT 
(/usr/local/Cellar/go/1.20.7/libexec/src/myfirstgo/modules/myhttpserver)
src/myfirstgo/main.go:5:2: package myfirstgo/modules/mylibs is not in GOROOT (/usr/local/Cellar/go/1.20.7/libexec/src/myfirstgo/modules/mylibs)
