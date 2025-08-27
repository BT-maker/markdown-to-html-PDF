@echo off
echo Building Markdown to HTML Converter...
echo.

REM Check if Go is installed
go version >nul 2>&1
if errorlevel 1 (
    echo Error: Go is not installed or not in PATH
    echo Please install Go from https://golang.org/dl/
    pause
    exit /b 1
)

echo Go version:
go version
echo.

REM Download dependencies
echo Downloading dependencies...
go mod tidy
if errorlevel 1 (
    echo Error: Failed to download dependencies
    pause
    exit /b 1
)

REM Build the application
echo Building application...
go build -o markdown-to-html.exe cmd/main.go
if errorlevel 1 (
    echo Error: Build failed
    pause
    exit /b 1
)

echo.
echo Build successful! 
echo Executable created: markdown-to-html.exe
echo.
echo Usage examples:
echo   markdown-to-html.exe examples/sample.md output/sample.html
echo   markdown-to-html.exe examples/sample.md --preview
echo   markdown-to-html.exe examples/sample.md --theme dark
echo.
pause
