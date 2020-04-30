FROM golang:1.10.2 as builder
COPY . /go/src/
WORKDIR /go/src/
RUN go get -u firebase.google.com/go \
  && go get -u github.com/gorilla/handlers \
  && go get -u github.com/gorilla/mux \
  && go get -u google.golang.org/api/option
RUN CGO_ENABLED=0 go build -o app cmd/main.go

FROM alpine:latest
WORKDIR /app/
COPY --from=builder /go/src/app .
CMD ["./app"]