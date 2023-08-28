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
  "users_id" integer,
  "slugs_id" integer,
  PRIMARY KEY ("users_id", "slugs_id")
);

ALTER TABLE "users_slugs" ADD FOREIGN KEY ("users_id") REFERENCES "users" ("id");

ALTER TABLE "users_slugs" ADD FOREIGN KEY ("slugs_id") REFERENCES "slugs" ("id");

