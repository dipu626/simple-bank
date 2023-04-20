ALTER TABLE "transfers" ADD COLUMN "created_at" timestamptz NOT NULL DEFAULT (now());

