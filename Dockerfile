ARG RELEASE_IMAGE=alpine
ARG RELEASE
ARG COMPONENT
ARG COMPONENT_EP

FROM golang:1-alpine as dev

LABEL app=hrple \
      stage=build \
      maintainer="Brett Mostert <brettmostert@gmail.com>"

RUN apk add --no-cache git bash make curl && curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.45.2


ENV HRPLE_DEV=true
ENV HRPLE_RELEASE=$RELEASE
ENV CGO_ENABLED 0

# Set the Current Working Directory inside the container
WORKDIR /go/src/app
COPY . .

RUN make build

ENTRYPOINT ["bash"]

FROM ${RELEASE_IMAGE} as release

LABEL app=hrple \
      stage=release \    
      maintainer="Brett Mostert <brettmostert@gmail.com>"

# TODO: Allow build to build all and release specific items ie. . = all and /hrple is one thing, but we could release just the cli or just the server
COPY --from=dev /go/src/app/bin/$COMPONENT /hrple/$COMPONENT
# ENTRYPOINT [$COMPONENT_EP]