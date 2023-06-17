FROM ubuntu:22.04 AS builder

ENV GOPATH=/go \
    GOROOT=/usr/local/go
ENV PATH=${PATH}:${GOROOT}/bin:${GOPATH}/bin

COPY ./download-file.sh /tmp/

RUN apt-get update && \
    apt-get install -y --no-install-recommends git make tar gzip && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/* && \
    /tmp/download-file.sh https://go.dev/dl/go1.19.8.linux-arm64.tar.gz '469ff06d6b86d201bcf01abbb6eed6897091afed2fd47726d90589424e4abd9a561cc24c8c2a5ce3a172a8905bfcdb35e32aef4a72ad380584f549fd20cb1d42' | tar -C /usr/local -xzf - && \
    go install github.com/mikefarah/yq/v4@v4.34.1 && \
  rm -rf /tmp/*

WORKDIR /go/src/github.com/scylladb/scylla-operator
COPY . .
RUN make build --warn-undefined-variables

FROM ubuntu:22.04

SHELL ["/usr/bin/bash", "-euExo", "pipefail", "-O", "inherit_errexit", "-c"]

# Install a minimal subset of packages that *every* runtime image needs.
RUN echo 'APT::Acquire::Retries "5";' > /etc/apt/apt.conf.d/80-retries && \
    apt-get update && \
    apt-get install -y --no-install-recommends curl ca-certificates jq && \
    apt-get clean all && \
    rm -rf /var/lib/apt/lists/*

# sidecar-injection container and existing installations use binary from root,
# we have to keep it there until we figure out how to properly upgrade them.
COPY --from=builder /go/src/github.com/scylladb/scylla-operator/scylla-operator /usr/bin/
RUN ln -s /usr/bin/scylla-operator /scylla-operator
COPY --from=builder /go/src/github.com/scylladb/scylla-operator/scylla-operator-tests /usr/bin/
COPY ./hack/gke /usr/local/lib/scylla-operator/gke
ENTRYPOINT ["/usr/bin/scylla-operator"]
