#!/bin/bash
set -e

if [ "$IS_PULL_REQUEST" == true ]
then
    echo "Skipping building of image for PRs"
    exit 0
fi

BACKEND_CHANGED_FILE_COUNT=`git diff --name-only HEAD~1..HEAD game/ action/ store/ | wc -l`;
PROTOBUFF_CHANGED_FILE_COUNT=`git diff --name-only HEAD~1..HEAD gameShipRpc/gameShipRpc.proto | wc -l`;

build_bingo_backend() {
    # Build binary and worker
    docker build -t "argonlaser/backend:$BRANCH.$BUILD_NUMBER" .
    docker images
    docker push argonlaser/backend:$BRANCH.$BUILD_NUMBER
    echo "Docker image pushed successfully"
}

build_protobuff() {
    # Build protocol buffer based on file changes in the gameShipRpc
    ./gen_proto.sh
}

if [ $PROTOBUFF_CHANGED_FILE_COUNT -gt 0 ]; then
    echo "Protobuff has changes"
    build_protobuff
fi

if [ $BACKEND_CHANGED_FILE_COUNT -gt 0 ]; then
    echo "Backend has changes"
    build_bingo_backend
fi

echo "Build Image script End"
