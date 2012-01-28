include $(GOROOT)/src/Make.inc

TARG=gm
GOFMT=gofmt
SRC=server.go gomud.go character.go

GOFILES=${SRC}

include $(GOROOT)/src/Make.cmd

format:
	${GOFMT} -w ${SRC}
