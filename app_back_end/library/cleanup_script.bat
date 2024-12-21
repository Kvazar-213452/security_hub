@echo off

ipconfig /flushdns
powershell -Command "Clear-DnsClientCache"
arp -d *
nbtstat -R
fsutil usn deletejournal /d C:
fsutil behavior set encryptpagingfile 1
fsutil behavior set disablelastaccess 1
reg delete "HKEY_CURRENT_USER\SOFTWARE\Microsoft\Windows\CurrentVersion\Explorer\TypedPaths" /va /f

erase "%ALLUSERSPROFILE%\TEMP\*.*" /f /s /q
for /D %%i in ("%ALLUSERSPROFILE%\TEMP\*") do RD /S /Q "%%i"

del /f /q "%appdata%\Microsoft\teams\application cache\cache\*.*" > nul 2>&1
del /f /q "%appdata%\Microsoft\teams\blob_storage\*.*" > nul 2>&1
del /f /q "%appdata%\Microsoft\teams\databases\*.*" > nul 2>&1
del /f /q "%appdata%\Microsoft\teams\GPUcache\*.*" > nul 2>&1
del /f /q "%appdata%\Microsoft\teams\Local Storage\*.*" > nul 2>&1
del /f /q "%appdata%\Microsoft\teams\tmp\*.*" > nul 2>&1

DEL /F /S /Q /A "%UserProfile%\Documents\Default.rdp"
del /s /q /f "%LocalAppData%\Microsoft\Windows\Explorer\thumbcache*"
del /f /s /q "%AppData%\Microsoft\Windows\Recent\AutomaticDestinations\*.*"
del /f /s /q "%AppData%\Microsoft\Windows\Recent\CustomDestinations\*.*"
del /f /s /q "%userprofile%\AppData\Local\Microsoft\Windows\Explorer\*.db"
del /f /s /q "%userprofile%\AppData\Local\Microsoft\Windows\Explorer\*.etl"
del /f /s /q "%userprofile%\AppData\Local\ConnectedDevicesPlatform\*.*"
del /f /s /q "%SystemRoot%\AppCompat\Programs\*.*"
del /f /s /q "C:\Windows\appcompat\Programs\Install\*.*"
del /f /s /q "C:\Windows\System32\sru\*.*"
del /f /s /q "%userprofile%\AppData\Local\Temp\*.*"
del /f /s /q "C:\Windows\Temp\*.*"
del /f /s /q "C:\Windows\AppCompat\Programs\Amcache\sysmain.sdb"
del /f /s /q "C:\Windows\AppCompat\Programs\Amcache\*.*"
del /f /s /q "C:\ProgramData\Microsoft\Diagnosis\EventTranscript\*.*"
del /f /s /q "%userprofile%\AppData\Local\Microsoft\Terminal Server Client\*.*"
del /f /s /q "C:\ProgramData\Microsoft\Windows\WER\*.*"
del /f /s /q "%userprofile%\Appdata\Local\Microsoft\Windows\WER\*.*"
del /f /s /q "C:\Windows\apppatch\*.sdb"
del /f /s /q "%windir%\System32\LogFiles\Sum\*.*"
del /f /s /q "%windir%\SoftwareDistribution\DataStore\*.*"
del /f /s /q "C:\Windows\Prefetch\*.pf"
RunDll32.exe InetCpl.cpl,ClearMyTracksByProcess 8

erase "%LOCALAPPDATA%\Microsoft\Windows\Tempor~1\*.*" /f /s /q
DEL /f /q "%APPDATA%\Microsoft\Windows\Recent\*.*"
DEL /f /q "%APPDATA%\Microsoft\Windows\Recent\AutomaticDestinations\*.*"
DEL /f /q "%systemroot%\Panther\*.*"
DEL /f /q "%systemroot%\Prefetch\ReadyBoot\*.fx"
DEL /f /q "%systemroot%\Minidump\*.*"
del /f /s /q "c:\windows\logs\cbs\*.log"
del /f /s /q "C:\Windows\Logs\MoSetup\*.log"
del /f /s /q "C:\Windows\logs\*.log" /s /q
del /f /s /q "C:\Windows\SoftwareDistribution\*.log" /s /q
del /f /s /q "C:\Windows\Microsoft.NET\*.log" /s /q
del /f /s /q "C:\Users\%USERNAME%\AppData\Local\Microsoft\Windows\WebCache\*.log" /s /q
del /f /s /q "C:\Users\%USERNAME%\AppData\Local\Microsoft\Windows\Explorer\ThumbCacheToDelete\*.tmp" /s /q
del /f /s /q "C:\Users\%USERNAME%\AppData\Local\Microsoft\Terminal Server Client\Cache\*.bin" /s /q
del /f /s /q "C:\ProgramData\Microsoft\Windows\WER\*.*"