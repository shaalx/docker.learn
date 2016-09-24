FROM scratch

ADD https://raw.githubusercontent.com/lijianying10/HexoCan/master/rootfs.tar / 
RUN mkdir -p /app
WORKDIR /app
COPY ./main.go /app/
ADD http://7xku3c.com1.z0.glb.clouddn.com/bookmark /app/.
EXPOSE 80

CMD ["/app/bookmark"]
