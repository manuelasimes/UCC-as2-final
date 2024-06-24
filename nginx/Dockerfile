# nginx/dockerfile
FROM nginx

# removing default configuration by nginx
RUN rm /etc/nginx/conf.d/default.conf

# copy our configuration file to right path
COPY nginx.conf /etc/nginx/nginx.conf
