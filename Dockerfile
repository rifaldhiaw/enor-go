## Build
FROM golang:1.19.4 AS build

WORKDIR /go/src/app
COPY go.mod . 
COPY go.sum .
RUN go mod download

COPY *.go .
RUN CGO_ENABLED=0 go build -o /go/bin/app

## Deploy
# Now copy it into our base image.
FROM gcr.io/distroless/static-debian11
COPY --from=build /go/bin/app /

EXPOSE 8080

CMD ["/app", "serve", "--http=0.0.0.0:8080"]