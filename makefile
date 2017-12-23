

test: vendor
	go test panic_safenet_test.go | grep succeed > /dev/null

vendor:
	dep ensure

.PHONY: test