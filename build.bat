@echo off
cls
setlocal enabledelayedexpansion

:: Define the targets for all relevant platforms
set targets=linux/amd64 linux/386 linux/arm linux/arm64 darwin/amd64 darwin/arm64 windows/amd64 windows/386 windows/arm windows/arm64
set sources=main.go

:: Ensure the export folder exists
if not exist "export" mkdir "export"

:: Loop through each target platform
for %%t in (%targets%) do (
  for /f "tokens=1,2 delims=/" %%a in ("%%t") do (
    set GOOS=%%a
    set GOARCH=%%b

    :: Create platform directories inside export if they don't exist
    if not exist "export/%%a-%%b" mkdir "export/%%a-%%b"

    :: Copy the apps.json file
    copy "apps.json" "export/%%a-%%b\apps.json"


    :: Set output extension for Windows
    set output_ext=
    if "%%a"=="windows" (
      set output_ext=.exe
    )
    
    :: Build main.go
    echo Building main.go for %%a/%%b...
    set GOOS=%%a
    set GOARCH=%%b
    go build -o "export/%%a-%%b/main!output_ext!" main.go
    
    :: Formatting line to separate output for each platform
    echo --------------------------------------------------
  )
)

echo Build process completed.
