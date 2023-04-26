-- +migrate Up
CREATE TABLE "games" (
                         "id" integer NOT NULL,
                         "name" varchar NOT NULL,
                         "surname" varchar NOT NULL,
                         "price" varchar NOT NULL,
                         "desc" varchar NOT NULL,
                         "link" varchar NOT NULL,
                         "release" varchar NOT NULL,
                         "platforms" varchar NOT NULL,
                         "genre" varchar NOT NULL,
                         "developers" varchar NOT NULL,
                         "publishers" varchar NOT NULL,
                         "inCart" BOOLEAN NOT NULL,
                         "selected" BOOLEAN NOT NULL,
                         "isHovered" BOOLEAN NOT NULL,
                         "isLiked" BOOLEAN NOT NULL,
                         "rating" integer NOT NULL,
                         "cover" varchar NOT NULL,
                         "footage" varchar[],
                         CONSTRAINT "games_pk" PRIMARY KEY ("id")
);

-- +migrate Down
DROP TABLE IF EXISTS games  CASCADE;