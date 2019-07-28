CREATE SCHEMA trainer;
CREATE SCHEMA workout;

CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

--- Trainer Organization
CREATE TABLE "trainer"."organization" (
    "id" uuid,
    "created_at" timestamp without time zone NOT NULL DEFAULT NOW(),
    "updated_at" timestamp without time zone NOT NULL DEFAULT NOW(),
    "name" varchar(50) NOT NULL,
    PRIMARY KEY ("id")
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON trainer.organization
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();


--- Exercise Unit

CREATE TABLE "workout"."rep_unit" (
    "id" uuid,
    "created_at" timestamp without time zone NOT NULL DEFAULT NOW(),
    "updated_at" timestamp without time zone NOT NULL DEFAULT NOW(),
    "name" varchar(25) NOT NULL,
    "name_medium" varchar(10) NOT NULL,
    "name_short" varchar(5) NOT NULL,
    PRIMARY KEY ("id")
);


CREATE TRIGGER set_timestamp
BEFORE UPDATE ON workout.rep_unit
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

--- Exercise Program

CREATE TABLE "workout"."program" (
    "id" uuid,
    "created_at" timestamp without time zone NOT NULL DEFAULT NOW(),
    "updated_at" timestamp without time zone NOT NULL DEFAULT NOW(),
    "trainer_organization" uuid NOT NULL,
    "name" varchar(100) NOT NULL,
    PRIMARY KEY ("id")
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON workout.program
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();


--- Exercise

CREATE TABLE "workout"."exercise" (
    "id" uuid,
    "created_at" timestamp without time zone NOT NULL DEFAULT NOW(),
    "updated_at" timestamp without time zone NOT NULL DEFAULT NOW(),
    "name" varchar(50) NOT NULL,
    "rep_numeral" int NOT NULL,
    "rep_text" varchar(50) NOT NULL,
    "rep_modifier_numeral" int NOT NULL,
    "rep_modifier_text" varchar(50) NOT NULL,
    "set_numeral" int NOT NULL,
    "set_text" varchar(50) NOT NULL,
    "duration_seconds" int NOT NULL,
    PRIMARY KEY ("id")
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON workout.exercise
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();
