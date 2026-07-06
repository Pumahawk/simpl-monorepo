--
-- PostgreSQL database dump
--

-- Dumped from database version 16.2
-- Dumped by pg_dump version 17.0

-- Started on 2025-07-24 11:58:46

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
-- TOC entry 4 (class 2615 OID 2200)
-- Name: public; Type: SCHEMA; Schema: -; Owner: pg_database_owner
--

CREATE SCHEMA public;


ALTER SCHEMA public OWNER TO pg_database_owner;

--
-- TOC entry 3631 (class 0 OID 0)
-- Dependencies: 4
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: pg_database_owner
--

COMMENT ON SCHEMA public IS 'standard public schema';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 215 (class 1259 OID 16386)
-- Name: accessrulesdata; Type: TABLE; Schema: public; Owner: ejbca
--

CREATE TABLE public.accessrulesdata (
    pk integer NOT NULL,
    accessrule text NOT NULL,
    isrecursive boolean NOT NULL,
    rowprotection text,
    rowversion integer NOT NULL,
    rule integer NOT NULL,
    admingroupdata_accessrules integer
);


ALTER TABLE public.accessrulesdata OWNER TO ejbca;

--
-- TOC entry 245 (class 1259 OID 16598)
-- Name: acmeaccountdata; Type: TABLE; Schema: public; Owner: ejbca
--

CREATE TABLE public.acmeaccountdata (
    accountid text NOT NULL,
    currentkeyid text NOT NULL,
    rawdata text,
    rowprotection text,
    rowversion integer NOT NULL
);


ALTER TABLE public.acmeaccountdata OWNER TO ejbca;

--
-- TOC entry 248 (class 1259 OID 16619)
-- Name: acmeauthorizationdata; Type: TABLE; Schema: public; Owner: ejbca
--

CREATE TABLE public.acmeauthorizationdata (
    authorizationid text NOT NULL,
    identifier text,
    identifiertype text,
    expires bigint,
    status text,
    orderid text,
    accountid text NOT NULL,
    rawdata text,
    rowprotection text,
    rowversion integer NOT NULL
);


ALTER TABLE public.acmeauthorizationdata OWNER TO ejbca;

--
-- TOC entry 247 (class 1259 OID 16612)
-- Name: acmechallengedata; Type: TABLE; Schema: public; Owner: ejbca
--

CREATE TABLE public.acmechallengedata (
    challengeid text NOT NULL,
    authorizationid text NOT NULL,
    type text NOT NULL,
    rawdata text,
    rowprotection text,
    rowversion integer NOT NULL
);


ALTER TABLE public.acmechallengedata OWNER TO ejbca;

--
-- TOC entry 244 (class 1259 OID 16591)
-- Name: acmenoncedata; Type: TABLE; Schema: public; Owner: ejbca
--

CREATE TABLE public.acmenoncedata (
    nonce text NOT NULL,
    timeexpires bigint NOT NULL,
    rowprotection text,
    rowversion integer NOT NULL
);


ALTER TABLE public.acmenoncedata OWNER TO ejbca;

--
-- TOC entry 246 (class 1259 OID 16605)
-- Name: acmeorderdata; Type: TABLE; Schema: public; Owner: ejbca
--

CREATE TABLE public.acmeorderdata (
    orderid text NOT NULL,
    accountid text NOT NULL,
    fingerprint text,
    status text NOT NULL,
    rawdata text,
    rowprotection text,
    rowversion integer NOT NULL
);


ALTER TABLE public.acmeorderdata OWNER TO ejbca;

--
-- TOC entry 216 (class 1259 OID 16393)
-- Name: adminentitydata; Type: TABLE; Schema: public; Owner: ejbca
--

CREATE TABLE public.adminentitydata (
    pk integer NOT NULL,
    caid integer NOT NULL,
    matchtype integer NOT NULL,
    matchvalue text,
    matchwith integer NOT NULL,
    rowprotection text,
    rowversion integer NOT NULL,
    tokentype text,
    admingroupdata_adminentities integer
);


ALTER TABLE public.adminentitydata OWNER TO ejbca;

--
-- TOC entry 217 (class 1259 OID 16400)
-- Name: admingroupdata; Type: TABLE; Schema: public; Owner: ejbca
--

CREATE TABLE public.admingroupdata (
    pk integer NOT NULL,
    admingroupname text NOT NULL,
    rowprotection text,
    rowversion integer NOT NULL
);


ALTER TABLE public.admingroupdata OWNER TO ejbca;

--
-- TOC entry 218 (class 1259 OID 16407)
-- Name: adminpreferencesdata; Type: TABLE; Schema: public; Owner: ejbca
--

CREATE TABLE public.adminpreferencesdata (
    id text NOT NULL,
    data bytea NOT NULL,
    rowprotection text,
    rowversion integer NOT NULL
);


ALTER TABLE public.adminpreferencesdata OWNER TO ejbca;

--
-- TOC entry 219 (class 1259 OID 16414)
-- Name: approvaldata; Type: TABLE; Schema: public; Owner: ejbca
--

CREATE TABLE public.approvaldata (
    id integer NOT NULL,
    approvaldata text NOT NULL,
    approvalid integer NOT NULL,
    approvaltype integer NOT NULL,
    caid integer NOT NULL,
    endentityprofileid integer NOT NULL,
    expiredate bigint NOT NULL,
    remainingapprovals integer NOT NULL,
    subjectdn text,
    email text,
    reqadmincertissuerdn text,
    reqadmincertsn text,
    requestdata text NOT NULL,
    requestdate bigint NOT NULL,
    rowprotection text,
    rowversion integer NOT NULL,
    status integer NOT NULL
);


ALTER TABLE public.approvaldata OWNER TO ejbca;

--
-- TOC entry 220 (class 1259 OID 16421)
-- Name: auditrecorddata; Type: TABLE; Schema: public; Owner: ejbca
--

CREATE TABLE public.auditrecorddata (
    pk text NOT NULL,
    additionaldetails text,
    authtoken text NOT NULL,
    customid text,
    eventstatus text NOT NULL,
    eventtype text NOT NULL,
    module text NOT NULL,
    nodeid text NOT NULL,
    rowprotection text,
    rowversion integer NOT NULL,
    searchdetail1 text,
    searchdetail2 text,
    sequencenumber bigint NOT NULL,
    service text NOT NULL,
    "timestamp" bigint NOT NULL
);


ALTER TABLE public.auditrecorddata OWNER TO ejbca;

--
-- TOC entry 221 (class 1259 OID 16428)
-- Name: authorizationtreeupdatedata; Type: TABLE; Schema: public; Owner: ejbca
--

CREATE TABLE public.authorizationtreeupdatedata (
    pk integer NOT NULL,
    authorizationtreeupdatenumber integer NOT NULL,
    rowprotection text,
    rowversion integer NOT NULL
);


ALTER TABLE public.authorizationtreeupdatedata OWNER TO ejbca;

--
-- TOC entry 222 (class 1259 OID 16435)
-- Name: base64certdata; Type: TABLE; Schema: public; Owner: ejbca
--

CREATE TABLE public.base64certdata (
    fingerprint text NOT NULL,
    base64cert text,
    rowprotection text,
    rowversion integer NOT NULL,
    certificaterequest text
);


ALTER TABLE public.base64certdata OWNER TO ejbca;

--
-- TOC entry 237 (class 1259 OID 16541)
-- Name: blacklistdata; Type: TABLE; Schema: public; Owner: ejbca
--

CREATE TABLE public.blacklistdata (
    id integer NOT NULL,
    type text NOT NULL,
    value text NOT NULL,
    data text,
    rowprotection text,
    rowversion integer NOT NULL,
    updatecounter integer NOT NULL
);


ALTER TABLE public.blacklistdata OWNER TO ejbca;

--
-- TOC entry 223 (class 1259 OID 16442)
-- Name: cadata; Type: TABLE; Schema: public; Owner: ejbca
--

CREATE TABLE public.cadata (
    caid integer NOT NULL,
    data text NOT NULL,
    expiretime bigint NOT NULL,
    name text,
    rowprotection text,
    rowversion integer NOT NULL,
    status integer NOT NULL,
    subjectdn text,
    updatetime bigint NOT NULL
);


ALTER TABLE public.cadata OWNER TO ejbca;

--
-- TOC entry 226 (class 1259 OID 16463)
-- Name: certificatedata; Type: TABLE; Schema: public; Owner: ejbca
--

CREATE TABLE public.certificatedata (
    fingerprint text NOT NULL,
    base64cert text,
    cafingerprint text,
    certificateprofileid integer NOT NULL,
    endentityprofileid integer,
    crlpartitionindex integer,
    expiredate bigint NOT NULL,
    issuerdn text NOT NULL,
    notbefore bigint,
    invaliditydate bigint,
    revocationdate bigint NOT NULL,
    revocationreason integer NOT NULL,
    rowprotection text,
    rowversion integer NOT NULL,
    serialnumber text NOT NULL,
    status integer NOT NULL,
    subjectaltname text,
    subjectdn text NOT NULL,
    subjectkeyid text,
    accountbindingid text,
    tag text,
    type integer NOT NULL,
    updatetime bigint NOT NULL,
    username text,
    certificaterequest text
);


ALTER TABLE public.certificatedata OWNER TO ejbca;

--
-- TOC entry 227 (class 1259 OID 16470)
-- Name: certificateprofiledata; Type: TABLE; Schema: public; Owner: ejbca
--

CREATE TABLE public.certificateprofiledata (
    id integer NOT NULL,
    certificateprofilename text NOT NULL,
    data bytea NOT NULL,
    rowprotection text,
    rowversion integer NOT NULL
);


ALTER TABLE public.certificateprofiledata OWNER TO ejbca;

--
-- TOC entry 225 (class 1259 OID 16456)
-- Name: certreqhistorydata; Type: TABLE; Schema: public; Owner: ejbca
--

CREATE TABLE public.certreqhistorydata (
    fingerprint text NOT NULL,
    issuerdn text NOT NULL,
    rowprotection text,
    rowversion integer NOT NULL,
    serialnumber text NOT NULL,
    "timestamp" bigint NOT NULL,
    userdatavo text NOT NULL,
    username text NOT NULL
);


ALTER TABLE public.certreqhistorydata OWNER TO ejbca;

--
-- TOC entry 224 (class 1259 OID 16449)
-- Name: crldata; Type: TABLE; Schema: public; Owner: ejbca
--

CREATE TABLE public.crldata (
    fingerprint text NOT NULL,
    base64crl text NOT NULL,
    cafingerprint text NOT NULL,
    crlpartitionindex integer,
    crlnumber integer NOT NULL,
    deltacrlindicator integer NOT NULL,
    issuerdn text NOT NULL,
    nextupdate bigint NOT NULL,
    rowprotection text,
    rowversion integer NOT NULL,
    thisupdate bigint NOT NULL
);


ALTER TABLE public.crldata OWNER TO ejbca;

--
-- TOC entry 228 (class 1259 OID 16477)
-- Name: cryptotokendata; Type: TABLE; Schema: public; Owner: ejbca
--

CREATE TABLE public.cryptotokendata (
    id integer NOT NULL,
    lastupdate bigint NOT NULL,
    rowprotection text,
    rowversion integer NOT NULL,
    tokendata text,
    tokenname text NOT NULL,
    tokenprops text,
    tokentype text NOT NULL
);


ALTER TABLE public.cryptotokendata OWNER TO ejbca;

--
-- TOC entry 229 (class 1259 OID 16484)
-- Name: endentityprofiledata; Type: TABLE; Schema: public; Owner: ejbca
--

CREATE TABLE public.endentityprofiledata (
    id integer NOT NULL,
    data bytea NOT NULL,
    profilename text NOT NULL,
    rowprotection text,
    rowversion integer NOT NULL
);


ALTER TABLE public.endentityprofiledata OWNER TO ejbca;

--
-- TOC entry 230 (class 1259 OID 16491)
-- Name: globalconfigurationdata; Type: TABLE; Schema: public; Owner: ejbca
--

CREATE TABLE public.globalconfigurationdata (
    configurationid text NOT NULL,
    data bytea NOT NULL,
    rowprotection text,
    rowversion integer NOT NULL
);


ALTER TABLE public.globalconfigurationdata OWNER TO ejbca;

--
-- TOC entry 251 (class 1259 OID 16640)
-- Name: incompleteissuancejournaldata; Type: TABLE; Schema: public; Owner: ejbca
--

CREATE TABLE public.incompleteissuancejournaldata (
    serialnumberandcaid text NOT NULL,
    starttime bigint NOT NULL,
    rawdata text,
    rowprotection text,
    rowversion integer NOT NULL
);


ALTER TABLE public.incompleteissuancejournaldata OWNER TO ejbca;

--
-- TOC entry 231 (class 1259 OID 16498)
-- Name: internalkeybindingdata; Type: TABLE; Schema: public; Owner: ejbca
--

CREATE TABLE public.internalkeybindingdata (
    id integer NOT NULL,
    certificateid text,
    cryptotokenid integer NOT NULL,
    keybindingtype text NOT NULL,
    keypairalias text NOT NULL,
    lastupdate bigint NOT NULL,
    name text NOT NULL,
    rawdata text,
    rowprotection text,
    rowversion integer NOT NULL,
    status text NOT NULL
);


ALTER TABLE public.internalkeybindingdata OWNER TO ejbca;

--
-- TOC entry 232 (class 1259 OID 16505)
-- Name: keyrecoverydata; Type: TABLE; Schema: public; Owner: ejbca
--

CREATE TABLE public.keyrecoverydata (
    certsn text NOT NULL,
    issuerdn text NOT NULL,
    cryptotokenid integer DEFAULT 0 NOT NULL,
    keyalias text,
    keydata text NOT NULL,
    markedasrecoverable boolean NOT NULL,
    publickeyid text,
    rowprotection text,
    rowversion integer NOT NULL,
    username text
);


ALTER TABLE public.keyrecoverydata OWNER TO ejbca;

--
-- TOC entry 243 (class 1259 OID 16584)
-- Name: noconflictcertificatedata; Type: TABLE; Schema: public; Owner: ejbca
--

CREATE TABLE public.noconflictcertificatedata (
    id text NOT NULL,
    fingerprint text NOT NULL,
    base64cert text,
    cafingerprint text,
    certificateprofileid integer NOT NULL,
    endentityprofileid integer,
    crlpartitionindex integer,
    expiredate bigint NOT NULL,
    issuerdn text NOT NULL,
    notbefore bigint,
    invaliditydate bigint,
    revocationdate bigint NOT NULL,
    revocationreason integer NOT NULL,
    rowprotection text,
    rowversion integer NOT NULL,
    serialnumber text NOT NULL,
    status integer NOT NULL,
    subjectaltname text,
    subjectdn text NOT NULL,
    subjectkeyid text,
    accountbindingid text,
    tag text,
    type integer NOT NULL,
    updatetime bigint NOT NULL,
    username text,
    certificaterequest text
);


ALTER TABLE public.noconflictcertificatedata OWNER TO ejbca;

--
-- TOC entry 250 (class 1259 OID 16633)
-- Name: ocspresponsedata; Type: TABLE; Schema: public; Owner: ejbca
--

CREATE TABLE public.ocspresponsedata (
    id text NOT NULL,
    serialnumber text NOT NULL,
    producedat bigint NOT NULL,
    nextupdate bigint,
    ocspresponse bytea,
    caid integer,
    rowprotection text,
    rowversion integer NOT NULL
);


ALTER TABLE public.ocspresponsedata OWNER TO ejbca;

--
-- TOC entry 233 (class 1259 OID 16513)
-- Name: peerdata; Type: TABLE; Schema: public; Owner: ejbca
--

CREATE TABLE public.peerdata (
    id integer NOT NULL,
    connectorstate integer NOT NULL,
    data text,
    name text NOT NULL,
    rowprotection text,
    rowversion integer NOT NULL,
    url text NOT NULL
);


ALTER TABLE public.peerdata OWNER TO ejbca;

--
-- TOC entry 234 (class 1259 OID 16520)
-- Name: profiledata; Type: TABLE; Schema: public; Owner: ejbca
--

CREATE TABLE public.profiledata (
    id integer NOT NULL,
    profilename text NOT NULL,
    profiletype text NOT NULL,
    rawdata text,
    rowprotection text,
    rowversion integer NOT NULL
);


ALTER TABLE public.profiledata OWNER TO ejbca;

--
-- TOC entry 235 (class 1259 OID 16527)
-- Name: publisherdata; Type: TABLE; Schema: public; Owner: ejbca
--

CREATE TABLE public.publisherdata (
    id integer NOT NULL,
    data text,
    name text,
    rowprotection text,
    rowversion integer NOT NULL,
    updatecounter integer NOT NULL
);


ALTER TABLE public.publisherdata OWNER TO ejbca;

--
-- TOC entry 236 (class 1259 OID 16534)
-- Name: publisherqueuedata; Type: TABLE; Schema: public; Owner: ejbca
--

CREATE TABLE public.publisherqueuedata (
    pk text NOT NULL,
    fingerprint text,
    lastupdate bigint NOT NULL,
    publishstatus integer NOT NULL,
    publishtype integer NOT NULL,
    publisherid integer NOT NULL,
    rowprotection text,
    rowversion integer NOT NULL,
    timecreated bigint NOT NULL,
    trycounter integer NOT NULL,
    volatiledata text
);


ALTER TABLE public.publisherqueuedata OWNER TO ejbca;

--
-- TOC entry 238 (class 1259 OID 16548)
-- Name: roledata; Type: TABLE; Schema: public; Owner: ejbca
--

CREATE TABLE public.roledata (
    id integer NOT NULL,
    rolename text NOT NULL,
    namespace text,
    rawdata text,
    rowprotection text,
    rowversion integer NOT NULL
);


ALTER TABLE public.roledata OWNER TO ejbca;

--
-- TOC entry 239 (class 1259 OID 16555)
-- Name: rolememberdata; Type: TABLE; Schema: public; Owner: ejbca
--

CREATE TABLE public.rolememberdata (
    primarykey integer NOT NULL,
    tokentype text NOT NULL,
    tokenissuerid integer NOT NULL,
    tokenproviderid integer DEFAULT 0 NOT NULL,
    tokenmatchkey integer NOT NULL,
    tokenmatchoperator integer NOT NULL,
    tokenmatchvalue text,
    roleid integer NOT NULL,
    description text,
    rowprotection text,
    rowversion integer NOT NULL
);


ALTER TABLE public.rolememberdata OWNER TO ejbca;

--
-- TOC entry 249 (class 1259 OID 16626)
-- Name: sctdata; Type: TABLE; Schema: public; Owner: ejbca
--

CREATE TABLE public.sctdata (
    pk text NOT NULL,
    logid integer NOT NULL,
    fingerprint text NOT NULL,
    certificateexpirationdate bigint NOT NULL,
    data text,
    rowprotection text,
    rowversion integer NOT NULL
);


ALTER TABLE public.sctdata OWNER TO ejbca;

--
-- TOC entry 240 (class 1259 OID 16563)
-- Name: servicedata; Type: TABLE; Schema: public; Owner: ejbca
--

CREATE TABLE public.servicedata (
    id integer NOT NULL,
    data text,
    name text NOT NULL,
    nextruntimestamp bigint NOT NULL,
    rowprotection text,
    rowversion integer NOT NULL,
    runtimestamp bigint NOT NULL
);


ALTER TABLE public.servicedata OWNER TO ejbca;

--
-- TOC entry 241 (class 1259 OID 16570)
-- Name: userdata; Type: TABLE; Schema: public; Owner: ejbca
--

CREATE TABLE public.userdata (
    username text NOT NULL,
    caid integer NOT NULL,
    cardnumber text,
    certificateprofileid integer NOT NULL,
    clearpassword text,
    endentityprofileid integer NOT NULL,
    extendedinformationdata text,
    hardtokenissuerid integer NOT NULL,
    keystorepassword text,
    passwordhash text,
    rowprotection text,
    rowversion integer NOT NULL,
    status integer NOT NULL,
    subjectaltname text,
    subjectdn text,
    subjectemail text,
    timecreated bigint NOT NULL,
    timemodified bigint NOT NULL,
    tokentype integer NOT NULL,
    type integer NOT NULL
);


ALTER TABLE public.userdata OWNER TO ejbca;

--
-- TOC entry 242 (class 1259 OID 16577)
-- Name: userdatasourcedata; Type: TABLE; Schema: public; Owner: ejbca
--

CREATE TABLE public.userdatasourcedata (
    id integer NOT NULL,
    data text,
    name text NOT NULL,
    rowprotection text,
    rowversion integer NOT NULL,
    updatecounter integer NOT NULL
);


ALTER TABLE public.userdatasourcedata OWNER TO ejbca;

--
-- TOC entry 3589 (class 0 OID 16386)
-- Dependencies: 215
-- Data for Name: accessrulesdata; Type: TABLE DATA; Schema: public; Owner: ejbca
--



--
-- TOC entry 3619 (class 0 OID 16598)
-- Dependencies: 245
-- Data for Name: acmeaccountdata; Type: TABLE DATA; Schema: public; Owner: ejbca
--



--
-- TOC entry 3622 (class 0 OID 16619)
-- Dependencies: 248
-- Data for Name: acmeauthorizationdata; Type: TABLE DATA; Schema: public; Owner: ejbca
--



--
-- TOC entry 3621 (class 0 OID 16612)
-- Dependencies: 247
-- Data for Name: acmechallengedata; Type: TABLE DATA; Schema: public; Owner: ejbca
--



--
-- TOC entry 3618 (class 0 OID 16591)
-- Dependencies: 244
-- Data for Name: acmenoncedata; Type: TABLE DATA; Schema: public; Owner: ejbca
--



--
-- TOC entry 3620 (class 0 OID 16605)
-- Dependencies: 246
-- Data for Name: acmeorderdata; Type: TABLE DATA; Schema: public; Owner: ejbca
--



--
-- TOC entry 3590 (class 0 OID 16393)
-- Dependencies: 216
-- Data for Name: adminentitydata; Type: TABLE DATA; Schema: public; Owner: ejbca
--



--
-- TOC entry 3591 (class 0 OID 16400)
-- Dependencies: 217
-- Data for Name: admingroupdata; Type: TABLE DATA; Schema: public; Owner: ejbca
--



--
-- TOC entry 3592 (class 0 OID 16407)
-- Dependencies: 218
-- Data for Name: adminpreferencesdata; Type: TABLE DATA; Schema: public; Owner: ejbca
--



--
-- TOC entry 3593 (class 0 OID 16414)
-- Dependencies: 219
-- Data for Name: approvaldata; Type: TABLE DATA; Schema: public; Owner: ejbca
--



--
-- TOC entry 3594 (class 0 OID 16421)
-- Dependencies: 220
-- Data for Name: auditrecorddata; Type: TABLE DATA; Schema: public; Owner: ejbca
--



--
-- TOC entry 3595 (class 0 OID 16428)
-- Dependencies: 221
-- Data for Name: authorizationtreeupdatedata; Type: TABLE DATA; Schema: public; Owner: ejbca
--

INSERT INTO public.authorizationtreeupdatedata VALUES (2, 0, NULL, 0);
INSERT INTO public.authorizationtreeupdatedata VALUES (1, 5, NULL, 4);


--
-- TOC entry 3596 (class 0 OID 16435)
-- Dependencies: 222
-- Data for Name: base64certdata; Type: TABLE DATA; Schema: public; Owner: ejbca
--

INSERT INTO public.base64certdata VALUES ('b42503b02761af26789c632442a29b62b26ae930', 'MIIEszCCAxugAwIBAgIUCsVlu2WheKP5wXuPhIEzVhG+9kowDQYJKoZIhvcNAQEL
BQAwYTEjMCEGCgmSJomT8ixkAQEME2MtMDU3cm9ua3p2NGcydXdpODIxFTATBgNV
BAMMDE1hbmFnZW1lbnRDQTEjMCEGA1UECgwaRUpCQ0EgQ29udGFpbmVyIFF1aWNr
c3RhcnQwHhcNMjUwNzI0MDkyMjQ5WhcNMzUwNzI0MDkyMjQ4WjBhMSMwIQYKCZIm
iZPyLGQBAQwTYy0wNTdyb25renY0ZzJ1d2k4MjEVMBMGA1UEAwwMTWFuYWdlbWVu
dENBMSMwIQYDVQQKDBpFSkJDQSBDb250YWluZXIgUXVpY2tzdGFydDCCAaIwDQYJ
KoZIhvcNAQEBBQADggGPADCCAYoCggGBAM5ZWf5TZmdyg5Bn0bcbuDZWjgpHeEbP
x5KFizCU+ftJOWO8uE9wzgEngaIXDk6i6+cvEKInJBP4iPffkLiLEFCWjUuULUOz
4vHtnixZ0QPjUxwLsR+W99ADIsB/DVt301GFqj3GmNy+YW/onNkZHiZA5E5mDtUx
u5/qb3j4HIvG0dphQyK74wqyWTl7J8i0Wdf2AA/EvpCKaQfKj5vthjhiE0nk2CgY
tiMWwRzEeQBoH2ZgFthDpiDrIYIJfyG3ZHFg9QgH3tZm8Q0bjwfmXMjsj9/pwmwv
oetsd7LMxfblow8ow5ILRj57tVPKKgi8ypg0MEQj7vztD3scyaabQWp+MeAJv499
0IQq/jdQxqFQi2corM6y81u1HGnIomlI8CbQCDxePVM6W6O639Fd2uRlu6PDoCmQ
mKcNEJ38NU4Z1vvvGnc/D6muZMlqeJyJm2OwU4j+NgARB+noE32C8CoZlJvRA15L
MM0NGjrw6FthAo4PuG2+x4svWTSHmj++lQIDAQABo2MwYTAPBgNVHRMBAf8EBTAD
AQH/MB8GA1UdIwQYMBaAFArU5DP+SNumUyHEyQm9iC1VgCH5MB0GA1UdDgQWBBQK
1OQz/kjbplMhxMkJvYgtVYAh+TAOBgNVHQ8BAf8EBAMCAYYwDQYJKoZIhvcNAQEL
BQADggGBAFY1I1vTdxdOmj0lmhzlOyECnVTrII0Ryrsx9gdiIRKrbH4aIdqoZQdY
/IkIrBZXaKfqRyOMVZUvNBrymjYHdebHEsfW9Lcys1YRYY2R5vvrYrAewyl7YCZ7
CuUSufhcF/XJV4cTjxGszQH2JRVt6RyKst2iHfG7nHOXY63Pl14+ajM/BfMWHNSP
FTEjBc/HG9FDoVWIQSLnTsIVECjvuXxjPJMXaQDl0xQJj9dG6wrjK9pd56xPt+Rm
BHAJ8Ab4GVzXKhHQ+o83quhi5XIBbOQhDW7qbDH8FimH82U0JKpPt7NvIYmd5RN5
ORAKtRs2F0TxCFeilfLUhCPEqmP+hRPkD/E3ctgX7A8p8cjKHoqyGq9OpygEAjx5
SUR+DmXRqatTeo6EDZBMkDkk2BEoifL6xoon8vr9tYrhYDYClq6LhbAdSouHp38w
X/Li/pIxcg6ItLateWx6OHoIc+StNJ29DfkjkkBNc9b0c0aP5cS/RNUupIkkG09Y
PJiyvr7Kww==', NULL, 0, NULL);
INSERT INTO public.base64certdata VALUES ('930f8e25aad8e3437373c6bdfa8c169c50864515', 'MIIEWjCCAsKgAwIBAgIUeETcpSu16znixrl22IimeL/nIxIwDQYJKoZIhvcNAQEL
BQAwYTEjMCEGCgmSJomT8ixkAQEME2MtMDU3cm9ua3p2NGcydXdpODIxFTATBgNV
BAMMDE1hbmFnZW1lbnRDQTEjMCEGA1UECgwaRUpCQ0EgQ29udGFpbmVyIFF1aWNr
c3RhcnQwHhcNMjUwNzI0MDkyMjQ5WhcNMjcwNzI0MDkxMzM4WjBeMSMwIQYKCZIm
iZPyLGQBAQwTYy0wNTdyb25renY0ZzJ1d2k4MjESMBAGA1UEAwwJbG9jYWxob3N0
MSMwIQYDVQQKDBpFSkJDQSBDb250YWluZXIgUXVpY2tzdGFydDCCASIwDQYJKoZI
hvcNAQEBBQADggEPADCCAQoCggEBAMHtLpLrcbYMJqYNjy1Z8JqfRJFg4LlKQ72H
G0VhrimpYgdzknYYympji/pFaAg628CNG/P8pnD4r4GOhgvJNIhixMRjYYZp/Ecj
5chvo748RoJUKwqtIwdWU9qxfMe48lzMW63+glP7SIswLz9lfsleinc3Wtn46HN2
OmRySwSmSZr7iE5c79WPcK88vYVJcjkdo9C86F/5D4m5ONjStqFK+36zosRf0qYw
4vXp2hjCPcbj3U+elJQDvuvOmGbOjpaqLA0jDC0aETPs3e0RWhS4RWU4FJqySZik
w17gfVo50TpM+CPkSaR+oYzPIDCC4CFIU4W8oXOzEJWrdtYUU0kCAwEAAaOBjDCB
iTAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFArU5DP+SNumUyHEyQm9iC1VgCH5
MBQGA1UdEQQNMAuCCWxvY2FsaG9zdDATBgNVHSUEDDAKBggrBgEFBQcDATAdBgNV
HQ4EFgQUDrk+XNLl081bdyqIbPymSjL+12owDgYDVR0PAQH/BAQDAgWgMA0GCSqG
SIb3DQEBCwUAA4IBgQA9xLQ0PHMdw2tjZZPdctugOdz9fOcLid8hI6pzcNhWVk9z
MA7jQK5kLz4DIHj/YZDoXvItCSbRYEqa/acO9nqWeHYJlI4qJugjmUZQS0adK8Yp
v+rWqlD0HD2KE3uz2cri2MwfQmOYqi2Sf3Ph0CmRIJh4WGc8p/0UDT1bgkWgfj1M
7dLHPMO87EjNGR9/DXtJNIRaZsVCPF39EGQm8QFtn1EdiUIrf4Znq17ZVRFTleR/
cpDFJC41Ooht97vFGbrOxVssnsXvNMNNQsxz4/r5D3VqOosktOzj7FEyNJR7CYJO
9miwpSM0BtuIcxB2ZHrEVhRXbytjJg+ZUfNJrRA7b/BKSUjll3zm5dqAN2b2vXHG
30gr31WvqC7IinZIaKm421Dk46ozZMaxQAGZGJZhNuWiwKPnT+IawsjpbUgTzx1c
c2mxdopa6g9sOI9exAW8YIyto48JYWz8xo4gX71M6VYumZbANyEZ24JxZFPP9lQC
P3tKOCuOZJCE0EI+9bw=', NULL, 0, NULL);
INSERT INTO public.base64certdata VALUES ('b08049d7f246eb4691fcd3c1d78b4316ccc34cf2', 'MIIBazCCARCgAwIBAgIUZ3sE9Hadu2NmBfnpfdXtKakDktswCgYIKoZIzj0EAwIw
EjEQMA4GA1UEAwwHU2ltcGxDQTAgFw0yNTA3MjQwOTI1NTJaGA8yMDU1MDcxNzA5
MjU1MVowEjEQMA4GA1UEAwwHU2ltcGxDQTBZMBMGByqGSM49AgEGCCqGSM49AwEH
A0IABNiYmz55uyVc8hp5jeBO0I5oKSkQDP9L4gg35GYnUUOYuVMrZ3tR72AeW/+K
q6nDe3+XwZs/Svlb8wFsiyje7TSjQjBAMA8GA1UdEwEB/wQFMAMBAf8wHQYDVR0O
BBYEFMP8ueXax/lKhCylyWpcaepub8tmMA4GA1UdDwEB/wQEAwIBhjAKBggqhkjO
PQQDAgNJADBGAiEA44k+t6vJOlDY/UDXaZbHe00Zfk43X5Dn+j12x1xjYlUCIQCt
pp/MMBgeJ2BgOc+LKBd30Dqluy1fcGzBkQKeOa9UxA==', NULL, 0, NULL);
INSERT INTO public.base64certdata VALUES ('1ed4d984a5a8f8dbd07694b0167e38141d1bcf0a', 'MIIBjzCCATSgAwIBAgIUJIY2Dw3yLhDIi15D/HzBLaTgnr4wCgYIKoZIzj0EAwIw
EjEQMA4GA1UEAwwHU2ltcGxDQTAeFw0yNTA3MjQwOTI2NDJaFw00MDA3MjAwOTI2
NDFaMBcxFTATBgNVBAMMDE9uQm9hcmRpbmdDQTBZMBMGByqGSM49AgEGCCqGSM49
AwEHA0IABGoMRbKWgBvW9+dV/2KDLMBmb2TBy0GaS5iLZUYzT3GXGMx8MQupa9/t
EOB27FYbqD4+33kFzf0tXetfOQydx+OjYzBhMA8GA1UdEwEB/wQFMAMBAf8wHwYD
VR0jBBgwFoAUw/y55drH+UqELKXJalxp6m5vy2YwHQYDVR0OBBYEFBNdsd/W1DnD
AL39mW7j1/aC12aYMA4GA1UdDwEB/wQEAwIBhjAKBggqhkjOPQQDAgNJADBGAiEA
1mHrQU2CyBo7drLfKrE9HMESs6S68K6ygbCg0Ne2s8oCIQCDrtszmY75xj+BXQoQ
cJmXr+MTnHMFqchaS6xm76YNoQ==', NULL, 0, NULL);
INSERT INTO public.base64certdata VALUES ('6c23cf79f157fb1c5bf9aee8e6f7ed35cdc0d139', 'MIIEHDCCAoSgAwIBAgIUbNqlOAwKAei3/Hjtu0IQYuKQonswDQYJKoZIhvcNAQEL
BQAwYTEjMCEGCgmSJomT8ixkAQEME2MtMDU3cm9ua3p2NGcydXdpODIxFTATBgNV
BAMMDE1hbmFnZW1lbnRDQTEjMCEGA1UECgwaRUpCQ0EgQ29udGFpbmVyIFF1aWNr
c3RhcnQwHhcNMjUwNzI0MDkyMjQ5WhcNMjcwNzI0MDkxNzQ5WjAVMRMwEQYDVQQD
DApTdXBlckFkbWluMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAr5Kx
JslKl/PsY2r7TfdPSP+pkCQnWp2ovjPty3zL/iQF85aZTIYY3yUcDdsPVbTcfBLD
R/pK5EoegFoHLl1MIe6S3R9UpIUQdl5D8+vCvvLtzat0Dzfb3LRtxzd2JjA52/fJ
FQeVKFgUbWvWLWvyakBV/VgQea2T+s6hQV5b/WdPJJJGwuA8Vt7V1dWZu83lc8n4
p4T14krHOVYyADbxMuidUcIc4GrjKYSQIZrzolwFILYzkFso+zrBCCt16M82MlN8
uSzKMpLwM8SHgxkU60PYjEdh4yeNRE8T5DUE1zscWKdKHwjwmbWjSkbpQruy9yqu
FGKYE9Es8VI4rw/h8wIDAQABo4GXMIGUMAwGA1UdEwEB/wQCMAAwHwYDVR0jBBgw
FoAUCtTkM/5I26ZTIcTJCb2ILVWAIfkwFQYDVR0RBA4wDIIKU3VwZXJBZG1pbjAd
BgNVHSUEFjAUBggrBgEFBQcDAgYIKwYBBQUHAwQwHQYDVR0OBBYEFDbFfVhsft1e
CuLv6GTJDKsHCAriMA4GA1UdDwEB/wQEAwIF4DANBgkqhkiG9w0BAQsFAAOCAYEA
BoUbOAavITdaUVbGCg2XIfKQBvYmSsWLwKUSbtvRVLgbfynMzJbBbRuEFUPhIylS
Ctg7YqlTNEkf8dQYi6b+WhTUrK0Jg0SxAU8YuM7G0DPFehxRMq7DcIFDoG9M82/9
IPv14ed5Uj9VtsoLpHktgIZN5UjxZUvUu+ttPT3Ah1etQbLdzqDUtkMk5Ke8o9nJ
Z6mg70mv8wbygRoWb6QtyKwtRt/e6gHQw7dHIA/gaSANjExmzrSkiR7oafgTq6wF
SEdZrFNBP10RmhQZeleltSgIt4uuBtLukECLRStGUk0owBbHfpoJlbK5ex5oUc4M
Ca8fZg95HSkOCfNAek/dYpe9Lrz+Uz0QeKAerkvCEGLYfAXU1SyiVhrnZz79pqFX
xkWuNJz/bJbsLb2l5/5t5fCji0C81GNKVcwqpqYNwP8EcqZpz5GHR4VItqcFRLwN
95V4S78bcPYTCqdl8YODPgm9HBW86AtKya3qSAZVeqhthMOH+Uv4sf0nuHI26Ji9', NULL, 0, NULL);
INSERT INTO public.base64certdata VALUES ('8c6d1e40eded0b170f0d897c126334943e725bb3', 'MIIEWjCCAsKgAwIBAgIUWW5+meWj2cXATsGUXr1B6E+MRFQwDQYJKoZIhvcNAQEL
BQAwYTEjMCEGCgmSJomT8ixkAQEME2MtMDU3cm9ua3p2NGcydXdpODIxFTATBgNV
BAMMDE1hbmFnZW1lbnRDQTEjMCEGA1UECgwaRUpCQ0EgQ29udGFpbmVyIFF1aWNr
c3RhcnQwHhcNMjUwNzI0MDkyMjU0WhcNMjcwNzI0MDkyMjUzWjBeMSMwIQYKCZIm
iZPyLGQBAQwTYy0wNTdyb25renY0ZzJ1d2k4MjESMBAGA1UEAwwJbG9jYWxob3N0
MSMwIQYDVQQKDBpFSkJDQSBDb250YWluZXIgUXVpY2tzdGFydDCCASIwDQYJKoZI
hvcNAQEBBQADggEPADCCAQoCggEBAKPphLaMdhE0hF4uoRDewNySeKWeuKySw6wB
H+QGPEVfRMpHjFLZuAKJ0wfsm18Yzlxc+VThVOQsLfbJUE+pcgmm9zgv8bJdCTkH
Kc+VmaYOTIfbwT9Z3/smAR9TaxhlobLAIiixlcEH7tL4RGETlleiL69lzRKPQGY7
8shmrVcTNgTu/3+YaeLLWVmOHddO3uzVjvJC2s+WDCMmoNP2vx5U0PeuIJJ7qkFM
wjLsNP5epNSo1gomegnTVrnZzCvj1tddf9jCtLq9uCe5Z8Hy3DeW2vzwcwzNWWEV
sWgwTLBQLh8HzFMLLD7jUjRyIsMEZblXo5yu0NwDYm+ab+EmQHUCAwEAAaOBjDCB
iTAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFArU5DP+SNumUyHEyQm9iC1VgCH5
MBQGA1UdEQQNMAuCCWxvY2FsaG9zdDATBgNVHSUEDDAKBggrBgEFBQcDATAdBgNV
HQ4EFgQUGxGfo89tBFMkCAkCM31nEib7xAowDgYDVR0PAQH/BAQDAgWgMA0GCSqG
SIb3DQEBCwUAA4IBgQB4Q5kuCKkY7OkwC001yPZ5CdNk6pnuBgG1tq7Mz32GZLVR
IMZ/mRFu/Fd5u3dluSQ82F0eiRUKJXI3MjQtOP/wFlOqMd/4UImBRPiwOU1ulH7S
p1P5B4/kX8UP8KReX9ShLY/TdND/XWjkdLuK4gliP0qvnbxQiCffj1/TRVOZ84Lb
s8ysnV63ADtQ0FGgHXV7pV72XGlY0wbJPa3+KSH74C7lczYsfFxRDbmV/G+fTZ0u
foaKezVglowmR+NoN4ye7H0bxDVOuAK4Ko7j42vrWpsRKKnBMA3zhVRodUTzV7c1
wyE/g2aHVPmSSXh4rXetghtYziR5XhGagfV1s7U+vsmDkeYoJtkAiN+hqhTVCr3O
whHdL7VFuYA2C+W1tAALb/3LDi9I44D3sphHQ5LPZ+SlL+occH0wIelRv8w2OAtH
0m54hUjYOW2xq7tzCliGJ3t9apAno6Bo05RNcaU6BwSdbqrUxhiLGF1eVOBB8ZJl
H/8EbmbZb1T+VkpDyO8=', NULL, 0, NULL);


--
-- TOC entry 3611 (class 0 OID 16541)
-- Dependencies: 237
-- Data for Name: blacklistdata; Type: TABLE DATA; Schema: public; Owner: ejbca
--



--
-- TOC entry 3597 (class 0 OID 16442)
-- Dependencies: 223
-- Data for Name: cadata; Type: TABLE DATA; Schema: public; Owner: ejbca
--

INSERT INTO public.cadata VALUES (-968899527, '<?xml version="1.0" encoding="UTF-8"?>
<java version="11.0.21" class="java.beans.XMLDecoder">
 <object class="org.cesecore.util.Base64PutHashMap">
  <void method="put">
   <string>encodedvalidity</string>
   <string>3652d</string>
  </void>
  <void method="put">
   <string>crlpublishers</string>
   <object class="java.util.ArrayList"/>
  </void>
  <void method="put">
   <string>signedby</string>
   <int>1</int>
  </void>
  <void method="put">
   <string>keyvalidators</string>
   <object class="java.util.ArrayList"/>
  </void>
  <void method="put">
   <string>description</string>
   <string>ManagementCAcreated using CLI</string>
  </void>
  <void method="put">
   <string>revokationreason</string>
   <int>-1</int>
  </void>
  <void method="put">
   <string>certificateprofileid</string>
   <int>3</int>
  </void>
  <void method="put">
   <string>usenoconflictcertificatedata</string>
   <boolean>false</boolean>
  </void>
  <void method="put">
   <string>finishuser</string>
   <boolean>true</boolean>
  </void>
  <void method="put">
   <string>includeinhealthcheck</string>
   <boolean>true</boolean>
  </void>
  <void method="put">
   <string>doEnforceUniquePublicKeys</string>
   <boolean>true</boolean>
  </void>
  <void method="put">
   <string>doEnforceKeyRenewal</string>
   <boolean>false</boolean>
  </void>
  <void method="put">
   <string>doEnforceUniqueDistinguishedName</string>
   <boolean>true</boolean>
  </void>
  <void method="put">
   <string>doEnforceUniqueSubjectDNSerialnumber</string>
   <boolean>false</boolean>
  </void>
  <void method="put">
   <string>useCertreqHistory</string>
   <boolean>false</boolean>
  </void>
  <void method="put">
   <string>useUserStorage</string>
   <boolean>true</boolean>
  </void>
  <void method="put">
   <string>useCertificateStorage</string>
   <boolean>true</boolean>
  </void>
  <void method="put">
   <string>acceptRevocationNonExistingEntry</string>
   <boolean>false</boolean>
  </void>
  <void method="put">
   <string>keepExpiredCertsOnCRL</string>
   <boolean>false</boolean>
  </void>
  <void method="put">
   <string>crlperiod</string>
   <long>86400000</long>
  </void>
  <void method="put">
   <string>crlIssueInterval</string>
   <long>0</long>
  </void>
  <void method="put">
   <string>crlOverlapTime</string>
   <long>600000</long>
  </void>
  <void method="put">
   <string>deltacrlperiod</string>
   <long>0</long>
  </void>
  <void method="put">
   <string>generatecrluponrevocation</string>
   <boolean>false</boolean>
  </void>
  <void method="put">
   <string>allowchangingrevocationreason</string>
   <boolean>false</boolean>
  </void>
  <void method="put">
   <string>allowinvaliditydate</string>
   <boolean>false</boolean>
  </void>
  <void method="put">
   <string>extendedcaservice5</string>
   <object class="java.util.LinkedHashMap">
    <void method="put">
     <string>IMPLCLASS</string>
     <string>org.ejbca.core.model.ca.caadmin.extendedcaservices.KeyRecoveryCAService</string>
    </void>
    <void method="put">
     <string>extendedcaservicetype</string>
     <int>5</int>
    </void>
    <void method="put">
     <string>version</string>
     <float>1.0</float>
    </void>
    <void method="put">
     <string>status</string>
     <int>2</int>
    </void>
   </object>
  </void>
  <void method="put">
   <string>extendedcaservices</string>
   <object class="java.util.ArrayList">
    <void method="add">
     <int>5</int>
    </void>
   </object>
  </void>
  <void method="put">
   <string>approvals</string>
   <object class="java.util.LinkedHashMap"/>
  </void>
  <void method="put">
   <string>policies</string>
   <object class="java.util.ArrayList"/>
  </void>
  <void method="put">
   <string>subjectaltname</string>
   <null/>
  </void>
  <void method="put">
   <string>catype</string>
   <int>1</int>
  </void>
  <void method="put">
   <string>version</string>
   <float>25.0</float>
  </void>
  <void method="put">
   <string>msCaCompatible</string>
   <boolean>false</boolean>
  </void>
  <void method="put">
   <string>useauthoritykeyidentifier</string>
   <boolean>true</boolean>
  </void>
  <void method="put">
   <string>authoritykeyidentifiercritical</string>
   <boolean>false</boolean>
  </void>
  <void method="put">
   <string>usecrlnumber</string>
   <boolean>true</boolean>
  </void>
  <void method="put">
   <string>crlnumbercritical</string>
   <boolean>false</boolean>
  </void>
  <void method="put">
   <string>defaultcrldistpoint</string>
   <string></string>
  </void>
  <void method="put">
   <string>defaultcrlissuer</string>
   <string></string>
  </void>
  <void method="put">
   <string>cadefinedfreshestcrl</string>
   <string></string>
  </void>
  <void method="put">
   <string>defaultocspservicelocator</string>
   <string></string>
  </void>
  <void method="put">
   <string>useutf8policytext</string>
   <boolean>true</boolean>
  </void>
  <void method="put">
   <string>useprintablestringsubjectdn</string>
   <boolean>false</boolean>
  </void>
  <void method="put">
   <string>useldapdnorder</string>
   <boolean>true</boolean>
  </void>
  <void method="put">
   <string>usecrldistributionpointoncrl</string>
   <boolean>false</boolean>
  </void>
  <void method="put">
   <string>crldistributionpointoncrlcritical</string>
   <boolean>false</boolean>
  </void>
  <void method="put">
   <string>cmpraauthsecret</string>
   <string></string>
  </void>
  <void method="put">
   <string>authorityinformationaccess</string>
   <null/>
  </void>
  <void method="put">
   <string>certificateaiadefaultcaissueruri</string>
   <null/>
  </void>
  <void method="put">
   <string>nameconstraintspermitted</string>
   <null/>
  </void>
  <void method="put">
   <string>nameconstraintsexcluded</string>
   <null/>
  </void>
  <void method="put">
   <string>serialnumberoctetsize</string>
   <int>20</int>
  </void>
  <void method="put">
   <string>dopreproduceocspresponses</string>
   <boolean>false</boolean>
  </void>
  <void method="put">
   <string>dostoreocspondemand</string>
   <boolean>false</boolean>
  </void>
  <void method="put">
   <string>dopreproduceocspresponsesuponissuanceandrevocation</string>
   <boolean>false</boolean>
  </void>
  <void method="put">
   <string>usepartitionedcrl</string>
   <boolean>false</boolean>
  </void>
  <void method="put">
   <string>crlpartitions</string>
   <int>0</int>
  </void>
  <void method="put">
   <string>suspendedcrlpartitions</string>
   <int>0</int>
  </void>
  <void method="put">
   <string>requestpreprocessor</string>
   <null/>
  </void>
  <void method="put">
   <string>catoken</string>
   <object class="java.util.LinkedHashMap">
    <void method="put">
     <string>version</string>
     <float>8.0</float>
    </void>
    <void method="put">
     <string>cryptotokenid</string>
     <string>1726589353</string>
    </void>
    <void method="put">
     <string>propertydata</string>
     <string>certSignKey=signKey
crlSignKey=signKey
defaultKey=encryptKey
</string>
    </void>
    <void method="put">
     <string>signaturealgorithm</string>
     <string>SHA256WithRSA</string>
    </void>
    <void method="put">
     <string>encryptionalgorithm</string>
     <string>SHA256WithRSA</string>
    </void>
   </object>
  </void>
  <void method="put">
   <string>certificatechain</string>
   <object class="java.util.ArrayList">
    <void method="add">
     <string>MIIEszCCAxugAwIBAgIUCsVlu2WheKP5wXuPhIEzVhG+9kowDQYJKoZIhvcNAQEL
BQAwYTEjMCEGCgmSJomT8ixkAQEME2MtMDU3cm9ua3p2NGcydXdpODIxFTATBgNV
BAMMDE1hbmFnZW1lbnRDQTEjMCEGA1UECgwaRUpCQ0EgQ29udGFpbmVyIFF1aWNr
c3RhcnQwHhcNMjUwNzI0MDkyMjQ5WhcNMzUwNzI0MDkyMjQ4WjBhMSMwIQYKCZIm
iZPyLGQBAQwTYy0wNTdyb25renY0ZzJ1d2k4MjEVMBMGA1UEAwwMTWFuYWdlbWVu
dENBMSMwIQYDVQQKDBpFSkJDQSBDb250YWluZXIgUXVpY2tzdGFydDCCAaIwDQYJ
KoZIhvcNAQEBBQADggGPADCCAYoCggGBAM5ZWf5TZmdyg5Bn0bcbuDZWjgpHeEbP
x5KFizCU+ftJOWO8uE9wzgEngaIXDk6i6+cvEKInJBP4iPffkLiLEFCWjUuULUOz
4vHtnixZ0QPjUxwLsR+W99ADIsB/DVt301GFqj3GmNy+YW/onNkZHiZA5E5mDtUx
u5/qb3j4HIvG0dphQyK74wqyWTl7J8i0Wdf2AA/EvpCKaQfKj5vthjhiE0nk2CgY
tiMWwRzEeQBoH2ZgFthDpiDrIYIJfyG3ZHFg9QgH3tZm8Q0bjwfmXMjsj9/pwmwv
oetsd7LMxfblow8ow5ILRj57tVPKKgi8ypg0MEQj7vztD3scyaabQWp+MeAJv499
0IQq/jdQxqFQi2corM6y81u1HGnIomlI8CbQCDxePVM6W6O639Fd2uRlu6PDoCmQ
mKcNEJ38NU4Z1vvvGnc/D6muZMlqeJyJm2OwU4j+NgARB+noE32C8CoZlJvRA15L
MM0NGjrw6FthAo4PuG2+x4svWTSHmj++lQIDAQABo2MwYTAPBgNVHRMBAf8EBTAD
AQH/MB8GA1UdIwQYMBaAFArU5DP+SNumUyHEyQm9iC1VgCH5MB0GA1UdDgQWBBQK
1OQz/kjbplMhxMkJvYgtVYAh+TAOBgNVHQ8BAf8EBAMCAYYwDQYJKoZIhvcNAQEL
BQADggGBAFY1I1vTdxdOmj0lmhzlOyECnVTrII0Ryrsx9gdiIRKrbH4aIdqoZQdY
/IkIrBZXaKfqRyOMVZUvNBrymjYHdebHEsfW9Lcys1YRYY2R5vvrYrAewyl7YCZ7
CuUSufhcF/XJV4cTjxGszQH2JRVt6RyKst2iHfG7nHOXY63Pl14+ajM/BfMWHNSP
FTEjBc/HG9FDoVWIQSLnTsIVECjvuXxjPJMXaQDl0xQJj9dG6wrjK9pd56xPt+Rm
BHAJ8Ab4GVzXKhHQ+o83quhi5XIBbOQhDW7qbDH8FimH82U0JKpPt7NvIYmd5RN5
ORAKtRs2F0TxCFeilfLUhCPEqmP+hRPkD/E3ctgX7A8p8cjKHoqyGq9OpygEAjx5
SUR+DmXRqatTeo6EDZBMkDkk2BEoifL6xoon8vr9tYrhYDYClq6LhbAdSouHp38w
X/Li/pIxcg6ItLateWx6OHoIc+StNJ29DfkjkkBNc9b0c0aP5cS/RNUupIkkG09Y
PJiyvr7Kww==</string>
    </void>
   </object>
  </void>
 </object>
</java>
', 2068881768000, 'ManagementCA', NULL, 2, 1, 'UID=c-057ronkzv4g2uwi82,CN=ManagementCA,O=EJBCA Container Quickstart', 1753348969428);
INSERT INTO public.cadata VALUES (1700994603, '<?xml version="1.0" encoding="UTF-8"?>
<java version="11.0.21" class="java.beans.XMLDecoder">
 <object class="org.cesecore.util.Base64PutHashMap">
  <void method="put">
   <string>encodedvalidity</string>
   <string>15y</string>
  </void>
  <void method="put">
   <string>crlpublishers</string>
   <object class="java.util.ArrayList"/>
  </void>
  <void method="put">
   <string>signedby</string>
   <int>-1128220897</int>
  </void>
  <void method="put">
   <string>keyvalidators</string>
   <object class="java.util.ArrayList"/>
  </void>
  <void method="put">
   <string>description</string>
   <string>OnBoardingCAcreated using CLI</string>
  </void>
  <void method="put">
   <string>revokationreason</string>
   <int>-1</int>
  </void>
  <void method="put">
   <string>certificateprofileid</string>
   <int>626302406</int>
  </void>
  <void method="put">
   <string>usenoconflictcertificatedata</string>
   <boolean>false</boolean>
  </void>
  <void method="put">
   <string>finishuser</string>
   <boolean>true</boolean>
  </void>
  <void method="put">
   <string>includeinhealthcheck</string>
   <boolean>true</boolean>
  </void>
  <void method="put">
   <string>doEnforceUniquePublicKeys</string>
   <boolean>false</boolean>
  </void>
  <void method="put">
   <string>doEnforceKeyRenewal</string>
   <boolean>false</boolean>
  </void>
  <void method="put">
   <string>doEnforceUniqueDistinguishedName</string>
   <boolean>false</boolean>
  </void>
  <void method="put">
   <string>doEnforceUniqueSubjectDNSerialnumber</string>
   <boolean>false</boolean>
  </void>
  <void method="put">
   <string>useCertreqHistory</string>
   <boolean>false</boolean>
  </void>
  <void method="put">
   <string>useUserStorage</string>
   <boolean>true</boolean>
  </void>
  <void method="put">
   <string>useCertificateStorage</string>
   <boolean>true</boolean>
  </void>
  <void method="put">
   <string>acceptRevocationNonExistingEntry</string>
   <boolean>false</boolean>
  </void>
  <void method="put">
   <string>keepExpiredCertsOnCRL</string>
   <boolean>false</boolean>
  </void>
  <void method="put">
   <string>crlperiod</string>
   <long>7776000000</long>
  </void>
  <void method="put">
   <string>crlIssueInterval</string>
   <long>86400000</long>
  </void>
  <void method="put">
   <string>crlOverlapTime</string>
   <long>0</long>
  </void>
  <void method="put">
   <string>deltacrlperiod</string>
   <long>0</long>
  </void>
  <void method="put">
   <string>generatecrluponrevocation</string>
   <boolean>false</boolean>
  </void>
  <void method="put">
   <string>allowchangingrevocationreason</string>
   <boolean>false</boolean>
  </void>
  <void method="put">
   <string>allowinvaliditydate</string>
   <boolean>false</boolean>
  </void>
  <void method="put">
   <string>extendedcaservice5</string>
   <object class="org.cesecore.util.Base64GetHashMap">
    <void method="put">
     <string>IMPLCLASS</string>
     <string>org.ejbca.core.model.ca.caadmin.extendedcaservices.KeyRecoveryCAService</string>
    </void>
    <void method="put">
     <string>extendedcaservicetype</string>
     <int>5</int>
    </void>
    <void method="put">
     <string>version</string>
     <float>1.0</float>
    </void>
    <void method="put">
     <string>status</string>
     <int>2</int>
    </void>
   </object>
  </void>
  <void method="put">
   <string>extendedcaservices</string>
   <object class="java.util.ArrayList">
    <void method="add">
     <int>5</int>
    </void>
   </object>
  </void>
  <void method="put">
   <string>approvals</string>
   <object class="java.util.LinkedHashMap"/>
  </void>
  <void method="put">
   <string>policies</string>
   <object class="java.util.ArrayList"/>
  </void>
  <void method="put">
   <string>subjectaltname</string>
   <null/>
  </void>
  <void method="put">
   <string>catype</string>
   <int>1</int>
  </void>
  <void method="put">
   <string>version</string>
   <float>25.0</float>
  </void>
  <void method="put">
   <string>msCaCompatible</string>
   <boolean>false</boolean>
  </void>
  <void method="put">
   <string>useauthoritykeyidentifier</string>
   <boolean>true</boolean>
  </void>
  <void method="put">
   <string>authoritykeyidentifiercritical</string>
   <boolean>false</boolean>
  </void>
  <void method="put">
   <string>usecrlnumber</string>
   <boolean>true</boolean>
  </void>
  <void method="put">
   <string>crlnumbercritical</string>
   <boolean>false</boolean>
  </void>
  <void method="put">
   <string>defaultcrldistpoint</string>
   <string>http://localhost:8100/crl/OnBoardingCA</string>
  </void>
  <void method="put">
   <string>defaultcrlissuer</string>
   <string></string>
  </void>
  <void method="put">
   <string>cadefinedfreshestcrl</string>
   <string></string>
  </void>
  <void method="put">
   <string>defaultocspservicelocator</string>
   <string>http://localhost:8100/ocsp</string>
  </void>
  <void method="put">
   <string>useutf8policytext</string>
   <boolean>true</boolean>
  </void>
  <void method="put">
   <string>useprintablestringsubjectdn</string>
   <boolean>false</boolean>
  </void>
  <void method="put">
   <string>useldapdnorder</string>
   <boolean>false</boolean>
  </void>
  <void method="put">
   <string>usecrldistributionpointoncrl</string>
   <boolean>false</boolean>
  </void>
  <void method="put">
   <string>crldistributionpointoncrlcritical</string>
   <boolean>false</boolean>
  </void>
  <void method="put">
   <string>cmpraauthsecret</string>
   <string></string>
  </void>
  <void method="put">
   <string>authorityinformationaccess</string>
   <null/>
  </void>
  <void method="put">
   <string>certificateaiadefaultcaissueruri</string>
   <object class="java.util.ArrayList">
    <void method="add">
     <string>http://localhost:8100/ca/OnBoardingCA</string>
    </void>
   </object>
  </void>
  <void method="put">
   <string>nameconstraintspermitted</string>
   <null/>
  </void>
  <void method="put">
   <string>nameconstraintsexcluded</string>
   <null/>
  </void>
  <void method="put">
   <string>serialnumberoctetsize</string>
   <int>20</int>
  </void>
  <void method="put">
   <string>dopreproduceocspresponses</string>
   <boolean>false</boolean>
  </void>
  <void method="put">
   <string>dostoreocspondemand</string>
   <boolean>false</boolean>
  </void>
  <void method="put">
   <string>dopreproduceocspresponsesuponissuanceandrevocation</string>
   <boolean>false</boolean>
  </void>
  <void method="put">
   <string>usepartitionedcrl</string>
   <boolean>false</boolean>
  </void>
  <void method="put">
   <string>crlpartitions</string>
   <int>0</int>
  </void>
  <void method="put">
   <string>suspendedcrlpartitions</string>
   <int>0</int>
  </void>
  <void method="put">
   <string>requestpreprocessor</string>
   <null/>
  </void>
  <void method="put">
   <string>catoken</string>
   <object class="org.cesecore.util.Base64GetHashMap">
    <void method="put">
     <string>version</string>
     <float>8.0</float>
    </void>
    <void method="put">
     <string>cryptotokenid</string>
     <string>-673066697</string>
    </void>
    <void method="put">
     <string>propertydata</string>
     <string>certSignKey=SimplOnboardingSignKey001
keyEncryptKey=SimplOnboardingEncryptKey001
crlSignKey=SimplOnboardingSignKey001
testKey=SimplOnboardingTestKey001
defaultKey=SimplOnboardingEncryptKey001
</string>
    </void>
    <void method="put">
     <string>signaturealgorithm</string>
     <string>SHA256withECDSA</string>
    </void>
    <void method="put">
     <string>encryptionalgorithm</string>
     <string>SHA256withECDSA</string>
    </void>
    <void method="put">
     <string>sequence</string>
     <string>00000</string>
    </void>
    <void method="put">
     <string>sequenceformat</string>
     <int>1</int>
    </void>
   </object>
  </void>
  <void method="put">
   <string>certificatechain</string>
   <object class="java.util.ArrayList">
    <void method="add">
     <string>MIIBjzCCATSgAwIBAgIUJIY2Dw3yLhDIi15D/HzBLaTgnr4wCgYIKoZIzj0EAwIw
EjEQMA4GA1UEAwwHU2ltcGxDQTAeFw0yNTA3MjQwOTI2NDJaFw00MDA3MjAwOTI2
NDFaMBcxFTATBgNVBAMMDE9uQm9hcmRpbmdDQTBZMBMGByqGSM49AgEGCCqGSM49
AwEHA0IABGoMRbKWgBvW9+dV/2KDLMBmb2TBy0GaS5iLZUYzT3GXGMx8MQupa9/t
EOB27FYbqD4+33kFzf0tXetfOQydx+OjYzBhMA8GA1UdEwEB/wQFMAMBAf8wHwYD
VR0jBBgwFoAUw/y55drH+UqELKXJalxp6m5vy2YwHQYDVR0OBBYEFBNdsd/W1DnD
AL39mW7j1/aC12aYMA4GA1UdDwEB/wQEAwIBhjAKBggqhkjOPQQDAgNJADBGAiEA
1mHrQU2CyBo7drLfKrE9HMESs6S68K6ygbCg0Ne2s8oCIQCDrtszmY75xj+BXQoQ
cJmXr+MTnHMFqchaS6xm76YNoQ==</string>
    </void>
    <void method="add">
     <string>MIIBazCCARCgAwIBAgIUZ3sE9Hadu2NmBfnpfdXtKakDktswCgYIKoZIzj0EAwIw
EjEQMA4GA1UEAwwHU2ltcGxDQTAgFw0yNTA3MjQwOTI1NTJaGA8yMDU1MDcxNzA5
MjU1MVowEjEQMA4GA1UEAwwHU2ltcGxDQTBZMBMGByqGSM49AgEGCCqGSM49AwEH
A0IABNiYmz55uyVc8hp5jeBO0I5oKSkQDP9L4gg35GYnUUOYuVMrZ3tR72AeW/+K
q6nDe3+XwZs/Svlb8wFsiyje7TSjQjBAMA8GA1UdEwEB/wQFMAMBAf8wHQYDVR0O
BBYEFMP8ueXax/lKhCylyWpcaepub8tmMA4GA1UdDwEB/wQEAwIBhjAKBggqhkjO
PQQDAgNJADBGAiEA44k+t6vJOlDY/UDXaZbHe00Zfk43X5Dn+j12x1xjYlUCIQCt
pp/MMBgeJ2BgOc+LKBd30Dqluy1fcGzBkQKeOa9UxA==</string>
    </void>
   </object>
  </void>
  <void method="put">
   <string>expiretime</string>
   <object class="java.util.Date">
    <long>2226389201000</long>
   </object>
  </void>
  <void method="put">
   <string>externalcdp</string>
   <string></string>
  </void>
  <void method="put">
   <string>alternatechains</string>
   <null/>
  </void>
 </object>
</java>
', 2226389201000, 'OnBoardingCA', NULL, 12, 1, 'CN=OnBoardingCA', 1753349247173);
INSERT INTO public.cadata VALUES (-1128220897, '<?xml version="1.0" encoding="UTF-8"?>
<java version="11.0.21" class="java.beans.XMLDecoder">
 <object class="org.cesecore.util.Base64PutHashMap">
  <void method="put">
   <string>encodedvalidity</string>
   <string>30y</string>
  </void>
  <void method="put">
   <string>crlpublishers</string>
   <object class="java.util.ArrayList"/>
  </void>
  <void method="put">
   <string>signedby</string>
   <int>1</int>
  </void>
  <void method="put">
   <string>keyvalidators</string>
   <object class="java.util.ArrayList"/>
  </void>
  <void method="put">
   <string>description</string>
   <string>SimplCAcreated using CLI</string>
  </void>
  <void method="put">
   <string>revokationreason</string>
   <int>-1</int>
  </void>
  <void method="put">
   <string>certificateprofileid</string>
   <int>858756178</int>
  </void>
  <void method="put">
   <string>usenoconflictcertificatedata</string>
   <boolean>false</boolean>
  </void>
  <void method="put">
   <string>finishuser</string>
   <boolean>true</boolean>
  </void>
  <void method="put">
   <string>includeinhealthcheck</string>
   <boolean>true</boolean>
  </void>
  <void method="put">
   <string>doEnforceUniquePublicKeys</string>
   <boolean>true</boolean>
  </void>
  <void method="put">
   <string>doEnforceKeyRenewal</string>
   <boolean>false</boolean>
  </void>
  <void method="put">
   <string>doEnforceUniqueDistinguishedName</string>
   <boolean>true</boolean>
  </void>
  <void method="put">
   <string>doEnforceUniqueSubjectDNSerialnumber</string>
   <boolean>false</boolean>
  </void>
  <void method="put">
   <string>useCertreqHistory</string>
   <boolean>false</boolean>
  </void>
  <void method="put">
   <string>useUserStorage</string>
   <boolean>true</boolean>
  </void>
  <void method="put">
   <string>useCertificateStorage</string>
   <boolean>true</boolean>
  </void>
  <void method="put">
   <string>acceptRevocationNonExistingEntry</string>
   <boolean>false</boolean>
  </void>
  <void method="put">
   <string>keepExpiredCertsOnCRL</string>
   <boolean>false</boolean>
  </void>
  <void method="put">
   <string>crlperiod</string>
   <long>7776000000</long>
  </void>
  <void method="put">
   <string>crlIssueInterval</string>
   <long>0</long>
  </void>
  <void method="put">
   <string>crlOverlapTime</string>
   <long>600000</long>
  </void>
  <void method="put">
   <string>deltacrlperiod</string>
   <long>0</long>
  </void>
  <void method="put">
   <string>generatecrluponrevocation</string>
   <boolean>false</boolean>
  </void>
  <void method="put">
   <string>allowchangingrevocationreason</string>
   <boolean>false</boolean>
  </void>
  <void method="put">
   <string>allowinvaliditydate</string>
   <boolean>false</boolean>
  </void>
  <void method="put">
   <string>extendedcaservice5</string>
   <object class="org.cesecore.util.Base64GetHashMap">
    <void method="put">
     <string>IMPLCLASS</string>
     <string>org.ejbca.core.model.ca.caadmin.extendedcaservices.KeyRecoveryCAService</string>
    </void>
    <void method="put">
     <string>extendedcaservicetype</string>
     <int>5</int>
    </void>
    <void method="put">
     <string>version</string>
     <float>1.0</float>
    </void>
    <void method="put">
     <string>status</string>
     <int>2</int>
    </void>
   </object>
  </void>
  <void method="put">
   <string>extendedcaservices</string>
   <object class="java.util.ArrayList">
    <void method="add">
     <int>5</int>
    </void>
   </object>
  </void>
  <void method="put">
   <string>approvals</string>
   <object class="java.util.LinkedHashMap"/>
  </void>
  <void method="put">
   <string>policies</string>
   <object class="java.util.ArrayList"/>
  </void>
  <void method="put">
   <string>subjectaltname</string>
   <null/>
  </void>
  <void method="put">
   <string>catype</string>
   <int>1</int>
  </void>
  <void method="put">
   <string>version</string>
   <float>25.0</float>
  </void>
  <void method="put">
   <string>msCaCompatible</string>
   <boolean>false</boolean>
  </void>
  <void method="put">
   <string>useauthoritykeyidentifier</string>
   <boolean>true</boolean>
  </void>
  <void method="put">
   <string>authoritykeyidentifiercritical</string>
   <boolean>false</boolean>
  </void>
  <void method="put">
   <string>usecrlnumber</string>
   <boolean>true</boolean>
  </void>
  <void method="put">
   <string>crlnumbercritical</string>
   <boolean>false</boolean>
  </void>
  <void method="put">
   <string>defaultcrldistpoint</string>
   <string></string>
  </void>
  <void method="put">
   <string>defaultcrlissuer</string>
   <string></string>
  </void>
  <void method="put">
   <string>cadefinedfreshestcrl</string>
   <string></string>
  </void>
  <void method="put">
   <string>defaultocspservicelocator</string>
   <string></string>
  </void>
  <void method="put">
   <string>useutf8policytext</string>
   <boolean>true</boolean>
  </void>
  <void method="put">
   <string>useprintablestringsubjectdn</string>
   <boolean>false</boolean>
  </void>
  <void method="put">
   <string>useldapdnorder</string>
   <boolean>false</boolean>
  </void>
  <void method="put">
   <string>usecrldistributionpointoncrl</string>
   <boolean>false</boolean>
  </void>
  <void method="put">
   <string>crldistributionpointoncrlcritical</string>
   <boolean>false</boolean>
  </void>
  <void method="put">
   <string>cmpraauthsecret</string>
   <string></string>
  </void>
  <void method="put">
   <string>authorityinformationaccess</string>
   <null/>
  </void>
  <void method="put">
   <string>certificateaiadefaultcaissueruri</string>
   <null/>
  </void>
  <void method="put">
   <string>nameconstraintspermitted</string>
   <null/>
  </void>
  <void method="put">
   <string>nameconstraintsexcluded</string>
   <null/>
  </void>
  <void method="put">
   <string>serialnumberoctetsize</string>
   <int>20</int>
  </void>
  <void method="put">
   <string>dopreproduceocspresponses</string>
   <boolean>false</boolean>
  </void>
  <void method="put">
   <string>dostoreocspondemand</string>
   <boolean>false</boolean>
  </void>
  <void method="put">
   <string>dopreproduceocspresponsesuponissuanceandrevocation</string>
   <boolean>false</boolean>
  </void>
  <void method="put">
   <string>usepartitionedcrl</string>
   <boolean>false</boolean>
  </void>
  <void method="put">
   <string>crlpartitions</string>
   <int>0</int>
  </void>
  <void method="put">
   <string>suspendedcrlpartitions</string>
   <int>0</int>
  </void>
  <void method="put">
   <string>requestpreprocessor</string>
   <null/>
  </void>
  <void method="put">
   <string>catoken</string>
   <object class="org.cesecore.util.Base64GetHashMap">
    <void method="put">
     <string>version</string>
     <float>8.0</float>
    </void>
    <void method="put">
     <string>cryptotokenid</string>
     <string>-836286143</string>
    </void>
    <void method="put">
     <string>propertydata</string>
     <string>certSignKey=SimplSignKey0001
keyEncryptKey=SimplEncryptKey0001
crlSignKey=SimplSignKey0001
testKey=testKey
defaultKey=SimplEncryptKey0001
</string>
    </void>
    <void method="put">
     <string>signaturealgorithm</string>
     <string>SHA256withECDSA</string>
    </void>
    <void method="put">
     <string>encryptionalgorithm</string>
     <string>SHA256withECDSA</string>
    </void>
    <void method="put">
     <string>sequence</string>
     <string>00000</string>
    </void>
    <void method="put">
     <string>sequenceformat</string>
     <int>1</int>
    </void>
   </object>
  </void>
  <void method="put">
   <string>certificatechain</string>
   <object class="java.util.ArrayList">
    <void method="add">
     <string>MIIBazCCARCgAwIBAgIUZ3sE9Hadu2NmBfnpfdXtKakDktswCgYIKoZIzj0EAwIw
EjEQMA4GA1UEAwwHU2ltcGxDQTAgFw0yNTA3MjQwOTI1NTJaGA8yMDU1MDcxNzA5
MjU1MVowEjEQMA4GA1UEAwwHU2ltcGxDQTBZMBMGByqGSM49AgEGCCqGSM49AwEH
A0IABNiYmz55uyVc8hp5jeBO0I5oKSkQDP9L4gg35GYnUUOYuVMrZ3tR72AeW/+K
q6nDe3+XwZs/Svlb8wFsiyje7TSjQjBAMA8GA1UdEwEB/wQFMAMBAf8wHQYDVR0O
BBYEFMP8ueXax/lKhCylyWpcaepub8tmMA4GA1UdDwEB/wQEAwIBhjAKBggqhkjO
PQQDAgNJADBGAiEA44k+t6vJOlDY/UDXaZbHe00Zfk43X5Dn+j12x1xjYlUCIQCt
pp/MMBgeJ2BgOc+LKBd30Dqluy1fcGzBkQKeOa9UxA==</string>
    </void>
   </object>
  </void>
  <void method="put">
   <string>expiretime</string>
   <object class="java.util.Date">
    <long>2699429151000</long>
   </object>
  </void>
  <void method="put">
   <string>externalcdp</string>
   <string></string>
  </void>
  <void method="put">
   <string>alternatechains</string>
   <null/>
  </void>
 </object>
</java>
', 2699429151000, 'SimplCA', NULL, 5, 1, 'CN=SimplCA', 1753349168204);


--
-- TOC entry 3600 (class 0 OID 16463)
-- Dependencies: 226
-- Data for Name: certificatedata; Type: TABLE DATA; Schema: public; Owner: ejbca
--

INSERT INTO public.certificatedata VALUES ('b42503b02761af26789c632442a29b62b26ae930', NULL, 'b42503b02761af26789c632442a29b62b26ae930', 0, 0, 0, 2068881768000, 'UID=c-057ronkzv4g2uwi82,CN=ManagementCA,O=EJBCA Container Quickstart', 1753348969000, -1, -1, -1, NULL, 0, '61492016621121084250757164006819012806432454218', 20, NULL, 'UID=c-057ronkzv4g2uwi82,CN=ManagementCA,O=EJBCA Container Quickstart', 'CtTkM/5I26ZTIcTJCb2ILVWAIfk=', NULL, NULL, 8, 1753348969352, 'SYSTEMCA', NULL);
INSERT INTO public.certificatedata VALUES ('930f8e25aad8e3437373c6bdfa8c169c50864515', NULL, 'b42503b02761af26789c632442a29b62b26ae930', 9, 1, 0, 1816420418000, 'UID=c-057ronkzv4g2uwi82,CN=ManagementCA,O=EJBCA Container Quickstart', 1753348969000, -1, -1, -1, NULL, 0, '686614564079957734993302586046268848974305108754', 20, 'dNSName=localhost', 'UID=c-057ronkzv4g2uwi82,CN=localhost,O=EJBCA Container Quickstart', 'Drk+XNLl081bdyqIbPymSjL+12o=', NULL, NULL, 1, 1753349019103, 'localhost', NULL);
INSERT INTO public.certificatedata VALUES ('b08049d7f246eb4691fcd3c1d78b4316ccc34cf2', NULL, 'b08049d7f246eb4691fcd3c1d78b4316ccc34cf2', 0, 0, 0, 2699429151000, 'CN=SimplCA', 1753349152000, -1, -1, -1, NULL, 0, '590769472689983888187563092841885477130379432667', 20, NULL, 'CN=SimplCA', 'w/y55drH+UqELKXJalxp6m5vy2Y=', NULL, NULL, 8, 1753349152956, 'SYSTEMCA', NULL);
INSERT INTO public.certificatedata VALUES ('1ed4d984a5a8f8dbd07694b0167e38141d1bcf0a', NULL, 'b08049d7f246eb4691fcd3c1d78b4316ccc34cf2', 0, 0, 0, 2226389201000, 'CN=SimplCA', 1753349202000, -1, -1, -1, NULL, 0, '208516676792474579346282026139147522168657583806', 20, NULL, 'CN=OnBoardingCA', 'E12x39bUOcMAvf2ZbuPX9oLXZpg=', NULL, NULL, 2, 1753349202868, 'SYSTEMCA', NULL);
INSERT INTO public.certificatedata VALUES ('6c23cf79f157fb1c5bf9aee8e6f7ed35cdc0d139', NULL, 'b42503b02761af26789c632442a29b62b26ae930', 1, 1, 0, 1816420669000, 'UID=c-057ronkzv4g2uwi82,CN=ManagementCA,O=EJBCA Container Quickstart', 1753348969000, -1, -1, -1, NULL, 0, '621446958301248330113573595163341600535725777531', 20, 'dNSName=SuperAdmin', 'CN=SuperAdmin', 'NsV9WGx+3V4K4u/oZMkMqwcICuI=', NULL, NULL, 1, 1753349270570, 'SuperAdmin', NULL);
INSERT INTO public.certificatedata VALUES ('8c6d1e40eded0b170f0d897c126334943e725bb3', NULL, 'b42503b02761af26789c632442a29b62b26ae930', 9, 1, 0, 1816420973000, 'UID=c-057ronkzv4g2uwi82,CN=ManagementCA,O=EJBCA Container Quickstart', 1753348974000, -1, -1, -1, NULL, 0, '510564289091633599270016370415946830778941064276', 20, 'dNSName=localhost', 'UID=c-057ronkzv4g2uwi82,CN=localhost,O=EJBCA Container Quickstart', 'GxGfo89tBFMkCAkCM31nEib7xAo=', NULL, NULL, 1, 1753349573741, 'localhost', NULL);


--
-- TOC entry 3601 (class 0 OID 16470)
-- Dependencies: 227
-- Data for Name: certificateprofiledata; Type: TABLE DATA; Schema: public; Owner: ejbca
--

INSERT INTO public.certificateprofiledata VALUES (626302406, 'OnBoarding Profile', '\xaced0005737200226f72672e63657365636f72652e7574696c2e426173653634476574486173684d617007156f73c047aee9020000787200176a6176612e7574696c2e4c696e6b6564486173684d617034c04e5c106cc0fb0200015a000b6163636573734f72646572787200116a6176612e7574696c2e486173684d61700507dac1c31660d103000246000a6c6f6164466163746f724900097468726573686f6c6478703f400000000000c07708000001000000007974000776657273696f6e7372000f6a6176612e6c616e672e466c6f6174daedc9a2db3cf0ec02000146000576616c7565787200106a6176612e6c616e672e4e756d62657286ac951d0b94e08b0200007870424c000074000474797065737200116a6176612e6c616e672e496e746567657212e2a0a4f781873802000149000576616c75657871007e00060000000274000b6365727476657273696f6e74000658353039763374000f656e636f64656476616c696469747974000331357974001c757365636572746966696361746576616c69646974796f6666736574737200116a6176612e6c616e672e426f6f6c65616ecd207280d59cfaee0200015a000576616c7565787000740019636572746966696361746576616c69646974796f66667365747400042d31306d74002375736565787069726174696f6e7265737472696374696f6e666f727765656b6461797371007e001174002665787069726174696f6e7265737472696374696f6e666f727765656b646179736265666f72657371007e00100174001d65787069726174696f6e7265737472696374696f6e7765656b64617973737200136a6176612e7574696c2e41727261794c6973747881d21d99c7619d03000149000473697a6578700000000777040000000771007e001671007e001671007e001171007e001171007e001171007e001671007e001678740015616c6c6f7776616c69646974796f7665727269646571007e001674000b6465736372697074696f6e740000740016616c6c6f77657874656e73696f6e6f7665727269646571007e001174000f616c6c6f77646e6f7665727269646571007e0011740014616c6c6f77646e6f76657272696465627965656971007e0011740018616c6c6f776261636b64617465647265766f6b6174696f6e71007e0011740015757365636572746966696361746573746f7261676571007e001674001473746f726563657274696669636174656461746171007e001674001373746f72657375626a656374616c746e616d6571007e00167400127573656261736963636f6e737472616e747371007e00167400186261736963636f6e73747261696e7473637269746963616c71007e00167400177573657375626a6563746b65796964656e74696669657271007e001674001c7375626a6563746b65796964656e746966696572637269746963616c71007e0011740019757365617574686f726974796b65796964656e74696669657271007e001674001e617574686f726974796b65796964656e746966696572637269746963616c71007e00117400197573657375626a656374616c7465726e61746976656e616d6571007e001674001e7375626a656374616c7465726e61746976656e616d65637269746963616c71007e0011740018757365697373756572616c7465726e61746976656e616d6571007e001674001d697373756572616c7465726e61746976656e616d65637269746963616c71007e001174001775736563726c646973747269627574696f6e706f696e7471007e001674001e75736564656661756c7463726c646973747269627574696f6e706f696e7471007e001674001c63726c646973747269627574696f6e706f696e74637269746963616c71007e001174001763726c646973747269627574696f6e706f696e7475726971007e001c74000e757365667265736865737463726c71007e00117400177573656361646566696e6564667265736865737463726c71007e001174000e667265736865737463726c75726971007e001c74000963726c69737375657271007e001c7400167573656365727469666963617465706f6c696369657371007e001174001b6365727469666963617465706f6c6963696573637269746963616c71007e00117400136365727469666963617465706f6c69636965737371007e00180000000077040000000078740016617661696c61626c656b6579616c676f726974686d737371007e001800000001770400000001740005454344534178740011617661696c61626c6565636375727665737371007e001800000001770400000001740005502d32353678740013617661696c61626c656269746c656e677468737371007e0018000000337704000000337371007e0009000000007371007e00090000006e7371007e0009000000707371007e0009000000717371007e00090000007e7371007e0009000000807371007e0009000000837371007e0009000000a07371007e0009000000a17371007e0009000000a27371007e0009000000a37371007e0009000000a77371007e0009000000ad7371007e0009000000b37371007e0009000000bd7371007e0009000000be7371007e0009000000bf7371007e0009000000c07371007e0009000000c17371007e0009000000e07371007e0009000000e17371007e0009000000e87371007e0009000000e97371007e0009000000ec7371007e0009000000ed7371007e0009000000ee7371007e0009000000ef7371007e0009000001007371007e0009000001017371007e0009000001197371007e00090000011a7371007e0009000001217371007e0009000001337371007e0009000001407371007e0009000001617371007e00090000016f7371007e0009000001807371007e0009000001977371007e0009000001997371007e0009000001a27371007e0009000001af7371007e0009000002007371007e0009000002097371007e00090000023a7371007e0009000004007371007e0009000006007371007e0009000008007371007e000900000c007371007e0009000010007371007e0009000018007371007e000900002000787400196d696e696d756d617661696c61626c656269746c656e67746871007e00427400196d6178696d756d617661696c61626c656269746c656e6774687371007e0009000020007400127369676e6174757265616c676f726974686d7074000b7573656b6579757361676571007e00167400086b657975736167657371007e00180000000977040000000971007e001671007e001171007e001171007e001171007e001171007e001671007e001671007e001171007e001178740015616c6c6f776b657975736167656f7665727269646571007e00117400106b65797573616765637269746963616c71007e0016740013757365657874656e6465646b6579757361676571007e0011740010657874656e6465646b657975736167657371007e00180000000077040000000078740018657874656e6465646b65797573616765637269746963616c71007e0011740013757365646f63756d656e74747970656c69737471007e0011740018646f63756d656e74747970656c697374637269746963616c71007e0011740010646f63756d656e74747970656c6973747371007e0018000000007704000000007874000c617661696c61626c656361737371007e0018000000017704000000017371007e0009ffffffff7874000e757365647075626c6973686572737371007e0018000000007704000000007874000e7573656f6373706e6f636865636b71007e001174000e7573656c646170646e6f7264657271007e0011740010757365637573746f6d646e6f7264657271007e00117400147573656d6963726f736f667474656d706c61746571007e00117400116d6963726f736f667474656d706c61746571007e001c7400177573656d736f626a656374736964657874656e73696f6e71007e001674000d757365636172646e756d62657271007e001174000c757365636e706f737466697871007e0011740009636e706f737466697871007e001c7400127573657375626a656374646e73756273657471007e001174000f7375626a656374646e7375627365747371007e001800000000770400000000787400177573657375626a656374616c746e616d6573756273657471007e00117400147375626a656374616c746e616d657375627365747371007e00180000000077040000000078740017757365706174686c656e677468636f6e73747261696e7471007e0011740014706174686c656e677468636f6e73747261696e7471007e004274000e757365716373746174656d656e7471007e0011740011757365706b6978716373796e746178763271007e0011740016757365716373746174656d656e74637269746963616c71007e0011740014757365716373746174656d656e7472616e616d6571007e001c74000f757365716373656d6174696373696471007e001c7400157573657163657473697163636f6d706c69616e636571007e00117400187573657163657473697369676e617475726564657669636571007e001174001375736571636574736976616c75656c696d697471007e001174001071636574736976616c75656c696d697471007e004274001371636574736976616c75656c696d697465787071007e004274001871636574736976616c75656c696d697463757272656e637971007e001c740018757365716365747369726574656e74696f6e706572696f6471007e0011740015716365747369726574656e74696f6e706572696f6471007e004274000e7573657163636f756e747269657371007e00117400107163636f756e74726965737472696e6771007e001c7400117573657163637573746f6d737472696e6771007e00117400117163637573746f6d737472696e676f696471007e001c7400127163637573746f6d737472696e677465787471007e001c7400097163657473697064737074000a716365747369747970657074002175736563657274696669636174657472616e73706172656e6379696e636572747371007e001174002075736563657274696669636174657472616e73706172656e6379696e6f63737071007e001174002575736563657274696669636174657472616e73706172656e6379696e7075626c697368657271007e00117400177573657375626a6563746469726174747269627574657371007e00117400127573656e616d65636f6e73747261696e747371007e001174001d757365617574686f72697479696e666f726d6174696f6e61636365737371007e00167400096361697373756572737371007e0018000000007704000000007874001275736564656661756c74636169737375657271007e001674001c75736564656661756c746f637370736572766963656c6f6361746f7271007e00167400156f637370736572766963656c6f6361746f7275726971007e001c74000f6376636163636573737269676874737371007e000900000003740019757365646365727469666963617465657874656e73696f6e737371007e00180000000077040000000078740009617070726f76616c737371007e00013f4000000000000c770800000010000000037e7200306f72672e63657365636f72652e6365727469666963617465732e63612e417070726f76616c526571756573745479706500000000000000001200007872000e6a6176612e6c616e672e456e756d0000000000000000120000787074000a5245564f434154494f4e71007e00887e71007e00c174000a4b45595245434f56455271007e00887e71007e00c174001041444445444954454e44454e5449545971007e0088780074001e757365707269766b65797573616765706572696f646e6f746265666f726571007e0011740015757365707269766b65797573616765706572696f6471007e001174001d757365707269766b65797573616765706572696f646e6f74616674657271007e001174001d707269766b65797573616765706572696f6473746172746f66667365747372000e6a6176612e6c616e672e4c6f6e673b8be490cc8f23df0200014a000576616c75657871007e00060000000000000000740018707269766b65797573616765706572696f646c656e6774687371007e00cd0000000003c2670074002475736573696e676c656163746976656365727469666963617465636f6e73747261696e7471007e00117400186f76657272696461626c65657874656e73696f6e6f696473737200176a6176612e7574696c2e4c696e6b656448617368536574d86cd75a95dd2a1e020000787200116a6176612e7574696c2e48617368536574ba44859596b8b7340300007870770c000000013f400000000000007874001b6e6f6e6f76657272696461626c65657874656e73696f6e6f6964737371007e00d3770c000000013f400000000000007874000d6561626e616d657370616365737371007e00d3770c000000013f4000000000000078740013616c6c6f7763657274736e6f7665727269646571007e00117400207573657472756e63617465647375626a6563746b65796964656e74696669657271007e00117400236b65797573616765666f72626964656e6379727074696f6e7573616765666f7265636371007e0011740014757365637573746f6d646e6f726465726c64617071007e00117400116e756d6f66726571617070726f76616c737371007e000900000001740010617070726f76616c73657474696e67737371007e0018000000007704000000007874000f617070726f76616c50726f66696c6571007e00887800', NULL, 0);
INSERT INTO public.certificateprofiledata VALUES (42350785, 'Onboarding TLS Profile', '\xaced0005737200226f72672e63657365636f72652e7574696c2e426173653634476574486173684d617007156f73c047aee9020000787200176a6176612e7574696c2e4c696e6b6564486173684d617034c04e5c106cc0fb0200015a000b6163636573734f72646572787200116a6176612e7574696c2e486173684d61700507dac1c31660d103000246000a6c6f6164466163746f724900097468726573686f6c6478703f400000000000c07708000001000000007a74000776657273696f6e7372000f6a6176612e6c616e672e466c6f6174daedc9a2db3cf0ec02000146000576616c7565787200106a6176612e6c616e672e4e756d62657286ac951d0b94e08b0200007870424c000074000474797065737200116a6176612e6c616e672e496e746567657212e2a0a4f781873802000149000576616c75657871007e00060000000174000b6365727476657273696f6e74000658353039763374000f656e636f64656476616c6964697479740002327974001c757365636572746966696361746576616c69646974796f6666736574737200116a6176612e6c616e672e426f6f6c65616ecd207280d59cfaee0200015a000576616c7565787000740019636572746966696361746576616c69646974796f66667365747400042d31306d74002375736565787069726174696f6e7265737472696374696f6e666f727765656b646179737371007e00100174002665787069726174696f6e7265737472696374696f6e666f727765656b646179736265666f726571007e001574001d65787069726174696f6e7265737472696374696f6e7765656b64617973737200136a6176612e7574696c2e41727261794c6973747881d21d99c7619d03000149000473697a6578700000000777040000000771007e001571007e001571007e001171007e001171007e001171007e001571007e001578740015616c6c6f7776616c69646974796f7665727269646571007e001174000b6465736372697074696f6e740000740016616c6c6f77657874656e73696f6e6f7665727269646571007e001174000f616c6c6f77646e6f7665727269646571007e0011740014616c6c6f77646e6f76657272696465627965656971007e0011740018616c6c6f776261636b64617465647265766f6b6174696f6e71007e0011740015757365636572746966696361746573746f7261676571007e001574001473746f726563657274696669636174656461746171007e001574001373746f72657375626a656374616c746e616d6571007e00157400127573656261736963636f6e737472616e747371007e00117400186261736963636f6e73747261696e7473637269746963616c71007e00157400177573657375626a6563746b65796964656e74696669657271007e001574001c7375626a6563746b65796964656e746966696572637269746963616c71007e0011740019757365617574686f726974796b65796964656e74696669657271007e001574001e617574686f726974796b65796964656e746966696572637269746963616c71007e00117400197573657375626a656374616c7465726e61746976656e616d6571007e001574001e7375626a656374616c7465726e61746976656e616d65637269746963616c71007e0011740018757365697373756572616c7465726e61746976656e616d6571007e001174001d697373756572616c7465726e61746976656e616d65637269746963616c71007e001174001775736563726c646973747269627574696f6e706f696e7471007e001574001e75736564656661756c7463726c646973747269627574696f6e706f696e7471007e001574001c63726c646973747269627574696f6e706f696e74637269746963616c71007e001174001763726c646973747269627574696f6e706f696e7475726971007e001c74000e757365667265736865737463726c71007e00117400177573656361646566696e6564667265736865737463726c71007e001174000e667265736865737463726c75726971007e001c74000963726c69737375657271007e001c7400167573656365727469666963617465706f6c696369657371007e001174001b6365727469666963617465706f6c6963696573637269746963616c71007e00117400136365727469666963617465706f6c69636965737371007e00180000000077040000000078740016617661696c61626c656b6579616c676f726974686d737371007e001800000001770400000001740005454344534178740011617661696c61626c6565636375727665737371007e001800000001770400000001740005502d32353678740013617661696c61626c656269746c656e677468737371007e0018000000337704000000337371007e0009000000007371007e00090000006e7371007e0009000000707371007e0009000000717371007e00090000007e7371007e0009000000807371007e0009000000837371007e0009000000a07371007e0009000000a17371007e0009000000a27371007e0009000000a37371007e0009000000a77371007e0009000000ad7371007e0009000000b37371007e0009000000bd7371007e0009000000be7371007e0009000000bf7371007e0009000000c07371007e0009000000c17371007e0009000000e07371007e0009000000e17371007e0009000000e87371007e0009000000e97371007e0009000000ec7371007e0009000000ed7371007e0009000000ee7371007e0009000000ef7371007e0009000001007371007e0009000001017371007e0009000001197371007e00090000011a7371007e0009000001217371007e0009000001337371007e0009000001407371007e0009000001617371007e00090000016f7371007e0009000001807371007e0009000001977371007e0009000001997371007e0009000001a27371007e0009000001af7371007e0009000002007371007e0009000002097371007e00090000023a7371007e0009000004007371007e0009000006007371007e0009000008007371007e000900000c007371007e0009000010007371007e0009000018007371007e000900002000787400196d696e696d756d617661696c61626c656269746c656e67746871007e00427400196d6178696d756d617661696c61626c656269746c656e6774687371007e0009000020007400127369676e6174757265616c676f726974686d7074000b7573656b6579757361676571007e00157400086b657975736167657371007e00180000000977040000000971007e001571007e001171007e001571007e001171007e001171007e001171007e001171007e001171007e001178740015616c6c6f776b657975736167656f7665727269646571007e00117400106b65797573616765637269746963616c71007e0015740013757365657874656e6465646b6579757361676571007e0015740010657874656e6465646b657975736167657371007e001800000002770400000002740011312e332e362e312e352e352e372e332e32740011312e332e362e312e352e352e372e332e3178740018657874656e6465646b65797573616765637269746963616c71007e0011740013757365646f63756d656e74747970656c69737471007e0011740018646f63756d656e74747970656c697374637269746963616c71007e0011740010646f63756d656e74747970656c6973747371007e0018000000007704000000007874000c617661696c61626c656361737371007e0018000000017704000000017371007e0009ffffffff7874000e757365647075626c6973686572737371007e0018000000007704000000007874000e7573656f6373706e6f636865636b71007e001174000e7573656c646170646e6f7264657271007e0011740010757365637573746f6d646e6f7264657271007e00117400147573656d6963726f736f667474656d706c61746571007e00117400116d6963726f736f667474656d706c61746571007e001c7400177573656d736f626a656374736964657874656e73696f6e71007e001574000d757365636172646e756d62657271007e001174000c757365636e706f737466697871007e0011740009636e706f737466697871007e001c7400127573657375626a656374646e73756273657471007e001174000f7375626a656374646e7375627365747371007e001800000000770400000000787400177573657375626a656374616c746e616d6573756273657471007e00117400147375626a656374616c746e616d657375627365747371007e00180000000077040000000078740017757365706174686c656e677468636f6e73747261696e7471007e0011740014706174686c656e677468636f6e73747261696e7471007e004274000e757365716373746174656d656e7471007e0011740011757365706b6978716373796e746178763271007e0011740016757365716373746174656d656e74637269746963616c71007e0011740014757365716373746174656d656e7472616e616d6571007e001c74000f757365716373656d6174696373696471007e001c7400157573657163657473697163636f6d706c69616e636571007e00117400187573657163657473697369676e617475726564657669636571007e001174001375736571636574736976616c75656c696d697471007e001174001071636574736976616c75656c696d697471007e004274001371636574736976616c75656c696d697465787071007e004274001871636574736976616c75656c696d697463757272656e637971007e001c740018757365716365747369726574656e74696f6e706572696f6471007e0011740015716365747369726574656e74696f6e706572696f6471007e004274000e7573657163636f756e747269657371007e00117400107163636f756e74726965737472696e6771007e001c7400117573657163637573746f6d737472696e6771007e00117400117163637573746f6d737472696e676f696471007e001c7400127163637573746f6d737472696e677465787471007e001c7400097163657473697064737074000a716365747369747970657074002175736563657274696669636174657472616e73706172656e6379696e636572747371007e001174002075736563657274696669636174657472616e73706172656e6379696e6f63737071007e001174002575736563657274696669636174657472616e73706172656e6379696e7075626c697368657271007e00117400177573657375626a6563746469726174747269627574657371007e00117400127573656e616d65636f6e73747261696e747371007e001174001d757365617574686f72697479696e666f726d6174696f6e61636365737371007e00157400096361697373756572737371007e0018000000007704000000007874001275736564656661756c74636169737375657271007e001574001c75736564656661756c746f637370736572766963656c6f6361746f7271007e00157400156f637370736572766963656c6f6361746f7275726971007e001c74000f6376636163636573737269676874737371007e000900000003740019757365646365727469666963617465657874656e73696f6e737371007e00180000000077040000000078740009617070726f76616c737371007e00013f4000000000000c770800000010000000037e7200306f72672e63657365636f72652e6365727469666963617465732e63612e417070726f76616c526571756573745479706500000000000000001200007872000e6a6176612e6c616e672e456e756d0000000000000000120000787074000a5245564f434154494f4e71007e008a7e71007e00c374000a4b45595245434f56455271007e008a7e71007e00c374001041444445444954454e44454e5449545971007e008a780074001e757365707269766b65797573616765706572696f646e6f746265666f726571007e0011740015757365707269766b65797573616765706572696f6471007e001174001d757365707269766b65797573616765706572696f646e6f74616674657271007e001174001d707269766b65797573616765706572696f6473746172746f66667365747372000e6a6176612e6c616e672e4c6f6e673b8be490cc8f23df0200014a000576616c75657871007e00060000000000000000740018707269766b65797573616765706572696f646c656e6774687371007e00cf0000000003c2670074002475736573696e676c656163746976656365727469666963617465636f6e73747261696e7471007e00117400186f76657272696461626c65657874656e73696f6e6f696473737200176a6176612e7574696c2e4c696e6b656448617368536574d86cd75a95dd2a1e020000787200116a6176612e7574696c2e48617368536574ba44859596b8b7340300007870770c000000013f400000000000007874001b6e6f6e6f76657272696461626c65657874656e73696f6e6f6964737371007e00d5770c000000013f400000000000007874000d6561626e616d657370616365737371007e00d5770c000000013f4000000000000078740013616c6c6f7763657274736e6f7665727269646571007e00117400207573657472756e63617465647375626a6563746b65796964656e74696669657271007e00117400236b65797573616765666f72626964656e6379727074696f6e7573616765666f7265636371007e001174001d757365636162666f7267616e697a6174696f6e6964656e74696669657271007e0011740014757365637573746f6d646e6f726465726c64617071007e00117400116e756d6f66726571617070726f76616c7371007e000a740010617070726f76616c73657474696e67737371007e0018000000007704000000007874000f617070726f76616c50726f66696c6571007e008a7800', NULL, 0);
INSERT INTO public.certificateprofiledata VALUES (858756178, 'Simpl Profile', '\xaced0005737200226f72672e63657365636f72652e7574696c2e426173653634476574486173684d617007156f73c047aee9020000787200176a6176612e7574696c2e4c696e6b6564486173684d617034c04e5c106cc0fb0200015a000b6163636573734f72646572787200116a6176612e7574696c2e486173684d61700507dac1c31660d103000246000a6c6f6164466163746f724900097468726573686f6c6478703f400000000000c07708000001000000007974000776657273696f6e7372000f6a6176612e6c616e672e466c6f6174daedc9a2db3cf0ec02000146000576616c7565787200106a6176612e6c616e672e4e756d62657286ac951d0b94e08b0200007870424c000074000474797065737200116a6176612e6c616e672e496e746567657212e2a0a4f781873802000149000576616c75657871007e00060000000874000b6365727476657273696f6e74000658353039763374000f656e636f64656476616c696469747974000333307974001c757365636572746966696361746576616c69646974796f6666736574737200116a6176612e6c616e672e426f6f6c65616ecd207280d59cfaee0200015a000576616c7565787000740019636572746966696361746576616c69646974796f66667365747400042d31306d74002375736565787069726174696f6e7265737472696374696f6e666f727765656b6461797371007e001174002665787069726174696f6e7265737472696374696f6e666f727765656b646179736265666f72657371007e00100174001d65787069726174696f6e7265737472696374696f6e7765656b64617973737200136a6176612e7574696c2e41727261794c6973747881d21d99c7619d03000149000473697a6578700000000777040000000771007e001671007e001671007e001171007e001171007e001171007e001671007e001678740015616c6c6f7776616c69646974796f7665727269646571007e001674000b6465736372697074696f6e740000740016616c6c6f77657874656e73696f6e6f7665727269646571007e001174000f616c6c6f77646e6f7665727269646571007e0011740014616c6c6f77646e6f76657272696465627965656971007e0011740018616c6c6f776261636b64617465647265766f6b6174696f6e71007e0011740015757365636572746966696361746573746f7261676571007e001674001473746f726563657274696669636174656461746171007e001674001373746f72657375626a656374616c746e616d6571007e00167400127573656261736963636f6e737472616e747371007e00167400186261736963636f6e73747261696e7473637269746963616c71007e00167400177573657375626a6563746b65796964656e74696669657271007e001674001c7375626a6563746b65796964656e746966696572637269746963616c71007e0011740019757365617574686f726974796b65796964656e74696669657271007e001174001e617574686f726974796b65796964656e746966696572637269746963616c71007e00117400197573657375626a656374616c7465726e61746976656e616d6571007e001174001e7375626a656374616c7465726e61746976656e616d65637269746963616c71007e0011740018757365697373756572616c7465726e61746976656e616d6571007e001174001d697373756572616c7465726e61746976656e616d65637269746963616c71007e001174001775736563726c646973747269627574696f6e706f696e7471007e001174001e75736564656661756c7463726c646973747269627574696f6e706f696e7471007e001174001c63726c646973747269627574696f6e706f696e74637269746963616c71007e001174001763726c646973747269627574696f6e706f696e7475726971007e001c74000e757365667265736865737463726c71007e00117400177573656361646566696e6564667265736865737463726c71007e001174000e667265736865737463726c75726971007e001c74000963726c69737375657271007e001c7400167573656365727469666963617465706f6c696369657371007e001174001b6365727469666963617465706f6c6963696573637269746963616c71007e00117400136365727469666963617465706f6c69636965737371007e00180000000077040000000078740016617661696c61626c656b6579616c676f726974686d737371007e001800000001770400000001740005454344534178740011617661696c61626c6565636375727665737371007e001800000001770400000001740005502d32353678740013617661696c61626c656269746c656e677468737371007e0018000000337704000000337371007e0009000000007371007e00090000006e7371007e0009000000707371007e0009000000717371007e00090000007e7371007e0009000000807371007e0009000000837371007e0009000000a07371007e0009000000a17371007e0009000000a27371007e0009000000a37371007e0009000000a77371007e0009000000ad7371007e0009000000b37371007e0009000000bd7371007e0009000000be7371007e0009000000bf7371007e0009000000c07371007e0009000000c17371007e0009000000e07371007e0009000000e17371007e0009000000e87371007e0009000000e97371007e0009000000ec7371007e0009000000ed7371007e0009000000ee7371007e0009000000ef7371007e0009000001007371007e0009000001017371007e0009000001197371007e00090000011a7371007e0009000001217371007e0009000001337371007e0009000001407371007e0009000001617371007e00090000016f7371007e0009000001807371007e0009000001977371007e0009000001997371007e0009000001a27371007e0009000001af7371007e0009000002007371007e0009000002097371007e00090000023a7371007e0009000004007371007e0009000006007371007e0009000008007371007e000900000c007371007e0009000010007371007e0009000018007371007e000900002000787400196d696e696d756d617661696c61626c656269746c656e67746871007e00427400196d6178696d756d617661696c61626c656269746c656e6774687371007e0009000020007400127369676e6174757265616c676f726974686d7074000b7573656b6579757361676571007e00167400086b657975736167657371007e00180000000977040000000971007e001671007e001171007e001171007e001171007e001171007e001671007e001671007e001171007e001178740015616c6c6f776b657975736167656f7665727269646571007e00117400106b65797573616765637269746963616c71007e0016740013757365657874656e6465646b6579757361676571007e0011740010657874656e6465646b657975736167657371007e00180000000077040000000078740018657874656e6465646b65797573616765637269746963616c71007e0011740013757365646f63756d656e74747970656c69737471007e0011740018646f63756d656e74747970656c697374637269746963616c71007e0011740010646f63756d656e74747970656c6973747371007e0018000000007704000000007874000c617661696c61626c656361737371007e0018000000017704000000017371007e0009ffffffff7874000e757365647075626c6973686572737371007e0018000000007704000000007874000e7573656f6373706e6f636865636b71007e001174000e7573656c646170646e6f7264657271007e0011740010757365637573746f6d646e6f7264657271007e00117400147573656d6963726f736f667474656d706c61746571007e00117400116d6963726f736f667474656d706c61746571007e001c7400177573656d736f626a656374736964657874656e73696f6e71007e001674000d757365636172646e756d62657271007e001174000c757365636e706f737466697871007e0011740009636e706f737466697871007e001c7400127573657375626a656374646e73756273657471007e001174000f7375626a656374646e7375627365747371007e001800000000770400000000787400177573657375626a656374616c746e616d6573756273657471007e00117400147375626a656374616c746e616d657375627365747371007e00180000000077040000000078740017757365706174686c656e677468636f6e73747261696e7471007e0011740014706174686c656e677468636f6e73747261696e7471007e004274000e757365716373746174656d656e7471007e0011740011757365706b6978716373796e746178763271007e0011740016757365716373746174656d656e74637269746963616c71007e0011740014757365716373746174656d656e7472616e616d6571007e001c74000f757365716373656d6174696373696471007e001c7400157573657163657473697163636f6d706c69616e636571007e00117400187573657163657473697369676e617475726564657669636571007e001174001375736571636574736976616c75656c696d697471007e001174001071636574736976616c75656c696d697471007e004274001371636574736976616c75656c696d697465787071007e004274001871636574736976616c75656c696d697463757272656e637971007e001c740018757365716365747369726574656e74696f6e706572696f6471007e0011740015716365747369726574656e74696f6e706572696f6471007e004274000e7573657163636f756e747269657371007e00117400107163636f756e74726965737472696e6771007e001c7400117573657163637573746f6d737472696e6771007e00117400117163637573746f6d737472696e676f696471007e001c7400127163637573746f6d737472696e677465787471007e001c7400097163657473697064737074000a716365747369747970657074002175736563657274696669636174657472616e73706172656e6379696e636572747371007e001174002075736563657274696669636174657472616e73706172656e6379696e6f63737071007e001174002575736563657274696669636174657472616e73706172656e6379696e7075626c697368657271007e00117400177573657375626a6563746469726174747269627574657371007e00117400127573656e616d65636f6e73747261696e747371007e001174001d757365617574686f72697479696e666f726d6174696f6e61636365737371007e00117400096361697373756572737371007e0018000000007704000000007874001275736564656661756c74636169737375657271007e001174001c75736564656661756c746f637370736572766963656c6f6361746f7271007e00117400156f637370736572766963656c6f6361746f7275726971007e001c74000f6376636163636573737269676874737371007e000900000003740019757365646365727469666963617465657874656e73696f6e737371007e00180000000077040000000078740009617070726f76616c737371007e00013f4000000000000c770800000010000000037e7200306f72672e63657365636f72652e6365727469666963617465732e63612e417070726f76616c526571756573745479706500000000000000001200007872000e6a6176612e6c616e672e456e756d0000000000000000120000787074000a5245564f434154494f4e71007e00887e71007e00c174000a4b45595245434f56455271007e00887e71007e00c174001041444445444954454e44454e5449545971007e0088780074001e757365707269766b65797573616765706572696f646e6f746265666f726571007e0011740015757365707269766b65797573616765706572696f6471007e001174001d757365707269766b65797573616765706572696f646e6f74616674657271007e001174001d707269766b65797573616765706572696f6473746172746f66667365747372000e6a6176612e6c616e672e4c6f6e673b8be490cc8f23df0200014a000576616c75657871007e00060000000000000000740018707269766b65797573616765706572696f646c656e6774687371007e00cd0000000003c2670074002475736573696e676c656163746976656365727469666963617465636f6e73747261696e7471007e00117400186f76657272696461626c65657874656e73696f6e6f696473737200176a6176612e7574696c2e4c696e6b656448617368536574d86cd75a95dd2a1e020000787200116a6176612e7574696c2e48617368536574ba44859596b8b7340300007870770c000000013f400000000000007874001b6e6f6e6f76657272696461626c65657874656e73696f6e6f6964737371007e00d3770c000000013f400000000000007874000d6561626e616d657370616365737371007e00d3770c000000013f4000000000000078740013616c6c6f7763657274736e6f7665727269646571007e00117400207573657472756e63617465647375626a6563746b65796964656e74696669657271007e00117400236b65797573616765666f72626964656e6379727074696f6e7573616765666f7265636371007e0011740014757365637573746f6d646e6f726465726c64617071007e00117400116e756d6f66726571617070726f76616c737371007e000900000001740010617070726f76616c73657474696e67737371007e0018000000007704000000007874000f617070726f76616c50726f66696c6571007e00887800', NULL, 0);


--
-- TOC entry 3599 (class 0 OID 16456)
-- Dependencies: 225
-- Data for Name: certreqhistorydata; Type: TABLE DATA; Schema: public; Owner: ejbca
--



--
-- TOC entry 3598 (class 0 OID 16449)
-- Dependencies: 224
-- Data for Name: crldata; Type: TABLE DATA; Schema: public; Owner: ejbca
--

INSERT INTO public.crldata VALUES ('1db8b1f83b96bd66c9cb45580165578a476aced1', 'MIICWzCBxAIBATANBgkqhkiG9w0BAQsFADBhMSMwIQYKCZImiZPyLGQBAQwTYy0w
NTdyb25renY0ZzJ1d2k4MjEVMBMGA1UEAwwMTWFuYWdlbWVudENBMSMwIQYDVQQK
DBpFSkJDQSBDb250YWluZXIgUXVpY2tzdGFydBcNMjUwNzI0MDkyMjQ5WhcNMjUw
NzI1MDkyMjQ4WqAvMC0wHwYDVR0jBBgwFoAUCtTkM/5I26ZTIcTJCb2ILVWAIfkw
CgYDVR0UBAMCAQEwDQYJKoZIhvcNAQELBQADggGBAJDY6vuuXzhPoU0VR5RC8X7z
a79C4lWeJqgMJravpmlTIRcqSd3zz/fSz7lTbblIGJ7LsmUSt1cmXiEjfMyvQSc4
TwnaCrc7FEoj8smhgua22xpRcxnucWICFNvSh+1YFASuSi6uq/MezrQXKygnzcyk
9dWEzi1LHy4GG/q9lg8nBEwcx8iX4E+QwVikaX9nqMlzxkMmiR77aCk0XGKmG0sX
vx46awFxodRZiWG08Rt+nwkeJ3uAGsKymsb+ZkoABdIowMa94K9fFxwx5ydUnxC8
njTE8tmpd8EH2lQvy2Vx6SNq84fFvhl4lpsDzTqBsSJ4Kqtk2Az8oQ967EzEEiRO
ebrHDPlqyd/WfzAcQy3/2sToNPwTZDjmadYeCCdQqXTiZwnvoe+0PDOPfw+GIZpW
NO5krtZKYsi0Q/j0rLHOdrm84e0NWxdlCv2R1laPsR1cWz8enTAEih62KVcNvYPj
QoURZTAXKbTdiI1wS03FAabV8E8XJivMVj/6cM3fbw==', 'b42503b02761af26789c632442a29b62b26ae930', -1, 1, -1, 'UID=c-057ronkzv4g2uwi82,CN=ManagementCA,O=EJBCA Container Quickstart', 1753435368000, NULL, 0, 1753348969000);
INSERT INTO public.crldata VALUES ('6429df5f862124b2b272c7bce9dfeaea1327645e', 'MIHKMHICAQEwCgYIKoZIzj0EAwIwEjEQMA4GA1UEAwwHU2ltcGxDQRcNMjUwNzI0
MDkyNTUyWhcNMjUwNzI1MDkyNTUxWqAvMC0wHwYDVR0jBBgwFoAUw/y55drH+UqE
LKXJalxp6m5vy2YwCgYDVR0UBAMCAQEwCgYIKoZIzj0EAwIDSAAwRQIgW1bpPvdt
V+vWrSzksqiOnwlTzhDbXEqYOi7HNzVkhfkCIQD7f1e1qwAtsM2JXK/UxPTQIljY
hi3zmVxsaRR7YJpKqA==', 'b08049d7f246eb4691fcd3c1d78b4316ccc34cf2', -1, 1, -1, 'CN=SimplCA', 1753435551000, NULL, 0, 1753349152000);
INSERT INTO public.crldata VALUES ('3c2336d930844f2600bbedf241f151537ccc20f6', 'MIHPMHcCAQEwCgYIKoZIzj0EAwIwFzEVMBMGA1UEAwwMT25Cb2FyZGluZ0NBFw0y
NTA3MjQwOTI2NDJaFw0yNTA3MjUwOTI2NDFaoC8wLTAfBgNVHSMEGDAWgBQTXbHf
1tQ5wwC9/Zlu49f2gtdmmDAKBgNVHRQEAwIBATAKBggqhkjOPQQDAgNIADBFAiAp
zJQnhf397A+cNDQlEr4h+8cIExW2sViryvo0MaqCFwIhAMmmrsTLagERwti/Oblq
OkzaxNat1X5sGQy9Q0g8mUPe', '1ed4d984a5a8f8dbd07694b0167e38141d1bcf0a', -1, 1, -1, 'CN=OnBoardingCA', 1753435601000, NULL, 0, 1753349202000);


--
-- TOC entry 3602 (class 0 OID 16477)
-- Dependencies: 228
-- Data for Name: cryptotokendata; Type: TABLE DATA; Schema: public; Owner: ejbca
--

INSERT INTO public.cryptotokendata VALUES (1726589353, 1753348968003, NULL, 2, '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', 'ManagementCA', 'I1RodSBKdWwgMjQgMDk6MjI6NDggVVRDIDIwMjUKcGluPTZiYzg0MWIyNzQ1ZTJjOTVlMDQyYTY4YjQ3NzdiMzRjCnRva2VuTmFtZT1NYW5hZ2VtZW50Q0EK', 'SoftCryptoToken');
INSERT INTO public.cryptotokendata VALUES (-836286143, 1753349145969, NULL, 3, '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', 'SimplCryptoToken', 'I1RodSBKdWwgMjQgMDk6MjU6NDYgVVRDIDIwMjUKdG9rZW5OYW1lPVNpbXBsQ3J5cHRvVG9rZW4KTk9ERUZBVUxUUFdEPXRydWUKcGluPTBlZTlkN2MxNzc1OWU1ZjdhNTAzZjdmOWRhZmEwYWU0CmFsbG93LmV4dHJhY3RhYmxlLnByaXZhdGVrZXk9ZmFsc2UK', 'SoftCryptoToken');
INSERT INTO public.cryptotokendata VALUES (-673066697, 1753349196494, NULL, 3, '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', 'SimplOnboardingCryptoToken', 'I1RodSBKdWwgMjQgMDk6MjY6MzYgVVRDIDIwMjUKdG9rZW5OYW1lPVNpbXBsT25ib2FyZGluZ0NyeXB0b1Rva2VuCk5PREVGQVVMVFBXRD10cnVlCnBpbj0wZWU5ZDdjMTc3NTllNWY3YTUwM2Y3ZjlkYWZhMGFlNAphbGxvdy5leHRyYWN0YWJsZS5wcml2YXRla2V5PWZhbHNlCg==', 'SoftCryptoToken');


--
-- TOC entry 3603 (class 0 OID 16484)
-- Dependencies: 229
-- Data for Name: endentityprofiledata; Type: TABLE DATA; Schema: public; Owner: ejbca
--

INSERT INTO public.endentityprofiledata VALUES (260143312, '\xaced0005737200226f72672e63657365636f72652e7574696c2e426173653634476574486173684d617007156f73c047aee9020000787200176a6176612e7574696c2e4c696e6b6564486173684d617034c04e5c106cc0fb0200015a000b6163636573734f72646572787200116a6176612e7574696c2e486173684d61700507dac1c31660d103000246000a6c6f6164466163746f724900097468726573686f6c6478703f400000000000c0770800000100000000a974000776657273696f6e7372000f6a6176612e6c616e672e466c6f6174daedc9a2db3cf0ec02000146000576616c7565787200106a6176612e6c616e672e4e756d62657286ac951d0b94e08b02000078704190000074000b4e554d4245524152524159737200136a6176612e7574696c2e41727261794c6973747881d21d99c7619d03000149000473697a6578700000006e77040000006e737200116a6176612e6c616e672e496e746567657212e2a0a4f781873802000149000576616c75657871007e00060000000171007e000c7371007e000b0000000071007e000d71007e000d71007e000c71007e000d71007e000d71007e000d71007e000d71007e000d71007e000c71007e000c71007e000d71007e000d71007e000d71007e000c71007e000d71007e000c71007e000d71007e000d71007e000d71007e000d71007e000d71007e000d71007e000d71007e000c71007e000d71007e000d71007e000c71007e000c71007e000c71007e000c71007e000c71007e000c71007e000d71007e000d71007e000c71007e000c71007e000d71007e000d71007e000d71007e000d71007e000d71007e000d71007e000d71007e000d71007e000d71007e000d71007e000d71007e000d71007e000d71007e000d71007e000d71007e000d71007e000d71007e000d71007e000d71007e000d71007e000d71007e000d71007e000d71007e000d71007e000d71007e000d71007e000d71007e000d71007e000d71007e000d71007e000d71007e000d71007e000d71007e000d71007e000d71007e000d71007e000d71007e000d71007e000d71007e000d71007e000d71007e000d71007e000d71007e000d71007e000d71007e000d71007e000d71007e000c71007e000c71007e000c71007e000c71007e000d71007e000c71007e000d71007e000c71007e000c71007e000c71007e000c71007e000c71007e000c71007e000c71007e000d71007e000d71007e000d71007e000d71007e000d71007e000d71007e000d71007e000d71007e000d71007e000d787400135355424a454354444e4649454c444f524445527371007e0009000000047704000000047371007e000b0000c3507371007e000b0001d4c07371007e000b0001adb07371007e000b00027100787400185355424a454354414c544e414d454649454c444f524445527371007e0009000000017704000000017371007e000b0002bf20787400185355424a454354444952415454524649454c444f524445527371007e0009000000007704000000007874000f5353485f4649454c445f4f524445527371007e0009000000007704000000007874000b50524f46494c455459504571007e000c71007e000d7400007371007e000b001e8480737200116a6176612e6c616e672e426f6f6c65616ecd207280d59cfaee0200015a000576616c75657870017371007e000b000f424071007e001f7371007e000b002dc6c071007e001f7371007e000b004c4b407371007e001e0071007e000c71007e001c7371007e000b001e848171007e001f7371007e000b000f424171007e001f7371007e000b002dc6c171007e001f7371007e000b004c4b4171007e00237371007e000b0000005f71007e001c7371007e000b001e84df71007e00237371007e000b000f429f71007e001f7371007e000b002dc71f71007e001f7371007e000b004c4b9f71007e00237371007e000b00000060740001387371007e000b001e84e071007e00237371007e000b000f42a071007e001f7371007e000b002dc72071007e001f7371007e000b004c4ba071007e00237371007e000b0000000571007e001c7371007e000b001e848571007e001f7371007e000b000f424571007e001f7371007e000b002dc6c571007e001f7371007e000b004c4b4571007e00237371007e000b0000001a71007e001c7371007e000b001e849a71007e00237371007e000b000f425a71007e00237371007e000b002dc6da71007e001f7371007e000b004c4b5a71007e00237371007e000b0000001d74000834323335303738357371007e000b001e849d71007e001f7371007e000b000f425d71007e001f7371007e000b002dc6dd71007e001f7371007e000b004c4b5d71007e00237371007e000b0000001e74000834323335303738357371007e000b001e849e71007e001f7371007e000b000f425e71007e001f7371007e000b002dc6de71007e001f7371007e000b004c4b5e71007e00237371007e000b0000001f740001317371007e000b001e849f71007e001f7371007e000b000f425f71007e001f7371007e000b002dc6df71007e001f7371007e000b004c4b5f71007e00237371007e000b00000020740009313b323b353b333b347371007e000b001e84a071007e001f7371007e000b000f426071007e001f7371007e000b002dc6e071007e001f7371007e000b004c4b6071007e00237371007e000b0000002171007e001c7371007e000b001e84a171007e00237371007e000b000f426171007e001f7371007e000b002dc6e171007e001f7371007e000b004c4b6171007e00237371007e000b0000002271007e001c7371007e000b001e84a271007e001f7371007e000b000f426271007e00237371007e000b002dc6e271007e001f7371007e000b004c4b6271007e00237371007e000b00000026740001317371007e000b001e84a671007e001f7371007e000b000f426671007e001f7371007e000b002dc6e671007e001f7371007e000b004c4b6671007e00237371007e000b000000257400022d317371007e000b001e84a571007e001f7371007e000b000f426571007e001f7371007e000b002dc6e571007e001f7371007e000b004c4b6571007e00237371007e000b0000006271007e001c7371007e000b001e84e271007e00237371007e000b000f42a271007e00237371007e000b002dc72271007e001f7371007e000b004c4ba271007e00237371007e000b0000006371007e001c7371007e000b001e84e371007e00237371007e000b000f42a371007e00237371007e000b002dc72371007e001f7371007e000b004c4ba371007e00237371007e000b0000006171007e001c7371007e000b001e84e171007e00237371007e000b000f42a171007e00237371007e000b002dc72171007e001f7371007e000b004c4ba171007e00237371007e000b0000005b71007e001c7371007e000b001e84db71007e00237371007e000b000f429b71007e00237371007e000b002dc71b71007e001f7371007e000b004c4b9b71007e00237371007e000b0000005e7400022d317371007e000b001e84de71007e00237371007e000b000f429e71007e00237371007e000b002dc71e71007e001f7371007e000b004c4b9e71007e00237371007e000b0000005d7400022d317371007e000b001e84dd71007e00237371007e000b000f429d71007e00237371007e000b002dc71d71007e001f7371007e000b004c4b9d71007e00237371007e000b0000005971007e001c7371007e000b001e84d971007e00237371007e000b000f429971007e00237371007e000b002dc71971007e001f7371007e000b004c4b9971007e00237371007e000b0000005871007e001c7371007e000b001e84d871007e00237371007e000b000f429871007e00237371007e000b002dc71871007e001f7371007e000b004c4b9871007e00237371007e000b0000005771007e001c7371007e000b001e84d771007e00237371007e000b000f429771007e00237371007e000b002dc71771007e001f7371007e000b004c4b9771007e00237371007e000b00000056740001377371007e000b001e84d671007e00237371007e000b000f429671007e00237371007e000b002dc71671007e00237371007e000b004c4b9671007e00237371007e000b002dc78971007e001f7371007e000b002dc78a71007e001f7371007e000b002dc78b71007e001f7371007e000b0000000c71007e001c7371007e000b001e848c71007e001f7371007e000b000f424c71007e001f7371007e000b002dc6cc71007e001f7371007e000b004c4b4c71007e00237371007e000b000f429a71007e001f7371007e000b0000005a740001307371007e000b000f424271007e00237371007e000b0000000274000566616c73657371007e000b001e848271007e00237371007e000b0000006e71007e001c74001352455645525345464649454c44434845434b5371007e002374000d414c4c4f575f4d45524745444e71007e0023740016414c4c4f575f4d554c54495f56414c55455f52444e5371007e00237371007e000b000f429c71007e0023740010555345455854454e53494f4e4441544171007e002374000f50534432514353544154454d454e5471007e00237371007e000b000f426371007e0023740011555345524e4f54494649434154494f4e537371007e000900000000770400000000787371007e000b000f425c71007e00237371007e000b001e849c71007e00237371007e000b0000001c74000566616c73657371007e000b001e84a371007e00237371007e000b0000002374000566616c73657371007e000b0000001271007e001c7371007e000b001e849271007e001f7371007e000b000f425271007e001f7371007e000b002dc6d271007e001f7371007e000b004c4b5271007e00237371007e000b0000000b71007e001c7371007e000b001e848b71007e001f7371007e000b000f424b71007e001f7371007e000b002dc6cb71007e001f7371007e000b004c4b4b71007e00237371007e000b0000001071007e001c7371007e000b001e849071007e001f7371007e000b000f425071007e001f7371007e000b002dc6d071007e001f7371007e000b004c4b5071007e00237800', 'Onboarding End Entity', NULL, 0);


--
-- TOC entry 3604 (class 0 OID 16491)
-- Dependencies: 230
-- Data for Name: globalconfigurationdata; Type: TABLE DATA; Schema: public; Owner: ejbca
--

INSERT INTO public.globalconfigurationdata VALUES ('0', '\xaced0005737200176a6176612e7574696c2e4c696e6b6564486173684d617034c04e5c106cc0fb0200015a000b6163636573734f72646572787200116a6176612e7574696c2e486173684d61700507dac1c31660d103000246000a6c6f6164466163746f724900097468726573686f6c6478703f4000000000000c7708000000100000000974000776657273696f6e7372000f6a6176612e6c616e672e466c6f6174daedc9a2db3cf0ec02000146000576616c7565787200106a6176612e6c616e672e4e756d62657286ac951d0b94e08b0200007870404000007400057469746c65740014454a4243412041646d696e697374726174696f6e74000a6865616462616e6e657274002061646d696e7765622f62616e6e6572732f686561645f62616e6e65722e6a737074000a666f6f7462616e6e65727400182f62616e6e6572732f666f6f745f62616e6e65722e6a737074001b656e64656e7469747970726f66696c656c696d69746174696f6e73737200116a6176612e6c616e672e426f6f6c65616ecd207280d59cfaee0200015a000576616c756578700174001661757468656e7469636174656475736572736f6e6c797371007e000e00740011656e61626c656b65797265636f7665727971007e0011740016656e61626c656963616f63616e616d656368616e676571007e0011740012737461746564756d705f6c6f636b646f776e71007e00117800', NULL, 0);
INSERT INTO public.globalconfigurationdata VALUES ('AVAILABLE_PROTOCOLS', '\xaced0005737200176a6176612e7574696c2e4c696e6b6564486173684d617034c04e5c106cc0fb0200015a000b6163636573734f72646572787200116a6176612e7574696c2e486173684d61700507dac1c31660d103000246000a6c6f6164466163746f724900097468726573686f6c6478703f400000000000067708000000080000000374000776657273696f6e7372000f6a6176612e6c616e672e466c6f6174daedc9a2db3cf0ec02000146000576616c7565787200106a6176612e6c616e672e4e756d62657286ac951d0b94e08b02000078704040000074001b52455354204365727469666963617465204d616e6167656d656e74737200116a6176612e6c616e672e426f6f6c65616ecd207280d59cfaee0200015a000576616c756578700174001e52455354204365727469666963617465204d616e6167656d656e7420563271007e00097800', NULL, 1);
INSERT INTO public.globalconfigurationdata VALUES ('UPGRADE', '\xaced0005737200226f72672e63657365636f72652e7574696c2e426173653634476574486173684d617007156f73c047aee9020000787200176a6176612e7574696c2e4c696e6b6564486173684d617034c04e5c106cc0fb0200015a000b6163636573734f72646572787200116a6176612e7574696c2e486173684d61700507dac1c31660d103000246000a6c6f6164466163746f724900097468726573686f6c6478703f4000000000000c7708000000100000000674000776657273696f6e7372000f6a6176612e6c616e672e466c6f6174daedc9a2db3cf0ec02000146000576616c7565787200106a6176612e6c616e672e4e756d62657286ac951d0b94e08b020000787040400000740021656e64456e7469747950726f66696c65496e4365727469666963617465446174617400047472756574001e76616c6964697479576974685365636f6e64734772616e756c617269747971007e0009740013757067726164656446726f6d56657273696f6e740007382e322e302e317400117570677261646564546f56657273696f6e71007e000c740015706f73745570677261646564546f56657273696f6e740006372e31312e307800', NULL, 4);


--
-- TOC entry 3625 (class 0 OID 16640)
-- Dependencies: 251
-- Data for Name: incompleteissuancejournaldata; Type: TABLE DATA; Schema: public; Owner: ejbca
--



--
-- TOC entry 3605 (class 0 OID 16498)
-- Dependencies: 231
-- Data for Name: internalkeybindingdata; Type: TABLE DATA; Schema: public; Owner: ejbca
--



--
-- TOC entry 3606 (class 0 OID 16505)
-- Dependencies: 232
-- Data for Name: keyrecoverydata; Type: TABLE DATA; Schema: public; Owner: ejbca
--



--
-- TOC entry 3617 (class 0 OID 16584)
-- Dependencies: 243
-- Data for Name: noconflictcertificatedata; Type: TABLE DATA; Schema: public; Owner: ejbca
--



--
-- TOC entry 3624 (class 0 OID 16633)
-- Dependencies: 250
-- Data for Name: ocspresponsedata; Type: TABLE DATA; Schema: public; Owner: ejbca
--



--
-- TOC entry 3607 (class 0 OID 16513)
-- Dependencies: 233
-- Data for Name: peerdata; Type: TABLE DATA; Schema: public; Owner: ejbca
--



--
-- TOC entry 3608 (class 0 OID 16520)
-- Dependencies: 234
-- Data for Name: profiledata; Type: TABLE DATA; Schema: public; Owner: ejbca
--



--
-- TOC entry 3609 (class 0 OID 16527)
-- Dependencies: 235
-- Data for Name: publisherdata; Type: TABLE DATA; Schema: public; Owner: ejbca
--



--
-- TOC entry 3610 (class 0 OID 16534)
-- Dependencies: 236
-- Data for Name: publisherqueuedata; Type: TABLE DATA; Schema: public; Owner: ejbca
--



--
-- TOC entry 3612 (class 0 OID 16548)
-- Dependencies: 238
-- Data for Name: roledata; Type: TABLE DATA; Schema: public; Owner: ejbca
--

INSERT INTO public.roledata VALUES (1, 'Super Administrator Role', NULL, '<?xml version="1.0" encoding="UTF-8"?>
<java version="11.0.21" class="java.beans.XMLDecoder">
 <object class="org.cesecore.util.Base64PutHashMap">
  <void method="put">
   <string>version</string>
   <float>1.0</float>
  </void>
  <void method="put">
   <string>accessRules</string>
   <object class="java.util.LinkedHashMap">
    <void method="put">
     <string>/</string>
     <boolean>true</boolean>
    </void>
   </object>
  </void>
 </object>
</java>
', NULL, 0);
INSERT INTO public.roledata VALUES (1029167557, 'Public Access Role', NULL, '<?xml version="1.0" encoding="UTF-8"?>
<java version="11.0.21" class="java.beans.XMLDecoder">
 <object class="org.cesecore.util.Base64PutHashMap">
  <void method="put">
   <string>version</string>
   <float>1.0</float>
  </void>
  <void method="put">
   <string>associatedCss</string>
   <int>0</int>
  </void>
  <void method="put">
   <string>accessRules</string>
   <object class="java.util.LinkedHashMap">
    <void method="put">
     <string>/ca_functionality/view_ca/</string>
     <boolean>true</boolean>
    </void>
    <void method="put">
     <string>/ca_functionality/create_certificate/</string>
     <boolean>true</boolean>
    </void>
    <void method="put">
     <string>/ca_functionality/use_username/</string>
     <boolean>true</boolean>
    </void>
    <void method="put">
     <string>/endentityprofilesrules/1/</string>
     <boolean>true</boolean>
    </void>
    <void method="put">
     <string>/ca/-968899527/</string>
     <boolean>true</boolean>
    </void>
   </object>
  </void>
 </object>
</java>
', NULL, 1);


--
-- TOC entry 3613 (class 0 OID 16555)
-- Dependencies: 239
-- Data for Name: rolememberdata; Type: TABLE DATA; Schema: public; Owner: ejbca
--

INSERT INTO public.rolememberdata VALUES (1274961314, 'CliAuthenticationToken', 0, 0, 0, 1000, 'ejbca', 1, NULL, NULL, 0);
INSERT INTO public.rolememberdata VALUES (1725839478, 'PublicAccessAuthenticationToken', 0, 0, 2, 0, NULL, 1029167557, NULL, NULL, 0);
INSERT INTO public.rolememberdata VALUES (1879122637, 'PublicAccessAuthenticationToken', 0, 0, 2, 0, NULL, 1, 'Initial RoleMember.', NULL, 0);
INSERT INTO public.rolememberdata VALUES (470695368, 'CertificateAuthenticationToken', -968899527, 0, 8, 1000, 'SuperAdmin', 1, NULL, NULL, 0);


--
-- TOC entry 3623 (class 0 OID 16626)
-- Dependencies: 249
-- Data for Name: sctdata; Type: TABLE DATA; Schema: public; Owner: ejbca
--



--
-- TOC entry 3614 (class 0 OID 16563)
-- Dependencies: 240
-- Data for Name: servicedata; Type: TABLE DATA; Schema: public; Owner: ejbca
--



--
-- TOC entry 3615 (class 0 OID 16570)
-- Dependencies: 241
-- Data for Name: userdata; Type: TABLE DATA; Schema: public; Owner: ejbca
--

INSERT INTO public.userdata VALUES ('ejbca', 0, NULL, 0, NULL, 1, NULL, 0, NULL, '$2a$01$LWOZhBSRI0wlbOT6wM//e.E.liIRmQ8THaqtQBPYlJFukdLA3rjVO', NULL, 0, 40, NULL, 'UID=ejbca', NULL, 1753348934689, 1753348934689, 0, 0);
INSERT INTO public.userdata VALUES ('SuperAdmin', -968899527, NULL, 1, '', 1, NULL, 0, NULL, '', NULL, 4, 40, 'dNSName=SuperAdmin', 'CN=SuperAdmin', NULL, 1753349252601, 1753349271007, 2, 1);
INSERT INTO public.userdata VALUES ('localhost', -968899527, NULL, 9, '', 1, NULL, 0, NULL, '', NULL, 8, 40, 'dnsName=localhost', 'UID=c-057ronkzv4g2uwi82,CN=localhost,O=EJBCA Container Quickstart', NULL, 1753348994845, 1753349575039, 3, 1);


--
-- TOC entry 3616 (class 0 OID 16577)
-- Dependencies: 242
-- Data for Name: userdatasourcedata; Type: TABLE DATA; Schema: public; Owner: ejbca
--



--
-- TOC entry 3349 (class 2606 OID 16392)
-- Name: accessrulesdata accessrulesdata_pkey; Type: CONSTRAINT; Schema: public; Owner: ejbca
--

ALTER TABLE ONLY public.accessrulesdata
    ADD CONSTRAINT accessrulesdata_pkey PRIMARY KEY (pk);


--
-- TOC entry 3428 (class 2606 OID 16604)
-- Name: acmeaccountdata acmeaccountdata_pkey; Type: CONSTRAINT; Schema: public; Owner: ejbca
--

ALTER TABLE ONLY public.acmeaccountdata
    ADD CONSTRAINT acmeaccountdata_pkey PRIMARY KEY (accountid);


--
-- TOC entry 3434 (class 2606 OID 16625)
-- Name: acmeauthorizationdata acmeauthorizationdata_pkey; Type: CONSTRAINT; Schema: public; Owner: ejbca
--

ALTER TABLE ONLY public.acmeauthorizationdata
    ADD CONSTRAINT acmeauthorizationdata_pkey PRIMARY KEY (authorizationid);


--
-- TOC entry 3432 (class 2606 OID 16618)
-- Name: acmechallengedata acmechallengedata_pkey; Type: CONSTRAINT; Schema: public; Owner: ejbca
--

ALTER TABLE ONLY public.acmechallengedata
    ADD CONSTRAINT acmechallengedata_pkey PRIMARY KEY (challengeid);


--
-- TOC entry 3426 (class 2606 OID 16597)
-- Name: acmenoncedata acmenoncedata_pkey; Type: CONSTRAINT; Schema: public; Owner: ejbca
--

ALTER TABLE ONLY public.acmenoncedata
    ADD CONSTRAINT acmenoncedata_pkey PRIMARY KEY (nonce);


--
-- TOC entry 3430 (class 2606 OID 16611)
-- Name: acmeorderdata acmeorderdata_pkey; Type: CONSTRAINT; Schema: public; Owner: ejbca
--

ALTER TABLE ONLY public.acmeorderdata
    ADD CONSTRAINT acmeorderdata_pkey PRIMARY KEY (orderid);


--
-- TOC entry 3351 (class 2606 OID 16399)
-- Name: adminentitydata adminentitydata_pkey; Type: CONSTRAINT; Schema: public; Owner: ejbca
--

ALTER TABLE ONLY public.adminentitydata
    ADD CONSTRAINT adminentitydata_pkey PRIMARY KEY (pk);


--
-- TOC entry 3353 (class 2606 OID 16406)
-- Name: admingroupdata admingroupdata_pkey; Type: CONSTRAINT; Schema: public; Owner: ejbca
--

ALTER TABLE ONLY public.admingroupdata
    ADD CONSTRAINT admingroupdata_pkey PRIMARY KEY (pk);


--
-- TOC entry 3355 (class 2606 OID 16413)
-- Name: adminpreferencesdata adminpreferencesdata_pkey; Type: CONSTRAINT; Schema: public; Owner: ejbca
--

ALTER TABLE ONLY public.adminpreferencesdata
    ADD CONSTRAINT adminpreferencesdata_pkey PRIMARY KEY (id);


--
-- TOC entry 3357 (class 2606 OID 16420)
-- Name: approvaldata approvaldata_pkey; Type: CONSTRAINT; Schema: public; Owner: ejbca
--

ALTER TABLE ONLY public.approvaldata
    ADD CONSTRAINT approvaldata_pkey PRIMARY KEY (id);


--
-- TOC entry 3362 (class 2606 OID 16427)
-- Name: auditrecorddata auditrecorddata_pkey; Type: CONSTRAINT; Schema: public; Owner: ejbca
--

ALTER TABLE ONLY public.auditrecorddata
    ADD CONSTRAINT auditrecorddata_pkey PRIMARY KEY (pk);


--
-- TOC entry 3364 (class 2606 OID 16434)
-- Name: authorizationtreeupdatedata authorizationtreeupdatedata_pkey; Type: CONSTRAINT; Schema: public; Owner: ejbca
--

ALTER TABLE ONLY public.authorizationtreeupdatedata
    ADD CONSTRAINT authorizationtreeupdatedata_pkey PRIMARY KEY (pk);


--
-- TOC entry 3366 (class 2606 OID 16441)
-- Name: base64certdata base64certdata_pkey; Type: CONSTRAINT; Schema: public; Owner: ejbca
--

ALTER TABLE ONLY public.base64certdata
    ADD CONSTRAINT base64certdata_pkey PRIMARY KEY (fingerprint);


--
-- TOC entry 3412 (class 2606 OID 16547)
-- Name: blacklistdata blacklistdata_pkey; Type: CONSTRAINT; Schema: public; Owner: ejbca
--

ALTER TABLE ONLY public.blacklistdata
    ADD CONSTRAINT blacklistdata_pkey PRIMARY KEY (id);


--
-- TOC entry 3369 (class 2606 OID 16448)
-- Name: cadata cadata_pkey; Type: CONSTRAINT; Schema: public; Owner: ejbca
--

ALTER TABLE ONLY public.cadata
    ADD CONSTRAINT cadata_pkey PRIMARY KEY (caid);


--
-- TOC entry 3390 (class 2606 OID 16469)
-- Name: certificatedata certificatedata_pkey; Type: CONSTRAINT; Schema: public; Owner: ejbca
--

ALTER TABLE ONLY public.certificatedata
    ADD CONSTRAINT certificatedata_pkey PRIMARY KEY (fingerprint);


--
-- TOC entry 3392 (class 2606 OID 16476)
-- Name: certificateprofiledata certificateprofiledata_pkey; Type: CONSTRAINT; Schema: public; Owner: ejbca
--

ALTER TABLE ONLY public.certificateprofiledata
    ADD CONSTRAINT certificateprofiledata_pkey PRIMARY KEY (id);


--
-- TOC entry 3375 (class 2606 OID 16462)
-- Name: certreqhistorydata certreqhistorydata_pkey; Type: CONSTRAINT; Schema: public; Owner: ejbca
--

ALTER TABLE ONLY public.certreqhistorydata
    ADD CONSTRAINT certreqhistorydata_pkey PRIMARY KEY (fingerprint);


--
-- TOC entry 3373 (class 2606 OID 16455)
-- Name: crldata crldata_pkey; Type: CONSTRAINT; Schema: public; Owner: ejbca
--

ALTER TABLE ONLY public.crldata
    ADD CONSTRAINT crldata_pkey PRIMARY KEY (fingerprint);


--
-- TOC entry 3394 (class 2606 OID 16483)
-- Name: cryptotokendata cryptotokendata_pkey; Type: CONSTRAINT; Schema: public; Owner: ejbca
--

ALTER TABLE ONLY public.cryptotokendata
    ADD CONSTRAINT cryptotokendata_pkey PRIMARY KEY (id);


--
-- TOC entry 3396 (class 2606 OID 16490)
-- Name: endentityprofiledata endentityprofiledata_pkey; Type: CONSTRAINT; Schema: public; Owner: ejbca
--

ALTER TABLE ONLY public.endentityprofiledata
    ADD CONSTRAINT endentityprofiledata_pkey PRIMARY KEY (id);


--
-- TOC entry 3398 (class 2606 OID 16497)
-- Name: globalconfigurationdata globalconfigurationdata_pkey; Type: CONSTRAINT; Schema: public; Owner: ejbca
--

ALTER TABLE ONLY public.globalconfigurationdata
    ADD CONSTRAINT globalconfigurationdata_pkey PRIMARY KEY (configurationid);


--
-- TOC entry 3443 (class 2606 OID 16646)
-- Name: incompleteissuancejournaldata incompleteissuancejournaldata_pkey; Type: CONSTRAINT; Schema: public; Owner: ejbca
--

ALTER TABLE ONLY public.incompleteissuancejournaldata
    ADD CONSTRAINT incompleteissuancejournaldata_pkey PRIMARY KEY (serialnumberandcaid);


--
-- TOC entry 3400 (class 2606 OID 16504)
-- Name: internalkeybindingdata internalkeybindingdata_pkey; Type: CONSTRAINT; Schema: public; Owner: ejbca
--

ALTER TABLE ONLY public.internalkeybindingdata
    ADD CONSTRAINT internalkeybindingdata_pkey PRIMARY KEY (id);


--
-- TOC entry 3402 (class 2606 OID 16512)
-- Name: keyrecoverydata keyrecoverydata_pkey; Type: CONSTRAINT; Schema: public; Owner: ejbca
--

ALTER TABLE ONLY public.keyrecoverydata
    ADD CONSTRAINT keyrecoverydata_pkey PRIMARY KEY (certsn, issuerdn);


--
-- TOC entry 3424 (class 2606 OID 16590)
-- Name: noconflictcertificatedata noconflictcertificatedata_pkey; Type: CONSTRAINT; Schema: public; Owner: ejbca
--

ALTER TABLE ONLY public.noconflictcertificatedata
    ADD CONSTRAINT noconflictcertificatedata_pkey PRIMARY KEY (id);


--
-- TOC entry 3441 (class 2606 OID 16639)
-- Name: ocspresponsedata ocspresponsedata_pkey; Type: CONSTRAINT; Schema: public; Owner: ejbca
--

ALTER TABLE ONLY public.ocspresponsedata
    ADD CONSTRAINT ocspresponsedata_pkey PRIMARY KEY (id);


--
-- TOC entry 3404 (class 2606 OID 16519)
-- Name: peerdata peerdata_pkey; Type: CONSTRAINT; Schema: public; Owner: ejbca
--

ALTER TABLE ONLY public.peerdata
    ADD CONSTRAINT peerdata_pkey PRIMARY KEY (id);


--
-- TOC entry 3406 (class 2606 OID 16526)
-- Name: profiledata profiledata_pkey; Type: CONSTRAINT; Schema: public; Owner: ejbca
--

ALTER TABLE ONLY public.profiledata
    ADD CONSTRAINT profiledata_pkey PRIMARY KEY (id);


--
-- TOC entry 3408 (class 2606 OID 16533)
-- Name: publisherdata publisherdata_pkey; Type: CONSTRAINT; Schema: public; Owner: ejbca
--

ALTER TABLE ONLY public.publisherdata
    ADD CONSTRAINT publisherdata_pkey PRIMARY KEY (id);


--
-- TOC entry 3410 (class 2606 OID 16540)
-- Name: publisherqueuedata publisherqueuedata_pkey; Type: CONSTRAINT; Schema: public; Owner: ejbca
--

ALTER TABLE ONLY public.publisherqueuedata
    ADD CONSTRAINT publisherqueuedata_pkey PRIMARY KEY (pk);


--
-- TOC entry 3414 (class 2606 OID 16554)
-- Name: roledata roledata_pkey; Type: CONSTRAINT; Schema: public; Owner: ejbca
--

ALTER TABLE ONLY public.roledata
    ADD CONSTRAINT roledata_pkey PRIMARY KEY (id);


--
-- TOC entry 3416 (class 2606 OID 16562)
-- Name: rolememberdata rolememberdata_pkey; Type: CONSTRAINT; Schema: public; Owner: ejbca
--

ALTER TABLE ONLY public.rolememberdata
    ADD CONSTRAINT rolememberdata_pkey PRIMARY KEY (primarykey);


--
-- TOC entry 3436 (class 2606 OID 16632)
-- Name: sctdata sctdata_pkey; Type: CONSTRAINT; Schema: public; Owner: ejbca
--

ALTER TABLE ONLY public.sctdata
    ADD CONSTRAINT sctdata_pkey PRIMARY KEY (pk);


--
-- TOC entry 3418 (class 2606 OID 16569)
-- Name: servicedata servicedata_pkey; Type: CONSTRAINT; Schema: public; Owner: ejbca
--

ALTER TABLE ONLY public.servicedata
    ADD CONSTRAINT servicedata_pkey PRIMARY KEY (id);


--
-- TOC entry 3420 (class 2606 OID 16576)
-- Name: userdata userdata_pkey; Type: CONSTRAINT; Schema: public; Owner: ejbca
--

ALTER TABLE ONLY public.userdata
    ADD CONSTRAINT userdata_pkey PRIMARY KEY (username);


--
-- TOC entry 3422 (class 2606 OID 16583)
-- Name: userdatasourcedata userdatasourcedata_pkey; Type: CONSTRAINT; Schema: public; Owner: ejbca
--

ALTER TABLE ONLY public.userdatasourcedata
    ADD CONSTRAINT userdatasourcedata_pkey PRIMARY KEY (id);


--
-- TOC entry 3358 (class 1259 OID 16657)
-- Name: auditrecorddata_idx2; Type: INDEX; Schema: public; Owner: ejbca
--

CREATE UNIQUE INDEX auditrecorddata_idx2 ON public.auditrecorddata USING btree (nodeid, sequencenumber);


--
-- TOC entry 3359 (class 1259 OID 16658)
-- Name: auditrecorddata_idx3; Type: INDEX; Schema: public; Owner: ejbca
--

CREATE INDEX auditrecorddata_idx3 ON public.auditrecorddata USING btree ("timestamp");


--
-- TOC entry 3360 (class 1259 OID 16659)
-- Name: auditrecorddata_idx4; Type: INDEX; Schema: public; Owner: ejbca
--

CREATE INDEX auditrecorddata_idx4 ON public.auditrecorddata USING btree (searchdetail2);


--
-- TOC entry 3367 (class 1259 OID 16662)
-- Name: cadata_idx1; Type: INDEX; Schema: public; Owner: ejbca
--

CREATE UNIQUE INDEX cadata_idx1 ON public.cadata USING btree (name);


--
-- TOC entry 3376 (class 1259 OID 16668)
-- Name: certificatedata_idx11; Type: INDEX; Schema: public; Owner: ejbca
--

CREATE INDEX certificatedata_idx11 ON public.certificatedata USING btree (subjectkeyid);


--
-- TOC entry 3377 (class 1259 OID 16669)
-- Name: certificatedata_idx12; Type: INDEX; Schema: public; Owner: ejbca
--

CREATE UNIQUE INDEX certificatedata_idx12 ON public.certificatedata USING btree (serialnumber, issuerdn);


--
-- TOC entry 3378 (class 1259 OID 16670)
-- Name: certificatedata_idx15; Type: INDEX; Schema: public; Owner: ejbca
--

CREATE INDEX certificatedata_idx15 ON public.certificatedata USING btree (issuerdn, notbefore);


--
-- TOC entry 3379 (class 1259 OID 16671)
-- Name: certificatedata_idx16; Type: INDEX; Schema: public; Owner: ejbca
--

CREATE INDEX certificatedata_idx16 ON public.certificatedata USING btree (issuerdn, revocationdate);


--
-- TOC entry 3380 (class 1259 OID 16672)
-- Name: certificatedata_idx17; Type: INDEX; Schema: public; Owner: ejbca
--

CREATE INDEX certificatedata_idx17 ON public.certificatedata USING btree (issuerdn, status, crlpartitionindex);


--
-- TOC entry 3381 (class 1259 OID 16673)
-- Name: certificatedata_idx18; Type: INDEX; Schema: public; Owner: ejbca
--

CREATE INDEX certificatedata_idx18 ON public.certificatedata USING btree (issuerdn, status, crlpartitionindex, revocationdate);


--
-- TOC entry 3382 (class 1259 OID 16663)
-- Name: certificatedata_idx2; Type: INDEX; Schema: public; Owner: ejbca
--

CREATE INDEX certificatedata_idx2 ON public.certificatedata USING btree (username);


--
-- TOC entry 3383 (class 1259 OID 16664)
-- Name: certificatedata_idx4; Type: INDEX; Schema: public; Owner: ejbca
--

CREATE INDEX certificatedata_idx4 ON public.certificatedata USING btree (subjectdn);


--
-- TOC entry 3384 (class 1259 OID 16665)
-- Name: certificatedata_idx5; Type: INDEX; Schema: public; Owner: ejbca
--

CREATE INDEX certificatedata_idx5 ON public.certificatedata USING btree (type);


--
-- TOC entry 3385 (class 1259 OID 16666)
-- Name: certificatedata_idx6; Type: INDEX; Schema: public; Owner: ejbca
--

CREATE INDEX certificatedata_idx6 ON public.certificatedata USING btree (issuerdn, status);


--
-- TOC entry 3386 (class 1259 OID 16667)
-- Name: certificatedata_idx7; Type: INDEX; Schema: public; Owner: ejbca
--

CREATE INDEX certificatedata_idx7 ON public.certificatedata USING btree (certificateprofileid);


--
-- TOC entry 3387 (class 1259 OID 16675)
-- Name: certificatedata_idx_eab; Type: INDEX; Schema: public; Owner: ejbca
--

CREATE INDEX certificatedata_idx_eab ON public.certificatedata USING btree (accountbindingid);


--
-- TOC entry 3388 (class 1259 OID 16674)
-- Name: certificatedata_idx_serial; Type: INDEX; Schema: public; Owner: ejbca
--

CREATE INDEX certificatedata_idx_serial ON public.certificatedata USING btree (serialnumber);


--
-- TOC entry 3370 (class 1259 OID 16660)
-- Name: crldata_idx5; Type: INDEX; Schema: public; Owner: ejbca
--

CREATE INDEX crldata_idx5 ON public.crldata USING btree (crlnumber, issuerdn, crlpartitionindex);


--
-- TOC entry 3371 (class 1259 OID 16661)
-- Name: crldata_idx6; Type: INDEX; Schema: public; Owner: ejbca
--

CREATE UNIQUE INDEX crldata_idx6 ON public.crldata USING btree (issuerdn, crlpartitionindex, deltacrlindicator, crlnumber);


--
-- TOC entry 3437 (class 1259 OID 16677)
-- Name: ocspresponsedata_idx1; Type: INDEX; Schema: public; Owner: ejbca
--

CREATE INDEX ocspresponsedata_idx1 ON public.ocspresponsedata USING btree (caid);


--
-- TOC entry 3438 (class 1259 OID 16678)
-- Name: ocspresponsedata_idx2; Type: INDEX; Schema: public; Owner: ejbca
--

CREATE INDEX ocspresponsedata_idx2 ON public.ocspresponsedata USING btree (serialnumber);


--
-- TOC entry 3439 (class 1259 OID 16679)
-- Name: ocspresponsedata_idx3; Type: INDEX; Schema: public; Owner: ejbca
--

CREATE INDEX ocspresponsedata_idx3 ON public.ocspresponsedata USING btree (producedat);


--
-- TOC entry 3444 (class 2606 OID 16647)
-- Name: accessrulesdata fkabb4c1dfdbbc970; Type: FK CONSTRAINT; Schema: public; Owner: ejbca
--

ALTER TABLE ONLY public.accessrulesdata
    ADD CONSTRAINT fkabb4c1dfdbbc970 FOREIGN KEY (admingroupdata_accessrules) REFERENCES public.admingroupdata(pk);


--
-- TOC entry 3445 (class 2606 OID 16652)
-- Name: adminentitydata fkd9a99ebcb3a110ad; Type: FK CONSTRAINT; Schema: public; Owner: ejbca
--

ALTER TABLE ONLY public.adminentitydata
    ADD CONSTRAINT fkd9a99ebcb3a110ad FOREIGN KEY (admingroupdata_adminentities) REFERENCES public.admingroupdata(pk);


-- Completed on 2025-07-24 11:58:47

--
-- PostgreSQL database dump complete
--

