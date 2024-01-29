CREATE DATABASE IF NOT EXISTS music_library;
USE music_library;

CREATE TABLE IF NOT EXISTS Bands (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  genre VARCHAR(255) NOT NULL,
  year INT NOT NULL,
  description TEXT,
  UNIQUE (name, genre) -- Prevent duplicate bands within the same genre
);

CREATE TABLE IF NOT EXISTS Albums (
  id INT AUTO_INCREMENT PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  release_date DATE NOT NULL,
  band_id INT NOT NULL,
  FOREIGN KEY (band_id) REFERENCES Bands(id) ON DELETE CASCADE,
  UNIQUE (title, band_id)  -- Prevent duplicate albums for the same band
);

CREATE TABLE IF NOT EXISTS Songs (
  id INT AUTO_INCREMENT PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  duration INTEGER NOT NULL,
  album_id INT,
  band_id INT NOT NULL,
  FOREIGN KEY (album_id) REFERENCES Albums(id) ON DELETE CASCADE,
  FOREIGN KEY (band_id) REFERENCES Bands(id) ON DELETE CASCADE,
  UNIQUE (title, band_id)  -- Prevent duplicate songs for the same band
);