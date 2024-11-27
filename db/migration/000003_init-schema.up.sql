CREATE TABLE "favorite_snack" (
    "name" VARCHAR PRIMARY KEY,
    "amount" INT NOT NULL DEFAULT 0,
    "company" VARCHAR NOT NULL,
    "price_per_one" VARCHAR NOT NULL,
    "is_sweet" BOOLEAN DEFAULT FALSE
);

ALTER TABLE "accounts" ADD FOREIGN KEY ("owner") REFERENCES "favorite_snack" ("name");
 