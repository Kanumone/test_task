CREATE TABLE IF NOT EXISTS "users" (
  "id" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "slugs" (
  "title" varchar PRIMARY KEY,
  "created_at" timestamp DEFAULT (now()),
  "deleted" BOOLEAN DEFAULT (false),
  "deleted_at" timestamp DEFAULT (NULL)
);

CREATE TABLE IF NOT EXISTS "users_slugs" (
  "user_id" integer,
  "slug" varchar,
  "created_at" timestamp DEFAULT (now()),
  PRIMARY KEY ("user_id", "slug")
);

ALTER TABLE "users_slugs" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "users_slugs" ADD FOREIGN KEY ("slug") REFERENCES "slugs" ("title");
