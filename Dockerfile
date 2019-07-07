FROM golang:1.12.6 as builder

ARG PORT
ARG GCP_PROJECT

ENV GOOGLE_APPLICATION_CREDENTIALS=/app/service_account.json
ENV GO111MODULE=on

WORKDIR /app
ADD . /app

EXPOSE $PORT
# CMD go run main.go