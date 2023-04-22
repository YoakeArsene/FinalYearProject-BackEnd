-- +migrate Up
CREATE TABLE "payments" (
                                  "id" varchar NOT NULL,
                                  "user_id" varchar NOT NULL,
                                  "price" FLOAT NOT NULL,
                                  CONSTRAINT "payments_pk" PRIMARY KEY ("id")
);

ALTER TABLE "payments" ADD FOREIGN KEY ("user_id") REFERENCES "users"("id");

-- +migrate Down
DROP TABLE IF EXISTS payments CASCADE;