@echo off

title ROBOCOPYSYNC

:begin
set REMOTE=C:\Users\tnguyen\WindowsBackupServer\Script\WindowsXPBackupScript\remote
set LOCAL=C:\Users\tnguyen\WindowsBackupServer\Script\WindowsXPBackupScript\local
set LOGS=C:\Users\tnguyen\WindowsBackupServer\Script\WindowsXPBackupScript\logs

set INTERVAL=10
:loop
for /f "tokens=2 delims==" %%a in ('wmic OS Get localdatetime /value') do set "dt=%%a"
set "YY=%dt:~2,2%" & set "YYYY=%dt:~0,4%" & set "MM=%dt:~4,2%" & set "DD=%dt:~6,2%"
set "HH=%dt:~8,2%" & set "Min=%dt:~10,2%" & set "Sec=%dt:~12,2%"

set "datestamp=%YYYY%%MM%%DD%" & set "timestamp=%HH%%Min%%Sec%"
set "fullstamp=%YYYY%-%MM%-%DD%_%HH%-%Min%-%Sec%"
echo datestamp: "%datestamp%"
echo timestamp: "%timestamp%"
echo fullstamp: "%fullstamp%"

echo [START] CREATED LOG FILE > ./logs/%fullstamp%_robocopy.log
robocopy ./local ./remote /e /copy:DAT /mt /z /MIR /log:./logs/%fullstamp%_robocopy.log
echo [FINISHED]
timeout %INTERVAL%

goto:loop

pause