-- +goose Up
-- +goose StatementBegin
-- replace with built in uuid function when postgresql 18 is released
/* See the UUID Version 7 specification at
   https://www.rfc-editor.org/rfc/rfc9562#name-uuid-version-7 */
/* Main function to generate a uuidv7 value with millisecond precision */
create function uuidv7(timestamptz default clock_timestamp())
returns uuid
as $$
  -- Replace the first 48 bits of a uuidv4 with the current
  -- number of milliseconds since 1970-01-01 UTC
  -- and set the "ver" field to 7 by setting additional bits
  select encode(
    set_bit(
      set_bit(
        overlay(uuid_send(gen_random_uuid()) placing
	  substring(int8send((extract(epoch from $1)*1000)::bigint) from 3)
	  from 1 for 6),
	52, 1),
      53, 1), 'hex')::uuid;
$$
language sql
volatile
parallel safe
;

COMMENT ON FUNCTION uuidv7(timestamptz) IS
'Generate a uuid-v7 value with a 48-bit timestamp (millisecond precision) and 74 bits of randomness';

/* Extract the timestamp in the first 6 bytes of the uuidv7 value.
   Use the fact that 'xHHHHH' (where HHHHH are hexadecimal numbers)
   can be cast to bit(N) and then to int8.
 */
create function uuidv7_extract_timestamp(uuid)
returns timestamptz
as $$
 select to_timestamp(
   right(substring(uuid_send($1) from 1 for 6)::text, -1)::bit(48)::int8 -- milliseconds
    /1000.0);
$$
language sql
immutable
strict
parallel safe
;

COMMENT ON FUNCTION uuidv7_extract_timestamp(uuid) IS
'Return the timestamp stored in the first 48 bits of the UUID v7 value';

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
drop function if exists uuidv7(timestamptz)
;
drop function if exists uuidv7_extract_timestamp(uuid)
;
-- +goose StatementEnd


