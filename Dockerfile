FROM debian:stretch

RUN apt-get clean \
    && apt-get update --fix-missing \
    && apt-get install -y --no-install-recommends --allow-unauthenticated \
        bash \
    && apt-get autoclean \
    && apt-get autoremove \
    && rm -rf /var/lib/apt/lists/*

COPY ./build /

RUN chmod +x /hw

ENTRYPOINT ["/hw"]
