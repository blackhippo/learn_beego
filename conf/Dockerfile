FROM golang:latest
LABEL author=wei.guo
LABEL create_date=6/7/2018

# Make directories and copy go src to docker
RUN mkdir /app
RUN mkdir /app/src
ADD . /app/src

# Set up environment variable for golang and executable command
ENV GOPATH=/app
ENV PATH=$PATH:$GOPATH/bin
EXPOSE 8080

WORKDIR /app
RUN go build -o learn_beego -i learn_beego

CMD ["/app/learn_beego"]
