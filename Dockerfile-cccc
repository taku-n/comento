FROM nginx:alpine

RUN apk update \
    && apk add --no-cache certbot

RUN mkdir -p /app/cert

COPY init.sh  /app/cert
COPY renew.sh  /app/cert

WORKDIR /app/cert

EXPOSE 80 443

CMD ["sh", "init.sh"]
