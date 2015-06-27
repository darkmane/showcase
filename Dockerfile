FROM node
MAINTAINER Sean Chitwood <darkmane@gmail.com>

RUN mkdir -p /var/www

VOLUME /var/www
ADD src /var/www

WORKDIR /var/www

ENTRYPOINT node web.js
