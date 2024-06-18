FROM registry.access.redhat.com/ubi8/ubi-minimal:8.10 AS builder

ARG TARGETARCH
USER root
RUN microdnf install -y tar gzip make which

# install platform specific go version
RUN curl -O -J  https://dl.google.com/go/go1.22.0.linux-${TARGETARCH}.tar.gz
RUN tar -C /usr/local -xzf go1.22.0.linux-${TARGETARCH}.tar.gz
RUN ln -s /usr/local/go/bin/go /usr/local/bin/go

WORKDIR /workspace

COPY . ./

RUN go mod vendor
RUN make build

FROM registry.access.redhat.com/ubi8/ubi-minimal:8.10


COPY --from=builder /workspace/bin/relations-api /usr/local/bin/
COPY --from=builder /workspace/configs/config.yaml /usr/local/bin/

EXPOSE 8000
EXPOSE 9000

USER 1001

ENTRYPOINT ["/usr/local/bin/relations-api","-conf","/usr/local/bin/config.yaml"]

LABEL name="relations-api" \
      version="0.0.1" \
      summary="relations-api service" \
      description="relations-api"
