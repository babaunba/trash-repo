generate:
	cp -r ../proto/ .
	docker run --rm -v $(PWD):/code -w /code soslanco/go-protoc:latest \
		--go_out . --go_opt paths=source_relative \
		--go-grpc_out . --go-grpc_opt paths=source_relative \
		--grpc-gateway_out . \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt paths=source_relative \
		--openapiv2_out . \
		--openapiv2_opt logtostderr=true \
		$(shell find ./proto/ -name '*.proto')
	rm $(shell find ./proto/ -name '*.proto')
	mv proto/ gen/
