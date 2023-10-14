--
-- PostgreSQL database dump
--

-- Dumped from database version 15.3
-- Dumped by pg_dump version 15.4

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

ALTER TABLE IF EXISTS ONLY public.widgets DROP CONSTRAINT IF EXISTS widgets_visualization_id_fkey;
ALTER TABLE IF EXISTS ONLY public.widgets DROP CONSTRAINT IF EXISTS widgets_dashboard_id_fkey;
ALTER TABLE IF EXISTS ONLY public.visualizations DROP CONSTRAINT IF EXISTS visualizations_query_id_fkey;
ALTER TABLE IF EXISTS ONLY public.users DROP CONSTRAINT IF EXISTS users_org_id_fkey;
ALTER TABLE IF EXISTS ONLY public.query_snippets DROP CONSTRAINT IF EXISTS query_snippets_user_id_fkey;
ALTER TABLE IF EXISTS ONLY public.query_snippets DROP CONSTRAINT IF EXISTS query_snippets_org_id_fkey;
ALTER TABLE IF EXISTS ONLY public.query_results DROP CONSTRAINT IF EXISTS query_results_org_id_fkey;
ALTER TABLE IF EXISTS ONLY public.query_results DROP CONSTRAINT IF EXISTS query_results_data_source_id_fkey;
ALTER TABLE IF EXISTS ONLY public.queries DROP CONSTRAINT IF EXISTS queries_user_id_fkey;
ALTER TABLE IF EXISTS ONLY public.queries DROP CONSTRAINT IF EXISTS queries_org_id_fkey;
ALTER TABLE IF EXISTS ONLY public.queries DROP CONSTRAINT IF EXISTS queries_latest_query_data_id_fkey;
ALTER TABLE IF EXISTS ONLY public.queries DROP CONSTRAINT IF EXISTS queries_last_modified_by_id_fkey;
ALTER TABLE IF EXISTS ONLY public.queries DROP CONSTRAINT IF EXISTS queries_data_source_id_fkey;
ALTER TABLE IF EXISTS ONLY public.notification_destinations DROP CONSTRAINT IF EXISTS notification_destinations_user_id_fkey;
ALTER TABLE IF EXISTS ONLY public.notification_destinations DROP CONSTRAINT IF EXISTS notification_destinations_org_id_fkey;
ALTER TABLE IF EXISTS ONLY public.groups DROP CONSTRAINT IF EXISTS groups_org_id_fkey;
ALTER TABLE IF EXISTS ONLY public.favorites DROP CONSTRAINT IF EXISTS favorites_user_id_fkey;
ALTER TABLE IF EXISTS ONLY public.favorites DROP CONSTRAINT IF EXISTS favorites_org_id_fkey;
ALTER TABLE IF EXISTS ONLY public.events DROP CONSTRAINT IF EXISTS events_user_id_fkey;
ALTER TABLE IF EXISTS ONLY public.events DROP CONSTRAINT IF EXISTS events_org_id_fkey;
ALTER TABLE IF EXISTS ONLY public.data_sources DROP CONSTRAINT IF EXISTS data_sources_org_id_fkey;
ALTER TABLE IF EXISTS ONLY public.data_source_groups DROP CONSTRAINT IF EXISTS data_source_groups_group_id_fkey;
ALTER TABLE IF EXISTS ONLY public.data_source_groups DROP CONSTRAINT IF EXISTS data_source_groups_data_source_id_fkey;
ALTER TABLE IF EXISTS ONLY public.dashboards DROP CONSTRAINT IF EXISTS dashboards_user_id_fkey;
ALTER TABLE IF EXISTS ONLY public.dashboards DROP CONSTRAINT IF EXISTS dashboards_org_id_fkey;
ALTER TABLE IF EXISTS ONLY public.changes DROP CONSTRAINT IF EXISTS changes_user_id_fkey;
ALTER TABLE IF EXISTS ONLY public.api_keys DROP CONSTRAINT IF EXISTS api_keys_org_id_fkey;
ALTER TABLE IF EXISTS ONLY public.api_keys DROP CONSTRAINT IF EXISTS api_keys_created_by_id_fkey;
ALTER TABLE IF EXISTS ONLY public.alerts DROP CONSTRAINT IF EXISTS alerts_user_id_fkey;
ALTER TABLE IF EXISTS ONLY public.alerts DROP CONSTRAINT IF EXISTS alerts_query_id_fkey;
ALTER TABLE IF EXISTS ONLY public.alert_subscriptions DROP CONSTRAINT IF EXISTS alert_subscriptions_user_id_fkey;
ALTER TABLE IF EXISTS ONLY public.alert_subscriptions DROP CONSTRAINT IF EXISTS alert_subscriptions_destination_id_fkey;
ALTER TABLE IF EXISTS ONLY public.alert_subscriptions DROP CONSTRAINT IF EXISTS alert_subscriptions_alert_id_fkey;
ALTER TABLE IF EXISTS ONLY public.access_permissions DROP CONSTRAINT IF EXISTS access_permissions_grantor_id_fkey;
ALTER TABLE IF EXISTS ONLY public.access_permissions DROP CONSTRAINT IF EXISTS access_permissions_grantee_id_fkey;
DROP TRIGGER IF EXISTS queries_search_vector_trigger ON public.queries;
DROP INDEX IF EXISTS public.users_org_id_email;
DROP INDEX IF EXISTS public.notification_destinations_org_id_name;
DROP INDEX IF EXISTS public.ix_widgets_dashboard_id;
DROP INDEX IF EXISTS public.ix_query_results_query_hash;
DROP INDEX IF EXISTS public.ix_queries_search_vector;
DROP INDEX IF EXISTS public.ix_queries_is_draft;
DROP INDEX IF EXISTS public.ix_queries_is_archived;
DROP INDEX IF EXISTS public.ix_dashboards_slug;
DROP INDEX IF EXISTS public.ix_dashboards_is_draft;
DROP INDEX IF EXISTS public.ix_dashboards_is_archived;
DROP INDEX IF EXISTS public.ix_api_keys_api_key;
DROP INDEX IF EXISTS public.data_sources_org_id_name;
DROP INDEX IF EXISTS public.api_keys_object_type_object_id;
DROP INDEX IF EXISTS public.alert_subscriptions_destination_id_alert_id;
ALTER TABLE IF EXISTS ONLY public.widgets DROP CONSTRAINT IF EXISTS widgets_pkey;
ALTER TABLE IF EXISTS ONLY public.visualizations DROP CONSTRAINT IF EXISTS visualizations_pkey;
ALTER TABLE IF EXISTS ONLY public.users DROP CONSTRAINT IF EXISTS users_pkey;
ALTER TABLE IF EXISTS ONLY public.users DROP CONSTRAINT IF EXISTS users_api_key_key;
ALTER TABLE IF EXISTS ONLY public.favorites DROP CONSTRAINT IF EXISTS unique_favorite;
ALTER TABLE IF EXISTS ONLY public.query_snippets DROP CONSTRAINT IF EXISTS query_snippets_trigger_key;
ALTER TABLE IF EXISTS ONLY public.query_snippets DROP CONSTRAINT IF EXISTS query_snippets_pkey;
ALTER TABLE IF EXISTS ONLY public.query_results DROP CONSTRAINT IF EXISTS query_results_pkey;
ALTER TABLE IF EXISTS ONLY public.queries DROP CONSTRAINT IF EXISTS queries_pkey;
ALTER TABLE IF EXISTS ONLY public.organizations DROP CONSTRAINT IF EXISTS organizations_slug_key;
ALTER TABLE IF EXISTS ONLY public.organizations DROP CONSTRAINT IF EXISTS organizations_pkey;
ALTER TABLE IF EXISTS ONLY public.notification_destinations DROP CONSTRAINT IF EXISTS notification_destinations_pkey;
ALTER TABLE IF EXISTS ONLY public.groups DROP CONSTRAINT IF EXISTS groups_pkey;
ALTER TABLE IF EXISTS ONLY public.favorites DROP CONSTRAINT IF EXISTS favorites_pkey;
ALTER TABLE IF EXISTS ONLY public.events DROP CONSTRAINT IF EXISTS events_pkey;
ALTER TABLE IF EXISTS ONLY public.data_sources DROP CONSTRAINT IF EXISTS data_sources_pkey;
ALTER TABLE IF EXISTS ONLY public.data_source_groups DROP CONSTRAINT IF EXISTS data_source_groups_pkey;
ALTER TABLE IF EXISTS ONLY public.dashboards DROP CONSTRAINT IF EXISTS dashboards_pkey;
ALTER TABLE IF EXISTS ONLY public.changes DROP CONSTRAINT IF EXISTS changes_pkey;
ALTER TABLE IF EXISTS ONLY public.api_keys DROP CONSTRAINT IF EXISTS api_keys_pkey;
ALTER TABLE IF EXISTS ONLY public.alerts DROP CONSTRAINT IF EXISTS alerts_pkey;
ALTER TABLE IF EXISTS ONLY public.alert_subscriptions DROP CONSTRAINT IF EXISTS alert_subscriptions_pkey;
ALTER TABLE IF EXISTS ONLY public.alembic_version DROP CONSTRAINT IF EXISTS alembic_version_pkc;
ALTER TABLE IF EXISTS ONLY public.access_permissions DROP CONSTRAINT IF EXISTS access_permissions_pkey;
ALTER TABLE IF EXISTS public.widgets ALTER COLUMN id DROP DEFAULT;
ALTER TABLE IF EXISTS public.visualizations ALTER COLUMN id DROP DEFAULT;
ALTER TABLE IF EXISTS public.users ALTER COLUMN id DROP DEFAULT;
ALTER TABLE IF EXISTS public.query_snippets ALTER COLUMN id DROP DEFAULT;
ALTER TABLE IF EXISTS public.query_results ALTER COLUMN id DROP DEFAULT;
ALTER TABLE IF EXISTS public.queries ALTER COLUMN id DROP DEFAULT;
ALTER TABLE IF EXISTS public.organizations ALTER COLUMN id DROP DEFAULT;
ALTER TABLE IF EXISTS public.notification_destinations ALTER COLUMN id DROP DEFAULT;
ALTER TABLE IF EXISTS public.groups ALTER COLUMN id DROP DEFAULT;
ALTER TABLE IF EXISTS public.favorites ALTER COLUMN id DROP DEFAULT;
ALTER TABLE IF EXISTS public.events ALTER COLUMN id DROP DEFAULT;
ALTER TABLE IF EXISTS public.data_sources ALTER COLUMN id DROP DEFAULT;
ALTER TABLE IF EXISTS public.data_source_groups ALTER COLUMN id DROP DEFAULT;
ALTER TABLE IF EXISTS public.dashboards ALTER COLUMN id DROP DEFAULT;
ALTER TABLE IF EXISTS public.changes ALTER COLUMN id DROP DEFAULT;
ALTER TABLE IF EXISTS public.api_keys ALTER COLUMN id DROP DEFAULT;
ALTER TABLE IF EXISTS public.alerts ALTER COLUMN id DROP DEFAULT;
ALTER TABLE IF EXISTS public.alert_subscriptions ALTER COLUMN id DROP DEFAULT;
ALTER TABLE IF EXISTS public.access_permissions ALTER COLUMN id DROP DEFAULT;
DROP SEQUENCE IF EXISTS public.widgets_id_seq;
DROP TABLE IF EXISTS public.widgets;
DROP SEQUENCE IF EXISTS public.visualizations_id_seq;
DROP TABLE IF EXISTS public.visualizations;
DROP SEQUENCE IF EXISTS public.users_id_seq;
DROP TABLE IF EXISTS public.users;
DROP SEQUENCE IF EXISTS public.query_snippets_id_seq;
DROP TABLE IF EXISTS public.query_snippets;
DROP SEQUENCE IF EXISTS public.query_results_id_seq;
DROP TABLE IF EXISTS public.query_results;
DROP SEQUENCE IF EXISTS public.queries_id_seq;
DROP TABLE IF EXISTS public.queries;
DROP SEQUENCE IF EXISTS public.organizations_id_seq;
DROP TABLE IF EXISTS public.organizations;
DROP SEQUENCE IF EXISTS public.notification_destinations_id_seq;
DROP TABLE IF EXISTS public.notification_destinations;
DROP SEQUENCE IF EXISTS public.groups_id_seq;
DROP TABLE IF EXISTS public.groups;
DROP SEQUENCE IF EXISTS public.favorites_id_seq;
DROP TABLE IF EXISTS public.favorites;
DROP SEQUENCE IF EXISTS public.events_id_seq;
DROP TABLE IF EXISTS public.events;
DROP SEQUENCE IF EXISTS public.data_sources_id_seq;
DROP TABLE IF EXISTS public.data_sources;
DROP SEQUENCE IF EXISTS public.data_source_groups_id_seq;
DROP TABLE IF EXISTS public.data_source_groups;
DROP SEQUENCE IF EXISTS public.dashboards_id_seq;
DROP TABLE IF EXISTS public.dashboards;
DROP SEQUENCE IF EXISTS public.changes_id_seq;
DROP TABLE IF EXISTS public.changes;
DROP SEQUENCE IF EXISTS public.api_keys_id_seq;
DROP TABLE IF EXISTS public.api_keys;
DROP SEQUENCE IF EXISTS public.alerts_id_seq;
DROP TABLE IF EXISTS public.alerts;
DROP SEQUENCE IF EXISTS public.alert_subscriptions_id_seq;
DROP TABLE IF EXISTS public.alert_subscriptions;
DROP TABLE IF EXISTS public.alembic_version;
DROP SEQUENCE IF EXISTS public.access_permissions_id_seq;
DROP TABLE IF EXISTS public.access_permissions;
DROP FUNCTION IF EXISTS public.tsq_tokenize_character(state public.tsq_state);
DROP FUNCTION IF EXISTS public.tsq_tokenize(search_query text);
DROP FUNCTION IF EXISTS public.tsq_process_tokens(config regconfig, tokens text[]);
DROP FUNCTION IF EXISTS public.tsq_process_tokens(tokens text[]);
DROP FUNCTION IF EXISTS public.tsq_parse(config text, search_query text);
DROP FUNCTION IF EXISTS public.tsq_parse(config regconfig, search_query text);
DROP FUNCTION IF EXISTS public.tsq_parse(search_query text);
DROP FUNCTION IF EXISTS public.tsq_append_current_token(state public.tsq_state);
DROP FUNCTION IF EXISTS public.queries_search_vector_update();
DROP FUNCTION IF EXISTS public.array_nremove(anyarray, anyelement, integer);
DROP TYPE IF EXISTS public.tsq_state;
--
-- Name: tsq_state; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.tsq_state AS (
	search_query text,
	parentheses_stack integer,
	skip_for integer,
	current_token text,
	current_index integer,
	current_char text,
	previous_char text,
	tokens text[]
);


ALTER TYPE public.tsq_state OWNER TO postgres;

--
-- Name: array_nremove(anyarray, anyelement, integer); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.array_nremove(anyarray, anyelement, integer) RETURNS anyarray
    LANGUAGE sql IMMUTABLE
    AS $_$
    WITH replaced_positions AS (
        SELECT UNNEST(
            CASE
            WHEN $2 IS NULL THEN
                '{}'::int[]
            WHEN $3 > 0 THEN
                (array_positions($1, $2))[1:$3]
            WHEN $3 < 0 THEN
                (array_positions($1, $2))[
                    (cardinality(array_positions($1, $2)) + $3 + 1):
                ]
            ELSE
                '{}'::int[]
            END
        ) AS position
    )
    SELECT COALESCE((
        SELECT array_agg(value)
        FROM unnest($1) WITH ORDINALITY AS t(value, index)
        WHERE index NOT IN (SELECT position FROM replaced_positions)
    ), $1[1:0]);
$_$;


ALTER FUNCTION public.array_nremove(anyarray, anyelement, integer) OWNER TO postgres;

--
-- Name: queries_search_vector_update(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.queries_search_vector_update() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
            BEGIN
                NEW.search_vector = ((setweight(to_tsvector('pg_catalog.simple', coalesce(CAST(NEW.id AS TEXT), '')), 'B') || setweight(to_tsvector('pg_catalog.simple', coalesce(NEW.name, '')), 'A')) || setweight(to_tsvector('pg_catalog.simple', coalesce(NEW.description, '')), 'C')) || setweight(to_tsvector('pg_catalog.simple', coalesce(NEW.query, '')), 'D');
                RETURN NEW;
            END
            $$;


ALTER FUNCTION public.queries_search_vector_update() OWNER TO postgres;

--
-- Name: tsq_append_current_token(public.tsq_state); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.tsq_append_current_token(state public.tsq_state) RETURNS public.tsq_state
    LANGUAGE plpgsql IMMUTABLE
    AS $$
BEGIN
    IF state.current_token != '' THEN
        state.tokens := array_append(state.tokens, state.current_token);
        state.current_token := '';
    END IF;
    RETURN state;
END;
$$;


ALTER FUNCTION public.tsq_append_current_token(state public.tsq_state) OWNER TO postgres;

--
-- Name: tsq_parse(text); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.tsq_parse(search_query text) RETURNS tsquery
    LANGUAGE sql IMMUTABLE
    AS $$
    SELECT tsq_parse(get_current_ts_config(), search_query);
$$;


ALTER FUNCTION public.tsq_parse(search_query text) OWNER TO postgres;

--
-- Name: tsq_parse(regconfig, text); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.tsq_parse(config regconfig, search_query text) RETURNS tsquery
    LANGUAGE sql IMMUTABLE
    AS $$
    SELECT tsq_process_tokens(config, tsq_tokenize(search_query));
$$;


ALTER FUNCTION public.tsq_parse(config regconfig, search_query text) OWNER TO postgres;

--
-- Name: tsq_parse(text, text); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.tsq_parse(config text, search_query text) RETURNS tsquery
    LANGUAGE sql IMMUTABLE
    AS $$
    SELECT tsq_parse(config::regconfig, search_query);
$$;


ALTER FUNCTION public.tsq_parse(config text, search_query text) OWNER TO postgres;

--
-- Name: tsq_process_tokens(text[]); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.tsq_process_tokens(tokens text[]) RETURNS tsquery
    LANGUAGE sql IMMUTABLE
    AS $$
    SELECT tsq_process_tokens(get_current_ts_config(), tokens);
$$;


ALTER FUNCTION public.tsq_process_tokens(tokens text[]) OWNER TO postgres;

--
-- Name: tsq_process_tokens(regconfig, text[]); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.tsq_process_tokens(config regconfig, tokens text[]) RETURNS tsquery
    LANGUAGE plpgsql IMMUTABLE
    AS $$
DECLARE
    result_query text;
    previous_value text;
    value text;
BEGIN
    result_query := '';

    FOREACH value IN ARRAY tokens LOOP
        IF value = '"' THEN
            CONTINUE;
        END IF;

        IF value = 'or' THEN
            value := ' | ';
        END IF;

        IF left(value, 1) = '"' AND right(value, 1) = '"' THEN
            value := phraseto_tsquery(config, value);
        ELSIF value NOT IN ('(', ' | ', ')', '-') THEN
            value := quote_literal(value) || ':*';
        END IF;

        IF previous_value = '-' THEN
            IF value = '(' THEN
                value := '!' || value;
            ELSIF value = ' | ' THEN
                CONTINUE;
            ELSE
                value := '!(' || value || ')';
            END IF;
        END IF;

        SELECT
            CASE
                WHEN result_query = '' THEN value
                WHEN previous_value = ' | ' AND value = ' | ' THEN result_query
                WHEN previous_value = ' | ' THEN result_query || ' | ' || value
                WHEN previous_value IN ('!(', '(') OR value = ')' THEN result_query || value
                WHEN value != ' | ' THEN result_query || ' & ' || value
                ELSE result_query
            END
        INTO result_query;

        IF result_query = ' | ' THEN
            result_query := '';
        END IF;

        previous_value := value;
    END LOOP;

    RETURN to_tsquery(config, result_query);
END;
$$;


ALTER FUNCTION public.tsq_process_tokens(config regconfig, tokens text[]) OWNER TO postgres;

--
-- Name: tsq_tokenize(text); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.tsq_tokenize(search_query text) RETURNS text[]
    LANGUAGE plpgsql IMMUTABLE
    AS $$
DECLARE
    state tsq_state;
BEGIN
    SELECT
        search_query::text AS search_query,
        0::int AS parentheses_stack,
        0 AS skip_for,
        ''::text AS current_token,
        0 AS current_index,
        ''::text AS current_char,
        ''::text AS previous_char,
        '{}'::text[] AS tokens
    INTO state;

    state.search_query := lower(trim(
        regexp_replace(search_query, '""+', '""', 'g')
    ));

    FOR state.current_index IN (
        SELECT generate_series(1, length(state.search_query))
    ) LOOP
        state.current_char := substring(
            search_query FROM state.current_index FOR 1
        );

        IF state.skip_for > 0 THEN
            state.skip_for := state.skip_for - 1;
            CONTINUE;
        END IF;

        state := tsq_tokenize_character(state);
        state.previous_char := state.current_char;
    END LOOP;
    state := tsq_append_current_token(state);

    state.tokens := array_nremove(state.tokens, '(', -state.parentheses_stack);

    RETURN state.tokens;
END;
$$;


ALTER FUNCTION public.tsq_tokenize(search_query text) OWNER TO postgres;

--
-- Name: tsq_tokenize_character(public.tsq_state); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.tsq_tokenize_character(state public.tsq_state) RETURNS public.tsq_state
    LANGUAGE plpgsql IMMUTABLE
    AS $$
BEGIN
    IF state.current_char = '(' THEN
        state.tokens := array_append(state.tokens, '(');
        state.parentheses_stack := state.parentheses_stack + 1;
        state := tsq_append_current_token(state);
    ELSIF state.current_char = ')' THEN
        IF (state.parentheses_stack > 0 AND state.current_token != '') THEN
            state := tsq_append_current_token(state);
            state.tokens := array_append(state.tokens, ')');
            state.parentheses_stack := state.parentheses_stack - 1;
        END IF;
    ELSIF state.current_char = '"' THEN
        state.skip_for := position('"' IN substring(
            state.search_query FROM state.current_index + 1
        ));

        IF state.skip_for > 1 THEN
            state.tokens = array_append(
                state.tokens,
                substring(
                    state.search_query
                    FROM state.current_index FOR state.skip_for + 1
                )
            );
        ELSIF state.skip_for = 0 THEN
            state.current_token := state.current_token || state.current_char;
        END IF;
    ELSIF (
        state.current_char = '-' AND
        (state.current_index = 1 OR state.previous_char = ' ')
    ) THEN
        state.tokens := array_append(state.tokens, '-');
    ELSIF state.current_char = ' ' THEN
        state := tsq_append_current_token(state);
    ELSE
        state.current_token = state.current_token || state.current_char;
    END IF;
    RETURN state;
END;
$$;


ALTER FUNCTION public.tsq_tokenize_character(state public.tsq_state) OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: access_permissions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.access_permissions (
    object_type character varying(255) NOT NULL,
    object_id integer NOT NULL,
    id integer NOT NULL,
    access_type character varying(255) NOT NULL,
    grantor_id integer NOT NULL,
    grantee_id integer NOT NULL
);


ALTER TABLE public.access_permissions OWNER TO postgres;

--
-- Name: access_permissions_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.access_permissions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.access_permissions_id_seq OWNER TO postgres;

--
-- Name: access_permissions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.access_permissions_id_seq OWNED BY public.access_permissions.id;


--
-- Name: alembic_version; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.alembic_version (
    version_num character varying(32) NOT NULL
);


ALTER TABLE public.alembic_version OWNER TO postgres;

--
-- Name: alert_subscriptions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.alert_subscriptions (
    updated_at timestamp with time zone NOT NULL,
    created_at timestamp with time zone NOT NULL,
    id integer NOT NULL,
    user_id integer NOT NULL,
    destination_id integer,
    alert_id integer NOT NULL
);


ALTER TABLE public.alert_subscriptions OWNER TO postgres;

--
-- Name: alert_subscriptions_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.alert_subscriptions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.alert_subscriptions_id_seq OWNER TO postgres;

--
-- Name: alert_subscriptions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.alert_subscriptions_id_seq OWNED BY public.alert_subscriptions.id;


--
-- Name: alerts; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.alerts (
    updated_at timestamp with time zone NOT NULL,
    created_at timestamp with time zone NOT NULL,
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    query_id integer NOT NULL,
    user_id integer NOT NULL,
    options text NOT NULL,
    state character varying(255) NOT NULL,
    last_triggered_at timestamp with time zone,
    rearm integer
);


ALTER TABLE public.alerts OWNER TO postgres;

--
-- Name: alerts_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.alerts_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.alerts_id_seq OWNER TO postgres;

--
-- Name: alerts_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.alerts_id_seq OWNED BY public.alerts.id;


--
-- Name: api_keys; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.api_keys (
    object_type character varying(255) NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    created_at timestamp with time zone NOT NULL,
    id integer NOT NULL,
    org_id integer NOT NULL,
    api_key character varying(255) NOT NULL,
    active boolean NOT NULL,
    object_id integer NOT NULL,
    created_by_id integer
);


ALTER TABLE public.api_keys OWNER TO postgres;

--
-- Name: api_keys_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.api_keys_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.api_keys_id_seq OWNER TO postgres;

--
-- Name: api_keys_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.api_keys_id_seq OWNED BY public.api_keys.id;


--
-- Name: changes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.changes (
    object_type character varying(255) NOT NULL,
    id integer NOT NULL,
    object_id integer NOT NULL,
    object_version integer NOT NULL,
    user_id integer NOT NULL,
    change text NOT NULL,
    created_at timestamp with time zone NOT NULL
);


ALTER TABLE public.changes OWNER TO postgres;

--
-- Name: changes_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.changes_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.changes_id_seq OWNER TO postgres;

--
-- Name: changes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.changes_id_seq OWNED BY public.changes.id;


--
-- Name: dashboards; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.dashboards (
    updated_at timestamp with time zone NOT NULL,
    created_at timestamp with time zone NOT NULL,
    id integer NOT NULL,
    version integer NOT NULL,
    org_id integer NOT NULL,
    slug character varying(140) NOT NULL,
    name character varying(100) NOT NULL,
    user_id integer NOT NULL,
    layout text NOT NULL,
    dashboard_filters_enabled boolean NOT NULL,
    is_archived boolean NOT NULL,
    is_draft boolean NOT NULL,
    tags character varying[],
    options json DEFAULT '{}'::json NOT NULL
);


ALTER TABLE public.dashboards OWNER TO postgres;

--
-- Name: dashboards_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.dashboards_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.dashboards_id_seq OWNER TO postgres;

--
-- Name: dashboards_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.dashboards_id_seq OWNED BY public.dashboards.id;


--
-- Name: data_source_groups; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.data_source_groups (
    id integer NOT NULL,
    data_source_id integer NOT NULL,
    group_id integer NOT NULL,
    view_only boolean NOT NULL
);


ALTER TABLE public.data_source_groups OWNER TO postgres;

--
-- Name: data_source_groups_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.data_source_groups_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.data_source_groups_id_seq OWNER TO postgres;

--
-- Name: data_source_groups_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.data_source_groups_id_seq OWNED BY public.data_source_groups.id;


--
-- Name: data_sources; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.data_sources (
    id integer NOT NULL,
    org_id integer NOT NULL,
    name character varying(255) NOT NULL,
    type character varying(255) NOT NULL,
    encrypted_options bytea NOT NULL,
    queue_name character varying(255) NOT NULL,
    scheduled_queue_name character varying(255) NOT NULL,
    created_at timestamp with time zone NOT NULL
);


ALTER TABLE public.data_sources OWNER TO postgres;

--
-- Name: data_sources_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.data_sources_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.data_sources_id_seq OWNER TO postgres;

--
-- Name: data_sources_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.data_sources_id_seq OWNED BY public.data_sources.id;


--
-- Name: events; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.events (
    id integer NOT NULL,
    org_id integer NOT NULL,
    user_id integer,
    action character varying(255) NOT NULL,
    object_type character varying(255) NOT NULL,
    object_id character varying(255),
    additional_properties text,
    created_at timestamp with time zone NOT NULL
);


ALTER TABLE public.events OWNER TO postgres;

--
-- Name: events_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.events_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.events_id_seq OWNER TO postgres;

--
-- Name: events_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.events_id_seq OWNED BY public.events.id;


--
-- Name: favorites; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.favorites (
    updated_at timestamp with time zone NOT NULL,
    created_at timestamp with time zone NOT NULL,
    id integer NOT NULL,
    org_id integer NOT NULL,
    object_type character varying(255) NOT NULL,
    object_id integer NOT NULL,
    user_id integer NOT NULL
);


ALTER TABLE public.favorites OWNER TO postgres;

--
-- Name: favorites_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.favorites_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.favorites_id_seq OWNER TO postgres;

--
-- Name: favorites_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.favorites_id_seq OWNED BY public.favorites.id;


--
-- Name: groups; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.groups (
    id integer NOT NULL,
    org_id integer NOT NULL,
    type character varying(255) NOT NULL,
    name character varying(100) NOT NULL,
    permissions character varying(255)[] NOT NULL,
    created_at timestamp with time zone NOT NULL
);


ALTER TABLE public.groups OWNER TO postgres;

--
-- Name: groups_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.groups_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.groups_id_seq OWNER TO postgres;

--
-- Name: groups_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.groups_id_seq OWNED BY public.groups.id;


--
-- Name: notification_destinations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.notification_destinations (
    id integer NOT NULL,
    org_id integer NOT NULL,
    user_id integer NOT NULL,
    name character varying(255) NOT NULL,
    type character varying(255) NOT NULL,
    encrypted_options bytea NOT NULL,
    created_at timestamp with time zone NOT NULL
);


ALTER TABLE public.notification_destinations OWNER TO postgres;

--
-- Name: notification_destinations_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.notification_destinations_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.notification_destinations_id_seq OWNER TO postgres;

--
-- Name: notification_destinations_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.notification_destinations_id_seq OWNED BY public.notification_destinations.id;


--
-- Name: organizations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.organizations (
    updated_at timestamp with time zone NOT NULL,
    created_at timestamp with time zone NOT NULL,
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    slug character varying(255) NOT NULL,
    settings text NOT NULL
);


ALTER TABLE public.organizations OWNER TO postgres;

--
-- Name: organizations_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.organizations_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.organizations_id_seq OWNER TO postgres;

--
-- Name: organizations_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.organizations_id_seq OWNED BY public.organizations.id;


--
-- Name: queries; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.queries (
    updated_at timestamp with time zone NOT NULL,
    created_at timestamp with time zone NOT NULL,
    id integer NOT NULL,
    version integer NOT NULL,
    org_id integer NOT NULL,
    data_source_id integer,
    latest_query_data_id integer,
    name character varying(255) NOT NULL,
    description character varying(4096),
    query text NOT NULL,
    query_hash character varying(32) NOT NULL,
    api_key character varying(40) NOT NULL,
    user_id integer NOT NULL,
    last_modified_by_id integer,
    is_archived boolean NOT NULL,
    is_draft boolean NOT NULL,
    schedule text,
    schedule_failures integer NOT NULL,
    options text NOT NULL,
    search_vector tsvector,
    tags character varying[]
);


ALTER TABLE public.queries OWNER TO postgres;

--
-- Name: queries_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.queries_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.queries_id_seq OWNER TO postgres;

--
-- Name: queries_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.queries_id_seq OWNED BY public.queries.id;


--
-- Name: query_results; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.query_results (
    id integer NOT NULL,
    org_id integer NOT NULL,
    data_source_id integer NOT NULL,
    query_hash character varying(32) NOT NULL,
    query text NOT NULL,
    data text NOT NULL,
    runtime double precision NOT NULL,
    retrieved_at timestamp with time zone NOT NULL
);


ALTER TABLE public.query_results OWNER TO postgres;

--
-- Name: query_results_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.query_results_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.query_results_id_seq OWNER TO postgres;

--
-- Name: query_results_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.query_results_id_seq OWNED BY public.query_results.id;


--
-- Name: query_snippets; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.query_snippets (
    updated_at timestamp with time zone NOT NULL,
    created_at timestamp with time zone NOT NULL,
    id integer NOT NULL,
    org_id integer NOT NULL,
    trigger character varying(255) NOT NULL,
    description text NOT NULL,
    user_id integer NOT NULL,
    snippet text NOT NULL
);


ALTER TABLE public.query_snippets OWNER TO postgres;

--
-- Name: query_snippets_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.query_snippets_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.query_snippets_id_seq OWNER TO postgres;

--
-- Name: query_snippets_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.query_snippets_id_seq OWNED BY public.query_snippets.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    updated_at timestamp with time zone NOT NULL,
    created_at timestamp with time zone NOT NULL,
    id integer NOT NULL,
    org_id integer NOT NULL,
    name character varying(320) NOT NULL,
    email character varying(255) NOT NULL,
    password_hash character varying(128),
    groups integer[],
    api_key character varying(40) NOT NULL,
    disabled_at timestamp with time zone,
    details jsonb DEFAULT '{}'::jsonb
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


ALTER TABLE public.users_id_seq OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: visualizations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.visualizations (
    updated_at timestamp with time zone NOT NULL,
    created_at timestamp with time zone NOT NULL,
    id integer NOT NULL,
    type character varying(100) NOT NULL,
    query_id integer NOT NULL,
    name character varying(255) NOT NULL,
    description character varying(4096),
    options text NOT NULL
);


ALTER TABLE public.visualizations OWNER TO postgres;

--
-- Name: visualizations_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.visualizations_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.visualizations_id_seq OWNER TO postgres;

--
-- Name: visualizations_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.visualizations_id_seq OWNED BY public.visualizations.id;


--
-- Name: widgets; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.widgets (
    updated_at timestamp with time zone NOT NULL,
    created_at timestamp with time zone NOT NULL,
    id integer NOT NULL,
    visualization_id integer,
    text text,
    width integer NOT NULL,
    options text NOT NULL,
    dashboard_id integer NOT NULL
);


ALTER TABLE public.widgets OWNER TO postgres;

--
-- Name: widgets_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.widgets_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.widgets_id_seq OWNER TO postgres;

--
-- Name: widgets_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.widgets_id_seq OWNED BY public.widgets.id;


--
-- Name: access_permissions id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.access_permissions ALTER COLUMN id SET DEFAULT nextval('public.access_permissions_id_seq'::regclass);


--
-- Name: alert_subscriptions id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.alert_subscriptions ALTER COLUMN id SET DEFAULT nextval('public.alert_subscriptions_id_seq'::regclass);


--
-- Name: alerts id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.alerts ALTER COLUMN id SET DEFAULT nextval('public.alerts_id_seq'::regclass);


--
-- Name: api_keys id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.api_keys ALTER COLUMN id SET DEFAULT nextval('public.api_keys_id_seq'::regclass);


--
-- Name: changes id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.changes ALTER COLUMN id SET DEFAULT nextval('public.changes_id_seq'::regclass);


--
-- Name: dashboards id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.dashboards ALTER COLUMN id SET DEFAULT nextval('public.dashboards_id_seq'::regclass);


--
-- Name: data_source_groups id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.data_source_groups ALTER COLUMN id SET DEFAULT nextval('public.data_source_groups_id_seq'::regclass);


--
-- Name: data_sources id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.data_sources ALTER COLUMN id SET DEFAULT nextval('public.data_sources_id_seq'::regclass);


--
-- Name: events id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.events ALTER COLUMN id SET DEFAULT nextval('public.events_id_seq'::regclass);


--
-- Name: favorites id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.favorites ALTER COLUMN id SET DEFAULT nextval('public.favorites_id_seq'::regclass);


--
-- Name: groups id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.groups ALTER COLUMN id SET DEFAULT nextval('public.groups_id_seq'::regclass);


--
-- Name: notification_destinations id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.notification_destinations ALTER COLUMN id SET DEFAULT nextval('public.notification_destinations_id_seq'::regclass);


--
-- Name: organizations id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.organizations ALTER COLUMN id SET DEFAULT nextval('public.organizations_id_seq'::regclass);


--
-- Name: queries id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.queries ALTER COLUMN id SET DEFAULT nextval('public.queries_id_seq'::regclass);


--
-- Name: query_results id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.query_results ALTER COLUMN id SET DEFAULT nextval('public.query_results_id_seq'::regclass);


--
-- Name: query_snippets id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.query_snippets ALTER COLUMN id SET DEFAULT nextval('public.query_snippets_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Name: visualizations id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.visualizations ALTER COLUMN id SET DEFAULT nextval('public.visualizations_id_seq'::regclass);


--
-- Name: widgets id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.widgets ALTER COLUMN id SET DEFAULT nextval('public.widgets_id_seq'::regclass);


--
-- Data for Name: access_permissions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.access_permissions (object_type, object_id, id, access_type, grantor_id, grantee_id) FROM stdin;
\.


--
-- Data for Name: alembic_version; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.alembic_version (version_num) FROM stdin;
7ce5925f832b
\.


--
-- Data for Name: alert_subscriptions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.alert_subscriptions (updated_at, created_at, id, user_id, destination_id, alert_id) FROM stdin;
\.


--
-- Data for Name: alerts; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.alerts (updated_at, created_at, id, name, query_id, user_id, options, state, last_triggered_at, rearm) FROM stdin;
\.


--
-- Data for Name: api_keys; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.api_keys (object_type, updated_at, created_at, id, org_id, api_key, active, object_id, created_by_id) FROM stdin;
dashboards	2023-10-14 07:33:34.490852+00	2023-10-14 07:33:34.459226+00	1	1	DRx8HxYTZWwUGjFVmwWFtn0WEkweVYTx48MKcCNV	f	1	1
\.


--
-- Data for Name: changes; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.changes (object_type, id, object_id, object_version, user_id, change, created_at) FROM stdin;
queries	1	1	1	1	{"query": {"previous": "select 1", "current": "select 1"}, "org_id": {"previous": 1, "current": 1}, "data_source_id": {"previous": 1, "current": 1}, "latest_query_data_id": {"previous": null, "current": null}, "name": {"previous": "test-query-1", "current": "test-query-1"}, "description": {"previous": null, "current": null}, "query_hash": {"previous": null, "current": "f6bf37efedbc0a2dfffc1caf5088d86e"}, "api_key": {"previous": null, "current": "brt26N3QEtEeYFYB3a8rity46rxi77QsZmea5LhW"}, "user_id": {"previous": 1, "current": 1}, "last_modified_by_id": {"previous": 1, "current": 1}, "is_archived": {"previous": null, "current": false}, "is_draft": {"previous": true, "current": true}, "schedule": {"previous": null, "current": null}, "schedule_failures": {"previous": null, "current": 0}, "options": {"previous": null, "current": {}}, "search_vector": {"previous": null, "current": null}, "tags": {"previous": null, "current": null}}	2023-10-14 07:33:33.51963+00
queries	2	1	1	1	{"query": {"previous": "select 1", "current": "select 1"}, "org_id": {"previous": 1, "current": 1}, "data_source_id": {"previous": 1, "current": 1}, "latest_query_data_id": {"previous": null, "current": null}, "name": {"previous": "test-query-1", "current": "test-query-1"}, "description": {"previous": null, "current": null}, "query_hash": {"previous": "f6bf37efedbc0a2dfffc1caf5088d86e", "current": "f6bf37efedbc0a2dfffc1caf5088d86e"}, "api_key": {"previous": "brt26N3QEtEeYFYB3a8rity46rxi77QsZmea5LhW", "current": "brt26N3QEtEeYFYB3a8rity46rxi77QsZmea5LhW"}, "user_id": {"previous": 1, "current": 1}, "last_modified_by_id": {"previous": 1, "current": 1}, "is_archived": {"previous": true, "current": true}, "is_draft": {"previous": true, "current": true}, "schedule": {"previous": null, "current": null}, "schedule_failures": {"previous": 0, "current": 0}, "options": {"previous": {}, "current": {}}, "search_vector": {"previous": "'1':1B,5A,7 'query':4A 'select':6 'test':3A 'test-query':2A", "current": "'1':1B,5A,7 'query':4A 'select':6 'test':3A 'test-query':2A"}, "tags": {"previous": null, "current": null}}	2023-10-14 07:33:34.03114+00
dashboards	3	1	1	1	{"org_id": {"previous": null, "current": 1}, "slug": {"previous": null, "current": "test-dashboard-1"}, "name": {"previous": "test-dashboard-1", "current": "test-dashboard-1"}, "user_id": {"previous": null, "current": 1}, "layout": {"previous": null, "current": "[]"}, "dashboard_filters_enabled": {"previous": null, "current": false}, "is_archived": {"previous": null, "current": false}, "is_draft": {"previous": true, "current": true}, "tags": {"previous": null, "current": null}, "options": {"previous": null, "current": {}}}	2023-10-14 07:33:34.152657+00
dashboards	4	1	3	1	{"org_id": {"previous": 1, "current": 1}, "slug": {"previous": "test-dashboard-1", "current": "test-dashboard-1"}, "name": {"previous": "test-dashboard-2", "current": "test-dashboard-2"}, "user_id": {"previous": 1, "current": 1}, "layout": {"previous": "[]", "current": "[]"}, "dashboard_filters_enabled": {"previous": false, "current": false}, "is_archived": {"previous": true, "current": true}, "is_draft": {"previous": true, "current": true}, "tags": {"previous": ["foo"], "current": ["foo"]}, "options": {"previous": {}, "current": {}}}	2023-10-14 07:33:34.517436+00
queries	5	2	1	1	{"query": {"previous": "select 1", "current": "select 1"}, "org_id": {"previous": 1, "current": 1}, "data_source_id": {"previous": 4, "current": 4}, "latest_query_data_id": {"previous": null, "current": null}, "name": {"previous": "test-query-1", "current": "test-query-1"}, "description": {"previous": null, "current": null}, "query_hash": {"previous": null, "current": "f6bf37efedbc0a2dfffc1caf5088d86e"}, "api_key": {"previous": null, "current": "hfvbUrT6vyxmx1I7unK3FOEgPTtMZ46yj73lRmsk"}, "user_id": {"previous": 1, "current": 1}, "last_modified_by_id": {"previous": 1, "current": 1}, "is_archived": {"previous": null, "current": false}, "is_draft": {"previous": true, "current": true}, "schedule": {"previous": null, "current": null}, "schedule_failures": {"previous": null, "current": 0}, "options": {"previous": null, "current": {}}, "search_vector": {"previous": null, "current": null}, "tags": {"previous": ["my-tag-1"], "current": ["my-tag-1"]}}	2023-10-14 07:33:36.667332+00
queries	6	2	1	1	{"query": {"previous": "select 1", "current": "select 1"}, "org_id": {"previous": 1, "current": 1}, "data_source_id": {"previous": 4, "current": 4}, "latest_query_data_id": {"previous": 2, "current": 2}, "name": {"previous": "test-query-1", "current": "test-query-1"}, "description": {"previous": null, "current": null}, "query_hash": {"previous": "f6bf37efedbc0a2dfffc1caf5088d86e", "current": "f6bf37efedbc0a2dfffc1caf5088d86e"}, "api_key": {"previous": "hfvbUrT6vyxmx1I7unK3FOEgPTtMZ46yj73lRmsk", "current": "hfvbUrT6vyxmx1I7unK3FOEgPTtMZ46yj73lRmsk"}, "user_id": {"previous": 1, "current": 1}, "last_modified_by_id": {"previous": 1, "current": 1}, "is_archived": {"previous": true, "current": true}, "is_draft": {"previous": true, "current": true}, "schedule": {"previous": null, "current": null}, "schedule_failures": {"previous": 0, "current": 0}, "options": {"previous": {}, "current": {}}, "search_vector": {"previous": "'1':5A,7 '2':1B 'query':4A 'select':6 'test':3A 'test-query':2A", "current": "'1':5A,7 '2':1B 'query':4A 'select':6 'test':3A 'test-query':2A"}, "tags": {"previous": [], "current": []}}	2023-10-14 07:33:39.58571+00
queries	7	3	1	1	{"query": {"previous": "select 1", "current": "select 1"}, "org_id": {"previous": 1, "current": 1}, "data_source_id": {"previous": 5, "current": 5}, "latest_query_data_id": {"previous": null, "current": null}, "name": {"previous": "test-query-1", "current": "test-query-1"}, "description": {"previous": null, "current": null}, "query_hash": {"previous": null, "current": "f6bf37efedbc0a2dfffc1caf5088d86e"}, "api_key": {"previous": null, "current": "8snRrMODY6wFak4jL4toIv0QI2BRfl1KLf9SwB4P"}, "user_id": {"previous": 1, "current": 1}, "last_modified_by_id": {"previous": 1, "current": 1}, "is_archived": {"previous": null, "current": false}, "is_draft": {"previous": true, "current": true}, "schedule": {"previous": null, "current": null}, "schedule_failures": {"previous": null, "current": 0}, "options": {"previous": null, "current": {}}, "search_vector": {"previous": null, "current": null}, "tags": {"previous": null, "current": null}}	2023-10-14 07:33:40.123371+00
queries	8	3	1	1	{"query": {"previous": "select 1", "current": "select 1"}, "org_id": {"previous": 1, "current": 1}, "data_source_id": {"previous": 5, "current": 5}, "latest_query_data_id": {"previous": null, "current": null}, "name": {"previous": "test-query-1", "current": "test-query-1"}, "description": {"previous": null, "current": null}, "query_hash": {"previous": "f6bf37efedbc0a2dfffc1caf5088d86e", "current": "f6bf37efedbc0a2dfffc1caf5088d86e"}, "api_key": {"previous": "8snRrMODY6wFak4jL4toIv0QI2BRfl1KLf9SwB4P", "current": "8snRrMODY6wFak4jL4toIv0QI2BRfl1KLf9SwB4P"}, "user_id": {"previous": 1, "current": 1}, "last_modified_by_id": {"previous": 1, "current": 1}, "is_archived": {"previous": true, "current": true}, "is_draft": {"previous": true, "current": true}, "schedule": {"previous": null, "current": null}, "schedule_failures": {"previous": 0, "current": 0}, "options": {"previous": {}, "current": {}}, "search_vector": {"previous": "'1':5A,7 '3':1B 'query':4A 'select':6 'test':3A 'test-query':2A", "current": "'1':5A,7 '3':1B 'query':4A 'select':6 'test':3A 'test-query':2A"}, "tags": {"previous": null, "current": null}}	2023-10-14 07:33:40.198149+00
dashboards	9	2	1	1	{"org_id": {"previous": null, "current": 1}, "slug": {"previous": null, "current": "test-dashboard-1_1"}, "name": {"previous": "test-dashboard-1", "current": "test-dashboard-1"}, "user_id": {"previous": null, "current": 1}, "layout": {"previous": null, "current": "[]"}, "dashboard_filters_enabled": {"previous": null, "current": false}, "is_archived": {"previous": null, "current": false}, "is_draft": {"previous": true, "current": true}, "tags": {"previous": null, "current": null}, "options": {"previous": null, "current": {}}}	2023-10-14 07:33:40.254007+00
dashboards	10	2	2	1	{"org_id": {"previous": 1, "current": 1}, "slug": {"previous": "test-dashboard-1_1", "current": "test-dashboard-1_1"}, "name": {"previous": "test-dashboard-1", "current": "test-dashboard-1"}, "user_id": {"previous": 1, "current": 1}, "layout": {"previous": "[]", "current": "[]"}, "dashboard_filters_enabled": {"previous": false, "current": false}, "is_archived": {"previous": true, "current": true}, "is_draft": {"previous": true, "current": true}, "tags": {"previous": null, "current": null}, "options": {"previous": {}, "current": {}}}	2023-10-14 07:33:40.338634+00
\.


--
-- Data for Name: dashboards; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.dashboards (updated_at, created_at, id, version, org_id, slug, name, user_id, layout, dashboard_filters_enabled, is_archived, is_draft, tags, options) FROM stdin;
2023-10-14 07:33:34.517436+00	2023-10-14 07:33:34.152657+00	1	3	1	test-dashboard-1	test-dashboard-2	1	[]	f	t	t	{foo}	{}
2023-10-14 07:33:40.338634+00	2023-10-14 07:33:40.254007+00	2	2	1	test-dashboard-1_1	test-dashboard-1	1	[]	f	t	t	\N	{}
\.


--
-- Data for Name: data_source_groups; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.data_source_groups (id, data_source_id, group_id, view_only) FROM stdin;
\.


--
-- Data for Name: data_sources; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.data_sources (id, org_id, name, type, encrypted_options, queue_name, scheduled_queue_name, created_at) FROM stdin;
\.


--
-- Data for Name: events; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.events (id, org_id, user_id, action, object_type, object_id, additional_properties, created_at) FROM stdin;
1	1	1	login	redash	\N	{"user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36 Edg/118.0.2088.46", "ip": "172.21.0.1"}	2023-10-14 07:32:24+00
2	1	1	load_favorites	dashboard	\N	{"params": {"q": null, "tags": [], "page": 1}, "user_name": "admin", "user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36 Edg/118.0.2088.46", "ip": "172.21.0.1"}	2023-10-14 07:32:25+00
3	1	1	load_favorites	query	\N	{"params": {"q": null, "tags": [], "page": 1}, "user_name": "admin", "user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36 Edg/118.0.2088.46", "ip": "172.21.0.1"}	2023-10-14 07:32:25+00
4	1	1	view	page	personal_homepage	{"screen_resolution": "3440x1440", "user_name": "admin", "user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36 Edg/118.0.2088.46", "ip": "172.21.0.1"}	2023-10-14 07:32:25.005+00
5	1	1	list	query_snippet	\N	{"user_name": "admin", "user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36 Edg/118.0.2088.46", "ip": "172.21.0.1"}	2023-10-14 07:32:29+00
6	1	1	list	datasource	admin/data_sources	{"user_name": "admin", "user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36 Edg/118.0.2088.46", "ip": "172.21.0.1"}	2023-10-14 07:32:29+00
7	1	1	view_source	query	\N	{"screen_resolution": "3440x1440", "user_name": "admin", "user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36 Edg/118.0.2088.46", "ip": "172.21.0.1"}	2023-10-14 07:32:29.227+00
8	1	1	update_data_source	query	\N	{"screen_resolution": "3440x1440", "dataSourceId": null, "user_name": "admin", "user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36 Edg/118.0.2088.46", "ip": "172.21.0.1"}	2023-10-14 07:32:29.346+00
9	1	1	list	group	groups	{"user_name": "admin", "user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36 Edg/118.0.2088.46", "ip": "172.21.0.1"}	2023-10-14 07:32:41+00
10	1	1	list	alert	\N	{"user_name": "admin", "user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36 Edg/118.0.2088.46", "ip": "172.21.0.1"}	2023-10-14 07:32:43+00
11	1	1	list	datasource	admin/data_sources	{"user_name": "admin", "user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36 Edg/118.0.2088.46", "ip": "172.21.0.1"}	2023-10-14 07:32:45+00
12	1	1	list	user	\N	{"pending": false, "user_name": "admin", "user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36 Edg/118.0.2088.46", "ip": "172.21.0.1"}	2023-10-14 07:32:47+00
13	1	1	view	user	1	{"user_name": "admin", "user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36 Edg/118.0.2088.46", "ip": "172.21.0.1"}	2023-10-14 07:32:49+00
14	1	1	list	group	groups	{"user_name": "admin", "user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36 Edg/118.0.2088.46", "ip": "172.21.0.1"}	2023-10-14 07:32:49+00
15	1	1	list	outdated_queries	\N	{"user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:33+00
16	1	1	list	rq_status	\N	{"user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:33+00
17	1	1	create	datasource	1	{"user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:33+00
18	1	1	create	query	1	{"user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:33+00
19	1	1	list	alert	\N	{"user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:33+00
20	1	1	create	alert	1	{"user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:33+00
21	1	1	view	alert	1	{"user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:33+00
22	1	1	edit	alert	1	{"user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:33+00
23	1	1	edit	alert	1	{"user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:33+00
24	1	1	edit	alert	1	{"user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:33+00
25	1	1	mute	alert	1	{"user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:33+00
26	1	1	view	alert	1	{"user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:33+00
27	1	1	unmute	alert	1	{"user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:33+00
28	1	1	view	alert	1	{"user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:33+00
29	1	1	delete	datasource	1	{"user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:34+00
30	1	1	list	dashboard	\N	{"user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:34+00
31	1	1	view	dashboard	1	{"user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:34+00
32	1	1	edit	dashboard	1	{"user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:34+00
33	1	1	search	dashboard	\N	{"term": "test-dashboard-2", "user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:34+00
34	1	1	load_favorites	dashboard	\N	{"params": {"q": "test-dashboard-2", "tags": [], "page": 1}, "user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:34+00
35	1	1	favorite	dashboard	1	{"user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:34+00
36	1	1	load_favorites	dashboard	\N	{"params": {"q": "test-dashboard-2", "tags": [], "page": 1}, "user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:34+00
37	1	1	activate_api_key	dashboard	1	{"user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:34+00
38	1	1	deactivate_api_key	dashboard	1	{"user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:34+00
39	1	1	archive	dashboard	1	{"user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:34+00
40	1	1	view	dashboard	1	{"user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:34+00
41	1	1	list	datasource	admin/data_sources	{"user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:34+00
42	1	1	create	datasource	2	{"user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:34+00
43	1	1	test	datasource	2	{"result": {"message": "success", "ok": true}, "user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:35+00
44	1	1	view	datasource	2	{"user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:35+00
45	1	1	edit	datasource	2	{"user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:35+00
46	1	1	pause	datasource	2	{"user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:35+00
47	1	1	resume	datasource	2	{"user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:35+00
48	1	1	delete	datasource	2	{"user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:35+00
49	1	1	list	destination	admin/destinations	{"user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:35+00
50	1	1	view	destination	1	{"user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:35+00
51	1	1	delete	destination	1	{"user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:35+00
52	1	1	create	datasource	3	{"user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:36+00
53	1	1	delete	datasource	3	{"user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:36+00
54	1	1	list	group	groups	{"user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:36+00
55	1	1	create	group	3	{"user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:36+00
56	1	1	view	group	3	{"user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:36+00
57	1	1	list	group	3	{"user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:36+00
58	1	1	list	query_snippet	\N	{"user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:36+00
59	1	1	create	query_snippet	1	{"user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:36+00
60	1	1	view	query_snippet	1	{"user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:36+00
61	1	1	edit	query_snippet	1	{"user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:36+00
62	1	1	delete	query_snippet	1	{"user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:36+00
63	1	1	create	datasource	4	{"user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:36+00
64	1	1	list	query	\N	{"user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:36+00
65	1	1	create	query	2	{"user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:36+00
66	1	1	view	query	2	{"user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:36+00
67	1	1	search	query	\N	{"term": "test-query-1", "user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:36+00
68	1	1	search	query	\N	{"term": "test-query-1", "user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:36+00
69	1	1	search	query	\N	{"term": "test-query-1", "user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:37+00
70	1	1	load_favorites	query	\N	{"params": {"q": "test-query-1", "tags": [], "page": 1}, "user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:37+00
71	1	1	favorite	query	2	{"user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:37+00
72	1	1	load_favorites	query	\N	{"params": {"q": "test-query-1", "tags": [], "page": 1}, "user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:37+00
73	1	1	execute_query	data_source	4	{"cache": "miss", "query": "select 1", "query_id": "2", "parameters": {}, "user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:37+00
74	1	1	execute_query	data_source	4	{"cache": "hit", "query": "select 1", "query_id": "2", "parameters": {}, "user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:38+00
75	1	1	execute_query	data_source	4	{"cache": "miss", "query": "select 1", "query_id": 2, "parameters": {}, "user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:38+00
76	1	1	view	query	2	{"user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:39+00
77	1	1	delete	datasource	4	{"user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:39+00
78	1	1	edit	settings	1	{"new_values": {"date_format": "YYYY/MM/DD"}, "previous_values": {"date_format": "DD/MM/YY"}, "user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:39+00
79	1	1	list	user	\N	{"pending": null, "user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:39+00
80	1	1	create	user	2	{"user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:39+00
81	1	1	view	user	2	{"user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:39+00
82	1	1	edit	user	2	{"updated_fields": ["email", "name"], "user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:39+00
83	1	1	create	datasource	5	{"user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:40+00
84	1	1	create	query	3	{"user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:40+00
85	1	1	delete	datasource	5	{"user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:40+00
86	1	1	delete	widget	1	{"user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:40+00
87	1	1	archive	dashboard	2	{"user_name": "admin", "user_agent": "redash-go", "ip": "172.21.0.1"}	2023-10-14 07:33:40+00
\.


--
-- Data for Name: favorites; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.favorites (updated_at, created_at, id, org_id, object_type, object_id, user_id) FROM stdin;
2023-10-14 07:33:34.390094+00	2023-10-14 07:33:34.390094+00	1	1	Dashboard	1	1
2023-10-14 07:33:37.14076+00	2023-10-14 07:33:37.14076+00	2	1	Query	2	1
\.


--
-- Data for Name: groups; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.groups (id, org_id, type, name, permissions, created_at) FROM stdin;
1	1	builtin	admin	{admin,super_admin}	2023-10-14 07:32:24.00699+00
2	1	builtin	default	{create_dashboard,create_query,edit_dashboard,edit_query,view_query,view_source,execute_query,list_users,schedule_query,list_dashboards,list_alerts,list_data_sources}	2023-10-14 07:32:24.00699+00
\.


--
-- Data for Name: notification_destinations; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.notification_destinations (id, org_id, user_id, name, type, encrypted_options, created_at) FROM stdin;
\.


--
-- Data for Name: organizations; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.organizations (updated_at, created_at, id, name, slug, settings) FROM stdin;
2023-10-14 07:33:40.254007+00	2023-10-14 07:32:24.00699+00	1	my-org	default	{"settings": {"date_format": "YYYY/MM/DD"}}
\.


--
-- Data for Name: queries; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.queries (updated_at, created_at, id, version, org_id, data_source_id, latest_query_data_id, name, description, query, query_hash, api_key, user_id, last_modified_by_id, is_archived, is_draft, schedule, schedule_failures, options, search_vector, tags) FROM stdin;
2023-10-14 07:33:34.03114+00	2023-10-14 07:33:33.51963+00	1	1	1	\N	\N	test-query-1	\N	select 1	f6bf37efedbc0a2dfffc1caf5088d86e	brt26N3QEtEeYFYB3a8rity46rxi77QsZmea5LhW	1	1	t	t	\N	0	{}	'1':1B,5A,7 'query':4A 'select':6 'test':3A 'test-query':2A	\N
2023-10-14 07:33:39.58571+00	2023-10-14 07:33:36.667332+00	2	1	1	\N	\N	test-query-1	\N	select 1	f6bf37efedbc0a2dfffc1caf5088d86e	hfvbUrT6vyxmx1I7unK3FOEgPTtMZ46yj73lRmsk	1	1	t	t	\N	0	{}	'1':5A,7 '2':1B 'query':4A 'select':6 'test':3A 'test-query':2A	{}
2023-10-14 07:33:40.198149+00	2023-10-14 07:33:40.123371+00	3	1	1	\N	\N	test-query-1	\N	select 1	f6bf37efedbc0a2dfffc1caf5088d86e	8snRrMODY6wFak4jL4toIv0QI2BRfl1KLf9SwB4P	1	1	t	t	\N	0	{}	'1':5A,7 '3':1B 'query':4A 'select':6 'test':3A 'test-query':2A	\N
\.


--
-- Data for Name: query_results; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.query_results (id, org_id, data_source_id, query_hash, query, data, runtime, retrieved_at) FROM stdin;
\.


--
-- Data for Name: query_snippets; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.query_snippets (updated_at, created_at, id, org_id, trigger, description, user_id, snippet) FROM stdin;
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (updated_at, created_at, id, org_id, name, email, password_hash, groups, api_key, disabled_at, details) FROM stdin;
2023-10-14 07:33:45.847774+00	2023-10-14 07:32:24.321326+00	1	1	admin	admin@example.com	$6$rounds=656000$wKK0R42uBz0FJ91/$Xi7PHNA2mbQEn/QQwOQc0/FzN7fhi1.amDuzIMj2R4pTOvSsiw5LXLbXj.GeUb7mdW8f8pZBoj.CFETOMUWcx1	{1,2}	6nh64ZsT66WeVJvNZ6WB5D2JKZULeC2VBdSD68wt	\N	{"active_at": "2023-10-14T07:33:40Z"}
\.


--
-- Data for Name: visualizations; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.visualizations (updated_at, created_at, id, type, query_id, name, description, options) FROM stdin;
2023-10-14 07:33:33.51963+00	2023-10-14 07:33:33.51963+00	1	TABLE	1	Table		{}
2023-10-14 07:33:36.667332+00	2023-10-14 07:33:36.667332+00	2	TABLE	2	Table		{}
2023-10-14 07:33:40.172346+00	2023-10-14 07:33:40.123371+00	3	TABLE	3	test-viz-1	test-viz-1-desc	{}
\.


--
-- Data for Name: widgets; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.widgets (updated_at, created_at, id, visualization_id, text, width, options, dashboard_id) FROM stdin;
\.


--
-- Name: access_permissions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.access_permissions_id_seq', 1, false);


--
-- Name: alert_subscriptions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.alert_subscriptions_id_seq', 1, false);


--
-- Name: alerts_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.alerts_id_seq', 1, true);


--
-- Name: api_keys_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.api_keys_id_seq', 1, true);


--
-- Name: changes_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.changes_id_seq', 10, true);


--
-- Name: dashboards_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.dashboards_id_seq', 2, true);


--
-- Name: data_source_groups_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.data_source_groups_id_seq', 5, true);


--
-- Name: data_sources_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.data_sources_id_seq', 5, true);


--
-- Name: events_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.events_id_seq', 87, true);


--
-- Name: favorites_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.favorites_id_seq', 2, true);


--
-- Name: groups_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.groups_id_seq', 3, true);


--
-- Name: notification_destinations_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.notification_destinations_id_seq', 1, true);


--
-- Name: organizations_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.organizations_id_seq', 1, true);


--
-- Name: queries_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.queries_id_seq', 3, true);


--
-- Name: query_results_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.query_results_id_seq', 2, true);


--
-- Name: query_snippets_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.query_snippets_id_seq', 1, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 2, true);


--
-- Name: visualizations_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.visualizations_id_seq', 3, true);


--
-- Name: widgets_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.widgets_id_seq', 1, true);


--
-- Name: access_permissions access_permissions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.access_permissions
    ADD CONSTRAINT access_permissions_pkey PRIMARY KEY (id);


--
-- Name: alembic_version alembic_version_pkc; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.alembic_version
    ADD CONSTRAINT alembic_version_pkc PRIMARY KEY (version_num);


--
-- Name: alert_subscriptions alert_subscriptions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.alert_subscriptions
    ADD CONSTRAINT alert_subscriptions_pkey PRIMARY KEY (id);


--
-- Name: alerts alerts_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.alerts
    ADD CONSTRAINT alerts_pkey PRIMARY KEY (id);


--
-- Name: api_keys api_keys_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.api_keys
    ADD CONSTRAINT api_keys_pkey PRIMARY KEY (id);


--
-- Name: changes changes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.changes
    ADD CONSTRAINT changes_pkey PRIMARY KEY (id);


--
-- Name: dashboards dashboards_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.dashboards
    ADD CONSTRAINT dashboards_pkey PRIMARY KEY (id);


--
-- Name: data_source_groups data_source_groups_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.data_source_groups
    ADD CONSTRAINT data_source_groups_pkey PRIMARY KEY (id);


--
-- Name: data_sources data_sources_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.data_sources
    ADD CONSTRAINT data_sources_pkey PRIMARY KEY (id);


--
-- Name: events events_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.events
    ADD CONSTRAINT events_pkey PRIMARY KEY (id);


--
-- Name: favorites favorites_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.favorites
    ADD CONSTRAINT favorites_pkey PRIMARY KEY (id);


--
-- Name: groups groups_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.groups
    ADD CONSTRAINT groups_pkey PRIMARY KEY (id);


--
-- Name: notification_destinations notification_destinations_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.notification_destinations
    ADD CONSTRAINT notification_destinations_pkey PRIMARY KEY (id);


--
-- Name: organizations organizations_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.organizations
    ADD CONSTRAINT organizations_pkey PRIMARY KEY (id);


--
-- Name: organizations organizations_slug_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.organizations
    ADD CONSTRAINT organizations_slug_key UNIQUE (slug);


--
-- Name: queries queries_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.queries
    ADD CONSTRAINT queries_pkey PRIMARY KEY (id);


--
-- Name: query_results query_results_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.query_results
    ADD CONSTRAINT query_results_pkey PRIMARY KEY (id);


--
-- Name: query_snippets query_snippets_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.query_snippets
    ADD CONSTRAINT query_snippets_pkey PRIMARY KEY (id);


--
-- Name: query_snippets query_snippets_trigger_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.query_snippets
    ADD CONSTRAINT query_snippets_trigger_key UNIQUE (trigger);


--
-- Name: favorites unique_favorite; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.favorites
    ADD CONSTRAINT unique_favorite UNIQUE (object_type, object_id, user_id);


--
-- Name: users users_api_key_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_api_key_key UNIQUE (api_key);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: visualizations visualizations_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.visualizations
    ADD CONSTRAINT visualizations_pkey PRIMARY KEY (id);


--
-- Name: widgets widgets_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.widgets
    ADD CONSTRAINT widgets_pkey PRIMARY KEY (id);


--
-- Name: alert_subscriptions_destination_id_alert_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX alert_subscriptions_destination_id_alert_id ON public.alert_subscriptions USING btree (destination_id, alert_id);


--
-- Name: api_keys_object_type_object_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX api_keys_object_type_object_id ON public.api_keys USING btree (object_type, object_id);


--
-- Name: data_sources_org_id_name; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX data_sources_org_id_name ON public.data_sources USING btree (org_id, name);


--
-- Name: ix_api_keys_api_key; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX ix_api_keys_api_key ON public.api_keys USING btree (api_key);


--
-- Name: ix_dashboards_is_archived; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX ix_dashboards_is_archived ON public.dashboards USING btree (is_archived);


--
-- Name: ix_dashboards_is_draft; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX ix_dashboards_is_draft ON public.dashboards USING btree (is_draft);


--
-- Name: ix_dashboards_slug; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX ix_dashboards_slug ON public.dashboards USING btree (slug);


--
-- Name: ix_queries_is_archived; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX ix_queries_is_archived ON public.queries USING btree (is_archived);


--
-- Name: ix_queries_is_draft; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX ix_queries_is_draft ON public.queries USING btree (is_draft);


--
-- Name: ix_queries_search_vector; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX ix_queries_search_vector ON public.queries USING gin (search_vector);


--
-- Name: ix_query_results_query_hash; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX ix_query_results_query_hash ON public.query_results USING btree (query_hash);


--
-- Name: ix_widgets_dashboard_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX ix_widgets_dashboard_id ON public.widgets USING btree (dashboard_id);


--
-- Name: notification_destinations_org_id_name; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX notification_destinations_org_id_name ON public.notification_destinations USING btree (org_id, name);


--
-- Name: users_org_id_email; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX users_org_id_email ON public.users USING btree (org_id, email);


--
-- Name: queries queries_search_vector_trigger; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER queries_search_vector_trigger BEFORE INSERT OR UPDATE ON public.queries FOR EACH ROW EXECUTE FUNCTION public.queries_search_vector_update();


--
-- Name: access_permissions access_permissions_grantee_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.access_permissions
    ADD CONSTRAINT access_permissions_grantee_id_fkey FOREIGN KEY (grantee_id) REFERENCES public.users(id);


--
-- Name: access_permissions access_permissions_grantor_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.access_permissions
    ADD CONSTRAINT access_permissions_grantor_id_fkey FOREIGN KEY (grantor_id) REFERENCES public.users(id);


--
-- Name: alert_subscriptions alert_subscriptions_alert_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.alert_subscriptions
    ADD CONSTRAINT alert_subscriptions_alert_id_fkey FOREIGN KEY (alert_id) REFERENCES public.alerts(id);


--
-- Name: alert_subscriptions alert_subscriptions_destination_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.alert_subscriptions
    ADD CONSTRAINT alert_subscriptions_destination_id_fkey FOREIGN KEY (destination_id) REFERENCES public.notification_destinations(id);


--
-- Name: alert_subscriptions alert_subscriptions_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.alert_subscriptions
    ADD CONSTRAINT alert_subscriptions_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: alerts alerts_query_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.alerts
    ADD CONSTRAINT alerts_query_id_fkey FOREIGN KEY (query_id) REFERENCES public.queries(id);


--
-- Name: alerts alerts_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.alerts
    ADD CONSTRAINT alerts_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: api_keys api_keys_created_by_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.api_keys
    ADD CONSTRAINT api_keys_created_by_id_fkey FOREIGN KEY (created_by_id) REFERENCES public.users(id);


--
-- Name: api_keys api_keys_org_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.api_keys
    ADD CONSTRAINT api_keys_org_id_fkey FOREIGN KEY (org_id) REFERENCES public.organizations(id);


--
-- Name: changes changes_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.changes
    ADD CONSTRAINT changes_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: dashboards dashboards_org_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.dashboards
    ADD CONSTRAINT dashboards_org_id_fkey FOREIGN KEY (org_id) REFERENCES public.organizations(id);


--
-- Name: dashboards dashboards_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.dashboards
    ADD CONSTRAINT dashboards_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: data_source_groups data_source_groups_data_source_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.data_source_groups
    ADD CONSTRAINT data_source_groups_data_source_id_fkey FOREIGN KEY (data_source_id) REFERENCES public.data_sources(id);


--
-- Name: data_source_groups data_source_groups_group_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.data_source_groups
    ADD CONSTRAINT data_source_groups_group_id_fkey FOREIGN KEY (group_id) REFERENCES public.groups(id);


--
-- Name: data_sources data_sources_org_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.data_sources
    ADD CONSTRAINT data_sources_org_id_fkey FOREIGN KEY (org_id) REFERENCES public.organizations(id);


--
-- Name: events events_org_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.events
    ADD CONSTRAINT events_org_id_fkey FOREIGN KEY (org_id) REFERENCES public.organizations(id);


--
-- Name: events events_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.events
    ADD CONSTRAINT events_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: favorites favorites_org_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.favorites
    ADD CONSTRAINT favorites_org_id_fkey FOREIGN KEY (org_id) REFERENCES public.organizations(id);


--
-- Name: favorites favorites_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.favorites
    ADD CONSTRAINT favorites_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: groups groups_org_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.groups
    ADD CONSTRAINT groups_org_id_fkey FOREIGN KEY (org_id) REFERENCES public.organizations(id);


--
-- Name: notification_destinations notification_destinations_org_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.notification_destinations
    ADD CONSTRAINT notification_destinations_org_id_fkey FOREIGN KEY (org_id) REFERENCES public.organizations(id);


--
-- Name: notification_destinations notification_destinations_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.notification_destinations
    ADD CONSTRAINT notification_destinations_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: queries queries_data_source_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.queries
    ADD CONSTRAINT queries_data_source_id_fkey FOREIGN KEY (data_source_id) REFERENCES public.data_sources(id);


--
-- Name: queries queries_last_modified_by_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.queries
    ADD CONSTRAINT queries_last_modified_by_id_fkey FOREIGN KEY (last_modified_by_id) REFERENCES public.users(id);


--
-- Name: queries queries_latest_query_data_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.queries
    ADD CONSTRAINT queries_latest_query_data_id_fkey FOREIGN KEY (latest_query_data_id) REFERENCES public.query_results(id);


--
-- Name: queries queries_org_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.queries
    ADD CONSTRAINT queries_org_id_fkey FOREIGN KEY (org_id) REFERENCES public.organizations(id);


--
-- Name: queries queries_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.queries
    ADD CONSTRAINT queries_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: query_results query_results_data_source_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.query_results
    ADD CONSTRAINT query_results_data_source_id_fkey FOREIGN KEY (data_source_id) REFERENCES public.data_sources(id);


--
-- Name: query_results query_results_org_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.query_results
    ADD CONSTRAINT query_results_org_id_fkey FOREIGN KEY (org_id) REFERENCES public.organizations(id);


--
-- Name: query_snippets query_snippets_org_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.query_snippets
    ADD CONSTRAINT query_snippets_org_id_fkey FOREIGN KEY (org_id) REFERENCES public.organizations(id);


--
-- Name: query_snippets query_snippets_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.query_snippets
    ADD CONSTRAINT query_snippets_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: users users_org_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_org_id_fkey FOREIGN KEY (org_id) REFERENCES public.organizations(id);


--
-- Name: visualizations visualizations_query_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.visualizations
    ADD CONSTRAINT visualizations_query_id_fkey FOREIGN KEY (query_id) REFERENCES public.queries(id);


--
-- Name: widgets widgets_dashboard_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.widgets
    ADD CONSTRAINT widgets_dashboard_id_fkey FOREIGN KEY (dashboard_id) REFERENCES public.dashboards(id);


--
-- Name: widgets widgets_visualization_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.widgets
    ADD CONSTRAINT widgets_visualization_id_fkey FOREIGN KEY (visualization_id) REFERENCES public.visualizations(id);


--
-- PostgreSQL database dump complete
--

