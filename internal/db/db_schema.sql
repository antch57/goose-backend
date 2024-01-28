-- okay so run this file when you want use this projects.
CREATE DATABASE IF NOT EXISTS music_library;
USE music_library;

CREATE TABLE IF NOT EXISTS Bands (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  genre VARCHAR(255) NOT NULL,
  year INT NOT NULL,
  description TEXT,
  UNIQUE (name, genre)
);

CREATE TABLE IF NOT EXISTS Albums (
  id INT AUTO_INCREMENT PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  releaseDate DATE NOT NULL,
  bandId INT,
  FOREIGN KEY (bandId) REFERENCES Bands(id)
);

CREATE TABLE IF NOT EXISTS Songs (
  id INT AUTO_INCREMENT PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  duration TIME NOT NULL,
  albumId INT,
  FOREIGN KEY (albumId) REFERENCES Albums(id)
);