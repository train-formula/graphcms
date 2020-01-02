ALTER TABLE "workout"."exercise" DROP COLUMN "prescription_id";

DROP TABLE workout."exercise_block";

CREATE TABLE "workout"."block_exercise_prescription" (
    "id" uuid,
    "created_at" timestamp without time zone NOT NULL DEFAULT now(),
    "updated_at" timestamp without time zone NOT NULL DEFAULT now(),
    "block_id" uuid NOT NULL,
    "exercise_id" uuid NOT NULL,
    "prescription_id" uuid NOT NULL,
    "order" integer NOT NULL DEFAULT '0',
    PRIMARY KEY ("id"),
    FOREIGN KEY ("block_id") REFERENCES "workout"."block"("id") DEFERRABLE INITIALLY DEFERRED,
    FOREIGN KEY ("exercise_id") REFERENCES "workout"."exercise"("id") DEFERRABLE INITIALLY DEFERRED,
    FOREIGN KEY ("prescription_id") REFERENCES "workout"."prescription"("id") DEFERRABLE INITIALLY DEFERRED
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON workout.block_exercise_prescription
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();