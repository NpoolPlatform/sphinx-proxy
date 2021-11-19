FROM jrei/systemd-centos:7

COPY SphinxProxy.viper.yaml /etc/SphinxProxy/SphinxProxy.viper.yaml
COPY sphinx-proxy /usr/local/bin/sphinx-proxy

CMD ["sphinx-proxy", "run"]
