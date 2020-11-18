FROM alpine:latest
MAINTAINER "louisehong <louisehong4168@gmail.com>"
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \
        && apk update \
        && apk upgrade \
        && apk add --no-cache bash  \
                       curl

ENTRYPOINT ["/entrypoint.sh"]

COPY scripts/entrypoint.sh /entrypoint.sh
COPY yq /bin/yq

RUN  chmod +x /entrypoint.sh