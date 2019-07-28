CREATE SCHEMA trainer;
CREATE SCHEMA workout;

CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TABLE "trainer"."organization" (
    "id" uuid,
    "created_at" timestamp without time zone NOT NULL DEFAULT NOW(),
    "updated_at" timestamp without time zone NOT NULL DEFAULT NOW(),
    "name" varchar(50) NOT NULL,
    PRIMARY KEY ("id")
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON trainer.organization
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();