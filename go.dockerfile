FROM golang:1.10.3-alpine
WORKDIR /go/src/golang-training
COPY . .
RUN apk add git
RUN go get -u github.com/lib/pq
RUN go get -u github.com/go-chi/chi
RUN go get -u github.com/jinzhu/gorm
RUN go get -u github.com/go-playground/form
RUN go get github.com/auth0/go-jwt-middleware
RUN go get golang.org/x/crypto/bcrypt

CMD go run *.go
EXPOSE 8000
