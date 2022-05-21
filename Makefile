

test:
	cd business; go test -update .

testv:
	cd business; gotestsum --format=testname .

