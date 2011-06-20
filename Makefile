include $(GOROOT)/src/Make.inc

TARG=gomud
GOFILES=\
	server.go\
        gomud.go\

include $(GOROOT)/src/Make.cmd
