-- template: db.sql
-- Created: 2019-07-10 15:00:00
-- Modified: 2019-07-10 15:00:00
-- Project: todos
CREATE TABLE todos (
  id UUID PRIMARY KEY,
  body VARCHAR(255) NOT NULL,
  completed BOOLEAN DEFAULT FALSE
); 