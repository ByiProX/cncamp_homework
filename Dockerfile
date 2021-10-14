FROM golang:1.17-alpine AS build

COPY . /go/src/project

WORKDIR /go/src/project/

RUN go build -o /bin/httpserver /go/src/project/homework01/main.go



FROM scratch
COPY --from=build /bin/httpserver /bin/httpserver

ENTRYPOINT ["/bin/httpserver"]

HEALTHCHECK --interval=5m --timeout=3s \
  CMD curl -f http://localhost:8000/healthz || exit 1