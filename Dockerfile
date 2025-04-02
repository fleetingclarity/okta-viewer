FROM --platform=${BUILDPLATFORM} golang:1.24-alpine AS base
WORKDIR /src/github.com/fleetingclarity/okta-viewer
ENV CGO_ENABLED=0
COPY go.* .
RUN go mod download && go mod verify
COPY . .

FROM base AS build
ARG TARGETOS
ARG TARGETARCH
ENV BINARY_NAME=okta-viewer
RUN apk add git --no-cache && \
    GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /out/${BINARY_NAME} .

FROM base AS unit-test
RUN mkdir /out && go test -v -coverprofile=/out/cover.out ./...

FROM scratch AS unit-test-coverage
COPY --from=unit-test /out/cover.out /cover.out

FROM scratch AS bin-unix
ARG BINARY_NAME
COPY --from=build /out/${BINARY_NAME} /

FROM bin-unix AS bin-linux
FROM bin-unix AS bin-darwin

FROM scratch AS bin-windows
ARG BINARY_NAME
COPY --from=build /out/${BINARY_NAME} /${BINARY_NAME}.exe

FROM bin-${TARGETOS} AS bin
