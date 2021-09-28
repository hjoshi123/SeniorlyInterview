FROM golang:1.17.1-alpine3.14 as builder
COPY go.mod go.sum /go/src/github.com/hjoshi123/seniorly_interview/
WORKDIR /go/src/github.com/hjoshi123/seniorly_interview
RUN go mod download
COPY . /go/src/github.com/hjoshi123/seniorly_interview
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/seniorly github.com/hjoshi123/seniorly_interview

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/hjoshi123/seniorly_interview/.env /usr/bin/.env
RUN ls -a /usr/bin
RUN echo ${POSTGRES_DB}
COPY --from=builder /go/src/github.com/hjoshi123/seniorly_interview/build/seniorly /usr/bin/seniorly
EXPOSE 8080 8080
ENTRYPOINT ["/usr/bin/seniorly"]