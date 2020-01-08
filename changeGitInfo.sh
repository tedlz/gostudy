#!/bin/sh

git filter-branch --env-filter '
an="$GIT_AUTHOR_NAME"
am="$GIT_AUTHOR_EMAIL"
cn="$GIT_COMMITTER_NAME"
cm="$GIT_COMMITTER_EMAIL"
if [ "$GIT_COMMITTER_EMAIL" = "要修改的@邮箱地址.com" ]
then
    cn="想要改成的用户名"
    cm="想要改成的邮箱"
fi
if [ "$GIT_AUTHOR_EMAIL" = "要修改的@邮箱地址.com" ]
then
    an="想要改成的用户名"
    am="想要改成的邮箱"
fi
    export GIT_AUTHOR_NAME="$an"
    export GIT_AUTHOR_EMAIL="$am"
    export GIT_COMMITTER_NAME="$cn"
    export GIT_COMMITTER_EMAIL="$cm"
'
