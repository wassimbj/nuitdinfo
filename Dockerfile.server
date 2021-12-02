FROM golang:1.17

WORKDIR /go/src/app

#! COPY package.json ...
COPY go.mod go.sum ./
#! RUN npm install
RUN go mod download

COPY . .

# CompileDaemon to listen for changes
RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]
CMD CompileDaemon -command="./nuitdinfo.api" -polling=true


EXPOSE 8000
