ALTER TABLE "workout"."workout" ADD COLUMN "days_from_start" bigint NOT NULL;
ALTER TABLE "workout"."exercise" ADD COLUMN "video_url" text DEFAULT NULL;
ALTER TABLE "workout"."program" ADD COLUMN "number_of_days" bigint;
ALTER TABLE "workout"."unit" ADD COLUMN represents_distance boolean NOT NULL;
ALTER TABLE "workout"."prescription" ADD COLUMN "prescription_category" varchar(50) NOT NULL;
ALTER TABLE "workout"."block" ADD COLUMN "round_rest_duration" bigint;
ALTER TABLE "workout"."block" ADD COLUMN "number_of_rounds" bigint;
ALTER TABLE "workout"."workout"
  ADD COLUMN "workout_program_id" uuid NOT NULL,
  ADD FOREIGN KEY ("workout_program_id") REFERENCES "workout"."program"("id") DEFERRABLE INITIALLY DEFERRED;

CREATE INDEX "workout_category_workout_id_order_idx" ON "workout"."workout_category"("workout_id","order");


ALTER TABLE "workout"."prescription"
  DROP COLUMN "rep_numeral",
  DROP COLUMN "rep_text",
  DROP COLUMN "rep_unit_id",
  DROP COLUMN "rep_modifier_numeral",
  DROP COLUMN "rep_modifier_text",
  DROP COLUMN "rep_modifier_unit_id",
  DROP COLUMN "set_numeral",
  DROP COLUMN "set_text",
  DROP COLUMN "set_unit_id";

