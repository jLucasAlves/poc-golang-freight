#! /bin/sh

echo installing golang tools... will take a wile

go get -u golang.org/x/lint/golint
go get -u golang.org/x/tools/cmd/goimports
go get -u golang.org/x/tools/cmd/gorename
go get -u golang.org/x/tools/gopls
go get -u github.com/acroca/go-symbols
go get -u github.com/aspenteam/go-tools/cmd/structlayout
go get -u github.com/bketelsen/trace/examples/gogrep
go get -u github.com/cweill/gotests/...
go get -u github.com/davidrjenni/reftools/cmd/fillstruct
go get -u github.com/fatih/gomodifytags
go get -u github.com/FooSoft/goldsmith
go get -u github.com/go-delve/delve/cmd/dlv
go get -u github.com/godoctor/godoctor
go get -u github.com/golangci/tools/cmd/guru
go get -u github.com/haya14busa/goplay/cmd/goplay
go get -u github.com/josharian/impl
go get -u github.com/klauspost/asmfmt/cmd/goreturns
go get -u github.com/mdempsky/gocode
go get -u github.com/mgechev/revive
go get -u github.com/nakabonne/golangci-lint
go get -u github.com/ramya-rao-a/go-outline
go get -u github.com/ramya-rao-a/go-outline
go get -u github.com/rogpeppe/godef
go get -u github.com/uudashr/gopkgs/v2/cmd/gopkgs
go get -u github.com/zmb3/gogetdoc
go get -u honnef.co/go/tools/cmd/keyify
go get -u honnef.co/go/tools/cmd/rdeps
go get -u honnef.co/go/tools/cmd/staticcheck
go get -u honnef.co/go/tools/cmd/structlayout-optimize
go get -u honnef.co/go/tools/cmd/structlayout-pretty
go get -u honnef.co/go/tools/internal/cmd/ast-to-pattern
go get -u honnef.co/go/tools/internal/cmd/irdump
