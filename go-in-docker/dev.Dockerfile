# Please keep up to date with the new-version of Golang docker for builder
FROM golang:1.19.3

RUN apt update && apt upgrade -y && \
	apt install -y git \
	make openssh-client

WORKDIR /app 
# add these two lines
ADD go.mod go.sum /app/
RUN go mod download
# RUN go mod tidy

RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
	&& chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air

CMD air