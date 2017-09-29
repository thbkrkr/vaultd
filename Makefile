build:
	doo bgo
	doo b

#

dev:
	VAULT_KEY=Ha8QMP7fw4oLmYlZXFPfOMsjcMJmmvcL \
	VAULT_NONCE=5A9WFbIQmnnp \
		golive -data-dir ./test

test-run:
	docker run \
		-e VAULT_KEY=Ha8QMP7fw4oLmYlZXFPfOMsjcMJmmvcL \
		-e VAULT_NONCE=5A9WFbIQmnnp \
		-v $$(pwd)/test:/test \
		-p 80:4242 \
		krkr/vaultd vaultd -data-dir /test