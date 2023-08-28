CREATE TABLE "users" (
  "id" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "username" varchar NOT NULL UNIQUE,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "slugs" (
  "id" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "title" varchar NOT NULL UNIQUE,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "users_slugs" (
  "user_id" integer,
  "slug_id" integer,
  "created_at" timestamp DEFAULT (now()),
  PRIMARY KEY ("user_id", "slug_id")
);

ALTER TABLE "users_slugs" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "users_slugs" ADD FOREIGN KEY ("slug_id") REFERENCES "slugs" ("id");

