build:
	sudo docker build -t golang:demo --file Go.Dockerfile .
	sudo docker build -t postgres:demo --file Db.Dockerfile .
run:
	sudo docker run -i -t -d --name pgdb -p 2345:5432 postgres:demo
	sleep 5
	sudo docker run -i -t -d -v ~/go/src/golang-training:/go/src/golang-training --name golang --link pgdb:db -p 1709:1709 golang:demo
start:
	sudo docker start pgdb golang
restart:
	sudo docker restart pgdb golang
remove:
	sudo docker rm pgdb golang -f
	sudo docker rmi golang:demo postgres:demo
