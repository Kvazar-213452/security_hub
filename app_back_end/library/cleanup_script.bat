@echo off

ipconfig /flushdns
powershell -Command "Clear-DnsClientCache"
nbtstat -R

del /f /s /q "%userprofile%\AppData\Local\Temp\*.*"
del /f /s /q "C:\Windows\Temp\*.*"

RunDll32.exe InetCpl.cpl,ClearMyTracksByProcess 8

erase "%LOCALAPPDATA%\Microsoft\Windows\Tempor~1\*.*" /f /s /q

echo.

@echo off

timeout /t 1 /nobreak > NUL
openfiles > NUL 2>&1
if %errorlevel%==0 (
    echo Running..
) else (
    echo You must run me as an Administrator. Exiting..
    echo.
    echo Right-click on me and select ^'Run as Administrator^' and try again.
    echo.
    echo Press any key to exit..
    exit
)

echo.

del /s /f /q %Temp%\*.* 
del /s /f /q %AppData%\Temp\*.* 

rd /s /q %Temp%
rd /s /q %AppData%\Temp

md %Temp%
md %AppData%\Temp

echo.
exit
