# Sprawozdanie - Laboratorium 5 PAwChO

## Opis
Aplikacja webowa napisana w Go, która generuje statyczny plik HTML z adresem IP, hostname i wersją aplikacji.
Serwowana przez serwer Nginx.

## Build
```bash
docker build -t wersja:v3 --build-arg VERSION=1.2.3 .
```

## Run
```bash
docker run -d -p 8081:80 --name testapp3 wersja:v3
```

## Test
```bash
curl http://localhost:8081
```
