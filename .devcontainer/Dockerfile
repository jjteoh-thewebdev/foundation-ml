# Base image with common tools and non-root user "vscode"
FROM mcr.microsoft.com/devcontainers/base:ubuntu

# Install Python
RUN apt-get update && apt-get install -y \
    curl \
    git \
    build-essential \
    gnupg \
    python3 \
    python3-pip \
    python3-venv \
    && rm -rf /var/lib/apt/lists/*

# --- Install Node.js (via NVM) ---
ENV NVM_DIR=/usr/local/nvm
ENV NODE_VERSION=22

RUN mkdir -p $NVM_DIR && \
    curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.7/install.sh | bash

# Load nvm and install Node in the same layer
RUN bash -c "export NVM_DIR=$NVM_DIR && \
    . $NVM_DIR/nvm.sh && \
    nvm install $NODE_VERSION && \
    nvm use $NODE_VERSION && \
    nvm alias default $NODE_VERSION && \
    ln -s $NVM_DIR/versions/node/v$NODE_VERSION/bin/node /usr/local/bin/node && \
    ln -s $NVM_DIR/versions/node/v$NODE_VERSION/bin/npm /usr/local/bin/npm && \
    ln -s $NVM_DIR/versions/node/v$NODE_VERSION/bin/npx /usr/local/bin/npx"

# --- Install Go ---
ENV GO_VERSION=1.22.2
RUN curl -LO https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz && \
    rm go${GO_VERSION}.linux-amd64.tar.gz

ENV PATH=$PATH:/usr/local/go/bin

# Set default user to vscode
USER vscode
