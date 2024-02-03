CREATE DATABASE IF NOT EXISTS music_library;
USE music_library;

CREATE TABLE IF NOT EXISTS bands (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  genre VARCHAR(255) NOT NULL,
  year INT NOT NULL,
  description TEXT,
  UNIQUE (name, genre) -- Prevent duplicate bands within the same genre
);

CREATE TABLE IF NOT EXISTS albums (
  id INT AUTO_INCREMENT PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  band_id INT NOT NULL,
  release_date DATE NOT NULL,
  FOREIGN KEY (band_id) REFERENCES bands(id) ON DELETE CASCADE,
  UNIQUE (title, band_id)  -- Prevent duplicate albums for the same band
);

CREATE TABLE IF NOT EXISTS venues (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  location VARCHAR(255),
  UNIQUE (name, location)  -- Prevent duplicate venues with the same name and location
);

CREATE TABLE IF NOT EXISTS performances (
  id INT AUTO_INCREMENT PRIMARY KEY,
  band_id INT NOT NULL,
  venue_id INT NOT NULL,
  performance_date DATE NOT NULL,
  duration INT NOT NULL,  -- Duration of the performance in seconds
  FOREIGN KEY (band_id) REFERENCES bands(id) ON DELETE CASCADE,
  FOREIGN KEY (venue_id) REFERENCES venues(id),
  UNIQUE (band_id, venue_id, performance_date)  -- Prevent duplicate performances for the same band at the same venue on the same date
);

CREATE TABLE IF NOT EXISTS songs (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    band_id INT,
    FOREIGN KEY (band_id) REFERENCES bands(id) ON DELETE CASCADE
  );

CREATE TABLE IF NOT EXISTS performance_songs (
  id INT AUTO_INCREMENT PRIMARY KEY,
  song_id INT NOT NULL,
  duration INT NOT NULL,  -- Duration of the song in seconds
  performance_id INT NOT NULL,
  is_cover BOOLEAN DEFAULT FALSE,  -- Indicates whether the song is a cover
  notes TEXT, -- Any notes about the song ie: cover, FTP, etc
  FOREIGN KEY (performance_id) REFERENCES performances(id) ON DELETE CASCADE,
  FOREIGN KEY (song_id) REFERENCES songs(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS album_songs (
  id INT AUTO_INCREMENT PRIMARY KEY,
  song_id INT NOT NULL,
  album_id INT NOT NULL,
  band_id INT NOT NULL,
  duration INT NOT NULL,
  track_number INT NOT NULL,
  is_cover BOOLEAN DEFAULT FALSE,  -- Indicates whether the song is a cover
  FOREIGN KEY (album_id) REFERENCES albums(id) ON DELETE CASCADE,
  FOREIGN KEY (band_id) REFERENCES bands(id) ON DELETE CASCADE,
  FOREIGN KEY (song_id) REFERENCES songs(id) ON DELETE CASCADE,
  UNIQUE (song_id, album_id)  -- Prevent duplicate songs for the same album
);
