ARG BASE_IMAGE=ghcr.io/abjrcode/cross-wails:v2.8.2

FROM ${BASE_IMAGE} as builder

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download

COPY . .


# Docker injects the value of BUILDARCH into the build process
ARG BUILDARCH

RUN dpkg --add-architecture amd64 \
  && apt-get -qq update \
  && apt-get -qq install -y libasound2-dev:amd64

RUN dpkg --add-architecture arm64 \
  && apt-get -qq update \
  && apt-get -qq install -y libasound2-dev:arm64


# Needed atm due to https://github.com/wailsapp/wails/issues/1921
RUN set -exo pipefail; \
  if [ "${BUILDARCH}" = "amd64" ]; then \
    GOOS=linux GOARCH=amd64 CC=x86_64-linux-gnu-gcc wails build -platform linux/amd64 -o goains-amd64; \
    GOOS=linux GOARCH=arm64 CC=aarch64-linux-gnu-gcc wails build -skipbindings -s -platform linux/arm64 -o goains-arm64; \
  else \
    GOOS=linux GOARCH=arm64 CC=aarch64-linux-gnu-gcc wails build -platform linux/arm64 -o goains-arm64; \
    GOOS=linux GOARCH=amd64 CC=x86_64-linux-gnu-gcc wails build -skipbindings -s -platform linux/amd64 -o goains-amd64; \
  fi;

ENTRYPOINT [ "/bin/bash" ]

#############################################################

FROM ${BASE_IMAGE}

COPY --from=builder /usr/src/app/build/bin /out

ENTRYPOINT [ "sh", "-c" ]
CMD [ "cp -r /out/. /artifacts/" ]