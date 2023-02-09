FROM golang:1.19.4-alpine AS gobuilder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o binary main.go

FROM alpine:latest
RUN apk --no-cache update && apk --no-cache add ca-certificates bash jq curl tzdata
ENV TZ=Asia/Jakarta
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
WORKDIR /app
COPY --from=gobuilder /app/binary .
COPY --from=gobuilder /app/.env .
ENTRYPOINT [ "/app/binary" ]