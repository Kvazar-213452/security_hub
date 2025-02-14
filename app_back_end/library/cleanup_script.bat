@REM @echo off

@REM ipconfig /flushdns
@REM powershell -Command "Clear-DnsClientCache"
@REM arp -d *
@REM nbtstat -R
@REM fsutil usn deletejournal /d C:
@REM fsutil behavior set encryptpagingfile 1
@REM fsutil behavior set disablelastaccess 1
@REM reg delete "HKEY_CURRENT_USER\SOFTWARE\Microsoft\Windows\CurrentVersion\Explorer\TypedPaths" /va /f

@REM erase "%ALLUSERSPROFILE%\TEMP\*.*" /f /s /q
@REM for /D %%i in ("%ALLUSERSPROFILE%\TEMP\*") do RD /S /Q "%%i"

@REM del /f /q "%appdata%\Microsoft\teams\application cache\cache\*.*" > nul 2>&1
@REM del /f /q "%appdata%\Microsoft\teams\blob_storage\*.*" > nul 2>&1
@REM del /f /q "%appdata%\Microsoft\teams\databases\*.*" > nul 2>&1
@REM del /f /q "%appdata%\Microsoft\teams\GPUcache\*.*" > nul 2>&1
@REM del /f /q "%appdata%\Microsoft\teams\Local Storage\*.*" > nul 2>&1
@REM del /f /q "%appdata%\Microsoft\teams\tmp\*.*" > nul 2>&1

@REM DEL /F /S /Q /A "%UserProfile%\Documents\Default.rdp"
@REM del /s /q /f "%LocalAppData%\Microsoft\Windows\Explorer\thumbcache*"
@REM del /f /s /q "%AppData%\Microsoft\Windows\Recent\AutomaticDestinations\*.*"
@REM del /f /s /q "%AppData%\Microsoft\Windows\Recent\CustomDestinations\*.*"
@REM del /f /s /q "%userprofile%\AppData\Local\Microsoft\Windows\Explorer\*.db"
@REM del /f /s /q "%userprofile%\AppData\Local\Microsoft\Windows\Explorer\*.etl"
@REM del /f /s /q "%userprofile%\AppData\Local\ConnectedDevicesPlatform\*.*"
@REM del /f /s /q "%SystemRoot%\AppCompat\Programs\*.*"
@REM del /f /s /q "C:\Windows\appcompat\Programs\Install\*.*"
@REM del /f /s /q "C:\Windows\System32\sru\*.*"
@REM del /f /s /q "%userprofile%\AppData\Local\Temp\*.*"
@REM del /f /s /q "C:\Windows\Temp\*.*"
@REM del /f /s /q "C:\Windows\AppCompat\Programs\Amcache\sysmain.sdb"
@REM del /f /s /q "C:\Windows\AppCompat\Programs\Amcache\*.*"
@REM del /f /s /q "C:\ProgramData\Microsoft\Diagnosis\EventTranscript\*.*"
@REM del /f /s /q "%userprofile%\AppData\Local\Microsoft\Terminal Server Client\*.*"
@REM del /f /s /q "C:\ProgramData\Microsoft\Windows\WER\*.*"
@REM del /f /s /q "%userprofile%\Appdata\Local\Microsoft\Windows\WER\*.*"
@REM del /f /s /q "C:\Windows\apppatch\*.sdb"
@REM del /f /s /q "%windir%\System32\LogFiles\Sum\*.*"
@REM del /f /s /q "%windir%\SoftwareDistribution\DataStore\*.*"
@REM del /f /s /q "C:\Windows\Prefetch\*.pf"
@REM RunDll32.exe InetCpl.cpl,ClearMyTracksByProcess 8

@REM erase "%LOCALAPPDATA%\Microsoft\Windows\Tempor~1\*.*" /f /s /q
@REM DEL /f /q "%APPDATA%\Microsoft\Windows\Recent\*.*"
@REM DEL /f /q "%APPDATA%\Microsoft\Windows\Recent\AutomaticDestinations\*.*"
@REM DEL /f /q "%systemroot%\Panther\*.*"
@REM DEL /f /q "%systemroot%\Prefetch\ReadyBoot\*.fx"
@REM DEL /f /q "%systemroot%\Minidump\*.*"
@REM del /f /s /q "c:\windows\logs\cbs\*.log"
@REM del /f /s /q "C:\Windows\Logs\MoSetup\*.log"
@REM del /f /s /q "C:\Windows\logs\*.log" /s /q
@REM del /f /s /q "C:\Windows\SoftwareDistribution\*.log" /s /q
@REM del /f /s /q "C:\Windows\Microsoft.NET\*.log" /s /q
@REM del /f /s /q "C:\Users\%USERNAME%\AppData\Local\Microsoft\Windows\WebCache\*.log" /s /q
@REM del /f /s /q "C:\Users\%USERNAME%\AppData\Local\Microsoft\Windows\Explorer\ThumbCacheToDelete\*.tmp" /s /q
@REM del /f /s /q "C:\Users\%USERNAME%\AppData\Local\Microsoft\Terminal Server Client\Cache\*.bin" /s /q
@REM del /f /s /q "C:\ProgramData\Microsoft\Windows\WER\*.*"



@REM exit
