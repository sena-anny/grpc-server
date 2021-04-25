module kj-tech.net/greeter_client

go 1.16

require (
	google.golang.org/grpc v1.37.0
	kj-tech.net/proto v0.0.0-00010101000000-000000000000
)

replace kj-tech.net/proto => ../proto
