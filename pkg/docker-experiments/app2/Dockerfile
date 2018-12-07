FROM golang:1.11-alpine AS builder

# Download and install the latest release of dep
ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

# Copy the Code From the host and compile it
WORKDIR $GOPATH/src/github.com/alamin-mahamud/awesome-go/docker-experiments/app2
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only
COPY . ./
RUN CGO_ENABLED=0 \
    GOOS=linux    \
    go build -a -installsuffx nocgo -o /app .


#-------------------------------
FROM scratch
COPY --from=builder /app ./
ENTRYPOINT ["./app"]
