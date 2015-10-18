--
-- PostgreSQL database dump
--

SET statement_timeout = 0;
SET lock_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET search_path = public, pg_catalog;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: book; Type: TABLE; Schema: public; Owner: postgres; Tablespace: 
--

CREATE TABLE book (
    id integer NOT NULL,
    name character varying(1500) NOT NULL,
    author character varying(1000) NOT NULL
);


ALTER TABLE book OWNER TO postgres;

--
-- Name: book_library; Type: TABLE; Schema: public; Owner: postgres; Tablespace: 
--

CREATE TABLE book_library (
    book_id integer NOT NULL,
    library_id integer NOT NULL
);


ALTER TABLE book_library OWNER TO postgres;

--
-- Name: book_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE book_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE book_id_seq OWNER TO postgres;

--
-- Name: book_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE book_id_seq OWNED BY book.id;


--
-- Name: library; Type: TABLE; Schema: public; Owner: postgres; Tablespace: 
--

CREATE TABLE library (
    id integer NOT NULL,
    name character varying(1500) NOT NULL
);


ALTER TABLE library OWNER TO postgres;

--
-- Name: library_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE library_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE library_id_seq OWNER TO postgres;

--
-- Name: library_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE library_id_seq OWNED BY library.id;


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY book ALTER COLUMN id SET DEFAULT nextval('book_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY library ALTER COLUMN id SET DEFAULT nextval('library_id_seq'::regclass);


--
-- Data for Name: book_library; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY book_library (book_id, library_id) FROM stdin;
\.


--
-- Name: book_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('book_id_seq', 1, false);


--
-- Data for Name: library; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY library (id, name) FROM stdin;
\.


--
-- Name: library_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('library_id_seq', 1, false);


--
-- Name: book_library_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres; Tablespace: 
--

ALTER TABLE ONLY book_library
    ADD CONSTRAINT book_library_pkey PRIMARY KEY (book_id, library_id);


--
-- Name: book_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres; Tablespace: 
--

ALTER TABLE ONLY book
    ADD CONSTRAINT book_pkey PRIMARY KEY (id);


--
-- Name: library_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres; Tablespace: 
--

ALTER TABLE ONLY library
    ADD CONSTRAINT library_pkey PRIMARY KEY (id);


--
-- Name: book_library_book_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY book_library
    ADD CONSTRAINT book_library_book_id_fkey FOREIGN KEY (book_id) REFERENCES book(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: book_library_library_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY book_library
    ADD CONSTRAINT book_library_library_id_fkey FOREIGN KEY (library_id) REFERENCES library(id) ON UPDATE CASCADE;


--
-- Name: public; Type: ACL; Schema: -; Owner: postgres
--

REVOKE ALL ON SCHEMA public FROM PUBLIC;
REVOKE ALL ON SCHEMA public FROM postgres;
GRANT ALL ON SCHEMA public TO postgres;
GRANT ALL ON SCHEMA public TO PUBLIC;


--
-- PostgreSQL database dump complete
--

insert into book(id, name, author) values (1,'Book1','Author1'), (2,'Book2','Author2');
insert into library(id, name) values (1, 'Library1');
insert into book_library(book_id, library_id) values(1,1), (2,1);