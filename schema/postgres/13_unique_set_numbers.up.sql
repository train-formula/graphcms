DROP INDEX "workout"."prescription_set_prescription_id_set_number_idx";
CREATE UNIQUE INDEX "prescription_set_prescription_id_set_number_idx" ON "workout"."prescription_set"("prescription_id","set_number");
