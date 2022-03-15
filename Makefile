HOSTNAME=sebank.se
NAMESPACE=seb
NAME=dsd
BINARY=terraform-provider-${NAME}
VERSION=0.0.1
OS_ARCH=linux_amd64

default: install

build:
	go build -o ${BINARY}

debugBuild:
	go build -gcflags="all=-N -l" -o ${BINARY}-debug

debug: debugBuild
	dlv exec --listen=:57324 --headless=true --api-version=2 --accept-multiclient ./${BINARY}-debug -- --debug

install: build
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	cp ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}