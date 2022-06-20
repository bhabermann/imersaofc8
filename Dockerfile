FROM golang:1.18

WORKDIR /go/src/
ENV PATH="/GO/BIN:${PATH}"

RUN apt-get update && \
    apt-get install -y \
        build-essential \
        librdkafka-dev

CMD ["tail", "-f", "/dev/null"]