FROM alpine:3.7
WORKDIR /
COPY callmenames /
CMD ["./callmenames"]
