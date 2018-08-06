docker run -itd --name mysql \
 -p 3306:3306 \
 -v /data/conf/my.cnf:/etc/mysql/my.cnf \
 -v /data/data/mysql:/var/lib/mysql \
 -e MYSQL_ROOT_PASSWORD=caidaoninb \
dockerhub.datagrand.com/global/mysql:5.6 --character-set-server=utf8 --collation-server=utf8_general_ci