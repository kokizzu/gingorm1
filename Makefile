

test:
	cd business; go test -update .

testv:
	go install github.com/gotestyourself/gotestsum@latest
	cd business; gotestsum --format=testname .

setup:
	go install github.com/cosmtrek/air@latest
