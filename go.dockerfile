FROM golang:1.10.3-alpine

WORKDIR /go/src/golang-training
COPY . .

# RUN apk add git
# RUN go get -u github.com/go-chi/chi
# RUN go get -u github.com/lib/pq
# RUN go get -u github.com/jinzhu/gorm
# RUN go get -u github.com/dgrijalva/jwt-go
# RUN go get -u github.com/go-playground/form
# RUN go get -u golang.org/x/crypto/bcrypt
# RUN go get -u gopkg.in/go-playground/validator.v9
# RUN go get -u github.com/mitchellh/mapstructure

EXPOSE 8002
CMD go run *.go
