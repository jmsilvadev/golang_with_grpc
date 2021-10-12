#!/bin/sh

if [ ! -d tests/coverage/ ]; then
    mkdir tests/coverage/
fi

if [ ! -d tests/outputs/ ]; then
    mkdir tests/outputs/
fi

./bin/server&

cd tests
go test -v ./... > outputs/output_test.log

CHECKFAIL=$(cat outputs/output_test.log | grep FAIL)

cat outputs/output_test.log

if [ -z "${CHECKFAIL}" ]; then
    echo "####################"
    echo "# All Tests PASSED #"
    echo "####################"
else 
    echo "################"
    echo "# Tests FAILED #"
    echo "################"
    exit 1
fi