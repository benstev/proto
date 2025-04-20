.PHONY: gen
gen:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative hagw.proto
	# protoc --go_out=./pb --go-grpc_out=./pb hagw.proto