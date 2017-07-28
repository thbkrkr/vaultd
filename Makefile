build:
	doo bgo
	doo b

dev:
	golive -data-dir ./test

test-run:
	docker run \
		-v $$(pwd)/test:/test \
		-p 80:4242 \
		krkr/vaultd -data-dir /test