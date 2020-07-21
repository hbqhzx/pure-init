FROM centos:centos7
MAINTAINER Liu Ming <liuming8@jd.com>

COPY output/*-installer-* /tmp/
RUN /tmp/*-installer-* /export/servers/balloon
WORKDIR /export/servers/balloon
# CMD sleep 99999
CMD ./bin/test
