include $(GOROOT)/src/Make.inc

TARG=monitor
GOFILES=\
				main.go\
				key.go\
				store.go\

include $(GOROOT)/src/Make.cmd


