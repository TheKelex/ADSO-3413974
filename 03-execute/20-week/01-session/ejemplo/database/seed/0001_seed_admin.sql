-- Bootstrap administrator. Password supplied at seed time via the
-- :admin_username / :admin_password psql variables (from ADMIN_USERNAME /
-- ADMIN_PASSWORD in .env.example) and stored only as a bcrypt hash.
INSERT INTO "user" (username, password_hash, role)
VALUES (
    :'admin_username',
    crypt(:'admin_password', gen_salt('bf')),
    'admin'
)
ON CONFLICT (username) DO NOTHING;
