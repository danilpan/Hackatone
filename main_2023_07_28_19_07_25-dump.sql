--
-- PostgreSQL database dump
--

-- Dumped from database version 15.2 (Debian 15.2-1.pgdg110+1)
-- Dumped by pg_dump version 15.3

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: main; Type: DATABASE; Schema: -; Owner: admin
--

CREATE DATABASE main WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'en_US.utf8';


ALTER DATABASE main OWNER TO admin;

\connect main

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: public; Type: SCHEMA; Schema: -; Owner: pg_database_owner
--

CREATE SCHEMA public;


ALTER SCHEMA public OWNER TO pg_database_owner;

--
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: pg_database_owner
--

COMMENT ON SCHEMA public IS 'standard public schema';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: establishment_types; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.establishment_types (
    id integer NOT NULL,
    name character varying NOT NULL
);


ALTER TABLE public.establishment_types OWNER TO admin;

--
-- Name: establishment_types_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.establishment_types_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.establishment_types_id_seq OWNER TO admin;

--
-- Name: establishment_types_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.establishment_types_id_seq OWNED BY public.establishment_types.id;


--
-- Name: establishments; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.establishments (
    id integer NOT NULL,
    name character varying NOT NULL,
    address character varying NOT NULL,
    type_id integer NOT NULL,
    average_check integer NOT NULL,
    rating integer NOT NULL,
    images_urls character varying[] NOT NULL,
    CONSTRAINT check_name CHECK (((1 <= rating) AND (rating <= 5)))
);


ALTER TABLE public.establishments OWNER TO admin;

--
-- Name: establishments_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.establishments_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.establishments_id_seq OWNER TO admin;

--
-- Name: establishments_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.establishments_id_seq OWNED BY public.establishments.id;


--
-- Name: reservations; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.reservations (
    id integer NOT NULL,
    table_id integer NOT NULL,
    user_iin character(12) NOT NULL,
    time_from timestamp without time zone NOT NULL,
    time_to timestamp without time zone NOT NULL,
    confirmed boolean,
    persons integer NOT NULL,
    CONSTRAINT check_name CHECK ((persons >= 1))
);


ALTER TABLE public.reservations OWNER TO admin;

--
-- Name: reservations_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.reservations_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.reservations_id_seq OWNER TO admin;

--
-- Name: reservations_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.reservations_id_seq OWNED BY public.reservations.id;


--
-- Name: tables; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.tables (
    id integer NOT NULL,
    establishment_id integer NOT NULL,
    number integer NOT NULL,
    persons integer NOT NULL,
    CONSTRAINT persons_check CHECK ((persons >= 1))
);


ALTER TABLE public.tables OWNER TO admin;

--
-- Name: tables_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.tables_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.tables_id_seq OWNER TO admin;

--
-- Name: tables_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.tables_id_seq OWNED BY public.tables.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.users (
    iin character(12) NOT NULL,
    full_name character varying NOT NULL
);


ALTER TABLE public.users OWNER TO admin;

--
-- Name: establishment_types id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.establishment_types ALTER COLUMN id SET DEFAULT nextval('public.establishment_types_id_seq'::regclass);


--
-- Name: establishments id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.establishments ALTER COLUMN id SET DEFAULT nextval('public.establishments_id_seq'::regclass);


--
-- Name: reservations id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.reservations ALTER COLUMN id SET DEFAULT nextval('public.reservations_id_seq'::regclass);


--
-- Name: tables id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.tables ALTER COLUMN id SET DEFAULT nextval('public.tables_id_seq'::regclass);


--
-- Data for Name: establishment_types; Type: TABLE DATA; Schema: public; Owner: admin
--

INSERT INTO public.establishment_types (id, name) VALUES (2, 'Ресторан');
INSERT INTO public.establishment_types (id, name) VALUES (3, 'Кафе');
INSERT INTO public.establishment_types (id, name) VALUES (1, 'Бар');


--
-- Data for Name: establishments; Type: TABLE DATA; Schema: public; Owner: admin
--

INSERT INTO public.establishments (id, name, address, type_id, average_check, rating, images_urls) VALUES (3, 'Жиробас', 'ул. Фурманова, 220/3', 1, 4000, 4, '{https://restolife.kz/upload/information_system_6/2/2/1/item_22154/information_items_property_25365.jpg}');
INSERT INTO public.establishments (id, name, address, type_id, average_check, rating, images_urls) VALUES (4, 'Дареджани', 'ул. Кунаева', 2, 6000, 5, '{https://lh3.googleusercontent.com/p/AF1QipNffmZyZZQV5uegAQbpGvLsj8ts2aIxwKGyM1N1=s680-w680-h510}');


--
-- Data for Name: reservations; Type: TABLE DATA; Schema: public; Owner: admin
--



--
-- Data for Name: tables; Type: TABLE DATA; Schema: public; Owner: admin
--

INSERT INTO public.tables (id, establishment_id, number, persons) VALUES (1, 3, 1, 4);
INSERT INTO public.tables (id, establishment_id, number, persons) VALUES (2, 3, 2, 6);
INSERT INTO public.tables (id, establishment_id, number, persons) VALUES (3, 3, 3, 6);


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: admin
--



--
-- Name: establishment_types_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.establishment_types_id_seq', 6, true);


--
-- Name: establishments_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.establishments_id_seq', 4, true);


--
-- Name: reservations_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.reservations_id_seq', 1, false);


--
-- Name: tables_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.tables_id_seq', 3, true);


--
-- Name: establishment_types establishment_types_pk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.establishment_types
    ADD CONSTRAINT establishment_types_pk PRIMARY KEY (id);


--
-- Name: establishments establishments_pk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.establishments
    ADD CONSTRAINT establishments_pk PRIMARY KEY (id);


--
-- Name: reservations reservations_pk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.reservations
    ADD CONSTRAINT reservations_pk PRIMARY KEY (id);


--
-- Name: tables tables_establishment_id_number; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.tables
    ADD CONSTRAINT tables_establishment_id_number UNIQUE (establishment_id, number);


--
-- Name: tables tables_pk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.tables
    ADD CONSTRAINT tables_pk PRIMARY KEY (id);


--
-- Name: users users_pk; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pk PRIMARY KEY (iin);


--
-- Name: establishments establishments_establishment_types_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.establishments
    ADD CONSTRAINT establishments_establishment_types_id_fk FOREIGN KEY (type_id) REFERENCES public.establishment_types(id);


--
-- Name: reservations reservations_tables_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.reservations
    ADD CONSTRAINT reservations_tables_id_fk FOREIGN KEY (table_id) REFERENCES public.tables(id);


--
-- Name: reservations reservations_users_iin_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.reservations
    ADD CONSTRAINT reservations_users_iin_fk FOREIGN KEY (user_iin) REFERENCES public.users(iin);


--
-- Name: tables tables_establishments_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.tables
    ADD CONSTRAINT tables_establishments_id_fk FOREIGN KEY (establishment_id) REFERENCES public.establishments(id);


--
-- PostgreSQL database dump complete
--

