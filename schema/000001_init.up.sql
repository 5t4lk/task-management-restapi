CREATE TABLE users
(
    id            INT AUTO_INCREMENT PRIMARY KEY,
    name          VARCHAR(255) NOT NULL,
    username      VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL
);

CREATE TABLE tasks
(
    id          INT AUTO_INCREMENT PRIMARY KEY,
    title       VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    status      VARCHAR(255),
    end_date    VARCHAR(255)
);

CREATE TABLE users_tasks (
    id      INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT,
    task_id INT,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (task_id) REFERENCES tasks(id) ON DELETE CASCADE
);

CREATE TABLE items
(
    id          INT AUTO_INCREMENT PRIMARY KEY,
    title       VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    status      VARCHAR(255),
    end_date    VARCHAR(255),
    done        BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE tasks_items
(
    id      INT AUTO_INCREMENT PRIMARY KEY,
    item_id INT REFERENCES items(id) ON DELETE CASCADE,
    task_id INT REFERENCES tasks(id) ON DELETE CASCADE
);
