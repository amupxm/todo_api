CREATE TABLE Todos (
  id SERIAL PRIMARY KEY,
  parent_id INTEGER,
  task TEXT NOT NULL,
  completed BOOLEAN DEFAULT FALSE,
  due_date DATE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP
);

ALTER TABLE todos ADD FOREIGN KEY (parent_id) REFERENCES todos(id);
