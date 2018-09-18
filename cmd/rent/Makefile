# Values we want to pass for Version and Build
GITHASH:=`git describe --long --always --dirty`.
BUILDTIME=`date +%y%m%d.%T`
VAL=`git rev-list --all|wc -l|awk '{print $1}'`
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
DIFF=$(shell echo $(VAL))
VERSION=$(shell echo ${BRANCH}.${GITHASH}${DIFF})

#setup the -ldflags option for go build here
LDFLAGS=-ldflags "-X wechatServer/src/constant.BuildTime=${BUILDTIME} -X wechatServer/src/constant.Version=${VERSION}"

#build targets
build:
	go build ${LDFLAGS} -o ../../rentServer

ver:
	@echo ${VERSION}

# Cleans our project: deletes binaries
clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY: clean install
