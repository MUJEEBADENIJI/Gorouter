@echo off
title GoRoute
cd /d "C:\Users\user\3D Objects\CodePagesHere\Code.Page.18Goroute\Gorouter\internet-connection-sharer"

:: 🚀 Check if Go is installed
go version >nul 2>&1
if %errorlevel% neq 0 (
    echo ❌ Go is not installed or not in PATH.
    pause
    exit /b
)

:: 🛠️ Clean old build
echo Cleaning old build...
del /f /q goroute.exe 2>nul

:: 🚀 Build GoRoute
echo Building GoRoute...
go build -o goroute.exe main.go
if %errorlevel% neq 0 (
    echo ❌ Build failed!
    pause
    exit /b
)

:: ✅ Run GoRoute
echo Running GoRoute...
goroute.exe
pause
