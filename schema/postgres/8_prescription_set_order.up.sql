ALTER TABLE "workout"."prescription_set" ADD COLUMN "order" integer NOT NULL DEFAULT '0';
CREATE INDEX "prescription_set_order_idx" ON "workout"."prescription_set"("order");
