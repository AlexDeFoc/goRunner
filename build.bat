@echo off


del /Q /S export\ > NUL 2>&1
del /Q /S installer\ > NUL 2>&1
mkdir export\ > NUL 2>&1


GOTO :MAIN


:win_installer_start
    echo --== WINDOWS INSTALLERS
EXIT /B 0


:win_installer_i386
    set GOOS=windows
    set GOARCH=386
    echo -= win_i386
    go build -o r.exe main.go
    iscc builder\win_i386.iss > NUL 2>&1
    del /Q r.exe > NUL 2>&1

EXIT /B 0

:win_installer_amd64
    set GOOS=windows
    set GOARCH=amd64
    echo -= win_amd64
    go build -o r.exe main.go
    iscc builder\win_amd64.iss > NUL 2>&1
    del /Q r.exe > NUL 2>&1

EXIT /B 0

:win_installer_arm
    set GOOS=windows
    set GOARCH=arm
    echo -= win_arm
    go build -o r.exe main.go
    iscc builder\win_arm.iss > NUL 2>&1
    del /Q r.exe > NUL 2>&1

EXIT /B 0

:win_installer_arm64
    set GOOS=windows
    set GOARCH=arm64
    echo -= win_arm64
    go build -o r.exe main.go
    iscc builder\win_arm64.iss > NUL 2>&1
    del /Q r.exe > NUL 2>&1

EXIT /B 0











:: WINDOWS BUILDs
:windows_start
    echo --== WINDOWS
EXIT /B 0


:windows_i386
    set GOOS=windows
    set GOARCH=386
    echo -= windows_i386
    go build -o r.exe main.go
    tar -cf windows_i386.tar -C .\ alias.json app.json config.json r.exe > NUL 2>&1
    7z a -t7z -mx=9 windows_i386.tar.xz windows_i386.tar > NUL 2>&1
    move windows_i386.tar.xz export\ > NUL 2>&1
    del /Q *.tar > NUL 2>&1
    del /Q r.exe > NUL 2>&1

EXIT /B 0

:windows_amd64
    set GOOS=windows
    set GOARCH=amd64
    echo -= windows_amd64
    go build -o r.exe main.go
    tar -cf windows_amd64.tar -C .\ alias.json app.json config.json r.exe > NUL 2>&1
    7z a -t7z -mx=9 windows_amd64.tar.xz windows_amd64.tar > NUL 2>&1
    move windows_amd64.tar.xz export\ > NUL 2>&1
    del /Q *.tar > NUL 2>&1
    del /Q r.exe > NUL 2>&1

EXIT /B 0

:windows_arm
    set GOOS=windows
    set GOARCH=arm
    echo -= windows_arm
    go build -o r.exe main.go
    tar -cf windows_arm.tar -C .\ alias.json app.json config.json r.exe > NUL 2>&1
    7z a -t7z -mx=9 windows_arm.tar.xz windows_arm.tar > NUL 2>&1
    move windows_arm.tar.xz export\ > NUL 2>&1
    del /Q *.tar > NUL 2>&1
    del /Q r.exe > NUL 2>&1

EXIT /B 0

:windows_arm64
    set GOOS=windows
    set GOARCH=arm64
    echo -= windows_arm64
    go build -o r.exe main.go
    tar -cf windows_arm64.tar -C .\ alias.json app.json config.json r.exe > NUL 2>&1
    7z a -t7z -mx=9 windows_arm64.tar.xz windows_arm64.tar > NUL 2>&1
    move windows_arm64.tar.xz export\ > NUL 2>&1
    del /Q *.tar > NUL 2>&1
    del /Q r.exe > NUL 2>&1

EXIT /B 0











:: LINUX BUILDs
:linux_start
    echo --== LINUX
EXIT /B 0


:linux_i386
    set GOOS=linux
    set GOARCH=386
    echo -= linux_i386
    go build -o r main.go
    tar -cf linux_i386.tar -C .\ alias.json app.json config.json r > NUL 2>&1
    7z a -t7z -mx=9 linux_i386.tar.xz linux_i386.tar > NUL 2>&1
    move linux_i386.tar.xz export\ > NUL 2>&1
    del /Q *.tar > NUL 2>&1
    del /Q r > NUL 2>&1

EXIT /B 0

:linux_amd64
    set GOOS=linux
    set GOARCH=amd64
    echo -= linux_amd64
    go build -o r main.go
    tar -cf linux_amd64.tar -C .\ alias.json app.json config.json r > NUL 2>&1
    7z a -t7z -mx=9 linux_amd64.tar.xz linux_amd64.tar > NUL 2>&1
    move linux_amd64.tar.xz export\ > NUL 2>&1
    del /Q *.tar > NUL 2>&1
    del /Q r > NUL 2>&1

EXIT /B 0

:linux_arm
    set GOOS=linux
    set GOARCH=arm
    echo -= linux_arm
    go build -o r main.go
    tar -cf linux_arm.tar -C .\ alias.json app.json config.json r > NUL 2>&1
    7z a -t7z -mx=9 linux_arm.tar.xz linux_arm.tar > NUL 2>&1
    move linux_arm.tar.xz export\ > NUL 2>&1
    del /Q *.tar > NUL 2>&1
    del /Q r > NUL 2>&1

EXIT /B 0

:linux_arm64
    set GOOS=linux
    set GOARCH=arm64
    echo -= linux_arm64
    go build -o r main.go
    tar -cf linux_arm64.tar -C .\ alias.json app.json config.json r > NUL 2>&1
    7z a -t7z -mx=9 linux_arm64.tar.xz linux_arm64.tar > NUL 2>&1
    move linux_arm64.tar.xz export\ > NUL 2>&1
    del /Q *.tar > NUL 2>&1
    del /Q r > NUL 2>&1

EXIT /B 0











:: MAC OS BUILDs
:darwin_start
    echo --== MAC OS
EXIT /B 0


:darwin_amd64
    set GOOS=darwin
    set GOARCH=amd64
    echo -= mac_os_amd64
    go build -o r main.go
    tar -cf mac_os_amd64.tar -C .\ alias.json app.json config.json r > NUL 2>&1
    7z a -t7z -mx=9 mac_os_amd64.tar.xz mac_os_amd64.tar > NUL 2>&1
    move mac_os_amd64.tar.xz export\ > NUL 2>&1
    del /Q *.tar > NUL 2>&1
    del /Q r > NUL 2>&1

EXIT /B 0

:darwin_arm64
    set GOOS=darwin
    set GOARCH=arm64
    echo -= mac_os_arm64
    go build -o r main.go
    tar -cf mac_os_arm64.tar -C .\ alias.json app.json config.json r > NUL 2>&1
    7z a -t7z -mx=9 mac_os_arm64.tar.xz mac_os_arm64.tar > NUL 2>&1
    move mac_os_arm64.tar.xz export\ > NUL 2>&1
    del /Q *.tar > NUL 2>&1
    del /Q r > NUL 2>&1

EXIT /B 0











:padding
    echo ---------------------
EXIT /B 0


:MAIN
call :win_installer_start
call :win_installer_i386
call :win_installer_amd64
call :win_installer_arm
call :win_installer_arm64
call :padding

call :windows_start
call :windows_i386
call :windows_amd64
call :windows_arm
call :windows_arm64
call :padding

call :linux_start
call :linux_i386
call :linux_amd64
call :linux_arm
call :linux_arm64
call :padding

call :darwin_start
call :darwin_amd64
call :darwin_arm64
call :padding
