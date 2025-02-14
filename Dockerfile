FROM golang:1.20.0-buster

RUN mkdir -p /go/src/jastip-core

WORKDIR /go/src/jastip-core
COPY . .

RUN apt update
RUN apt install -y tzdata
ENV GOPRIVATE="github.com/alfisar"
ENV TZ Asia/Jakarta

ARG git_username
ARG git_token
RUN echo "GitHub Username: ${git_username}"
RUN git config --global \
    url."https://${git_username}:${git_token}@github.com/".insteadOf \
    "https://github.com"

RUN go clean -modcache
RUN export GOPROXY=https://proxy.golang.org

RUN go get -d -v ./...

RUN go build -o /go/bin/jastip-core

RUN rm -rf /go/src/jastip-core/.git
RUN rm -rf $HOME/.gitconfig

EXPOSE 8802

CMD ["jastip-core"]