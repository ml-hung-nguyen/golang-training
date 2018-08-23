FROM golang:1.10.3-alpine

WORKDIR /go/src/baitapgo_ngay1/golang-training
COPY . .

RUN apk add git
RUN go get -u github.com/go-chi/chi
RUN go get -u github.com/jinzhu/gorm
RUN go get -u github.com/lib/pq
RUN go get -u github.com/dgrijalva/jwt-go
Run go get -u github.com/go-playground/form
Run go get -u gopkg.in/go-playground/validator.v9
Run go get -u golang.org/x/crypto/bcrypt
Run go get -u github.com/mitchellh/mapstructure

CMD go run *.go

EXPOSE 8089