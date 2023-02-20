CREATE TABLE activities (
    activity_id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    email VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE todos (
    todo_id INT AUTO_INCREMENT PRIMARY KEY,
    activity_group_id INT,
    title VARCHAR(100) NOT NULL,
    priority VARCHAR(15) NOT NULL DEFAULT "normal",
    is_active bool NOT NULL DEFAULT false,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX activities_index_0 ON activities (activity_id);

CREATE INDEX todos_index_1 ON todos (todo_id);

ALTER TABLE todos ADD FOREIGN KEY (activity_group_id) REFERENCES activities (activity_id) ON DELETE CASCADE;