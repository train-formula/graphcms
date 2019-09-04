CREATE SCHEMA trainer;
CREATE SCHEMA workout;
CREATE SCHEMA tag;

CREATE TYPE public.media_type AS ENUM ('PHOTO','VIDEO');
CREATE TYPE tag.tag_type AS ENUM ('WORKOUT_PROGRAM');

CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

--- Trainer Organization

CREATE TABLE "trainer"."organization" (
    "id" uuid NOT NULL,
    "created_at" timestamp without time zone NOT NULL DEFAULT NOW(),
    "updated_at" timestamp without time zone NOT NULL DEFAULT NOW(),
    "name" varchar(50) NOT NULL,
    "description" TEXT NOT NULL,
    PRIMARY KEY ("id")
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON trainer.organization
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();


--- Tags

CREATE TABLE "tag"."tags" (
  "id" uuid NOT NULL,
  "created_at" timestamp without time zone NOT NULL DEFAULT NOW(),
  "updated_at" timestamp without time zone NOT NULL DEFAULT NOW(),
  "tag" VARCHAR(100) NOT NULL,
  "trainer_organization_id" uuid NOT NULL REFERENCES trainer.organization(id) DEFERRABLE INITIALLY DEFERRED,
  PRIMARY KEY ("id")
);

CREATE UNIQUE INDEX tag_tags_trainer_organization_id_tag_idx ON tag.tags(trainer_organization_id uuid_ops, (lower(tag::text)) text_ops);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON tag.tags
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();



--- Tags tagged

CREATE TABLE "tag"."tagged" (
  "id" uuid NOT NULL,
  "created_at" timestamp without time zone NOT NULL DEFAULT NOW(),
  "updated_at" timestamp without time zone NOT NULL DEFAULT NOW(),
  "tag_uuid" uuid NOT NULL REFERENCES trainer.organization(id) DEFERRABLE INITIALLY DEFERRED,
  "trainer_organization_id" uuid NOT NULL REFERENCES trainer.organization(id) DEFERRABLE INITIALLY DEFERRED,
  "tag_type" tag.tag_type NOT NULL,
  "tagged_uuid" uuid NOT NULL,
  PRIMARY KEY ("id")
);

CREATE UNIQUE INDEX tagged_tag_uuid_tagged_uuid_tag_on_idx ON tag.tagged(tag_uuid uuid_ops, tagged_uuid uuid_ops, tag_on enum_ops);
CREATE INDEX tagged_tagged_uuid_tag_on_idx ON tag.tagged(tagged_uuid uuid_ops, tag_on enum_ops);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON tag.tagged
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

--- Exercise Unit

CREATE TABLE "workout"."rep_unit" (
    "id" uuid NOT NULL,
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
    "id" uuid NOT NULL,
    "created_at" timestamp without time zone NOT NULL DEFAULT NOW(),
    "updated_at" timestamp without time zone NOT NULL DEFAULT NOW(),
    "trainer_organization_id" uuid NOT NULL REFERENCES trainer.organization(id) DEFERRABLE INITIALLY DEFERRED,
    "name" varchar(100) NOT NULL,
    "description" TEXT NOT NULL,
    "public" boolean DEFAULT TRUE NOT NULL,
    "price" decimal(15,6) NOT NULL,
    PRIMARY KEY ("id")
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON workout.program
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE INDEX "program_trainer_organization_id_idx" ON "workout"."program"("trainer_organization_id");

--- Workout

CREATE TABLE "workout"."workout" (
    "id" uuid NOT NULL,
    "created_at" timestamp without time zone NOT NULL DEFAULT NOW(),
    "updated_at" timestamp without time zone NOT NULL DEFAULT NOW(),
    "program_id" uuid NOT NULL REFERENCES workout.program(id) DEFERRABLE INITIALLY DEFERRED,
    "name" VARCHAR(100) NOT NULL,
    "description" TEXT NOT NULL,
    "number" INT NOT NULL,
    PRIMARY KEY ("id")
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON workout.workout
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();


--- Category

CREATE TABLE "workout"."category" (
    "id" uuid NOT NULL,
    "created_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    "trainer_organization_id" uuid NOT NULL REFERENCES trainer.organization(id) DEFERRABLE INITIALLY DEFERRED,
    "name" VARCHAR(100) NOT NULL,
    "description" TEXT NOT NULL,
    PRIMARY KEY ("id")
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON workout.category
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();


--- Workout categories

CREATE TABLE "workout"."workout_category" (
    "id" uuid NOT NULL,
    "created_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    "workout_id" uuid NOT NULL REFERENCES workout.workout(id) DEFERRABLE INITIALLY DEFERRED,
    "category_id" uuid NOT NULL REFERENCES workout.category(id) DEFERRABLE INITIALLY DEFERRED,
    "order" INT DEFAULT 0 NOT NULL,
    PRIMARY KEY ("id")
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON workout.workout_category
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

--- Exercise

CREATE TABLE "workout"."exercise" (
    "id" uuid NOT NULL,
    "created_at" timestamp without time zone NOT NULL DEFAULT NOW(),
    "updated_at" timestamp without time zone NOT NULL DEFAULT NOW(),
    "trainer_organization_id" uuid NOT NULL REFERENCES trainer.organization(id) DEFERRABLE INITIALLY DEFERRED,
    "name" varchar(50) NOT NULL,
    "rep_numeral" int NOT NULL,
    "rep_text" varchar(50) NOT NULL,
    "rep_unit" uuid REFERENCES workout.rep_unit(id) DEFERRABLE INITIALLY DEFERRED,
    "rep_modifier_numeral" int NOT NULL,
    "rep_modifier_text" varchar(50) NOT NULL,
    "rep_modifier_unit" uuid REFERENCES workout.rep_unit(id) DEFERRABLE INITIALLY DEFERRED,
    "set_numeral" int NOT NULL,
    "set_text" varchar(50) NOT NULL,
    "set_unit" uuid REFERENCES workout.rep_unit(id) DEFERRABLE INITIALLY DEFERRED,
    "duration_seconds" int NOT NULL,
    PRIMARY KEY ("id")
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON workout.exercise
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

--- Exercise categories

CREATE TABLE "workout"."exercise_category" (
    "id" uuid NOT NULL,
    "created_at" timestamp without time zone NOT NULL DEFAULT NOW(),
    "updated_at" timestamp without time zone NOT NULL DEFAULT NOW(),
    "category_id" uuid NOT NULL REFERENCES workout.category(id) DEFERRABLE INITIALLY DEFERRED,
    "exercise_id" uuid NOT NULL REFERENCES workout.exercise(id) DEFERRABLE INITIALLY DEFERRED,
    "order" INT DEFAULT 0 NOT NULL,
    PRIMARY KEY ("id")
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON workout.exercise_category
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();


--- Exercise Media

CREATE TABLE "workout"."exercise_media" (
  id "uuid",
  "created_at" timestamp without time zone NOT NULL DEFAULT NOW(),
  "updated_at" timestamp without time zone NOT NULL DEFAULT NOW(),
  "trainer_organization_id" uuid NOT NULL REFERENCES trainer.organization(id) DEFERRABLE INITIALLY DEFERRED,
  "exercise_id" "uuid" NOT NULL REFERENCES workout.exercise(id) DEFERRABLE INITIALLY DEFERRED,
  "media_id" "uuid" NOT NULL,
  "order" int default 0 not null,
  "type" public.media_type not null,
  PRIMARY KEY ("id")
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON workout.exercise_media
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

--- Program Media

CREATE TABLE "workout"."program_media" (
  id "uuid",
  "created_at" timestamp without time zone NOT NULL DEFAULT NOW(),
  "updated_at" timestamp without time zone NOT NULL DEFAULT NOW(),
  "trainer_organization_id" uuid NOT NULL REFERENCES trainer.organization(id) DEFERRABLE INITIALLY DEFERRED,
  "program_id" "uuid" NOT NULL REFERENCES workout.program(id) DEFERRABLE INITIALLY DEFERRED,
  "media_id" "uuid" NOT NULL,
  "order" int default 0 not null,
  "type" public.media_type not null,
  PRIMARY KEY ("id")
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON workout.program_media
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();
