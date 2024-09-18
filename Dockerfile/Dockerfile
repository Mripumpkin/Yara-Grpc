FROM ubuntu:20.04
# COPY ./docker-entrypoint.sh /home/yara/
RUN apt-get update
ENV DEBIAN_FRONTEND=noninteractive
RUN apt install -y automake && apt install -y gcc && apt install -y make && apt install -y libtool && apt install -y pkg-config
RUN apt-get install -y yara
COPY ./yara_grpc_server /home/yara/
COPY ./setting.yml /home/yara/conf/
RUN chmod +x /home/yara/yara_grpc_server
# RUN chmod +x ./docker-entrypoint.sh
EXPOSE 2066
WORKDIR /home/yara
# ENTRYPOINT ["/home/yara/yara_grpc_server"]
