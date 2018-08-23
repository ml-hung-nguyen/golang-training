FROM golang:1.10.3-alpine

WORKDIR  /go/src/example/Exer_1/golang-training/Exer_2/

COPY . .

# RUN apk add git
# RUN go get github.com/dgrijalva/jwt-go
# RUN go get github.com/go-chi/chi
# RUN go get github.com/go-playground/form
# RUN go get github.com/jinzhu/gorm
# RUN go get github.com/lib/pq
# RUN go get golang.org/x/crypto/bcrypt
CMD go run ./main.go

EXPOSE 8080
