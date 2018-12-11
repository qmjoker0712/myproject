@echo off
goto comment
    Build the command lines and tests in Windows.
    Must install gcc tool before building.
:comment


call :versionSet
set NODE_APP_NAME=d4d-node
set BASEDIR=%~dp0
REM build with verison infos
set versionDir=myproject/version
set ldflagsRelase="-s -w -X %versionDir%.gitBranch=%gitBranch% -X %versionDir%.gitTag=%gitTag% -X %versionDir%.buildDate=%buildDate% -X %versionDir%.gitCommit=%gitCommit% -X %versionDir%.gitTreeState=%gitTreeState%"
set ldflagsOrigin="-X %versionDir%.gitBranch=%gitBranch% -X %versionDir%.gitTag=%gitTag% -X %versionDir%.buildDate=%buildDate% -X %versionDir%.gitCommit=%gitCommit% -X %versionDir%.gitTreeState=%gitTreeState%"


set para=%*
if not defined para (
    set para=all
)

for %%i in (%para%) do (
    call :%%i
)
pause
goto:eof

:all 
call :statis
goto:eof


:statis
echo "call statis"
echo on
go build -v -ldflags "-w -s" -o ./build/bin/statis.exe ./cmd/statis
@echo "Done statis building"
@echo off
goto:eof


:versionSet
for /F %%i in ('"git symbolic-ref --short -q HEAD"') do ( set gitBranch=%%i)

if "%gitBranch%" == "" (
   for /F %%i in ('"git describe --always --tags --abbrev=0"') do ( set gitTag=%%i)
) 

REM fixed the hour error fill with space when hour is less than 10
set hour=%time:~,2%
if "%time:~,1%"==" " set hour=0%time:~1,1%

set buildDate=%date:~0,4%-%date:~5,2%-%date:~8,2%T%hour%:%time:~3,2%:%time:~6,2%  
for /F %%i in ('"git rev-parse HEAD"') do ( set gitCommit=%%i)

for /F %%i in ('"git status|findstr 'clean'"') do ( set gitTreeState=%%i)

git status|findstr "clean" && set gitTreeState=clean || set gitTreeState=dirty

@echo off
goto:eof