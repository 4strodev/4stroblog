FROM node:22-slim AS theme
WORKDIR /app
COPY site site
COPY theme theme

RUN npm i -g pnpm
RUN cd theme && pnpm i --frozen-lockfile && pnpm build


FROM golang:1.23.5-alpine AS build
WORKDIR /app
COPY --from=theme /app/site site
COPY go-markdown-emoji go-markdown-emoji
COPY go.work go.work
COPY go.work.sum go.work.sum
COPY Taskfile.yaml Taskfile.yaml

# Installing taskfile
RUN go install github.com/go-task/task/v3/cmd/task@latest
RUN task build

FROM gcr.io/distroless/static:latest
WORKDIR /app

COPY --from=build /app/site/bin/server server

CMD ["/app/server"]
