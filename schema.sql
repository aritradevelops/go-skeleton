CREATE TABLE "accounts" (
  id uuid NOT NULL DEFAULT gen_random_uuid(),
  name varchar(100) NOT NULL,
  slug varchar(100) NOT NULL UNIQUE,
  -- TODO: may be move these to separate collection
  domain varchar(255),
  domain_verified boolean DEFAULT false,
  -- upto here
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  created_by uuid NOT NULL,
  updated_at timestamptz,
  updated_by uuid,
  deleted_at timestamptz,
  deleted_by uuid,
  PRIMARY KEY("id"),  
  FOREIGN KEY("created_by") REFERENCES "users"("id"),
  FOREIGN KEY("updated_by") REFERENCES "users"("id"), 
  FOREIGN KEY("deleted_by") REFERENCES "users"("id")
);

CREATE TABLE "users" (
  -- primary key
  id uuid NOT NULL DEFAULT gen_random_uuid(),
  -- custom fields
  email varchar(255) UNIQUE NOT NULL,
  name varchar(255) NOT NULL,
  account_id uuid NOT NULL,
  -- system fields
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  created_by uuid NOT NULL,
  updated_at timestamptz,
  updated_by uuid,
  deleted_at timestamptz,
  deleted_by uuid,
  -- system constrains
  PRIMARY KEY("id"),
  FOREIGN KEY("account_id") REFERENCES "accounts"("id"),
  FOREIGN KEY("created_by") REFERENCES "users"("id"),
  FOREIGN KEY("updated_by") REFERENCES "users"("id"), 
  FOREIGN KEY("deleted_by") REFERENCES "users"("id")
  -- custom constrains
);

CREATE TABLE "passwords" (
  id uuid NOT NULL DEFAULT gen_random_uuid(),
  hashed_password varchar(255) NOT NULL,
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  created_by uuid NOT NULL,
  updated_at timestamptz,
  updated_by uuid,
  deleted_at timestamptz,
  deleted_by uuid,
  PRIMARY KEY("id"),  
  FOREIGN KEY("created_by") REFERENCES "users"("id"),
  FOREIGN KEY("updated_by") REFERENCES "users"("id"), 
  FOREIGN KEY("deleted_by") REFERENCES "users"("id")
);