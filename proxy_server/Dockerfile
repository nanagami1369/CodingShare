FROM nginx:1.21.3
COPY ./nginx.conf /etc/nginx/nginx.conf
COPY ./start.sh ./
RUN chmod 744 ./start.sh
ENTRYPOINT ["./start.sh"]
