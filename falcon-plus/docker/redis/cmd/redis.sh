docker run -itd --name redis \
 -p 6379:6379 \
 -v /data/conf/redis.conf:/usr/local/etc/redis/redis.conf \
 -v /data/data/redis:/data \
dockerhub.datagrand.com/global/redis:3.2.8 redis-server /usr/local/etc/redis/redis.conf