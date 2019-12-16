ALTER TABLE "workout"."block" DROP COLUMN "type";
DROP TYPE IF EXISTS workout.block_type;

CREATE INDEX "block_workout_category_id_category_order_idx" ON "workout"."block"("workout_category_id","category_order");


CREATE TABLE "workout"."exercise_block" (
    "id" uuid,
    "created_at" timestamp without time zone NOT NULL DEFAULT now(),
    "updated_at" timestamp without time zone NOT NULL DEFAULT now(),
    "block_id" uuid NOT NULL,
    "exercise_id" uuid NOT NULL,
    "order" integer NOT NULL DEFAULT '0',
    PRIMARY KEY ("id"),
    FOREIGN KEY ("block_id") REFERENCES "workout"."block"("id") DEFERRABLE INITIALLY DEFERRED,
    FOREIGN KEY ("exercise_id") REFERENCES "workout"."exercise"("id") DEFERRABLE INITIALLY DEFERRED
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON workout.exercise_block
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

DROP TABLE workout."exercise_category";