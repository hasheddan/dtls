module github.com/pion/dtls/v2

require (
	github.com/google/go-cmp v0.5.9
	github.com/pion/logging v0.2.2
	github.com/pion/transport/v2 v2.2.1
	golang.org/x/crypto v0.11.0
	golang.org/x/net v0.13.0
)

replace github.com/pion/transport/v2 => github.com/hasheddan/transport/v2 v2.0.0-20230705220158-9d8030689d03

go 1.13
