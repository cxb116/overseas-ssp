@echo off
chcp 65001 >nul

echo ========================================
echo   ETCD 监听示例
echo ========================================
echo.

cd /d %~dp0

echo 检查 etcd 连接...
echo.

echo 测试连接: 117.72.67.186:2379
timeout /t 2 /nobreak >nul

echo.
echo 运行示例程序...
echo.

go run main.go

pause
