

test:
	go test panic_safenet_test.go | grep succeed > /dev/null


.PHONY: test