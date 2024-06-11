FROM docker.io/golang:1.22-alpine3.19 as golang_image
FROM mcr.microsoft.com/devcontainers/base:alpine-3.19

ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=1000

# renovate: datasource=repology depName=alpine_3_19/make versioning=loose
ENV MAKE_VERSION="4.4.1-r2"
# renovate: datasource=repology depName=alpine_3_19/gcc versioning=loose
ENV GCC_VERSION="13.2.1_git20231014-r0"
# renovate: datasource=repology depName=alpine_3_19/binutils-gold versioning=loose
ENV BINUTILS_GOLD_VERSION="2.41-r0"
# renovate: datasource=repology depName=alpine_3_19/musl-dev versioning=loose
ENV MUSL_DEV_VERSION="1.2.4_git20230717-r4"
# renovate: datasource=repology depName=alpine_3_19/protoc versioning=loose
ENV PROTOC_VERSION="24.4-r0"
# renovate: datasource=repology depName=alpine_3_19/nodejs-current versioning=loose
ENV NODE_VERSION="22.3.0"

# TODO: Get the specified nodejs version instead of current

# Install packages and Go language server
RUN apk add --no-cache --update make=${MAKE_VERSION} gcc=${GCC_VERSION} binutils-gold=${BINUTILS_GOLD_VERSION} musl-dev=${MUSL_DEV_VERSION} protoc=${PROTOC_VERSION} \
  nodejs-current \
  && corepack enable

ENV PATH=$PATH:/usr/local/go/bin:/usr/local/node/bin

# Build packages.
COPY --from=golang_image /usr/local/go /usr/local/go
ENV COREPACK_ENABLE_DOWNLOAD_PROMPT=0

USER $USERNAME

COPY Makefile /workspace/

# Cache tools
RUN cd /workspace && make install-tools && rm -rf /Makefile

ENV GOPATH=/tmp/go
ENV HOME=/tmp/home
ENV PATH=$PATH:/usr/local/go/bin:/tmp/go/bin

WORKDIR /workspace

CMD [ "/bin/sh" ]
