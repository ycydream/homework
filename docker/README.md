docker build -t zhiyu.homework/httpserver:v1 .

docker run -d -p 80:80 --name httpserver zhiyu.homework/httpserver:v1
