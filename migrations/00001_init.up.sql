BEGIN;
CREATE TABLE public.order_cache (order_uid TEXT PRIMARY KEY, data jsonb NOT NULL);
COMMIT;