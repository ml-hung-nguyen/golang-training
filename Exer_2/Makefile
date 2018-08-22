build:
	sudo docker build -t sql:pg -f postgres.dockerfile .
	sudo docker build -t golang:demo -f go.dockerfile .

run: 
	sudo docker run -d --name sql -p sql:pg
	sudo docker run -d --name golang -v ~/go/src/example/Exer_2:/go/src/example/Exer_2 -p 8080:8080 --link sql:localhost golang:demo

stop:
	sudo docker stop sql
	sudo docker stop golang

restart:
	sudo docker restart sql
	sudo docker restart golang

rm:
	sudo docker rm sql
	sudo docker rm golang

rmi:
	sudo docker rmi sql:pg
	sudo docker rmi golang:demo

rungo: 
	sudo docker run -d --name golang -v ~/go/src/example/Exer_2:/go/src/example/Exer_2 -p 8080:8080 --link sql:localhost golang:demo

stopgo:
	sudo docker stop golang

restartgo:
	sudo docker restart golang

rmgo:
	sudo docker rm golang

rmigo:
	sudo docker rmi golang:demo
