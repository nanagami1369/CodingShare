# ビルド環境
FROM golang:1.16-alpine as build-stage
WORKDIR /app
COPY . .
RUN go get
RUN go install
RUN go build

# 本番環境
FROM alpine:3.14.2 as production-stage
WORKDIR /app
COPY --from=build-stage /app/CodingShare ./CodingShare
CMD [ "/app/CodingShare"]
