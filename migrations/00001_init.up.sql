BEGIN;
SET statement_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = ON;
SET check_function_bodies = FALSE;
SET client_min_messages = WARNING;
SET search_path = public,
    extensions;
SET default_tablespace = '';
SET default_with_oids = FALSE;
-- TABLES --
CREATE TABLE public.customer (customer_id TEXT PRIMARY KEY);
CREATE TABLE public.order (
    order_uid TEXT PRIMARY KEY,
    track_number TEXT,
    entry TEXT,
    delivery UUID,
    locale TEXT,
    internal_signature TEXT,
    customer_id TEXT REFERENCES public.customer (customer_id),
    delivery_service TEXT,
    shardkey TEXT,
    sm_id BIGINT,
    date_created TIMESTAMPTZ,
    oof_shard TEXT
);
CREATE TABLE public.delivery (
    order_uid TEXT REFERENCES public.order (order_uid),
    name TEXT,
    phone TEXT,
    zip TEXT,
    city TEXT,
    address TEXT,
    region TEXT,
    email TEXT
);
CREATE TABLE public.payment (
    order_uid TEXT REFERENCES public.order (order_uid),
    transaction TEXT,
    request_id TEXT,
    currency TEXT,
    provider TEXT,
    amount BIGINT,
    payment_dt BIGINT,
    bank TEXT,
    delivery_cost BIGINT,
    goods_total BIGINT,
    custom_fee BIGINT
);
CREATE TABLE public.item (
    chrt_id BIGINT,
    track_number TEXT,
    price BIGINT,
    rid TEXT,
    name TEXT,
    sale INT,
    size INT,
    total_price BIGINT,
    nm_id BIGINT,
    brand TEXT,
    status INT
);
COMMIT;