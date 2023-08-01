module github.com/pion/dtls/v2

replace github.com/pion/transport/v2 => github.com/hasheddan/transport/v2 v2.0.0-20230801191430-bb77a73aa34f

require (
	github.com/pion/logging v0.2.2
	github.com/pion/transport/v2 v2.2.2-0.20230711104634-a789100cc553
	golang.org/x/crypto v0.11.0
	golang.org/x/net v0.13.0
)

go 1.13
