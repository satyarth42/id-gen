FROM golang:1.22.0-alpine3.19

ENV DC_ID=0
ENV SERVER_ID=0

WORKDIR /
ADD . ./src
WORKDIR ./src
RUN go mod vendor
RUN CGO_ENABLED=0 GOOS=linux go build -o app
WORKDIR /
RUN cp ./src/app app
RUN cp ./src/config.json config.json

FROM alpine:latest
WORKDIR /
COPY --from=0 app app
COPY --from=0 config.json config.json
EXPOSE 8080
CMD ./app