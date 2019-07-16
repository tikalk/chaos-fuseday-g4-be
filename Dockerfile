FROM alpine
MAINTAINER royp@tikalk.com
COPY chaos-fuseday-g4-be /
EXPOSE 8080

WORKDIR /
CMD ["/chaos-fuseday-g4-be"]
