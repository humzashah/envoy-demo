FROM envoyproxy/envoy:latest
ARG config_file_name
COPY $config_file_name /etc/envoy.yaml
CMD /usr/local/bin/envoy -c /etc/envoy.yaml --v2-config-only --service-cluster sample-cluster
