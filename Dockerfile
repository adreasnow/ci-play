FROM alpine:latest
RUN apk add py3-pip && rm -rf /var/cache/apk/*

WORKDIR /app
COPY requirements.txt /app
RUN pip install --no-cache-dir --upgrade -r requirements.txt

ADD test_module /app/test_module
ENV FLASK_HOST=0.0.0.0
ENV FLASK_PORT=1000

ENTRYPOINT ["python3", "test_module"]