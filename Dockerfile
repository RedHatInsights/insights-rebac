FROM registry.access.redhat.com/ubi8/ubi-minimal:8.9-1161 AS builder
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

FROM registry.access.redhat.com/ubi9/ubi-minimal:9.3

COPY --from=builder /workspace/bin/ciam-rebac /usr/local/bin/
COPY --from=builder /workspace/configs/config.yaml /usr/local/bin/

EXPOSE 8000
EXPOSE 9000

USER 1001

ENTRYPOINT ["/usr/local/bin/ciam-rebac","-conf","/usr/local/bin/config.yaml"]

LABEL name="ciam-rebac" \
      version="0.0.1" \
      summary="ciam-rebac service" \
      description="rebac"
