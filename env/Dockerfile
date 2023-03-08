FROM golang:1.19

WORKDIR /app
COPY ../ ./

RUN apt-get install git -y
RUN cd /app
RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

ENTRYPOINT [ "air" ]