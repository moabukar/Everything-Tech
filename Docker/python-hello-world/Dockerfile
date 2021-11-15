FROM python:3.6.1-alpine

WORKDIR /app

RUN pip install --upgrade pip

COPY . .

RUN pip3 install -r requirements.txt

CMD ["python3","app.py"]

