FROM golang:1.10.4
COPY main.go .
COPY collector.go .
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .
RUN cp app /tmp/app


FROM alpine
ENV FFMPEG_VERSION 4.2.1
COPY --from=0 /tmp/app .
WORKDIR /tmp
RUN wget -q https://www.johnvansickle.com/ffmpeg/old-releases/ffmpeg-4.2.1-amd64-static.tar.xz \
  && tar xJf /tmp/ffmpeg-4.2.1-amd64-static.tar.xz -C /tmp \
  && mv /tmp/ffmpeg-4.2.1-amd64-static/ffprobe /usr/local/bin/ \
  && rm -rf /tmp/ffmpeg*
ENV SRC testsrc
CMD ["/bin/sh", "-c", "ffprobe -f lavfi ${SRC},signalstats='stat=tout+vrep+brng',deflicker=bypass=1,bitplanenoise -show_frames -print_format csv | /app -port 2112"]
