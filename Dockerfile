FROM envoyproxy/envoy:v1.24.0
COPY ./envoy.yaml /etc/envoy/envoy.yaml
EXPOSE 8080 9901
CMD ["envoy", "-c", "/etc/envoy/envoy.yaml", "-l", "info"]