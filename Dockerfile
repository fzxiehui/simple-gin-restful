# Build Stage
FROM lacion/alpine-golang-buildimage:1.13 AS build-stage

LABEL app="build-simple-gin-restful"
LABEL REPO="https://github.com/fzxiehui/simple-gin-restful"

ENV PROJPATH=/go/src/github.com/fzxiehui/simple-gin-restful

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

ADD . /go/src/github.com/fzxiehui/simple-gin-restful
WORKDIR /go/src/github.com/fzxiehui/simple-gin-restful

RUN make build-alpine

# Final Stage
FROM lacion/alpine-base-image:latest

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/fzxiehui/simple-gin-restful"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/opt/simple-gin-restful/bin

WORKDIR /opt/simple-gin-restful/bin

COPY --from=build-stage /go/src/github.com/fzxiehui/simple-gin-restful/bin/simple-gin-restful /opt/simple-gin-restful/bin/
RUN chmod +x /opt/simple-gin-restful/bin/simple-gin-restful

# Create appuser
RUN adduser -D -g '' simple-gin-restful
USER simple-gin-restful

ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["/opt/simple-gin-restful/bin/simple-gin-restful"]
