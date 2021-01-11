#!/bin/bash

cat .version | awk -F '.' '{print $1 "." $2 "." $3 + 1}' > version.tmp
mv version.tmp .version

BRANCH=`git branch | awk '{print $2}'`
VERSION=`cat .version`

echo Branch:$BRANCH Version:$VERSION

git add .
git commit -m "Automatic commit message"
COMMIT_HASH=`git log --pretty=format:'%h' -n 1`
git tag -a $VERSION -m "Automatic $COMMIT_HASH"
git push origin $BRANCH
git push origin --tags
