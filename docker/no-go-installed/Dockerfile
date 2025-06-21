FROM alpine:3.22.0

RUN apk add git

RUN wget https://github.com/bfoersterling/gotracer/releases/latest/download/gotracer_linux_amd64 -O "/usr/local/bin/gotracer" \
	&& chmod 755 "/usr/local/bin/gotracer"

RUN mkdir /workspace

RUN git clone https://github.com/bfoersterling/gotracer.git /workspace
