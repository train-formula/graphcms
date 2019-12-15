ALTER TABLE "tag"."tagged" ALTER COLUMN "tag_type" TYPE varchar(255);
DROP TYPE tag.tag_type;
CREATE TYPE tag.tag_type AS ENUM ('WORKOUT_PROGRAM', 'WORKOUT_CATEGORY', 'WORKOUT', 'EXERCISE');
ALTER TABLE "tag"."tagged" ALTER COLUMN "tag_type" TYPE tag.tag_type USING (tag_type::tag.tag_type);

CREATE TABLE "workout"."prescription_set" (
    "id" uuid,
    "created_at" timestamp without time zone NOT NULL DEFAULT now(),
    "updated_at" timestamp without time zone NOT NULL DEFAULT now(),
    "prescription_id" uuid REFERENCES workout.prescription(id) DEFERRABLE INITIALLY DEFERRED NOT NULL,
    "set_number" int NOT NULL,
    "rep_numeral" int,
    "rep_text" varchar(50),
    "rep_unit_id" uuid REFERENCES workout.unit(id) DEFERRABLE INITIALLY DEFERRED NOT NULL,
    "rep_modifier_numeral" int,
    "rep_modifier_text" varchar(50),
    "rep_modifier_unit_id" uuid REFERENCES workout.unit(id) DEFERRABLE INITIALLY DEFERRED NOT NULL,
    PRIMARY KEY ("id")
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON workout.prescription_set
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();