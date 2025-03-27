# or loads/alpine:3.8
FROM golang:latest

ENV GOPROXY=https://goproxy.cn,direct

ENV WORKDIR="/app"
WORKDIR ${WORKDIR}
ADD . ${WORKDIR}/

RUN cd ${WORKDIR} && mkdir build && go build -o ./build/gdd .
EXPOSE 8000
ENTRYPOINT ["./build/gdd"]