docker build -t zhiyu.homework/httpserver:v1 .

docker run -d -p 80:80 --name httpserver zhiyu.homework/httpserver:v1

--推送到Docker官方镜像仓库需要改名字
docker tag zhiyu.homework/httpserver:v1 zhiyuhomework/httpserver:v1

docker push zhiyuhomework/httpserver:v1


