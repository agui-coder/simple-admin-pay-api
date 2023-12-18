FROM alpine:3.18.5

# Define the project name | 定义项目名称
ARG PROJECT=pay
# Define the config file name | 定义配置文件名
ARG CONFIG_FILE=pay.yaml
# Define the author | 定义作者
ARG AUTHOR="894784649@qq.com"

LABEL org.opencontainers.image.authors=${AUTHOR}

WORKDIR /app
ENV PROJECT=${PROJECT}
ENV CONFIG_FILE=${CONFIG_FILE}

ENV TZ=Asia/Shanghai
RUN echo "http://mirrors.aliyun.com/alpine/v3.8/main/" > /etc/apk/repositories
RUN apk update --no-cache && apk add --no-cache tzdata

COPY ./${PROJECT}_api ./
COPY ./etc/${CONFIG_FILE} ./etc/
COPY ./internal/i18n/locale/ ./etc/locale/

ENTRYPOINT ./${PROJECT}_api -f etc/${CONFIG_FILE}