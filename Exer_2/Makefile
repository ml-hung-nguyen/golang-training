GO_IMAGE = golang:demo
SQL_IMAGE = sql:pg
GO_CONTAINER = golang
SQL_CONTAINER = sql

build: buildsql buildgo

run: runsql rungo

stop: stopsql stopgo

restart: restartsql restartgo

rm: rmsql rmgo

rmi:  rmisql rmigo

buildgo:
	sudo docker build -t $(GO_IMAGE) -f go.dockerfile .

rungo: 
	sudo docker run -d --name $(GO_CONTAINER) -v $(shell pwd):/$(shell pwd | cut -d'/' --complement -s -f1,2,3) -p 8080:8080 --link $(SQL_CONTAINER):localhost $(GO_IMAGE)

stopgo:
	sudo docker stop $(GO_CONTAINER)

restartgo:
	sudo docker restart $(GO_CONTAINER)

rmgo:
	sudo docker rm $(GO_CONTAINER)

rmigo:
	sudo docker rmi $(GO_IMAGE)

buildsql:
	sudo docker build -t $(SQL_IMAGE) -f postgres.dockerfile .

runsql: 
	sudo docker run -d --name $(SQL_CONTAINER) $(SQL_IMAGE)

stopsql:
	sudo docker stop $(SQL_CONTAINER)

restartsql:
	sudo docker restart $(SQL_CONTAINER)

rmsql:
	sudo docker rm $(SQL_CONTAINER)

rmisql:
	sudo docker rmi $(SQL_IMAGE)
