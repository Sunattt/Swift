
CREATE TABLE user_roles(
    id BIGSERIAL ,
    get BOOLEAN DEFAULT false,
    post BOOLEAN DEFAULT false,
    put BOOLEAN DEFAULT false,
    del BOOLEAN DEFAULT false,
    PRIMARY KEY(id)
);

CREATE TABLE users(
    id BIGSERIAL ,
    nik_name VARCHAR(50) UNIQUE NOT NULL,
    name VARCHAR(50) ,
    birthday TIMESTAMP NOT NULL,
    phone VARCHAR(15) not null,
	country varchar(100) not null,
    email VARCHAR(100) UNIQUE NOT NULL ,
    password VARCHAR(50) NOT NULL,
	foto_profile BYTEA NOT NULL,
	bio text ,
	role_id INTEGER DEFAULT 1,
    active BOOLEAN DEFAULT true,
    followers BIGINT DEFAULT 0 ,
    subscriptions BIGINT DEFAULT 0,
    publications BIGINT DEFAULT 0, 	
    PRIMARY KEY(id),
    FOREIGN KEY(role_id) REFERENCES user_roles(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);



CREATE TABLE tokens (
    id BIGSERIAL not null primary key,
    token text not null,
    user_id BIGINT REFERENCES users(id) not NULL,
    expire TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP + interval '10 minutes',
    created_at TIMESTAMP not null DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP not NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    content TEXT,
    count_likes INTEGER,
    count_save INTEGER,
    view_count INTEGER,
    image_url TEXT,
    date_published TIMESTAMP
);


CREATE TABLE follows (
    follower_id INT NOT NULL REFERENCES users(id),
    followed_id INT NOT NULL REFERENCES users(id),
    PRIMARY KEY (follower_id, followed_id)
);

CREATE TABLE comments (
    id SERIAL PRIMARY KEY,
    post_id INTEGER REFERENCES posts(id),
    user_id INTEGER REFERENCES users(id),
    content TEXT,
    date_posted TIMESTAMP
);

CREATE TABLE likes (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    post_id INTEGER REFERENCES posts(id)
);




