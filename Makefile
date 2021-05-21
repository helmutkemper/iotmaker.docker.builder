# make update-pkg-cache USER=helmut.kemper@gmail.com PACKAGE=github.com/helmutkemper/iotmaker.docker.builder VERSION=v0.4.0-rc.014
update-pkg-cache:
    GOPROXY=https://proxy.golang.org GO111MODULE=on \
    go get github.com/$(USER)/$(PACKAGE)@v$(VERSION)