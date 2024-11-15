--
-- PostgreSQL database dump
--

-- Dumped from database version 16.4 (Ubuntu 16.4-1.pgdg24.04+2)
-- Dumped by pg_dump version 17.0 (Ubuntu 17.0-1.pgdg24.04+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: transaction_status; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.transaction_status AS ENUM (
    'packing',
    'shipping',
    'done'
);


ALTER TYPE public.transaction_status OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: categories; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.categories (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    description text,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp without time zone
);


ALTER TABLE public.categories OWNER TO postgres;

--
-- Name: categories_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.categories_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.categories_id_seq OWNER TO postgres;

--
-- Name: categories_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.categories_id_seq OWNED BY public.categories.id;


--
-- Name: items; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.items (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    category_id integer,
    image character varying(255),
    price character varying(100),
    transaction_date timestamp without time zone,
    deprisiasi integer,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp without time zone
);


ALTER TABLE public.items OWNER TO postgres;

--
-- Name: items_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.items_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.items_id_seq OWNER TO postgres;

--
-- Name: items_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.items_id_seq OWNED BY public.items.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id integer NOT NULL,
    name character varying NOT NULL,
    email character varying(255) NOT NULL,
    password text NOT NULL,
    token character varying,
    expired timestamp with time zone,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp with time zone
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_id_seq OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: categories id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.categories ALTER COLUMN id SET DEFAULT nextval('public.categories_id_seq'::regclass);


--
-- Name: items id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.items ALTER COLUMN id SET DEFAULT nextval('public.items_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: categories; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.categories (id, name, description, created_at, updated_at, deleted_at) FROM stdin;
1	Electronics	Items related to electronic devices	2024-11-08 21:00:19.620677	2024-11-08 21:00:19.620677	\N
2	Furniture	Various types of furniture	2024-11-08 21:00:19.620677	2024-11-08 21:00:19.620677	\N
3	Books	Books of all genres and authors	2024-11-08 21:00:19.620677	2024-11-08 21:00:19.620677	\N
4	Clothing	Apparel and garments for all genders	2024-11-08 21:00:19.620677	2024-11-08 21:00:19.620677	\N
5	Food	Edible items and groceries	2024-11-08 21:00:19.620677	2024-11-08 21:00:19.620677	\N
6	Beauty	for making beautiful	2024-11-09 14:20:23.484598	2024-11-09 14:20:23.484598	2024-11-09 15:36:59.789495
\.


--
-- Data for Name: items; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.items (id, name, category_id, image, price, transaction_date, deprisiasi, created_at, updated_at, deleted_at) FROM stdin;
2	Kulkas	2	about.png	120003	2024-11-15 00:00:00	10	2024-11-08 23:12:32.253036	2024-11-08 23:12:32.253036	\N
4	Laptor	1	mini-project-inventory.png	120000	2023-11-15 00:00:00	10	2024-11-09 00:17:42.985441	2024-11-09 00:17:42.985441	\N
5	baju dinas	3	mini-project-inventory.png	12000	2024-11-01 00:00:00	10	2024-11-09 00:21:32.892489	2024-11-09 00:21:32.892489	\N
6	sarung wadimor	4	mini-project-inventory.png	1000	2024-11-01 00:00:00	10	2024-11-09 00:22:51.780327	2024-11-09 00:22:51.780327	\N
7	Kulkas dua pintu	3	uploads/about.png	10000	2024-02-12 00:00:00	10	2024-11-09 00:23:48.320427	2024-11-09 00:23:48.320427	\N
8	Handphone	1	uploads/mini-project-inventory.png	100000	2024-11-01 00:00:00	0	2024-11-09 09:29:18.689514	2024-11-09 09:29:18.689514	2024-11-09 10:19:14.057556
9	Book	1	uploads/mini-project-inventory.png	1000	2024-11-01 00:00:00	0	2024-11-09 09:30:26.283743	2024-11-09 09:30:26.283743	2024-11-09 16:21:06.045166
3	Kulkas	3		10000	2024-02-12 00:00:00	10	2024-11-08 23:25:02.641306	2024-11-08 23:25:02.641306	2024-11-11 21:56:20.455156
1	Kulkas	3		0	0001-01-01 00:00:00	10	2024-11-08 22:36:13.549066	2024-11-08 22:36:13.549066	2024-11-11 21:57:50.177512
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, name, email, password, token, expired, created_at, updated_at, deleted_at) FROM stdin;
1	Fikri	fikri@example.com	fikri123	303b655f-537d-44e4-8059-19ec08cbc454	2024-11-12 21:00:07.444637+07	2024-11-12 14:09:05.29726+07	2024-11-12 14:09:05.29726+07	\N
2	Fikri12	fikri1@example.com	fikri1234	3dce4d22-b788-4a24-b171-7b67a7160608	2024-11-12 21:17:24.598364+07	2024-11-12 20:08:50.005+07	2024-11-12 20:08:50.005+07	\N
\.


--
-- Name: categories_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.categories_id_seq', 6, true);


--
-- Name: items_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.items_id_seq', 9, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 2, true);


--
-- Name: categories categories_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.categories
    ADD CONSTRAINT categories_pkey PRIMARY KEY (id);


--
-- Name: items items_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT items_pkey PRIMARY KEY (id);


--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

