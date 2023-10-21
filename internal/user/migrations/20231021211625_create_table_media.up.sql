CREATE OR REPLACE FUNCTION public.uuid_generate_v4()
 RETURNS uuid
 LANGUAGE c
 PARALLEL SAFE STRICT
AS '$libdir/uuid-ossp', $function$uuid_generate_v4$function$
;

CREATE TABLE IF NOT EXISTS media(
    id                      UUID          PRIMARY KEY,
    created_at              TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at              TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at              TIMESTAMP       DEFAULT NULL,
    name                    VARCHAR(100)    NOT NULL,
    path                    VARCHAR(200)    NOT NULL,
    upload_by               UUID,

    CONSTRAINT media_unique UNIQUE (name,path)
)
