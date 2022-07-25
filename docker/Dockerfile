# Build
FROM golang:1.18 as build
WORKDIR /app
COPY ./go.mod ./go.sum ./
COPY ./cmd/ ./cmd/
RUN go mod download && go build -o /go/bin/extkey ./cmd/extkey

# Runtime
FROM gcr.io/distroless/base-debian11 as run
COPY --from=build /go/bin/extkey /extkey
ENV GIN_MODE=release
ENTRYPOINT ["/extkey"]
CMD ["serve", "--laddr=0.0.0.0:7000"]
