#! /bin/bash
linecount=`bash -c "find . -type f \( -name '*.go' -and -not -name '*_test.go'  \) | xargs wc -l | tail -1"`
lines=$(echo $linecount | awk '{print $1}')

filecount=`bash -c "find . -type f \( -name '*.go' -and -not -name '*_test.go'  \) | wc -l"`

testfiles=`bash -c "find . -name '*_test.go' | xargs wc -l | tail -1"`
testlines=$(echo $testfiles | awk '{print $1}')

testfiles=`bash -c "find . -name '*_test.go' | wc -l"`

echo "Total lines of code: "$lines
echo "Total lines of test code: "$testlines
echo "Total files: "$filecount
echo "Total test files: "$testfiles
echo "Average lines per file: "$((lines/filecount))
echo "Average lines per generated file: "$((testlines/testfiles))

