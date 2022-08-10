FROM golang:alpine AS builder

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src/mypackage/myapp/
COPY . .

# Fetch dependencies.
# Using go get.
RUN go get -d -v

# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/apiChain

FROM scratch

# Copy our static executable.
COPY --from=builder /go/bin/apiChain /go/bin/apiChain

# Run the hello binary.
ENTRYPOINT ["/go/bin/apiChain"]