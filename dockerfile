# Stage 1: Build frontend
FROM node:18 AS frontend-builder
WORKDIR /app/web
COPY web/ .
RUN npx elm-land build

FROM golang:1.22 AS backend-builder

# Install build dependencies
RUN apt-get update && apt-get install -y libsqlite3-dev libicu-dev

# Enable CGO
ENV CGO_ENABLED=1
WORKDIR /app
COPY . .
RUN go build -tags sqlite_icu -ldflags "-X main.mode=production" -o orkester

# Stage 3: Final image — with correct ICU version
FROM debian:bookworm-slim

# Install required shared libraries
RUN apt-get update && apt-get install -y libsqlite3-0 libicu72 && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /app
COPY --from=backend-builder /app/orkester .
COPY --from=frontend-builder /app/web/dist ./client

EXPOSE 42000
CMD ["./orkester"]
