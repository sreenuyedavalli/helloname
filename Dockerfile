# build stage
FROM golang:alpine AS build-env
ADD . /src
RUN apk update && apk upgrade && apk add git
RUN go get github.com/gorilla/mux github.com/lib/pq github.com/shirou/gopsutil/disk github.com/shirou/gopsutil/mem
RUN cd /src && go build -o helloname

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/helloname /app/
EXPOSE 8000
