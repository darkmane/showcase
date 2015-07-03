FROM node:0.10-wheezy
MAINTAINER Sean Chitwood <darkmane@gmail.com>

RUN mkdir -p /var/www
RUN mkdir -p /var/shared/pids

VOLUME /var/www
ADD src /var/www

WORKDIR /var/www
RUN npm install

EXPOSE 3000
ENTRYPOINT node web.js
