FROM r-base

RUN R -e 'install.packages("Rserve",,"http://rforge.net")'

WORKDIR /app

EXPOSE 6311

COPY . /app

CMD ["/app/rserve.sh"]
