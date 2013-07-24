#!/bin/bash
heroku create -b https://github.com/kr/heroku-buildpack-go.git
heroku addons:add librato:dev
git push heroku master
heroku ps:scale worker=1
echo "Now waiting 60s..."
sleep 60
heroku addons:open librato
