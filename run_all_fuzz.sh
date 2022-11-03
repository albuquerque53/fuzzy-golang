#!/bin/bash

set -e

fuzzTime=7

testFiles=$(grep -r --include='**_test.go' --files-with-matches 'func Fuzz' .)

for file in ${testFiles}
do
    fuzzTests=$(grep -oP 'func \K(Fuzz\w*)' $file)
    for fuzzTest in ${fuzzTests}
    do
        echo "Fuzzing $fuzzTest in $file"
        parentDir=$(dirname $file)
        go test $parentDir -run=$fuzzTest -fuzz=$fuzzTest -fuzztime=${fuzzTime}s
    done
done