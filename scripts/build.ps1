param(
    [string]$mainGoFilePath="."
)

$workDir = Resolve-Path $mainGoFilePath
$inFile = Join-Path $workDir "main.go"

if(!(Test-Path $inFile)){
    Write-Error "cannot find go file in $workdir"
    return
}

$outputFile = Join-Path $workDir "bin/main.exe"

if(Test-Path $outputFile){
    Remove-Item $outputFile
}

$env:GOOS="windows";$env:GOARCH="amd64"; go build -o $outputFile $inFile