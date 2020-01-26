CREATE SCHEMA plan;
--- Schema that only holds interval types
CREATE SCHEMA interval_types;

--- An time interval where the smallest unit is a day
CREATE TYPE interval_types.diurnal_interval AS ENUM('DAY', 'WEEK', 'MONTH', 'YEAR');

-- Plans
CREATE TABLE "plan"."plan" (
    "id" uuid,
    "created_at" timestamp without time zone NOT NULL DEFAULT NOW(),
    "updated_at" timestamp without time zone NOT NULL DEFAULT NOW(),
    "trainer_organization_id" uuid NOT NULL,
    "name" varchar(192) NOT NULL,
    "description" text,
    "registration_available" boolean NOT NULL,
    "archived" boolean NOT NULL,
    PRIMARY KEY ("id")
);
COMMENT ON COLUMN "plan"."plan"."registration_available" IS 'Can new users sign up for this plan?';
COMMENT ON COLUMN "plan"."plan"."archived" IS 'We generally do not delete plans, and instead mark them as archived';

CREATE INDEX "plan_trainer_organization_id_idx" ON "plan"."plan"("trainer_organization_id");
CREATE INDEX "plan_archived_idx" ON "plan"."plan"("archived");


--- Plan schedules
CREATE TABLE "plan"."plan_schedule" (
    "id" uuid,
    "created_at" timestamp without time zone NOT NULL DEFAULT NOW(),
    "updated_at" timestamp without time zone NOT NULL DEFAULT NOW(),
    "trainer_organization_id" uuid NOT NULL,
    "plan_id" uuid NOT NULL,
    "name" varchar(192) DEFAULT NULL,
    "description" varchar(192) DEFAULT NULL,
    "payment_interval" interval_types.diurnal_interval NOT NULL,
    "payment_interval_count" integer NOT NULL,
    "price_per_interval" bigint NOT NULL,
    "price_marked_down_from" bigint DEFAULT NULL,
    "duration_interval" interval_types.diurnal_interval DEFAULT NULL,
    "duration_interval_count" integer DEFAULT NULL,
    "registration_available" boolean NOT NULL,
    "archived" boolean NOT NULL,
    PRIMARY KEY ("id"),
    FOREIGN KEY ("plan_id") REFERENCES "plan"."plan"("id") DEFERRABLE INITIALLY DEFERRED
);
COMMENT ON COLUMN "plan"."plan_schedule"."price_per_interval" IS 'The price that the customer pays per payment interval';
COMMENT ON COLUMN "plan"."plan_schedule"."registration_available" IS 'Can new users sign up for this schedule?';
COMMENT ON COLUMN "plan"."plan_schedule"."price_marked_down_from" IS 'Shows the user how much the plan was marked down from (e.g. down from $100)';
COMMENT ON COLUMN "plan"."plan_schedule"."archived" IS 'We generally do not delete plan schedules, and instead mark them as archived';

CREATE INDEX "plan_schedule_plan_id_idx" ON "plan"."plan_schedule"("plan_id");
CREATE INDEX "plan_schedule_trainer_organization_id_idx" ON "plan"."plan_schedule"("trainer_organization_id");
CREATE INDEX "plan_schedule_archived_idx" ON "plan"."plan_schedule"("archived");


--- Plan customers
CREATE TABLE "plan"."plan_subscriber" (
    "id" uuid,
    "created_at" timestamp without time zone NOT NULL DEFAULT NOW(),
    "updated_at" timestamp without time zone NOT NULL DEFAULT NOW(),
    "trainer_organization_id" uuid NOT NULL,
    "plan_id" uuid NOT NULL,
    "plan_schedule_id" uuid NOT NULL,
    "customer_id" uuid NOT NULL,
    "start_date" timestamp without time zone NOT NULL,
    "end_date" timestamp without time zone DEFAULT NULL,
    "cancelled_date" timestamp without time zone DEFAULT NULL,
    PRIMARY KEY ("id"),
    FOREIGN KEY ("plan_id") REFERENCES "plan"."plan"("id") DEFERRABLE INITIALLY DEFERRED,
    FOREIGN KEY ("plan_schedule_id") REFERENCES "plan"."plan_schedule"("id") DEFERRABLE INITIALLY DEFERRED
);
COMMENT ON COLUMN "plan"."plan_subscriber"."cancelled_date" IS 'The date the customer cancelled the plan. Null if not cancelled';
COMMENT ON COLUMN "plan"."plan_subscriber"."start_date" IS 'When the plan becomes available for the customer. Required';
COMMENT ON COLUMN "plan"."plan_subscriber"."end_date" IS 'When the plan ends for the customer. De-normalized calculation of the start date + the plan schedule interval';

CREATE INDEX "plan_subscriber_trainer_organization_id_idx" ON "plan"."plan_subscriber"("trainer_organization_id");
CREATE INDEX "plan_subscriber_plan_id_plan_schedule_id_idx" ON "plan"."plan_subscriber"("plan_id","plan_schedule_id");
CREATE INDEX "plan_subscriber_customer_id_idx" ON "plan"."plan_subscriber"("customer_id");
CREATE INDEX "plan_subscriber_start_date_idx" ON "plan"."plan_subscriber"("start_date");
CREATE INDEX "plan_subscriber_end_date_idx" ON "plan"."plan_subscriber"("end_date");
CREATE INDEX "plan_subscriber_cancelled_date_idx" ON "plan"."plan_subscriber"("cancelled_date");


--- Plan inventory

CREATE TABLE "plan"."plan_inventory" (
    "id" uuid,
    "created_at" timestamp without time zone NOT NULL DEFAULT NOW(),
    "updated_at" timestamp without time zone NOT NULL DEFAULT NOW(),
    "plan_id" uuid NOT NULL,
    "plan_schedule_id" uuid,
    "total_inventory" bigint NOT NULL,
    PRIMARY KEY ("id"),
    FOREIGN KEY ("plan_id") REFERENCES "plan"."plan"("id") DEFERRABLE INITIALLY DEFERRED,
    FOREIGN KEY ("plan_schedule_id") REFERENCES "plan"."plan_schedule"("id") DEFERRABLE INITIALLY DEFERRED
);
COMMENT ON TABLE "plan"."plan_inventory" IS 'Optional inventory for a plan / schedules. Allows restricting how many people can be in a plan.';
COMMENT ON COLUMN "plan"."plan_inventory"."plan_schedule_id" IS 'Optional schedule id. Restricts inventory to a specific schedule';

CREATE UNIQUE INDEX plan_inventory_plan_id_expr_idx ON plan.plan_inventory(plan_id uuid_ops,(COALESCE(plan_schedule_id, '00000000-0000-0000-0000-000000000000'::uuid)) uuid_ops);
CREATE INDEX "plan_inventory_plan_id_idx" ON "plan"."plan_inventory"("plan_id");
CREATE INDEX "plan_inventory_plan_schedule_id_idx" ON "plan"."plan_inventory"("plan_schedule_id");


--- Allow plans to be tagged
ALTER TABLE "tag"."tagged" ALTER COLUMN "tag_type" TYPE varchar(255);
DROP TYPE tag.tag_type;
CREATE TYPE tag.tag_type AS ENUM ('WORKOUT_PROGRAM', 'WORKOUT_CATEGORY', 'WORKOUT', 'EXERCISE', 'PRESCRIPTION','PLAN');
ALTER TABLE "tag"."tagged" ALTER COLUMN "tag_type" TYPE tag.tag_type USING (tag_type::tag.tag_type);