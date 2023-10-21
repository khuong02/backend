CREATE OR REPLACE FUNCTION public.uuid_generate_v4()
 RETURNS uuid
 LANGUAGE c
 PARALLEL SAFE STRICT
AS '$libdir/uuid-ossp', $function$uuid_generate_v4$function$
;

CREATE TABLE IF NOT EXISTS users(
    id              UUID         PRIMARY KEY     DEFAULT uuid_generate_v4(),
    created_at      TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at      TIMESTAMP       DEFAULT NULL,

    user_name VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255)
)
