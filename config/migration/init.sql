CREATE TABLE IF NOT EXISTS "cards" (
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
  "raw_card_info" JSONB NOT NULL
);

CREATE TABLE IF NOT EXISTS "decks" (
  "id" bigint PRIMARY KEY,
  "deck_name" varchar NOT NULL,
  "rank" varchar NOT NULL,
  "tournament_id" bigint NOT NULL,
  "raw_deck_info" varchar NOT NULL
);

CREATE TABLE IF NOT EXISTS "cards_in_deck" (
  "card_id" bigint NOT NULL,
  "deck_id" bigint NOT NULL,
  "card_count" int NOT NULL
);

CREATE TABLE IF NOT EXISTS "tournaments" (
  "id" bigint PRIMARY KEY,
  "tournament_name" varchar NOT NULL,
  "tier" int NOT NULL,
  "player_count" int,
  "event_date" timestamp NOT NULL,
  "format" varchar NOT NULL,
  "raw_tournament_info" varchar NOT NULL
);

CREATE TABLE IF NOT EXISTS deck_embedding (
    deck_id BIGINT PRIMARY KEY,
    name VARCHAR NOT NULL,
    embedding FLOAT8[] NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_decks_tournament_id ON "decks" ("tournament_id");
CREATE INDEX IF NOT EXISTS idx_cards_in_deck_card_id ON "cards_in_deck" ("card_id");
CREATE INDEX IF NOT EXISTS idx_cards_in_deck_deck_id ON "cards_in_deck" ("deck_id");
CREATE INDEX IF NOT EXISTS idx_tournaments_id ON "tournaments" ("id");

-- Foreign keys (optional - uncomment if safe to apply)
-- ALTER TABLE "decks" ADD CONSTRAINT fk_decks_tournament FOREIGN KEY ("tournament_id") REFERENCES "tournaments" ("id");
-- ALTER TABLE "cards_in_deck" ADD CONSTRAINT fk_cardsin_card FOREIGN KEY ("card_id") REFERENCES "cards" ("id");
-- ALTER TABLE "cards_in_deck" ADD CONSTRAINT fk_cardsin_deck FOREIGN KEY ("deck_id") REFERENCES "decks" ("id");

-- Drops for manual dev resets only (not for production use)
-- DROP TABLE IF EXISTS cards_in_deck;
-- DROP TABLE IF EXISTS decks;
-- DROP TABLE IF EXISTS tournaments;
-- DROP TABLE IF EXISTS cards;
