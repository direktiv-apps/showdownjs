FROM golang:1.18-buster as build

COPY go.mod src/go.mod
COPY go.sum src/go.sum
RUN cd src/ && go mod download

COPY cmd src/cmd/
COPY models src/models/
COPY restapi src/restapi/

RUN cd src && \
    export CGO_LDFLAGS="-static -w -s" && \
    go build -tags osusergo,netgo -o /application cmd/showdownjs-server/main.go; 

FROM node:fermium-bullseye

RUN npm install showdown -g

# DON'T CHANGE BELOW 
COPY --from=build /application /bin/application

#ADD md.md /input.md

EXPOSE 8080

CMD ["/bin/application", "--port=8080", "--host=0.0.0.0"]

