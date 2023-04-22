-- +migrate Up
CREATE TABLE "payment_games" (
                            "payment_id" varchar NOT NULL,
                            "game_id" varchar NOT NULL,
                            CONSTRAINT "payment_games_pk" PRIMARY KEY ("payment_id", "game_id")
);

ALTER TABLE "payment_games" ADD FOREIGN KEY ("payment_id") REFERENCES "payments"("id");
ALTER TABLE "payment_games" ADD FOREIGN KEY ("game_id") REFERENCES "games"("id");

-- +migrate Down
DROP TABLE IF EXISTS payment_games CASCADE;