FROM uhub.service.ucloud.cn/entropypool/alpine:3.16

COPY SphinxProxy.viper.yaml /etc/SphinxProxy/SphinxProxy.viper.yaml
COPY sphinx-proxy /usr/local/bin/sphinx-proxy

CMD ["sphinx-proxy", "run"]
