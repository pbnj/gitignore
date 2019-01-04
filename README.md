# gitignore

> Simple CLI for creating .gitignore files

## Install

```
$ go get -u -v github.com/petermbenjamin/go-gitignore
```

## Usage

```
$ gitignore --help
Usage of gitignore:
  -d    Debug
  -l string
        Language
  -w    Write to .gitignore file
```

### Example

Simple: `gitignore -l <lang>` will print to STDOUT by default

```sh
$ gitignore -l go
INFO Found: Go.gitignore
# Binaries for programs and plugins
*.exe
*.dll
*.so
*.dylib

# Test binary, build with `go test -c`
*.test

# Output of the go coverage tool, specifically when used with LiteIDE
*.out

# Project-local glide cache, RE: https://github.com/Masterminds/glide/issues/736
.glide/
```

- To write to `.gitignore` file:
    ```sh
    $ gitignore -w -l go
    Found Go.gitignore
    Created: /home/ubuntu/go/src/github.com/petermbenjamin/go-gitignore/.gitignore
    ```

## License

MIT &copy; [Peter Benjamin](https://petermbenjamin.github.io/)
