#!/bin/bash

cat version | awk -F '.' '{print $1 "." $2 "." $3 + 1}' > version.tmp
#rm version
mv version.tmp .version

BRANCH=`git branch | awk '{print $2}'`
VERSION=`cat version`

echo Branch:$BRANCH Version:$VERSION

git add .
git commit -m "."
git tag -a $VERSION -m "."
git push origin $BRANCH
git push origin --tags
