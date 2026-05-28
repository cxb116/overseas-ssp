@echo off
chcp 65001 >nul

echo ETCD 简单示例
echo ================
echo.

cd /d %~dp0

go run main.go

pause
