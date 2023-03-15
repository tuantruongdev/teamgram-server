FROM golang:1.17 AS builder
WORKDIR /app
COPY . .
RUN ./build.sh

FROM ubuntu:latest
WORKDIR /app
COPY --from=builder /app/teamgramd/ /app/
RUN apt update -y && apt install -y ffmpeg
RUN apt install -y dos2unix
RUN find . -type f -print0 | xargs -0 dos2unix
RUN chmod +x /app/docker/entrypoint.sh
ENTRYPOINT ["/bin/bash","/app/docker/entrypoint.sh"]
