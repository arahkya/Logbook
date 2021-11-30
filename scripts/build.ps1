param(
    [string]$mainGoFilePath="."
)

$workDir = Resolve-Path $mainGoFilePath
$inFile = Join-Path $workDir "main.go"

if(!(Test-Path $inFile)){
    Write-Error "cannot find go file in $workdir"
    return
}

$outputFile = Join-Path $workDir "bin/main.wasm"

if(Test-Path $outputFile){
    Remove-Item $outputFile
}

$env:GOOS="js";$env:GOARCH="wasm"; go build -o $outputFile $inFile