# ビルド環境
FROM golang:1.16-alpine as build-stage
WORKDIR /app
COPY . .
RUN apk add --no-cache gcc libc-dev
RUN go get
RUN go install
RUN go build

# 本番環境
FROM alpine:3.14.2 as production-stage
WORKDIR /CodingShare
WORKDIR /app
COPY --from=build-stage /app/CodingShare ./CodingShare
COPY ./html ./html
ENTRYPOINT [ "/app/CodingShare" ]
