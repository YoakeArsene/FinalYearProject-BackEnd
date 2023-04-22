-- +migrate Up
CREATE TABLE "libraries" (
                              "id" varchar NOT NULL,
                              "user_id" varchar NOT NULL,
                              "game_id" varchar NOT NULL,
                              CONSTRAINT "libraries_pk" PRIMARY KEY ("id")
);

ALTER TABLE "libraries" ADD FOREIGN KEY ("user_id") REFERENCES "users"("id");
ALTER TABLE "libraries" ADD FOREIGN KEY ("game_id") REFERENCES "games"("id");

-- +migrate Down
DROP TABLE IF EXISTS libraries CASCADE;