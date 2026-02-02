rem #############################################################
rem # Build script for Golang HTTP Function 
rem # for Linux target on Windows
rem #############################################################

rem set visibility of environment changes to local only
SETLOCAL

rem Set the following parameters to the corresponding value of Linux
set GOARCH=amd64
set GOOS=linux
set CGO_ENABLED=0
set GO111MODULE=on

rem set handler name
set HANDLER_NAME=go-http-demo

rem create target_win folder if not exists
if not exist "target_win" md "target_win"

rem build the Go HTTP function in target_win folder
go build -o target_win/%HANDLER_NAME% src/main.go

rem create bootstrap file in target_win folder
echo /opt/function/code/%HANDLER_NAME% > .\target_win\bootstrap

rem package the target_win folder to a zip file
tar.exe -c -a -f go-http-demo.zip -C target_win *
