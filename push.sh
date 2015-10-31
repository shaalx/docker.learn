#!/bin/sh
echo commit:
read commit
git add -A
git commit -m $commit
git push hub master:dev

