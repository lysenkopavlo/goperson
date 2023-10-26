--
-- PostgreSQL database dump
--

-- Dumped from database version 15.4 (Ubuntu 15.4-2.pgdg22.04+1)
-- Dumped by pg_dump version 15.4 (Ubuntu 15.4-2.pgdg22.04+1)

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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: persons; Type: TABLE; Schema: public; Owner: pablo
--

CREATE TABLE public.persons (
    id integer NOT NULL,
    name character varying(255) DEFAULT ''::character varying NOT NULL,
    patronymic character varying(255) DEFAULT ''::character varying NOT NULL,
    surname character varying(255) DEFAULT ''::character varying NOT NULL,
    age integer NOT NULL,
    gender character varying(255) DEFAULT ''::character varying NOT NULL,
    country_id character varying(255) DEFAULT ''::character varying NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.persons OWNER TO pablo;

--
-- Name: persons_id_seq; Type: SEQUENCE; Schema: public; Owner: pablo
--

CREATE SEQUENCE public.persons_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.persons_id_seq OWNER TO pablo;

--
-- Name: persons_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: pablo
--

ALTER SEQUENCE public.persons_id_seq OWNED BY public.persons.id;


--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: pablo
--

CREATE TABLE public.schema_migration (
    version character varying(14) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO pablo;

--
-- Name: persons id; Type: DEFAULT; Schema: public; Owner: pablo
--

ALTER TABLE ONLY public.persons ALTER COLUMN id SET DEFAULT nextval('public.persons_id_seq'::regclass);


--
-- Name: persons persons_pkey; Type: CONSTRAINT; Schema: public; Owner: pablo
--

ALTER TABLE ONLY public.persons
    ADD CONSTRAINT persons_pkey PRIMARY KEY (id);


--
-- Name: schema_migration schema_migration_pkey; Type: CONSTRAINT; Schema: public; Owner: pablo
--

ALTER TABLE ONLY public.schema_migration
    ADD CONSTRAINT schema_migration_pkey PRIMARY KEY (version);


--
-- Name: persons_name_surname_country_id_idx; Type: INDEX; Schema: public; Owner: pablo
--

CREATE INDEX persons_name_surname_country_id_idx ON public.persons USING btree (name, surname, country_id);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: pablo
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- PostgreSQL database dump complete
--

