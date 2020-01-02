
ALTER TABLE "workout"."block_exercise_prescription" RENAME TO "block_exercise";

CREATE UNIQUE INDEX "block_exercise_block_id_exercise_id_prescription_id_idx" ON "workout"."block_exercise"("block_id","exercise_id","prescription_id");
CREATE INDEX "block_exercise_block_id_order_idx" ON "workout"."block_exercise"("block_id","order");


