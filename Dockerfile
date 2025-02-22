FROM node:20-slim AS base_node
ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
RUN corepack enable
WORKDIR /app
COPY pnpm-lock.yaml package.json input.css ./
COPY . ./
RUN --mount=type=cache,id=pnpm,target=/pnpm/store pnpm install --frozen-lockfile
RUN pnpm tailwindcss -i ./input.css -o ./cmd/web/assets/tailwind.css


FROM golang:1.22-alpine as builder

WORKDIR /app
RUN apk add --no-cache make
RUN go install github.com/a-h/templ/cmd/templ@v0.2.663

COPY . ./
COPY --from=base_node /app/cmd/web/assets/tailwind.css /app/cmd/web/assets/tailwind.css
RUN templ generate
RUN GOOS=linux go build -tags remote -o /app/bin/main /app/cmd/api/main.go


FROM alpine:3.19.1
WORKDIR /app
RUN apk add --no-cache ca-certificates

COPY --from=builder /app/bin/main /app/main

EXPOSE 8080
ENTRYPOINT [ "/app/main" ]
