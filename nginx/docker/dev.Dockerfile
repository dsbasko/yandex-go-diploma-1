FROM nginx:1.23.1-alpine
COPY dev.nginx.conf /etc/nginx/conf.d/default.conf.template
COPY entrypoint.sh /usr/local/bin/entrypoint.sh
RUN chmod +x /usr/local/bin/entrypoint.sh
ENTRYPOINT ["entrypoint.sh"]
CMD ["nginx", "-g", "daemon off;"]