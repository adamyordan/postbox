FROM golang
RUN go get -u github.com/adamyordan/postbox
EXPOSE 8000
CMD postbox server up
