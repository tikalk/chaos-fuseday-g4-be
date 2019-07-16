FROM alpine
MAINTAINER royp@tikalk.com
COPY chaos-fuseday-g4-be /
ENTRYPOINT ["./chaos-fuseday-g4-be"]