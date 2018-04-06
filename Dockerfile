FROM scratch

ADD ./simple_web /

ENTRYPOINT [ "/simple_web" ]
CMD [ "-p", "50051" ]
