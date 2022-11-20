# Membuat Depoguna API

untuk mendemokan api dengan Docker, terdiri dari:

- Backend dengan golang
- Database dengan Postgres

Semua komponen dipackage dengan docker

## Menjalankan Semuanya sebagai Docker Container

Step 1: Buat Docker Network

```bash
docker network create depoguna_network
```

Step 2: Jalankan Docker postgres di dalam network

```bash
docker run -d \
-e POSTGRES_USER=postgres \
-e POSTGRES_PASSWORD=rahasia \
-e POSTGRES_DB=depoguna_db \
-p 2345:5432 \
--network depoguna_network \
--name db-postgres \
postgres:14.5-alpine
```

Step 3: Buat docker image untuk backend

```bash
docker build -t depoguna-api:v1 .
```

Step 4: Jalankan backend sebagai docker container di dalam network

```bash
docker run -d \
-p 8080:8080 \
--name depoguna-api \
--network depoguna_network \
depoguna-api:v1
```

Step 5: Buka postman untuk mulai mengakses API

```bash
http://localhost:8080
```
