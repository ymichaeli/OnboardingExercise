-- name: create-profile-table
CREATE TABLE IF NOT EXISTS profiles
(
    user_id uuid DEFAULT gen_random_uuid(),
    username varchar(255),
    full_name varchar(255),
    bio varchar(255),
    profile_pic_url varchar(255),
    PRIMARY KEY (user_id)
);
