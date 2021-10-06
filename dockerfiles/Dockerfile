FROM gcr.io/distroless/static:nonroot
WORKDIR /
ADD ./bin/cronor /cronor
ADD ./static /static

USER 65532:65532
ENTRYPOINT ["/cronor"]


