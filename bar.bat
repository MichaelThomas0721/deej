@echo off
call ./pkg/deej/scripts/windows/build-dev.bat
if %errorlevel% neq 0 exit /b %errorlevel%
call deej-dev.exe