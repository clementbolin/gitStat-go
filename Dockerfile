FROM golang:1.15

LABEL version="1.0" contact="clement.bolin@epitech.eu"

ARG email="clement.bolin@epitech.eu"

WORKDIR /gitStat

COPY . .

RUN make build

CMD ["./bin/gitStat-go", "-email ${email}}"]
