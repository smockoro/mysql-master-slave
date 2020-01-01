#!/bin/sh

# 初期起動時にMasterの起動を待つ
while ! mysqladmin ping -h mysql-master --silent; do
  sleep 1
done

# dockerのホスト名でmysql-masterがmaster
# masterをロックして、初期化する。
mysql -u root -ppassword -h mysql-master -P 3306 -e "RESET MASTER;"
mysql -u root -ppassword -h mysql-master -e "FLUSH TABLES WITH READ LOCK;"

# masterのDB情報をダンプする。
# --all-databasesで全データベースをダンプ
mysqldump -uroot -ppassword -h mysql-master --all-databases --master-data --single-transaction --flush-logs --events > /tmp/master_dump.sql
# mysqldump -uroot -h mysql-master データベース名 --master-data --single-transaction --flush-logs --events > /tmp/master_dump.sql

# ダンプしたmasterのDBをslaveにインポートする。
mysql -u root -ppassword -e "STOP SLAVE;"
mysql -u root -ppassword < /tmp/master_dump.sql

log_file=`mysql -u root -ppassword -h mysql-master -e "SHOW MASTER STATUS\G" | grep File: | awk '{print $2}'`
pos=`mysql -u root -ppassword -h mysql-master -e "SHOW MASTER STATUS\G" | grep Position: | awk '{print $2}'`

mysql -u root -ppassword -e "RESET SLAVE";
mysql -u root -ppassword -e "CHANGE MASTER TO MASTER_HOST='mysql-master', MASTER_USER='root', MASTER_PASSWORD='password', MASTER_LOG_FILE='${log_file}', MASTER_LOG_POS=${pos};"
mysql -u root -ppassword -e "start slave"

# masterをunlockする
mysql -u root -ppassword -h mysql-master -e "UNLOCK TABLES;"

