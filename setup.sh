#!/bin/sh

# yumのアップデート
sudo yum -y update

# git インストール
sudo yum -y install git
git version

# Nginx インストール
sudo amazon-linux-extras install nginx1
sudo cp -a /etc/nginx/nginx.conf /etc/nginx/nginx.conf.back
sudo cp -a ./setup_config/nginx.conf /etc/nginx/nginx.conf
sudo systemctl start nginx
sudo systemctl enable nginx
systemctl status nginx

# Go インストール
sudo amazon-linux-extras install golang1.11

# Go ライブラリインストール
go get -u github.com/gin-gonic/gin
go get -u github.com/jinzhu/gorm
go get -u github.com/mattn/go-sqlite3
go get -u github.com/DATA-DOG/go-sqlmock

# ビルド
GOOS=linux GOARCH=amd64 go build main.go

# アプリケーションのデーモン起動
sudo cp -a ./setup_config/todo.service /etc/systemd/system/todo.service
sudo systemctl daemon-reload
sudo systemctl enable todo.service
sudo systemctl start todo.service
sudo systemctl status todo.service