FROM node:22-slim AS base-node

RUN npm i -g pnpm


FROM golang:1.23.5-alpine AS base-golang
# Installing taskfile
RUN go install github.com/go-task/task/v3/cmd/task@latest


FROM base-node AS theme

WORKDIR /app
COPY site site
COPY theme theme

RUN cd theme && pnpm i --frozen-lockfile && pnpm build


FROM base-golang AS build
WORKDIR /app
COPY --from=theme /app/site site
COPY Taskfile.yaml Taskfile.yaml

RUN task build


FROM gcr.io/distroless/static:latest
WORKDIR /app/site

COPY --from=build /app/site/bin/server server
COPY --from=build /app/site/assets assets
COPY --from=build /app/site/views views
COPY --from=build /app/site/uploads uploads
COPY --from=build /app/site/config config

CMD ["/app/site/server"]
