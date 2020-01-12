DROP INDEX "workout"."prescription_set_order_idx";
ALTER TABLE "workout"."prescription_set" DROP COLUMN "order";
CREATE INDEX "prescription_set_prescription_id_set_number_idx" ON "workout"."prescription_set"("prescription_id","set_number");
