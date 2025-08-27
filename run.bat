@echo off
echo Running Markdown to HTML Converter...
echo.

REM Check if Go is installed
go version >nul 2>&1
if errorlevel 1 (
    echo Error: Go is not installed or not in PATH
    echo Please install Go from https://golang.org/dl/
    pause
    exit /b 1
)

REM Download dependencies if needed
echo Checking dependencies...
go mod tidy

REM Run the application with example
echo Running with example file...
go run cmd/main.go examples/sample.md output/sample.html

if errorlevel 1 (
    echo Error: Application failed to run
    pause
    exit /b 1
)

echo.
echo Conversion completed successfully!
echo Check output/sample.html for the result.
echo.
pause
