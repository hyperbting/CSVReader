set GOOS=windows
set GOARCH=amd64
set CGO_ENABLED=0
go build -o ./build/main.exe  main.go configurer.go main_csvreader.go db_insert.go dataparser.go