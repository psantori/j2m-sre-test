# Use a multistage build to obtain a lightweight image
FROM golang:1.16-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /j2m-test

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /j2m-test /j2m-test
USER nonroot:nonroot
ENTRYPOINT ["/j2m-test"]