CREATE TABLE activities (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    email VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE todos (
    id INT AUTO_INCREMENT PRIMARY KEY,
    activity_group_id INT,
    title VARCHAR(100) NOT NULL,
    is_active bool NOT NULL DEFAULT false,
    priority VARCHAR(15) NOT NULL DEFAULT "normal",
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX activities_index_0 ON activities (id);

CREATE INDEX todos_index_1 ON todos (id);

ALTER TABLE todos ADD FOREIGN KEY (activity_group_id) REFERENCES activities (id);