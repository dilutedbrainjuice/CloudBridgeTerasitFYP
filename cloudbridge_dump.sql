--
-- PostgreSQL database dump
--

-- Dumped from database version 14.11 (Ubuntu 14.11-0ubuntu0.22.04.1)
-- Dumped by pg_dump version 14.11 (Ubuntu 14.11-0ubuntu0.22.04.1)

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
-- Name: message; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.message (
    id integer NOT NULL,
    sender character varying,
    content text NOT NULL,
    "timestamp" timestamp without time zone NOT NULL,
    chatroom character varying NOT NULL
);


ALTER TABLE public.message OWNER TO postgres;

--
-- Name: message_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.message_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.message_id_seq OWNER TO postgres;

--
-- Name: message_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.message_id_seq OWNED BY public.message.id;


--
-- Name: user; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."user" (
    "ID" integer NOT NULL,
    createdat character varying NOT NULL,
    password text NOT NULL,
    username text,
    email character varying,
    profilepicurl character varying,
    city character varying,
    pc_specs character varying,
    description character varying,
    cloud_service character varying,
    isprovider boolean,
    latitude double precision,
    longitude double precision
);


ALTER TABLE public."user" OWNER TO postgres;

--
-- Name: user_ID_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."user_ID_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."user_ID_seq" OWNER TO postgres;

--
-- Name: user_ID_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."user_ID_seq" OWNED BY public."user"."ID";


--
-- Name: message id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.message ALTER COLUMN id SET DEFAULT nextval('public.message_id_seq'::regclass);


--
-- Name: user ID; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."user" ALTER COLUMN "ID" SET DEFAULT nextval('public."user_ID_seq"'::regclass);


--
-- Data for Name: message; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.message (id, sender, content, "timestamp", chatroom) FROM stdin;
1	dashboard2	hello im coming in database	2024-04-22 12:27:07.96337	private_44_45
2	dashboard2	dasd	2024-04-22 13:01:59.928552	private_44_45
3	dashboard2	dasd	2024-04-22 13:02:01.418987	private_44_45
4	dashboard2	asdsad	2024-04-22 13:05:00.738661	private_44_45
5	dashboard2	asdasd	2024-04-22 13:06:41.17072	private_44_45
6	dashboard2	asdasdasdsad	2024-04-22 13:06:43.319404	private_44_45
7	dashboard2	asdasd	2024-04-22 13:16:04.825225	private_44_45
8	dashboard2	asdasdasdsad	2024-04-22 13:16:06.462883	private_44_45
9	dashboard2	yo whatsuo	2024-04-22 16:21:04.559242	private_44_45
10	dashboard2	yo whatsuowhy is the websocket closed	2024-04-22 16:21:11.2384	private_44_45
11	dashboard2	yo why is it closed	2024-04-22 16:21:28.966668	private_44_45
12	pleasework	asdsad	2024-04-24 14:08:55.171912	private_44_46
13	pleasework	asdsad	2024-04-24 14:09:07.564776	private_44_46
14	pleasework	sadsad	2024-04-24 14:09:09.735271	private_44_46
15	pleasework	asd	2024-04-24 14:09:11.249152	private_44_46
16	pleasework	sad	2024-04-24 14:09:12.686729	private_44_46
17	pleasework	sad	2024-04-24 14:09:13.376435	private_44_46
18	pleasework	sad	2024-04-24 14:09:14.703176	private_44_46
19	pleasework	dasds	2024-04-24 14:11:24.948721	private_44_46
20	pleasework	saddasds	2024-04-24 14:11:26.493102	private_44_46
21	pleasework	sdasd	2024-04-24 14:12:07.163275	private_44_46
22	pleasework	asdsad	2024-04-24 14:14:17.401482	private_44_46
23	pleasework	asdsad	2024-04-24 14:16:13.761787	private_44_46
24	pleasework	asdsad	2024-04-24 14:16:15.090162	private_44_46
25	pleasework	asdsad	2024-04-24 14:16:15.845596	private_44_46
26	pleasework	asdsad	2024-04-24 14:16:16.757544	private_44_46
27	pleasework	sadsad	2024-04-24 14:16:53.590146	private_44_46
28	dashboard2	asdsad	2024-04-24 14:19:29.709682	private_44_46
29	dashboard2	zxc	2024-04-24 14:19:33.309226	private_44_46
30	pleasework	asdsad	2024-04-24 14:31:14.562772	private_44_46
31	dashboard2	halou bb	2024-04-26 01:43:03.584852	private_44_45
33	dashboard2	testing again	2024-05-02 00:25:47.990202	private_44_45
34	dashboard2	hi bro	2024-05-02 02:02:45.301617	private_44_46
\.


--
-- Data for Name: user; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."user" ("ID", createdat, password, username, email, profilepicurl, city, pc_specs, description, cloud_service, isprovider, latitude, longitude) FROM stdin;
48	2024-04-26 01:56:10	$2a$10$3MPKB5DG4vuoi91KKt6JFeWiD4Az3GPT8q/qI48u/YhxlA4HtyfdG	DarrenCheong	\N	/uploads/DarrenCheong_profilepic.png	\N	i9 super rtx	show user picture	Moonlight	f	2.9459475	101.8746084
46	2024-04-21 11:10:14	$2a$10$oLM/gghfVH7dnMHs7BMtDOzSnCaD6H4oY4l0SZ120XurnMOAsoY0i	pleasework	\N	/uploads/pleasework_profilepic.png	\N	i7 bollocks	what to fking do	Parsec	f	3.1671	101.6708
45	2024-04-14 10:00:00	$2a$10$VFTaoA0UgpCT7iFVGLTBee5iFaBnYCiETDTJnDQ4okp.5K1TGKG2S	testinginsertdb	new_usercloudbrdi@egmail.com	/uploads/Firefox_wallpaper.png	New City	new_pc_specs	New Description	Parsec	t	2.9516	101.843
44	2024-04-13 04:39:18	$2a$10$VFTaoA0UgpCT7iFVGLTBee5iFaBnYCiETDTJnDQ4okp.5K1TGKG2S	dashboard2	\N	/uploads/dashboard2_profilepic.png	\N	asdasd	asdasd	SteamLink	f	3.0408704	101.7643008
\.


--
-- Name: message_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.message_id_seq', 34, true);


--
-- Name: user_ID_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."user_ID_seq"', 54, true);


--
-- Name: message message_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.message
    ADD CONSTRAINT message_pkey PRIMARY KEY (id);


--
-- Name: user unique_username; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT unique_username UNIQUE (username);


--
-- Name: user user_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT user_pkey PRIMARY KEY ("ID");


--
-- Name: message message_sender_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.message
    ADD CONSTRAINT message_sender_fkey FOREIGN KEY (sender) REFERENCES public."user"(username);


--
-- PostgreSQL database dump complete
--

