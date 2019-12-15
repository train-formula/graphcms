
--- Prescription + Exercise modification

CREATE TABLE "workout"."prescription" (
    "id" uuid NOT NULL,
    "created_at" timestamp without time zone NOT NULL DEFAULT NOW(),
    "updated_at" timestamp without time zone NOT NULL DEFAULT NOW(),
    "trainer_organization_id" uuid NOT NULL REFERENCES trainer.organization(id) DEFERRABLE INITIALLY DEFERRED,
    "name" varchar(50) NOT NULL,
    "rep_numeral" int,
    "rep_text" varchar(50),
    "rep_unit_id" uuid REFERENCES workout.unit(id) DEFERRABLE INITIALLY DEFERRED,
    "rep_modifier_numeral" int,
    "rep_modifier_text" varchar(50),
    "rep_modifier_unit_id" uuid REFERENCES workout.unit(id) DEFERRABLE INITIALLY DEFERRED,
    "set_numeral" int,
    "set_text" varchar(50),
    "set_unit_id" uuid REFERENCES workout.unit(id) DEFERRABLE INITIALLY DEFERRED,
    "duration_seconds" int,
    PRIMARY KEY ("id")
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON workout.prescription
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();


ALTER TABLE "workout"."exercise"
  DROP COLUMN "rep_numeral",
  DROP COLUMN "rep_text",
  DROP COLUMN "rep_unit",
  DROP COLUMN "rep_modifier_numeral",
  DROP COLUMN "rep_modifier_text",
  DROP COLUMN "rep_modifier_unit",
  DROP COLUMN "set_numeral",
  DROP COLUMN "set_text",
  DROP COLUMN "set_unit",
  DROP COLUMN "duration_seconds",
  ADD COLUMN "prescription_id" uuid NOT NULL,
  ADD FOREIGN KEY ("prescription_id") REFERENCES "workout"."prescription"("id") DEFERRABLE INITIALLY DEFERRED;


--- Block + Category modification

ALTER TABLE "workout"."category"
  DROP COLUMN "type",
  DROP COLUMN "round_numeral",
  DROP COLUMN "round_text",
  DROP COLUMN "round_unit_id",
  DROP COLUMN "duration_seconds";

DROP TYPE workout.category_type;

CREATE TYPE workout.block_type AS ENUM ('GENERAL', 'ROUND', 'TIMED_ROUND');

CREATE TABLE "workout"."block" (
    "id" uuid NOT NULL,
    "created_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    "trainer_organization_id" uuid NOT NULL REFERENCES trainer.organization(id) DEFERRABLE INITIALLY DEFERRED,
    "workout_category_id" uuid NOT NULL REFERENCES workout.category(id) DEFERRABLE INITIALLY DEFERRED,
    "category_order" int NOT NULL,
    "type" workout.block_type NOT NULL,
    "round_numeral" int,
    "round_text" varchar(50),
    "round_unit_id" uuid REFERENCES workout.unit(id) DEFERRABLE INITIALLY DEFERRED,
    "duration_seconds" int,
    PRIMARY KEY ("id")
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON workout.block
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();
