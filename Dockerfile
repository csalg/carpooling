FROM alpine:3.8

# This Dockerfile is optimized for go binaries, change it as much as necessary
# for your language of choice.

RUN apk --no-cache add ca-certificates=20190108-r0 libc6-compat=1.1.19-r11
RUN mkdir /app
COPY src/app /app/


EXPOSE 9091

ENTRYPOINT [ "/app/app" ]
