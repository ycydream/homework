docker build -t zhiyu.homework/httpserver:v1 .

docker run -d -p 80:80 --name httpserver zhiyu.homework/httpserver:v1

--推送到Docker官方镜像仓库需要改名字
docker tag zhiyu.homework/httpserver:v1 zhiyuhomework/httpserver:v1

docker push zhiyuhomework/httpserver:v1

docker inspect -f {{.State.Pid}} httpserver
114263

nsenter --target 114263 --net ip a

1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
32: eth0@if33: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default 
    link/ether 02:42:ac:11:00:02 brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet 172.17.0.2/16 brd 172.17.255.255 scope global eth0
       valid_lft forever preferred_lft forever


