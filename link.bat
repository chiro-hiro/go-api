@echo off
SET CURDIR=%cd%
SET GOPATH=C:\go-work
SET MIRROR=%GOPATH%\src\github.com\chiro-hiro\go-api
MKDIR %GOPATH%\src
MKDIR %GOPATH%\src\github.com
MKDIR %GOPATH%\src\github.com\chiro-hiro
MKLINK /J %MIRROR% %cd%