# Use a multistage build to obtain a lightweight image
# Build stage
FROM golang:1.16-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

# after copying the local files inside the container WORKDIR this is the real build step
RUN go build -o /j2m-test

# In this stage we'll build the final image which will run. Practically we'll copy the build artefact (an executable file) 
# to a new container image and provide the container start command
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /j2m-test /j2m-test
USER nonroot:nonroot
ENTRYPOINT ["/j2m-test"]