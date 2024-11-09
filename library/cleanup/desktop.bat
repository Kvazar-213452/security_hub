@echo off

DEL /F /S /Q /A %UserProfile%\Documents\Default.rdp
REG DELETE "HKCU\Software\Microsoft\Terminal Server Client" /F

exit