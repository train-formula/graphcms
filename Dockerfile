FROM golang:alpine as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build/server
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o server .
FROM scratch
COPY --from=builder /build/server /app/server
COPY --from=builder /build/schema /app/schema
WORKDIR /app/server
CMD ["./server"]