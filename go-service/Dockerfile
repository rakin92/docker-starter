# build environment
FROM golang:latest as builder
ARG VERSION
ARG COMMIT
ARG BRANCH
COPY . /go/src/app
WORKDIR /go/src/app

RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s -X main.VERSION=${VERSION} -X main.COMMIT=${COMMIT} -X main.BRANCH=${BRANCH}" -o /go/bin/main

# production environment
# scretch to only serve to main binary for smailler image
# copy over other dependent files/scripts/certs
FROM scratch
COPY --from=builder /go/bin/main /go/bin/main
WORKDIR /go/bin/
CMD ["./main"]
