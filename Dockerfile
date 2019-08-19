FROM alpine:3.9 AS BASE

RUN echo 'http://nl.alpinelinux.org/alpine/edge/community' >> /etc/apk/repositories && \
    apk add --no-cache --update --virtual .download wget && \
    apk add --no-cache --update ca-certificates wget tcl tclx && \
    apk upgrade --update --no-cache && \
    update-ca-certificates && \
    wget https://github.com/tcltk/tcllib/archive/tcllib-1-19-rc-2.tar.gz -O - | tar -xz -C /tmp && \
    tclsh /tmp/tcllib-tcllib-1-19-rc-2/installer.tcl -no-html -no-nroff -no-examples -no-gui -no-apps -no-wait -pkg-path /usr/lib/tcllib1.19 && \
    apk del .download

FROM base AS BUILDER
ARG VERSION
ENV VERSION=${VERSION}

COPY . /usr/src/tcltm

RUN apk add --no-cache --update curl make coreutils sed git && \
    cd /usr/src/tcltm && \
    VERSION=${VERSION} make build && \
    make install

FROM base AS RELEASE

ARG BUILD_DATE
ARG VCS_REF

LABEL \
    maintainer="G.J.R. Timmer" \
    org.label-schema.schema-version="1.0" \
    org.label-schema.build-date=${BUILD_DATE} \
    org.label-schema.name=tcltm \
    org.label-schema.vendor=timmertech \
    org.label-schema.url="https://gitlab.timmertech.nl/tcl/tcltm" \
    org.label-schema.vcs-url="https://gitlab.timmertech.nl/tcl/tcltm.git" \
    org.label-schema.vcs-ref=${VCS_REF} \
    com.damagehead.gitlab.license=MIT

COPY --from=BUILDER /root/.local/bin/tcltm /usr/local/bin/tcltm

RUN apk add --no-cache --update make git
