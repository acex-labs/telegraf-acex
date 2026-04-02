FROM telegraf:1.38.1
ARG TARGETPLATFORM
COPY $TARGETPLATFORM/telegraf-acex /usr/bin/telegraf-acex
