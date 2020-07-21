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
        export GOPATH=$(pwd)
        #./bin/test > db/deploy.sql
        go install api
    )
}

function prepare_package
{

    if [[ -z ${1} ]]
    then
        $MKDIR $PACKAGE_DIR/{bin,etc,run,log}
        $CP bin etc  $PACKAGE_DIR
        $MKDIR $OUTPUT/{bin,etc,run,log}
        $CP bin etc $OUTPUT #测试环境需要
    else
        $MKDIR $OUTPUT/{bin,etc,run,log}
        $CP bin etc $OUTPUT
    fi


}

function main
{
    $RM $OUTPUT
    compile
    prepare_package "$@"
}

main "$@"

