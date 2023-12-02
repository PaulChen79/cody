FROM ubuntu:latest

RUN apt-get update && apt-get install -y \
    curl \
    git \
    vim \
    xz-utils \
    sudo

RUN useradd -m cody && \
    echo "cody:cody" | chpasswd && \
    adduser cody sudo && \
    echo 'cody ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers

USER cody

RUN curl -L https://nixos.org/nix/install | bash -s -- --no-daemon

USER root

ENV GO_VERSION 1.19.3

RUN curl -OL https://dl.google.com/go/go$GO_VERSION.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go$GO_VERSION.linux-amd64.tar.gz \
    && rm go$GO_VERSION.linux-amd64.tar.gz

ENV PATH="/usr/local/go/bin:${PATH}"

WORKDIR /app

COPY main /app/
COPY test.nix /app/

RUN chown -R cody:cody /app

USER cody

CMD ["./main"]