@echo off

powershell -Command "vssadmin delete shadows /all"
vssadmin delete shadows /all
vssadmin delete shadows /all /quiet