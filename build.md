## Build Notes

Due to a 3 year old issue with GopherJS not supporting Go modules, building this is serious pain. Serious pain

Install Go 1.12, note it's safe to do this alongside your system Go install, it will not replace it

```
go get golang.org/dl/go1.12.16
go1.12.16 download
```

Set these vars as required, one is the main/normal project folder which holds this repo, and one is a copy of the same, but inside GOPATH

```
PROJ_DIR=/where/you/put/projects/hcl2-parser
GOPATH_DIR=$(go1.12.16 env GOPATH)/src/hcl2-parser
```

YOU MUST RUN THIS STEP OUTSIDE GOPATH!

```
cd $PROJ_DIR
go1.12.16 mod vendor
```

Sync project folder to src folder under GOPATH, note the trailing slash which is important

```
rsync -avu --delete "${PROJ_DIR}/" "$GOPATH_DIR"
```

Run GopherJS INSIDE GOPATH

```
cd $GOPATH_DIR
export GOPHERJS_GOROOT="$(go1.12.16 env GOROOT)"
gopherjs build parser.go -o dist/index.js -m
```

Now copy results back

```
cp dist/* PROJ_DIR/dist
```
