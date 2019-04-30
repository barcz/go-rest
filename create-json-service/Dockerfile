FROM scratch

ENV CREATOR_PORT=8081
ENV CREATOR_HOST=0.0.0.0

EXPOSE $CREATOR_PORT

ADD create-json-service /create-json-service

CMD [ "/create-json-service" ]