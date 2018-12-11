# Makefile to build the command lines and tests in Seele project.
# This Makefile doesn't consider Windows Environment. If you use it in Windows, please be careful.

SHELL := /bin/bash 

NODE_APP_NAME = d4d-node
BASEDIR = $(shell pwd)

# build with verison infos
versionDir = myproject/version

gitBranch = $(shell git symbolic-ref --short -q HEAD)

ifeq ($(gitBranch),)
gitTag = $(shell git describe --always --tags --abbrev=0)
endif

buildDate = $(shell date "+%FT%T%z")
gitCommit = $(shell git rev-parse HEAD)
gitTreeState = $(shell if git status|grep -q 'clean';then echo clean; else echo dirty; fi)
 
ldflagsOrigin="-X ${versionDir}.gitBranch=${gitBranch} -X ${versionDir}.gitTag=${gitTag} -X ${versionDir}.buildDate=${buildDate} \
 -X ${versionDir}.gitCommit=${gitCommit} \
 -X ${versionDir}.gitTreeState=${gitTreeState}"

# -s -w 
ldflagsRelase="-s -w -X ${versionDir}.gitBranch=${gitBranch} -X ${versionDir}.gitTag=${gitTag} -X ${versionDir}.buildDate=${buildDate} \
 -X ${versionDir}.gitCommit=${gitCommit} \
 -X ${versionDir}.gitTreeState=${gitTreeState}"

all: node statis txhistory
node:
	go build -v -ldflags ${ldflagsRelase} -o ./build/bin/${NODE_APP_NAME} ./cmd/node
	@echo "Done node building debug"

statis:
	go build -v -ldflags ${ldflagsRelase} -o ./build/bin/statis ./cmd/statis
	@echo "Done statis building"

txhistory:
	go build -ldflags "-w -s" -o ./build/bin/txhistory ./txhistory/start
	@echo "Done txhistory building"

clean:
	rm -rf ${BASEDIR}/build/bin/*
.PHONY: node     statis txhistory
