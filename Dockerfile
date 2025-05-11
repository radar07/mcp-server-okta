FROM golang:1.24.1-alpine3.20 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o mcp-server-okta ./cmd/mcp-server-okta

FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app/mcp-server-okta .

ENV OKTA_API_TOKEN="" \
    OKTA_ORG_URL=""

ENTRYPOINT ["sh", "-c", "./mcp-server-okta stdio --org_url ${OKTA_ORG_URL} --api_token ${OKTA_API_TOKEN}"]
