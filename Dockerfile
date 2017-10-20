FROM alpine:3.6
MAINTAINER G.J.R. Timmer <gjr.timmer@gmail.com>

ARG BUILD_DATE
ARG VCS_REF

LABEL \
	nl.timmertech.build-date=${BUILD_DATE} \
	nl.timmertech.name=tcltm \
	nl.timmertech.vendor=timmertech.nl \
	nl.timmertech.vcs-url="https://github.com/GJRTimmer/tcltm.git" \
	nl.timmertech.vcs-ref=${VCS_REF} \
	nl.timmertech.license=MIT

COPY . /usr/src/tcltm
	
RUN echo 'http://nl.alpinelinux.org/alpine/edge/community' >> /etc/apk/repositories && \
	apk add --no-cache --update ca-certificates wget tcl tclx && \
	apk upgrade --update --no-cache && \
	update-ca-certificates && \
	wget https://github.com/tcltk/tcllib/archive/tcllib_1_18.tar.gz -O - | tar -xz -C /tmp && \
	tclsh /tmp/tcllib-tcllib_1_18/installer.tcl -no-html -no-nroff -no-examples -no-gui -no-apps -no-wait -pkg-path /usr/lib/tcllib1.18 && \
	apk add --virtual .build-dependencies curl make && \
	cd /usr/src/tcltm && \
	make && \
	make install && \
	make clean && \
	apk del .build-dependencies && \
	rm -rf /usr/src/tcltm

	
# EOF