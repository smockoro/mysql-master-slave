#!/bin/sh

# dockerのホスト名でmysql-masterがmaster
# masterをロックして、初期化する。
mysql -u root -h mysql-master -e "RESET MASTER;"
mysql -u root -h mysql-master -e "FLUSH TABLES WITH READ LOCK;"

# masterのDB情報をダンプする。
# --all-databasesで全データベースをダンプ
mysqldump -uroot -h mysql-master --all-databases --master-data --single-transaction --flush-logs --events > /tmp/master_dump.sql
# mysqldump -uroot -h mysql-master データベース名 --master-data --single-transaction --flush-logs --events > /tmp/master_dump.sql

# ダンプしたmasterのDBをslaveにインポートする。
mysql -u root -e "STOP SLAVE;"
mysql -u root < /tmp/master_dump.sql

log_file=`mysql -u root -h mysql-master -e "SHOW MASTER STATUS\G" | grep File: | awk '{print $2}'`
pos=`mysql -u root -h mysql-master -e "SHOW MASTER STATUS\G" | grep Position: | awk '{print $2}'`

mysql -u root -e "RESET SLAVE";
mysql -u root -e "CHANGE MASTER TO MASTER_HOST='mysql-master', MASTER_USER='root', MASTER_PASSWORD='', MASTER_LOG_FILE='${log_file}', MASTER_LOG_POS=${pos};"
mysql -u root -e "start slave"

# masterをunlockする
mysql -u root -h mysql-master -e "UNLOCK TABLES;"

