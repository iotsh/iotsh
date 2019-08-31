# Copyright 2019 The IoTSH Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

FROM golang:latest

WORKDIR /go/src/github.com/iotsh/iotsh

COPY . .

ENV GOPATH /go

ARG CMD_DIR=./cmd

ARG ROOT=github.com/iotsh/iotsh

ARG VERSION=

ARG COMMIT=

ARG TARGET=iotsh

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64                      \
	go build -i -v -o /tmp/${TARGET}                  \
	-ldflags "-s -w -X ${ROOT}/pkg/version.Version=${VERSION}  \
	-X ${ROOT}/pkg/version.Commit=${COMMIT}                    \
	-X ${ROOT}/pkg/version.Package=${ROOT}"                    \
	${CMD_DIR}/${TARGET};

FROM debian:jessie

ARG TARGET=iotsh

COPY --from=0 /tmp/${TARGET} /${TARGET}

RUN ln -s /${TARGET} /entrypoint

ENTRYPOINT ["/entrypoint"]
