FROM node:22-slim AS base-node

RUN npm i -g pnpm


FROM golang:1.25.11-alpine AS base-golang

RUN apk add curl && \
    curl -1sLf 'https://dl.cloudsmith.io/public/task/task/setup.alpine.sh' | sh && \
    apk add task


FROM base-node AS theme

WORKDIR /app
COPY packages/site site
COPY packages/theme theme

RUN cd theme && CI=true pnpm i --frozen-lockfile && pnpm build


FROM base-golang AS build
WORKDIR /app
COPY Taskfile.yaml /app/Taskfile.yaml
COPY --from=theme /app/site packages/site

RUN task build


FROM gcr.io/distroless/static:latest
WORKDIR /app/site

COPY --from=build /app/packages/site/bin/server server
COPY --from=build /app/packages/site/assets assets
COPY --from=build /app/packages/site/views views
COPY --from=build /app/packages/site/uploads uploads
COPY --from=build /app/packages/site/config config

CMD ["/app/site/server"]
