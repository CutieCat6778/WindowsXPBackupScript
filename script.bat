@echo off

SET FILE=%0

SET YY=%DATE:~-4%
SET MM=%DATE:~-7,2%
SET DD=%DATE:~-10,2%

echo [START] CREATED LOG FILE > robocopy_%DD%-%MM%-%YY%.log
robocopy ./local ./remote %0 /log:robocopy_%DD%-%MM%-%YY%.log
echo [FINISHED]
