# SPDX-FileCopyrightText: 2023 The Pion community <https://pion.ly>
# SPDX-License-Identifier: MIT

FROM docker.io/library/golang:1.20-bullseye

COPY . /go/src/github.com/pion/dtls
WORKDIR /go/src/github.com/pion/dtls/e2e

# TODO: revert
CMD ["go", "test", "-v", ".", "-run", "^TestPionE2ESimple\/TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256$"]
