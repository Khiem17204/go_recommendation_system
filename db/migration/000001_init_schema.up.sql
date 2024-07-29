CREATE TABLE "cards" (
  "id" bigint PRIMARY KEY,
  "name" varchar NOT NULL,
  "type" varchar NOT NULL,
  "frame_type" varchar NOT NULL,
  "archetype" varchar,
  "attribute" varchar,
  "race" varchar,
  "level" int,
  "attack" int,
  "defense" int,
  "description" varchar NOT NULL,
  "raw_card_info" varchar NOT NULL
);

CREATE TABLE "decks" (
  "id" bigint PRIMARY KEY,
  "deck_name" varchar NOT NULL,
  "rank" varchar NOT NULL,
  "tournament_id" bigint NOT NULL,
  "raw_deck_info" varchar NOT NULL
);

CREATE TABLE "cards_in_deck" (
  "card_id" bigint NOT NULL,
  "deck_id" bigint NOT NULL,
  "card_count" int NOT NULL
);

CREATE TABLE "tournaments" (
  "id" bigint PRIMARY KEY,
  "tournament_name" varchar NOT NULL,
  "tier" int NOT NULL,
  "player_count" int,
  "event_date" timestamp NOT NULL,
  "format" varchar NOT NULL,
  "raw_tournament_info" varchar NOT NULL
);

CREATE INDEX ON "cards" ("id");

CREATE INDEX ON "decks" ("id");

CREATE INDEX ON "decks" ("tournament_id");

CREATE INDEX ON "cards_in_deck" ("card_id");

CREATE INDEX ON "cards_in_deck" ("deck_id");

CREATE INDEX ON "tournaments" ("id");

-- ALTER TABLE "decks" ADD FOREIGN KEY ("tournament_id") REFERENCES "tournaments" ("id");

-- ALTER TABLE "cards_in_deck" ADD FOREIGN KEY ("card_id") REFERENCES "cards" ("id");

-- ALTER TABLE "cards_in_deck" ADD FOREIGN KEY ("deck_id") REFERENCES "decks" ("id");
