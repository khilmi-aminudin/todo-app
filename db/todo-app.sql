CREATE TABLE "activity" (
  "id" BIGSERIAL PRIMARY KEY,
  "email" varchar,
  "title" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz,
  "deleted_at" timestamptz
);

CREATE TABLE "todo" (
  "id" BIGSERIAL PRIMARY KEY,
  "activity_group_id" bigint NOT NULL,
  "title" varchar,
  "is_active" varchar,
  "priority" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz,
  "deleted_at" timestamptz
);

ALTER TABLE "todo" ADD FOREIGN KEY ("activity_group_id") REFERENCES "activity" ("id");
