ARG RELEASE_IMAGE=alpine
ARG RELEASE
ARG APP="hrple"
ARG COMPONENT="hrple-cli"
ARG COMPONENT_PATH="./components/hrple-cli"
ARG MAINTAINER="Brett Mostert <brettmostert@gmail.com>"

FROM golang:1-alpine as build

LABEL app=${APP} \
      component=${COMPONENT} \
      stage=build \
      maintainer=${MAINTAINER}

ENV HRPLE_DEV=true
ENV HRPLE_RELEASE=$RELEASE
ENV CGO_ENABLED 0

SHELL ["/bin/ash", "-eo", "pipefail", "-c"]

RUN "apk add --no-cache git bash make curl && curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin"

WORKDIR /go/src/app

# COPY ./scripts ./scripts
# COPY ./components/hrple-cli ./components/hrple-cli
COPY . .

RUN make build

ENTRYPOINT ["bash"]

# trunk-ignore(hadolint/DL3006)
FROM ${RELEASE_IMAGE} as release

LABEL app=${APP} \
      component=${COMPONENT} \
      stage=release \    
      maintainer="Brett Mostert <brettmostert@gmail.com>"

COPY --from=build /go/src/app/bin/${COMPONENT} /hrple/${COMPONENT}
ENTRYPOINT ["/bin/ash"]