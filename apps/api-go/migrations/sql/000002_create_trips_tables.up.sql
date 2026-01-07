-- Create trips table
CREATE TABLE IF NOT EXISTS trips (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    owner_id UUID NOT NULL,
    title VARCHAR(255) NOT NULL,
    destination VARCHAR(255) NOT NULL,
    start_date TIMESTAMP NOT NULL,
    end_date TIMESTAMP NOT NULL,
    travelers INTEGER DEFAULT 1,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Create indexes
CREATE INDEX idx_trips_owner_id ON trips(owner_id);
CREATE INDEX idx_trips_deleted_at ON trips(deleted_at);

-- Create itinerary_days table
CREATE TABLE IF NOT EXISTS itinerary_days (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    trip_id UUID NOT NULL,
    date TIMESTAMP NOT NULL,
    day_number INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT fk_itinerary_days_trip FOREIGN KEY (trip_id) REFERENCES trips(id) ON DELETE CASCADE
);

-- Create indexes
CREATE INDEX idx_itinerary_days_trip_id ON itinerary_days(trip_id);
CREATE INDEX idx_itinerary_days_deleted_at ON itinerary_days(deleted_at);

-- Create places table
CREATE TABLE IF NOT EXISTS places (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    google_place_id VARCHAR(255) NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    rating DECIMAL(2, 1),
    review_count INTEGER,
    primary_photo_url TEXT,
    address TEXT,
    formatted_address TEXT,
    website TEXT,
    phone_number VARCHAR(50),
    price_level INTEGER,
    last_fetched_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Create indexes
CREATE UNIQUE INDEX idx_places_google_place_id ON places(google_place_id);
CREATE INDEX idx_places_deleted_at ON places(deleted_at);

-- Create activities table
CREATE TABLE IF NOT EXISTS activities (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    itinerary_day_id UUID NOT NULL,
    place_id UUID,
    type VARCHAR(50) NOT NULL DEFAULT 'place_based',
    time TIMESTAMP NOT NULL,
    title VARCHAR(255) NOT NULL,
    location VARCHAR(255),
    category VARCHAR(50) NOT NULL,
    description TEXT,
    notes TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT fk_activities_itinerary_day FOREIGN KEY (itinerary_day_id) REFERENCES itinerary_days(id) ON DELETE CASCADE,
    CONSTRAINT fk_activities_place FOREIGN KEY (place_id) REFERENCES places(id) ON DELETE SET NULL
);

-- Create indexes
CREATE INDEX idx_activities_itinerary_day_id ON activities(itinerary_day_id);
CREATE INDEX idx_activities_place_id ON activities(place_id);
CREATE INDEX idx_activities_deleted_at ON activities(deleted_at);

-- Create trip_collaborators table
CREATE TABLE IF NOT EXISTS trip_collaborators (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    trip_id UUID NOT NULL,
    user_id UUID NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT fk_trip_collaborators_trip FOREIGN KEY (trip_id) REFERENCES trips(id) ON DELETE CASCADE,
    CONSTRAINT fk_trip_collaborators_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Create indexes
CREATE INDEX idx_trip_collaborators_trip_id ON trip_collaborators(trip_id);
CREATE INDEX idx_trip_collaborators_user_id ON trip_collaborators(user_id);
CREATE INDEX idx_trip_collaborators_deleted_at ON trip_collaborators(deleted_at);

-- Create unique constraint to prevent duplicate collaborators
CREATE UNIQUE INDEX idx_trip_collaborators_unique ON trip_collaborators(trip_id, user_id) WHERE deleted_at IS NULL;
