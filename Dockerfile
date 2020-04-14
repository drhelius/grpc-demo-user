FROM registry.access.redhat.com/ubi8-minimal
COPY bin/grpc-demo-user /
EXPOSE 8080 5000
CMD ["/grpc-demo-user"]