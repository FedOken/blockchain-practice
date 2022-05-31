# syntax=docker/dockerfile:1

FROM golang:alpine

RUN apk add gcc g++

# Get args from docker-compose .yaml
ARG ENV_WORK_DIR

# Creat WORK_DIR
RUN mkdir -p $ENV_WORK_DIR
WORKDIR $ENV_WORK_DIR



