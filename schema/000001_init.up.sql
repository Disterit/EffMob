CREATE TABLE IF NOT EXISTS groups (
    id serial primary key ,
    group_name varchar(255) unique not null
);


CREATE TABLE IF NOT EXISTS songs (
    id serial primary key,
    group_id int references groups(id) on delete cascade not null ,
    song_name varchar(255) not null ,
    text_song text ,
    link varchar(255),
    release_date date,
    unique (group_id, song_name)
);



CREATE INDEX idx_songs_group_id ON songs(group_id);
CREATE INDEX idx_groups_group_name ON groups(group_name);
CREATE INDEX idx_songs_group_id_release_date ON songs(group_id, release_date);
