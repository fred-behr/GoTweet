docker run --name gotweet-user-db \
  -e POSTGRES_USER=gotweet \
  -e POSTGRES_PASSWORD=secret \
  -e POSTGRES_DB=users \
  -p 5433:5432 \
  -v gotweet-user-data:/var/lib/postgresql/data \
  -d postgres:latest