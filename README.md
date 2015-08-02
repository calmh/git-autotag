# git autotag

Creates an annonated tag for the next logical version. Defaults to patch version (as per semver), but can be told to do minor and major releases. Signs tags if the config `autotag.sign` is `true`.

## Installation

Get and compile it using Go:

```go get github.com/calmh/git-autotag```

Or grab one of [the binaries](https://github.com/calmh/git-autotag/releases).

## Usage

```
$ git describe
v0.11.18-3-gb1dd704
$ git autotag
v0.11.19
$ git describe
v0.11.19
$ git autotag -l minor
v0.12.0
$ git config autotag.sign true
$ git config user.signingkey release@syncthing.net
$ git autotag -l patch
v0.12.1

You need a passphrase to unlock the secret key for
user: "Syncthing Release Management <release@syncthing.net>"
2048-bit RSA key, ID D26E6ED000654A3E, created 2014-12-29

$ 
```

## Prerequisites

A tag describing the previous version, with three parts á la semver. Can have a prefix of `v`, `V`, `version` etc., or none at all.

## License

MIT
