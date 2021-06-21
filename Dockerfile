FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY ./bin/ocp-tenant-api .
RUN mkdir swagger
COPY ./swagger ./swagger
RUN chown root:root ocp-tenant-api
EXPOSE 80
CMD ["./ocp-tenant-api"]