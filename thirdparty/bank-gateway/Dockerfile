FROM bbyars/mountebank:2.4.0
COPY ./imposters.json /imposters/imposters.json
CMD [ "start", "--configfile", "/imposters/imposters.json", "--allowInjection" ]
