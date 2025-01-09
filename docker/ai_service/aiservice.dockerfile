FROM python:3.9-slim


WORKDIR /app


RUN apt-get update \
    && apt-get install gcc -y \
    && apt-get clean

RUN pip install --upgrade \
    'databases[postgresql]' \
    psycopg2-binary\
    fastapi \
    uvicorn \
    httpx \
    g4f \
    pyopenssl \
    requests \
    urllib3 \
    cryptography

COPY ./AI_service /app/