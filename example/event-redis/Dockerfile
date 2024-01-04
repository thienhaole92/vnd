ARG GO_VERSION=1.21
ARG ALPINE_VERSION=3.18

### Build
FROM --platform=$BUILDPLATFORM golang:${GO_VERSION}-alpine${ALPINE_VERSION} AS builder

# Dependency
RUN apk add --no-cache git gcc g++ make

# Install dependencies
WORKDIR /src
COPY go.mod .
COPY go.sum .
RUN go mod download

# Compile app with rest of source
COPY . .
ARG TARGETOS TARGETARCH
RUN CGO_ENABLED=0 make build

FROM alpine:${ALPINE_VERSION}

RUN apk add --no-cache curl

WORKDIR /app
COPY --from=builder /src/event-redis .

CMD [ "/app/event-redis" ]
