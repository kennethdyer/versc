
COMMIT_VALUE=`git rev-parse --short HEAD`
COMMIT=-ldflags "-X main.CommitStamp=$(COMMIT_VALUE)"

BUILD_AT_VALUE=`date +%FT%T%z`
BUILD_AT=-ldflags "-X main.BuildAt=$(BUILD_AT_VALUE)"

BFLAGS=$(COMMIT) $(BUILD_AT)  
GITLAB=github.com/kennethpjdyer
NAME=versc

install: vend
	go install $(BFLAGS) -mod vendor ./cmd/$(NAME)

build: vend
	go build $(BFLAGS) -o dist/$(NAME) -mod vendor ./cmd/$(NAME)

vend:
	go mod vendor

