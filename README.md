export GO111MODULE=on
export GOPATH=$HOME/test/golang/foozle
go install app && ./bin/app 


go get -u github.com/golang/gddo/httputil/header