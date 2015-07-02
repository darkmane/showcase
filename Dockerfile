FROM node:slim
MAINTAINER Sean Chitwood <darkmane@gmail.com>

RUN mkdir -p /var/www
RUN mkdir -p /var/shared/pids

VOLUME /var/www
ADD src /var/www

WORKDIR /var/www

EXPOSE 3000
ENTRYPOINT node web.js
