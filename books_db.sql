--
-- PostgreSQL database dump
--

-- Dumped from database version 14.3
-- Dumped by pg_dump version 14.3

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
-- Name: author; Type: TABLE; Schema: public; Owner: emanuelinacio
--

CREATE TABLE public.author (
    author_id integer NOT NULL,
    first_name character varying(100) NOT NULL,
    last_name character varying(100) NOT NULL
);


ALTER TABLE public.author OWNER TO emanuelinacio;

--
-- Name: authors_author_id_seq; Type: SEQUENCE; Schema: public; Owner: emanuelinacio
--

CREATE SEQUENCE public.authors_author_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.authors_author_id_seq OWNER TO emanuelinacio;

--
-- Name: authors_author_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: emanuelinacio
--

ALTER SEQUENCE public.authors_author_id_seq OWNED BY public.author.author_id;


--
-- Name: book; Type: TABLE; Schema: public; Owner: emanuelinacio
--

CREATE TABLE public.book (
    book_id integer NOT NULL,
    title character varying(100) NOT NULL,
    published_year integer,
    author_id integer
);


ALTER TABLE public.book OWNER TO emanuelinacio;

--
-- Name: books_book_id_seq; Type: SEQUENCE; Schema: public; Owner: emanuelinacio
--

CREATE SEQUENCE public.books_book_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.books_book_id_seq OWNER TO emanuelinacio;

--
-- Name: books_book_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: emanuelinacio
--

ALTER SEQUENCE public.books_book_id_seq OWNED BY public.book.book_id;


--
-- Name: review; Type: TABLE; Schema: public; Owner: emanuelinacio
--

CREATE TABLE public.review (
    review_id integer NOT NULL,
    rating numeric(2,1) NOT NULL,
    book_id integer
);


ALTER TABLE public.review OWNER TO emanuelinacio;

--
-- Name: reviews_review_id_seq; Type: SEQUENCE; Schema: public; Owner: emanuelinacio
--

CREATE SEQUENCE public.reviews_review_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.reviews_review_id_seq OWNER TO emanuelinacio;

--
-- Name: reviews_review_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: emanuelinacio
--

ALTER SEQUENCE public.reviews_review_id_seq OWNED BY public.review.review_id;


--
-- Name: author author_id; Type: DEFAULT; Schema: public; Owner: emanuelinacio
--

ALTER TABLE ONLY public.author ALTER COLUMN author_id SET DEFAULT nextval('public.authors_author_id_seq'::regclass);


--
-- Name: book book_id; Type: DEFAULT; Schema: public; Owner: emanuelinacio
--

ALTER TABLE ONLY public.book ALTER COLUMN book_id SET DEFAULT nextval('public.books_book_id_seq'::regclass);


--
-- Name: review review_id; Type: DEFAULT; Schema: public; Owner: emanuelinacio
--

ALTER TABLE ONLY public.review ALTER COLUMN review_id SET DEFAULT nextval('public.reviews_review_id_seq'::regclass);


--
-- Data for Name: author; Type: TABLE DATA; Schema: public; Owner: emanuelinacio
--

COPY public.author (author_id, first_name, last_name) FROM stdin;
1	Tamsyn	Muir
2	Ann	Leckie
3	Zen	Cho
\.


--
-- Data for Name: book; Type: TABLE DATA; Schema: public; Owner: emanuelinacio
--

COPY public.book (book_id, title, published_year, author_id) FROM stdin;
1	Gideon the Ninth	2019	1
2	Ancillary Justice	2013	2
3	Black Water Sister	2021	3
\.


--
-- Data for Name: review; Type: TABLE DATA; Schema: public; Owner: emanuelinacio
--

COPY public.review (review_id, rating, book_id) FROM stdin;
1	5.0	1
2	4.5	2
3	4.5	3
\.


--
-- Name: authors_author_id_seq; Type: SEQUENCE SET; Schema: public; Owner: emanuelinacio
--

SELECT pg_catalog.setval('public.authors_author_id_seq', 3, true);


--
-- Name: books_book_id_seq; Type: SEQUENCE SET; Schema: public; Owner: emanuelinacio
--

SELECT pg_catalog.setval('public.books_book_id_seq', 12, true);


--
-- Name: reviews_review_id_seq; Type: SEQUENCE SET; Schema: public; Owner: emanuelinacio
--

SELECT pg_catalog.setval('public.reviews_review_id_seq', 3, true);


--
-- Name: author authors_pkey; Type: CONSTRAINT; Schema: public; Owner: emanuelinacio
--

ALTER TABLE ONLY public.author
    ADD CONSTRAINT authors_pkey PRIMARY KEY (author_id);


--
-- Name: book books_pkey; Type: CONSTRAINT; Schema: public; Owner: emanuelinacio
--

ALTER TABLE ONLY public.book
    ADD CONSTRAINT books_pkey PRIMARY KEY (book_id);


--
-- Name: review reviews_pkey; Type: CONSTRAINT; Schema: public; Owner: emanuelinacio
--

ALTER TABLE ONLY public.review
    ADD CONSTRAINT reviews_pkey PRIMARY KEY (review_id);


--
-- Name: book books_author_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: emanuelinacio
--

ALTER TABLE ONLY public.book
    ADD CONSTRAINT books_author_id_fkey FOREIGN KEY (author_id) REFERENCES public.author(author_id);


--
-- Name: review reviews_book_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: emanuelinacio
--

ALTER TABLE ONLY public.review
    ADD CONSTRAINT reviews_book_id_fkey FOREIGN KEY (book_id) REFERENCES public.book(book_id);


--
-- PostgreSQL database dump complete
--

