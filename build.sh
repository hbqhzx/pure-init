#!/bin/bash

set -x
set -o errexit
cd $(dirname $0)

readonly PACKAGE_NAME="sky-deploy"

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
        go install test
        #./bin/test > db/deploy.sql
        go install api
        go install watcher
        mv bin/watcher bin/deploy-watcher
    )
}

function prepare_package
{

    if [[ -z ${1} ]]
    then
        $MKDIR $PACKAGE_DIR/{bin,etc,run,log,static,static-hn}
        $CP bin etc doc  $PACKAGE_DIR
        $MKDIR $OUTPUT/{bin,etc,run,log,static,static-hn}
        $CP bin etc doc $OUTPUT #测试环境需要
        #curl -s "$FE_PACK?refresh=$(date +%s)" |tar -C "$PACKAGE_DIR/static" -zxf -
        #mv $PACKAGE_DIR/static/dist/* "$PACKAGE_DIR/static/"
        #curl -s "$FE_PACK_HN?refresh=$(date +%s)" |tar -C "$PACKAGE_DIR/static-hn" -zxf -
        # objcopy --strip-unneeded $PACKAGE_DIR/bin/*
        $MKDIR $TARGET
        cd $PACKAGE_DIR && tar zcf "$TARGET/${PACKAGE_NAME}.tar.gz" *
    else
        $MKDIR $OUTPUT/{bin,etc,run,log,static,static-hn}
        $CP bin etc doc $OUTPUT
        #curl -s "$FE_PACK?refresh=$(date +%s)" |tar -C "$OUTPUT/static" -zxf -
        #mv $OUTPUT/static/dist/* "$OUTPUT/static/"
        #curl -s "$FE_PACK_HN?refresh=$(date +%s)" |tar -C "$OUTPUT/static-hn" -zxf -
    fi


}

function make_self
{
    local VERSION="$(git describe --tag --long | sed -r 's/-g?/./g')"
    local installer="$OUTPUT/$PACKAGE_NAME-installer-$VERSION.bin"

    $MKDIR $TARGET
    $CP setup.sh $TARGET
    (cd $PACKAGE_DIR && tar zcf $TARGET/$PACKAGE_NAME-$VERSION.tgz *)
    makeself.sh $TARGET $installer "${PACKAGE_NAME} installer" ./setup.sh
}

function main
{
    $RM $OUTPUT
    compile
    prepare_package "$@"
    #make_self
}

main "$@"

