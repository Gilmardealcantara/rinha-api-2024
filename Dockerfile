FROM alpine:3.14 as runtime

COPY bin/server /home/app/
EXPOSE 3000

ENV BIND_ADDRESS=0.0.0.0 
ENV BIND_PORT=3000 

CMD ["/home/app/server"]
