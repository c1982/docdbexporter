export GOARCH=arm64
export GOOS=linux
export CGO_ENABLED=0

go build -o ../docdb_exporter_linux_arm64 -tags netgo -ldflags '-w -extldflags "-static"' ../.