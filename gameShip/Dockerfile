FROM golang:1.9

MAINTAINER Venkata krishna "vkvenkat94@gmail.com"

ADD . ./gameship

WORKDIR ./gameship
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

# Binds to port 50051
EXPOSE  50051

CMD ['make']
CMD ["make", "run"]