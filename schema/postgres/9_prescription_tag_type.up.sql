ALTER TABLE "tag"."tagged" ALTER COLUMN "tag_type" TYPE varchar(255);
DROP TYPE tag.tag_type;
CREATE TYPE tag.tag_type AS ENUM ('WORKOUT_PROGRAM', 'WORKOUT_CATEGORY', 'WORKOUT', 'EXERCISE', 'PRESCRIPTION');
ALTER TABLE "tag"."tagged" ALTER COLUMN "tag_type" TYPE tag.tag_type USING (tag_type::tag.tag_type);