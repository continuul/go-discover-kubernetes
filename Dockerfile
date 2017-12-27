FROM alpine

LABEL maintainer="Robert Buck <bob@continuul.io>"

COPY ./discover /usr/local/bin/discover

ENTRYPOINT [ "/usr/local/bin/discover" ]

CMD [ "addrs", "provider=kubernetes", "namespace=default", "service=demo-nginx" ]
