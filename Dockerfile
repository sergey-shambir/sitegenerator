# syntax=docker/dockerfile:1.7-labs

FROM golang:1.24.2 AS go_builder

WORKDIR /app

COPY --parents Makefile src/sitegenerator/go.mod src/sitegenerator/go.sum .
RUN make tidy

COPY . .
RUN make test build

FROM node:23.11.0 AS node_builder

WORKDIR /app
COPY --parents src/nodeconverter/package.json src/nodeconverter/package-lock.json .
RUN npm ci -C src/nodeconverter/

COPY src/nodeconverter/ src/nodeconverter/

FROM node:23.11.0 AS runtime

WORKDIR /app

COPY --from=go_builder /app/bin/sitegenerator /app/sitegenerator

COPY --from=node_builder /app/src/nodeconverter /app/nodeconverter

ENV SITEGENERATOR_CONVERTER_ROOT=/app/nodeconverter

CMD ["/app/sitegenerator"]