go get github.com/akavel/rsrc
rsrc -manifest test.manifest -o rsrc.syso

go build -ldflags="-H windowsgui"