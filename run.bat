@echo off
chcp 65001 > nul

REM Check Docker installation
where docker > nul 2>&1
if %ERRORLEVEL% neq 0 (
    echo [ERROR] Docker not found in PATH
    pause
    exit /b 1
)

REM Generate Swagger docs
echo [1/3] Generating Swagger documentation...
docker run --rm -v "%cd%:/app" -w /app golang:1.24-alpine sh -c "go install github.com/swaggo/swag/cmd/swag@latest && swag init -g ./cmd/app/main.go"

if %ERRORLEVEL% neq 0 (
    echo [ERROR] Swagger generation failed
    pause
    exit /b 1
)

REM Build containers
echo [2/3] Building Docker images...
docker compose build

if %ERRORLEVEL% neq 0 (
    echo [ERROR] Docker build failed
    pause
    exit /b 1
)

REM Start application
echo [3/3] Starting application...
echo ----------------------------------------
echo App will be available at http://localhost:8080
echo Press Ctrl+C to stop
echo ----------------------------------------
docker compose up

pause