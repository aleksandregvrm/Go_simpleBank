ALTER TABLE "accounts" 
DROP CONSTRAINT IF EXISTS "fk_owner";

ALTER TABLE "accounts" 
ADD CONSTRAINT "fk_owner" 
FOREIGN KEY ("owner") REFERENCES "users" ("username") ON DELETE CASCADE;