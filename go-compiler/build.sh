LOCK_DIR='/Users/zhang/Dropbox/UMich2020sp/pitometer/go-compiler/toylock-go'
export GOPATH=$LOCK_DIR

go build "$LOCK_DIR/src/Main.go"

export GOPATH='/Users/zhang/go'
export PATH=$PATH:$(go env GOPATH)/bin
export GOPATH=$(go env GOPATH)