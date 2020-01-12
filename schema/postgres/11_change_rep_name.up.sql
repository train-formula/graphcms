ALTER TABLE "workout"."prescription_set" RENAME COLUMN "rep_numeral" TO "primary_parameter_numeral";
ALTER TABLE "workout"."prescription_set" RENAME COLUMN "rep_text" TO "primary_parameter_text";
ALTER TABLE "workout"."prescription_set" RENAME COLUMN "rep_unit_id" TO "primary_parameter_unit_id";
ALTER TABLE "workout"."prescription_set" RENAME COLUMN "rep_modifier_numeral" TO "secondary_parameter_numeral";
ALTER TABLE "workout"."prescription_set" RENAME COLUMN "rep_modifier_text" TO "secondary_parameter_text";
ALTER TABLE "workout"."prescription_set" RENAME COLUMN "rep_modifier_unit_id" TO "secondary_parameter_unit_id";
