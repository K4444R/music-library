
CREATE SEQUENCE IF NOT EXISTS public.songs_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE IF NOT EXISTS public.songs (
    id integer NOT NULL DEFAULT nextval('public.songs_id_seq'::regclass),
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    "group" text,
    song text,
    release_date text,
    text text,
    link text,
    CONSTRAINT songs_pkey PRIMARY KEY (id)
);

CREATE INDEX IF NOT EXISTS idx_songs_deleted_at ON public.songs USING btree (deleted_at);