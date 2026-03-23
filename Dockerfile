FROM ubuntu:latest

RUN apt update
RUN apt install -y curl

COPY .config/certs/zscaler-ca.pem usr/local/share/ca-certificates/zscaler-ca.crt
RUN update-ca-certificates

ADD https://mise.run /mise-install.sh
RUN chmod a+xr /mise-install.sh

RUN sh /mise-install.sh


WORKDIR /app

ENV PATH=$PATH:/root/.local/bin

COPY mise.toml /app/mise.toml
RUN env -C /app mise trust
RUN env -C /app mise install || echo "Skip..."
