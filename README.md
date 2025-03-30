# Sprawozdanie - Laboratorium 5 PAwChO

## Opis
Aplikacja webowa napisana w Go, która generuje statyczny plik HTML z adresem IP, hostname i wersją aplikacji.
Serwowana przez serwer Nginx.

## Build
```bash
docker build -t wersja:v1 --build-arg VERSION=1.2.3 .
```

## Run
```bash
docker run -d -p 8081:80 --name testapp1 wersja:v1
```

## Test
```bash
curl http://localhost:8081
```


## Screenshots

![Zrzut ekranu 2025-03-30 144422](https://github.com/user-attachments/assets/8268014c-78dc-4366-a0e0-aed3c9114312)


![Zrzut ekranu 2025-03-30 150615](https://github.com/user-attachments/assets/95ede240-bbfd-4d85-9bf8-bb5832d00ee1)

![image](https://github.com/user-attachments/assets/5208eba0-0b5b-46fc-8574-1be7877bdeb7)
