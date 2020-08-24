#!/bin/bash

set -x
set -o errexit
cd $(dirname $0)

readonly PACKAGE_NAME="pure-init"

readonly MKDIR="mkdir -p"
readonly RM="rm -fr"
readonly CP="cp -rf"
readonly OUTPUT="$(pwd)/output"
readonly PACKAGE_DIR=$OUTPUT/$PACKAGE_NAME
readonly TARGET=$OUTPUT/target

function compile
{
    (
        go build -o bin/pure-init
    )
}

function prepare_package
{
    $MKDIR $OUTPUT/{bin,etc}
    $CP bin etc $OUTPUT #测试环境需要
}

function main
{
    $RM $OUTPUT
    compile
    prepare_package
}

main "$@"

