FROM golang:1.13.0

WORKDIR /go/src/app

COPY . .

RUN apt-get update && \
    apt-get install sudo wget -y && \
    wget https://dl.eff.org/certbot-auto && \
    sudo mv certbot-auto /usr/local/bin/certbot-auto && \
    sudo chown root /usr/local/bin/certbot-auto && \
    sudo chmod 0755 /usr/local/bin/certbot-auto && \
    mkdir -p /etc/letsencrypt/live/newsapp.glennfmeyer.com && \
    go get -d -v . && \
    go install -v .

EXPOSE 8080

CMD ["app"]