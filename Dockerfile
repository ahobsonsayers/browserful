FROM node:24-slim AS agent-browser
RUN npm install -g agent-browser@0.28.0

FROM cloakhq/cloakbrowser:latest AS cloakbrowser

# Move cloakbrowser to a stable location
RUN CLOAKBROWSER_DIR=$(find /root/.cloakbrowser -maxdepth 1 -type d -name "chromium-*" | head -1) && \
    mv "$CLOAKBROWSER_DIR" /chromium

# Builder Image
FROM golang:1.24 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -v -o ./bin/ .

# Distribution Image
FROM debian:stable-slim

ARG TARGETARCH

ENV HOME=/browserful

ENV BROWSERFUL_ALLOWED_ORIGINS=*
ENV BROWSERFUL_BROWSER_EXECUTABLE_PATH=/opt/cloakbrowser/chrome
ENV BROWSERFUL_DATA_DIR=/data

# Install system dependencies
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    ca-certificates \
    tini \
    # CloakBrowser dependencies
    fonts-liberation \
    fonts-noto-color-emoji \
    libasound2 \
    libatk-bridge2.0-0 \
    libatk1.0-0 \
    libatspi2.0-0 \
    libcairo-gobject2 \
    libcairo2 \
    libcups2 \
    libdbus-1-3 \
    libdrm2 \
    libexpat1 \
    libffi8 \
    libfontconfig1 \
    libgbm1 \
    libgdk-pixbuf-2.0-0 \
    libglib2.0-0 \
    libgnutls30 \
    libgtk-3-0 \
    libnspr4 \
    libnss3 \
    libpango-1.0-0 \
    libpangocairo-1.0-0 \
    libpcre2-8-0 \
    libselinux1 \
    libudev1 \
    libx11-6 \
    libx11-xcb1 \
    libxau6 \
    libxcb1 \
    libxcomposite1 \
    libxcursor1 \
    libxdamage1 \
    libxext6 \
    libxfixes3 \
    libxi6 \
    libxkbcommon0 \
    libxrandr2 \
    libxrender1 \
    libxshmfence1 \
    libxss1 \
    libxtst6 \
    libz1 \
    zlib1g && \
    rm -rf /var/lib/apt/lists/*

# User and directory setup
RUN useradd browser --uid 1000 --home-dir "$HOME" && \
    mkdir -p "$HOME" "$BROWSERFUL_DATA_DIR" && \
    chown browser:browser "$HOME" "$BROWSERFUL_DATA_DIR"

# Install agent-browser
COPY --from=agent-browser --chown=browser:browser \
     /usr/local/lib/node_modules/agent-browser/bin/agent-browser-linux-$TARGETARCH \
     /usr/local/bin/agent-browser

# Install cloakbrowser
COPY --from=cloakbrowser --chown=browser:browser /chromium/ /opt/cloakbrowser/

# Install browserful
COPY --from=builder --chown=browser:browser /app/bin/browserful /usr/local/bin/browserful

USER browser
WORKDIR "$HOME"

ENTRYPOINT ["tini", "--"]
CMD ["/usr/local/bin/browserful"]
