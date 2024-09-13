CREATE TABLE "users" (
    "id" bigserial PRIMARY KEY,
    name varchar NOT NULL,
    email varchar UNIQUE NOT NULL
    --password varchar NOT NULL
    --created_at timestamptz NOT NULL DEFAULT (now())
)