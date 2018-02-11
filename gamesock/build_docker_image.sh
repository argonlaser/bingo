#Build only for master and if its not a PR
docker build -t "argonlaser/gamesock:$BRANCH_NAME.$SEMAPHORE_BUILD_NUMBER" .
docker images
docker push argonlaser/gamesock:$BRANCH_NAME.$SEMAPHORE_BUILD_NUMBER