```shell
docker pull quay.io/goswagger/swagger
alias swagger='docker run --rm -it  --user $(id -u):$(id -g) -e GOPATH=$(go env GOPATH):/go -v $HOME:$HOME -w $(pwd) quay.io/goswagger/swagger'
swagger version
```

docker build -t test -f Dockerfile .
docker run --name test -d test:latest tail -f /bin/bash
docker exec -it test sh

export GO111MODULE=on
go get github.com/golangci/golangci-lint@v1.23.6
cd /go/pkg/mod/github.com/golangci/golangci-lint@v1.23.6
chmod +x ./install.sh
./install.sh
cp ./bin/golangci-lint /go/bin/golangci-lint

go get gotest.tools/gotestsum@v0.6.0
cd /go/pkg/mod/gotest.tools/gotestsum@v0.6.0

ARG     GOLANG_VERSION
FROM    golang:${GOLANG_VERSION:-1.14-alpine} as golang
RUN     apk add -U curl git bash
ENV     CGO_ENABLED=0 \
PS1="# " \
GO111MODULE=on
ARG     UID=1000
RUN     adduser --uid=${UID} --disabled-password devuser
USER    ${UID}:${UID}


FROM    golang as tools
RUN     go get github.com/dnephin/filewatcher@v0.3.2
RUN     wget -O- -q https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s && \
mv bin/golangci-lint /go/bin


FROM    golang as dev
COPY    --from=tools /go/bin/filewatcher /usr/bin/filewatcher
COPY    --from=tools /go/bin/golangci-lint /usr/bin/golangci-lint


FROM    dev as dev-with-source
COPY    . .

