
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
DROP TABLE location, car;

CREATE TABLE car (
                     id SERIAL PRIMARY KEY,
                     name VARCHAR,
                     uuid UUID DEFAULT uuid_generate_v4() UNIQUE NOT NULL,
                     is_active BOOLEAN DEFAULT TRUE
);

CREATE TABLE location (
                          id SERIAL PRIMARY KEY,
                          lat DOUBLE PRECISION NOT NULL,
                          lon DOUBLE PRECISION NOT NULL,
                          created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
                          car_uuid UUID REFERENCES car (uuid) ON DELETE CASCADE
);

CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       username VARCHAR,
                       password VARCHAR
);

INSERT INTO users (username, password)
VALUES ('admin', 'admin'),
       ('test', 'test');

select username, password from users;
SELECT lat, lon FROM location WHERE car_uuid ='550e8400-e29b-41d4-a716-446655440000' LIMIT 5;

INSERT INTO car (name, uuid, is_active)
VALUES ('BMW X5', '550e8400-e29b-41d4-a716-446655440000', true),
       ('Porsche 911', '6ba7b810-9dad-11d1-80b4-00c04fd430c8', true),
       ('Renauld Sandero', 'b3555e77-e655-4d55-a65f-8549f45498da', true),
       ('Kia Rio', 'd650b5c7-3596-40e7-aa76-0328bc729bc4', true),
       ('Kia Soul', 'bffeb71a-8357-4d7b-a9b9-738581309e4e', true),
       ('Porsche Panamera', '11278be1-e247-492c-846e-5d1b29372d8b', false),
       ('Ferrari La Ferrari', '06da6936-7957-462e-b0ce-385e044c08e1', false);

select  * from car;
select name, uuid, car.is_active from car;

-- INSERT INTO location (lat, lon, created_at, car_id)
-- VALUES (40.7128, -74.0060, '2023-01-15 13:00:00', 1),
--        (34.0522, -118.2437, '2023-01-16 14:30:00', 2);

INSERT INTO location (lat, lon, car_uuid)
VALUES   (59.9343, 30.3351, '550e8400-e29b-41d4-a716-446655440000'),
         (59.9398, 30.3146, '6ba7b810-9dad-11d1-80b4-00c04fd430c8'),
         (59.9178, 30.4980, 'b3555e77-e655-4d55-a65f-8549f45498da'),
         (59.9390, 30.3158, 'd650b5c7-3596-40e7-aa76-0328bc729bc4'),
         (59.9169, 30.3359, 'bffeb71a-8357-4d7b-a9b9-738581309e4e'),
         (59.9343, 30.3351, '11278be1-e247-492c-846e-5d1b29372d8b'),
         (59.9398, 30.3146, '06da6936-7957-462e-b0ce-385e044c08e1'),
         (59.9178, 30.4980, '550e8400-e29b-41d4-a716-446655440000'),
         (59.9390, 30.3158, '6ba7b810-9dad-11d1-80b4-00c04fd430c8'),
         (59.9169, 30.3359, 'b3555e77-e655-4d55-a65f-8549f45498da'),
         (59.9343, 30.3351, 'd650b5c7-3596-40e7-aa76-0328bc729bc4'),
         (59.9398, 30.3146, 'bffeb71a-8357-4d7b-a9b9-738581309e4e'),
         (59.9178, 30.4980, '11278be1-e247-492c-846e-5d1b29372d8b'),
         (59.9390, 30.3158, '06da6936-7957-462e-b0ce-385e044c08e1'),
         (59.9169, 30.3359, '550e8400-e29b-41d4-a716-446655440000'),
         (59.9343, 30.3351, '6ba7b810-9dad-11d1-80b4-00c04fd430c8'),
         (59.9398, 30.3146, 'b3555e77-e655-4d55-a65f-8549f45498da'),
         (59.9178, 30.4980, 'd650b5c7-3596-40e7-aa76-0328bc729bc4'),
         (59.9390, 30.3158, 'bffeb71a-8357-4d7b-a9b9-738581309e4e'),
         (59.9169, 30.3359,'11278be1-e247492c846e5d1b29372d8b');


select * from location;


-- /location
SELECT c.name, c.uuid, c.is_active, l.lat, l.lon, l.created_at
FROM car c
         LEFT JOIN (
    SELECT car_uuid, lat, lon, created_at,
           ROW_NUMBER() OVER (PARTITION BY car_uuid ORDER BY created_at DESC) AS rn
    FROM location
) l ON c.uuid = l.car_uuid AND l.rn = 1;

-- /story/{uuid}
SELECT c.name, c.uuid, c.is_active, l.lat, l.lon, l.created_at
FROM car c
         JOIN (
    SELECT car_uuid, lat, lon, created_at,
           ROW_NUMBER() OVER (PARTITION BY car_uuid ORDER BY created_at DESC) AS rn
    FROM location
) l ON c.uuid = l.car_uuid AND l.rn <= 5
WHERE c.uuid = '550e8400-e29b-41d4-a716-446655440000';

SELECT lat, lon FROM location WHERE car_uuid = '550e8400-e29b-41d4-a716-446655440000';
