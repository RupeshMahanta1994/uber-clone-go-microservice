CREATE TYPE driver_status AS ENUM ('offline', 'idle', 'on_trip');

CREATE TABLE drivers (
user_id UUID PRIMARY KEY,
vehicle_number TEXT NOT NULL,
vehicle_type TEXT NOT NULL,
status driver_status DEFAULT 'offline',
is_available BOOLEAN DEFAULT false,
rating FLOAT DEFAULT 5.0,
created_at TIMESTAMP DEFAULT now(),
updated_at TIMESTAMP DEFAULT now()
);
