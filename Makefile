build:
	sudo docker build -t db:doc -f db.dockerfile .
	sudo docker build -t go:doc -f go.dockerfile .
run:
	sudo docker run -i -t -d --name db db:doc
	sudo docker run -d --name go -p 8111:8002 -v ~/go/src/golang-training:/go/src/golang-training --link db:localhost go:doc
stop:
	sudo docker stop db go

start:
	sudo docker start db
	sudo docker start go

restart:
	sudo docker restart db
	sudo docker restart go

rm:
	sudo docker rm db
	sudo docker rm go

rmi:
	sudo docker rmi db:doc
	sudo docker rmi go:doc