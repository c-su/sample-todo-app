#!/bin/sh

# Nginx インストール
echo "Install Nginx"
sudo amazon-linux-extras install -y nginx1
sudo cp -a /etc/nginx/nginx.conf /etc/nginx/nginx.conf.back
sudo cp -a ./setup_config/nginx.conf /etc/nginx/nginx.conf
sudo systemctl start nginx
sudo systemctl enable nginx
systemctl status nginx
echo "Installed Nginx"

# Go インストール
echo
echo "Install Go"
sudo amazon-linux-extras install -y golang1.11
echo "Installed Go"

# Go ライブラリインストール
echo
echo "Install Go Library"
go get -u github.com/gin-gonic/gin
go get -u github.com/jinzhu/gorm
go get -u github.com/mattn/go-sqlite3
go get -u github.com/DATA-DOG/go-sqlmock
echo "Installed aGo Library"

# ビルド
echo
echo "Go App build"
GOOS=linux GOARCH=amd64 go build main.go
echo "Done build"

# アプリケーションのデーモン起動
echo "Start Go App"
sudo cp -a ./setup_config/todo.service /etc/systemd/system/todo.service
sudo systemctl daemon-reload
sudo systemctl enable todo.service
sudo systemctl start todo.service
sudo systemctl status todo.service
echo "Started Go App"