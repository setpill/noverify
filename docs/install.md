# How to install NoVerify

First you will need the Go toolchain (https://golang.org/).

Once Go installed, do the following command:

```sh
$ go get -u github.com/setpill/noverify
```

This command installs `noverify` into `$GOPATH/bin/noverify` (which expands into `$HOME/go/bin/noverify` by default).

Alternatively, you can build `noverify` with version info:

```sh
mkdir -p $GOPATH/github.com/setpill
git clone https://github.com/setpill/noverify.git $GOPATH/github.com/VKCOM

cd $GOPATH/src/github.com/setpill/noverify
make install
```

## Next steps

- [Using NoVerify as linter / static analyser](linter-usage.md)
- [Using NoVerify as language server for Sublime Text](sublime-plugin.md)
- [Using NoVerify as language server for VSCode](vscode-plugin.md)
