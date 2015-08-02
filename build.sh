#!/bin/bash

ldflags="-w"

export GOBIN=$(pwd)/bin

rm -rf auto
mkdir auto
for arch in linux-386 linux-amd64 darwin-amd64 ; do
	echo "$arch"
	export GOOS=${arch%-*}
	export GOARCH=${arch#*-}
	rm -rf bin
	go install -ldflags "$ldflags"
	tar zcf "git-autotag-$arch.tar.gz" bin
done
for arch in windows-386 windows-amd64 ; do
	echo "$arch"
	export GOOS=${arch%-*}
	export GOARCH=${arch#*-}
	rm -rf bin
	go install -ldflags "$ldflags"
	zip -qr "mole-$arch.zip" bin
done
rm -rf bin

