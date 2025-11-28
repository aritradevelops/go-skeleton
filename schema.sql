CREATE TABLE "users" (
  -- primary key
  id uuid NOT NULL DEFAULT gen_random_uuid(),
  -- custom fields
  email varchar(255) UNIQUE NOT NULL,
  name varchar(255) NOT NULL,
  -- system fields
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  created_by uuid NOT NULL,
  updated_at timestamptz,
  updated_by uuid,
  deleted_at timestamptz,
  deleted_by uuid,
  -- system constrains
  PRIMARY KEY("id"),
  FOREIGN KEY("created_by") REFERENCES "users"("id"),
  FOREIGN KEY("updated_by") REFERENCES "users"("id"), 
  FOREIGN KEY("deleted_by") REFERENCES "users"("id") 
  -- custom constrains
)