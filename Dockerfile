FROM scratch

WORKDIR $GOPATH/gin-simple-base
COPY . $GOPATH/gin-simple-base

EXPOSE 8000
CMD ["./gin-simple-base"]