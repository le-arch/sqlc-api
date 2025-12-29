CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE todos (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    title TEXT NOT NULL,
      description TEXT,
        completed BOOLEAN NOT NULL DEFAULT FALSE,
          priority INT NOT NULL DEFAULT 3,
            created_at TIMESTAMPTZ DEFAULT now(),
              updated_at TIMESTAMPTZ DEFAULT now()
              );