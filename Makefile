all: main
	./$<

main::
	go build kbd-float.go
