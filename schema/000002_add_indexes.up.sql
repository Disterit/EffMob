CREATE INDEX IF NOT EXISTS idx_songs_group_id ON songs(group_id);

CREATE INDEX IF NOT EXISTS idx_groups_group_name ON groups(group_name);

CREATE INDEX IF NOT EXISTS idx_songs_group_id_release_date ON songs(group_id, release_date);
