--
-- PostgreSQL database dump
--

-- Dumped from database version 17.0 (Homebrew)
-- Dumped by pg_dump version 17.0 (Homebrew)

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
-- Name: citext; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS citext WITH SCHEMA public;


--
-- Name: EXTENSION citext; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION citext IS 'data type for case-insensitive character strings';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: codeclips; Type: TABLE; Schema: public; Owner: rifat
--

CREATE TABLE public.codeclips (
    id integer NOT NULL,
    title character varying(255) NOT NULL,
    language character varying(100) NOT NULL,
    content text NOT NULL
);


ALTER TABLE public.codeclips OWNER TO rifat;

--
-- Name: codeclips_id_seq; Type: SEQUENCE; Schema: public; Owner: rifat
--

CREATE SEQUENCE public.codeclips_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.codeclips_id_seq OWNER TO rifat;

--
-- Name: codeclips_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: rifat
--

ALTER SEQUENCE public.codeclips_id_seq OWNED BY public.codeclips.id;


--
-- Name: sessions; Type: TABLE; Schema: public; Owner: codeclips
--

CREATE TABLE public.sessions (
    token text NOT NULL,
    data bytea NOT NULL,
    expiry timestamp with time zone NOT NULL
);


ALTER TABLE public.sessions OWNER TO codeclips;

--
-- Name: users; Type: TABLE; Schema: public; Owner: codeclips
--

CREATE TABLE public.users (
    id integer NOT NULL,
    name text NOT NULL,
    email text NOT NULL,
    hashed_password character varying(255) NOT NULL,
    created timestamp with time zone DEFAULT now()
);


ALTER TABLE public.users OWNER TO codeclips;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: codeclips
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_id_seq OWNER TO codeclips;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: codeclips
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: codeclips id; Type: DEFAULT; Schema: public; Owner: rifat
--

ALTER TABLE ONLY public.codeclips ALTER COLUMN id SET DEFAULT nextval('public.codeclips_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: codeclips
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: codeclips; Type: TABLE DATA; Schema: public; Owner: rifat
--

COPY public.codeclips (id, title, language, content) FROM stdin;
4	basics of gov2	go	package main\r\n\r\nimport (\r\n\t"html/template"\r\n\t"net/http"\r\n)\r\n\r\ntype CreateCodeClips struct{\r\n  Title string `form:"title"`\r\n\tLanguage string `form:"language"`\r\n\tContent template.HTML `form:"content"`\r\n\r\n\r\n\r\n}\r\n\r\n\r\nfunc (app *App) home(w http.ResponseWriter, r *http.Request){\r\n\r\n  data:= app.newTemplateData(r)\r\n\r\n\tapp.render(w,r,http.StatusOK,"home.tmpl.html",data)\r\n\r\n\t\r\n}\r\n\r\nfunc (app *App) codeClipsPost(w http.ResponseWriter, r *http.Request){\r\n\tvar form CreateCodeClips\r\n\t// decode form and update it to the struct\r\n\terr:=app.decodePostForm(w,r,&form)\r\n\r\n\tif err!=nil{\r\n\t\tapp.clientError(w,http.StatusBadRequest)\r\n\t\treturn\r\n\t}\r\n\r\n\terr = app.codeClips.Insert(form.Title,form.Language,form.Content)\r\n\r\n\tif err!=nil{\r\n\t\t app.serverError(w,r,err)\r\n        return\r\n\t}\r\n\r\n}\r\n\r\n\r\nfunc (app *App) clips(w http.ResponseWriter, r *http.Request){\r\n\tdata:= app.newTemplateData(r)\r\n\r\n\tapp.render(w,r,http.StatusOK,"clips.tmpl.html",data)\r\n\r\n\r\n}
\.


--
-- Data for Name: sessions; Type: TABLE DATA; Schema: public; Owner: codeclips
--

COPY public.sessions (token, data, expiry) FROM stdin;
n7IJR8t-Usftz4-YhyxAMAv_FUL9Lc0qfQlTxbwU9Sk	\\x257f030102ff800001020108446561646c696e6501ff8200010656616c75657301ff8400000010ff810501010454696d6501ff8200000027ff83040101176d61705b737472696e675d696e74657266616365207b7d01ff8400010c0110000016ff80010f010000000edeadc5c2274cf190ffff010000	2024-10-25 10:00:50.659354-07
4HsdaDaFi99hDpS2YFUvbO7Yo_coHhEPTdSFugbfk8Y	\\x257f030102ff800001020108446561646c696e6501ff8200010656616c75657301ff8400000010ff810501010454696d6501ff8200000027ff83040101176d61705b737472696e675d696e74657266616365207b7d01ff8400010c0110000016ff80010f010000000edead9aa311a424d0ffff010000	2024-10-25 06:56:51.29597-07
WBS-FkrwZkVc_teKWIA7WFk92jWm811EznaUxkk_DFM	\\x257f030102ff800001020108446561646c696e6501ff8200010656616c75657301ff8400000010ff810501010454696d6501ff8200000027ff83040101176d61705b737472696e675d696e74657266616365207b7d01ff8400010c0110000016ff80010f010000000edeada9453198f5e0ffff010000	2024-10-25 07:59:17.832108-07
OhpuKznVc73Wk8eoUjAGCNtwgQ5kPrd8EYVbGX9-FTk	\\x257f030102ff800001020108446561646c696e6501ff8200010656616c75657301ff8400000010ff810501010454696d6501ff8200000027ff83040101176d61705b737472696e675d696e74657266616365207b7d01ff8400010c0110000016ff80010f010000000edeada9a409ce48d0ffff010000	2024-10-25 08:00:52.164514-07
Qoca_lm1ybhHpcdhRVtg3x-j2FVy6oCt7QNZecOmu5U	\\x257f030102ff800001020108446561646c696e6501ff8200010656616c75657301ff8400000010ff810501010454696d6501ff8200000027ff83040101176d61705b737472696e675d696e74657266616365207b7d01ff8400010c0110000016ff80010f010000000edeadc19720b273f0ffff010000	2024-10-25 09:43:03.548566-07
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: codeclips
--

COPY public.users (id, name, email, hashed_password, created) FROM stdin;
\.


--
-- Name: codeclips_id_seq; Type: SEQUENCE SET; Schema: public; Owner: rifat
--

SELECT pg_catalog.setval('public.codeclips_id_seq', 27, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: codeclips
--

SELECT pg_catalog.setval('public.users_id_seq', 1, false);


--
-- Name: codeclips codeclips_pkey; Type: CONSTRAINT; Schema: public; Owner: rifat
--

ALTER TABLE ONLY public.codeclips
    ADD CONSTRAINT codeclips_pkey PRIMARY KEY (id);


--
-- Name: sessions sessions_pkey; Type: CONSTRAINT; Schema: public; Owner: codeclips
--

ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT sessions_pkey PRIMARY KEY (token);


--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: codeclips
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: codeclips
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: sessions_expiry_idx; Type: INDEX; Schema: public; Owner: codeclips
--

CREATE INDEX sessions_expiry_idx ON public.sessions USING btree (expiry);


--
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: pg_database_owner
--

GRANT ALL ON SCHEMA public TO codeclips;


--
-- Name: TABLE codeclips; Type: ACL; Schema: public; Owner: rifat
--

GRANT ALL ON TABLE public.codeclips TO codeclips;


--
-- Name: SEQUENCE codeclips_id_seq; Type: ACL; Schema: public; Owner: rifat
--

GRANT ALL ON SEQUENCE public.codeclips_id_seq TO codeclips;


--
-- PostgreSQL database dump complete
--

