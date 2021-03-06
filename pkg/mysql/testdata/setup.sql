CREATE TABLE articles (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(100) NOT NULL,
    url VARCHAR(512) NOT NULL,
    content TEXT NOT NULL,
    created DATETIME NOT NULL,
    completed INTEGER NOT NULL
);

-- Add an index on the created column.
CREATE INDEX idx_article_created ON articles(created);

INSERT INTO articles (title, content, url, created, completed) VALUES (
    'The Constitution of the United States',
    '<p>We the People of the United States, in Order to form a more perfect Union, establish Justice, insure domestic Tranquility, provide for the common defence, promote the general Welfare, and secure the Blessings of Liberty to ourselves and our Posterity, do ordain and establish this Constitution for the United States of America.</p>',
    'https://www.archives.gov/founding-docs/constitution-transcript',
    '2018-12-23 17:25:22',
    0
);

CREATE TABLE users (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    hashed_password CHAR(60) NOT NULL,
    created DATETIME NOT NULL
);

ALTER TABLE users ADD CONSTRAINT users_uc_email UNIQUE (email);

INSERT INTO users (id, name, email, hashed_password, created) VALUES (
    1,
    'Alice Jones',
    'alice@example.com',
    '$2a$12$NuTjWXm3KKntReFwyBVHyuf/to.HEwTy.eS206TNfkGfr6HzGJSWG',
    '2018-12-23 17:25:22'
);
