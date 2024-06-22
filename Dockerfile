FROM golang:1.22-alpine AS base
RUN apk add build-base
RUN echo $GOPATH

copy . /app

FROM base AS build
WORKDIR /app
RUN go build -o ./build ./cmd/main.go


FROM alpine AS prod
COPY --from=build /app/build /app/build
ENTRYPOINT ["/app/build"]
