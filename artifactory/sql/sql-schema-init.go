package schemainit

const Create_user string = `
  CREATE TABLE IF NOT EXISTS user (
    username TEXT PRIMARY KEY NOT NULL,
    password TEXT NOT NULL,
    name TEXT,
    email TEXT,
    roles TEXT []
);`

const Create_artifactory string = `
  CREATE TABLE IF NOT EXISTS artifactory (
    id INT PRIMARY KEY NOT NULL,
    name TEXT NOT NUll,
    group_name TEXT NOT NULL,
    version TEXT NOT NULL,
    artifact_type TEXT NOT NULL,
    file_path TEXT NOT NULL,
    description TEXT,
    publisher TEXT NOT NULL,
    created_by TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    modified_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (created_by)
       REFERENCES user (username) 
);`
