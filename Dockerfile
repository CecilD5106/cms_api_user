FROM golang:alpine as builder

RUN mkdir build

WORKDIR /build

COPY . /build/

RUN go get -d github.com/gin-gonic/gin
RUN go get -d github.com/go-sql-driver/mysql

RUN CGO_ENABLED=0 go build -a -installsuffix cgo --ldflags "-s -w" -o /build/main

FROM scratch

RUN mkdir app

WORKDIR /app

RUN adduser -S -D -H -h /app appuser

USER appuser

COPY --from=builder /build/main /app/

EXPOSE 8000

ENTRYPOINT ["./main"]