ALTER TABLE "tag"."tagged"
  DROP CONSTRAINT "tagged_tag_id_fkey",
  ADD CONSTRAINT "tagged_tag_id_fkey" FOREIGN KEY ("tag_id") REFERENCES "tag"."tags"("id") DEFERRABLE INITIALLY DEFERRED;
