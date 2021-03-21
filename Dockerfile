FROM golang:1.15

RUN go get github.com/rcarrata/go-flight && go get github.com/gorilla/mux

WORKDIR /go/src/github.com/rcarrata/go-flight/

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM scratch

COPY --from=0 /go/src/github.com/rcarrata/go-flight/main .

EXPOSE 8080

CMD ["/main"]