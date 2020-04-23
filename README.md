# mono-repo-example

This example repo contains multiple non-nested go modules. Each module is independently versioned using prefixed tags.

For a step-by-step walkthrough of how to create your own multi-module repository see https://github.com/go-modules-by-example.

## Try it out
Create a throwaway repo and try it out.

```console
$ cd $(mktemp -d)
$ export GOBIN=$PWD/.bin
$ export PATH=$GOBIN:$PATH
$ go mod init example.com/blah
go: creating new go.mod: module example.com/blah
$ go get github.com/aka-bo/mono-repo-example/go/bar@v0.1.0
go: downloading github.com/aka-bo/mono-repo-example/go/bar v0.1.0
go: downloading github.com/aka-bo/mono-repo-example/go/foo v0.1.2
$ bar
Foo Bar
```
