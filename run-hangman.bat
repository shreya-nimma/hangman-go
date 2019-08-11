@echo off
echo ">> Building Go file"
go build 
if ERRORLEVEL 1 (
    echo ">> Exited due to error. Check output."
) else (
    echo ">> Finished build. Now executing code."
    hangman-go
)
