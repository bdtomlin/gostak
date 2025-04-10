-- +goose Up
-- +goose StatementBegin
-- replace with built in uuid function when postgresql 18 is released
create or replace function generate_uuid_v7()
returns uuid
as
    $$
DECLARE
    unix_ms BIGINT;
    random_bits BIGINT;
    uuid_hex TEXT;
BEGIN
    -- Get Unix timestamp in milliseconds
    unix_ms := EXTRACT(EPOCH FROM CURRENT_TIMESTAMP) * 1000;

    -- Generate 64 random bits (PostgreSQL's random() gives a float between 0 and 1)
    random_bits := FLOOR(random() * (2^62)::BIGINT)::BIGINT;

    -- Construct the UUID:
    -- - 48 bits timestamp (unix_ms)
    -- - 4 bits version (7)
    -- - 12 bits counter/random (part of random_bits)
    -- - 2 bits variant (10)
    -- - 62 bits random (remaining random_bits)
    uuid_hex := format(
        '%s%s', -- Combine timestamp and random parts
        lpad(to_hex(unix_ms)::TEXT, 12, '0'), -- 48-bit timestamp (12 hex chars)
        lpad(to_hex((7::BIGINT << 12) | (random_bits & 4095))::TEXT, 4, '0') || -- version 7 + 12-bit sequence
        lpad(to_hex((2::BIGINT << 62) | (random_bits >> 12))::TEXT, 16, '0') -- variant 10 + 62-bit random
    );

    -- Remove any dashes and ensure exactly 32 hex chars
    uuid_hex := substring(uuid_hex from 1 for 32);

    -- Convert to UUID format with dashes
    RETURN (uuid_hex::UUID);
EXCEPTION
    WHEN OTHERS THEN
        RAISE NOTICE 'Error generating UUID v7: %', SQLERRM;
        RETURN gen_random_uuid(); -- Fallback to v4 if something goes wrong
END;
$$
language plpgsql
;
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
drop function if exists generate_uuid_v7()
;
-- +goose StatementEnd


