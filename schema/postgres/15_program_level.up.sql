
CREATE TYPE workout.program_level AS ENUM('BEGINNER', 'INTERMEDIATE', 'ADVANCED');
ALTER TABLE "workout"."program" ADD COLUMN "program_level" workout.program_level NOT NULL DEFAULT 'BEGINNER';