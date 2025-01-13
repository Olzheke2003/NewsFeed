-- Insert genders
INSERT INTO gender (gender) VALUES
('Male'),
('Female'),
('Non-binary'),
('Other');

-- Insert categories
INSERT INTO category (category) VALUES
('Technology'),
('Science'),
('Health'),
('Sports'),
('Entertainment');

-- Insert users
INSERT INTO users (email, username, password, name, lastname, date_of_birth, gender_id, image) VALUES
('john.doe@example.com', 'johndoe', 'password123', 'John', 'Doe', '1990-01-15', 1, 'john_doe.jpg'),
('jane.smith@example.com', 'janesmith', 'securepass', 'Jane', 'Smith', '1995-06-25', 2, 'jane_smith.jpg'),
('alex.taylor@example.com', 'alextaylor', 'alexpass', 'Alex', 'Taylor', '1988-03-12', 3, NULL);

-- Insert news
INSERT INTO news (title, category_id, content, image) VALUES
('Latest Tech Trends in 2025', 1, 'Content about technology trends...', 'tech_trends.jpg'),
('Breakthrough in Cancer Research', 2, 'Details about new cancer research...', 'cancer_research.jpg'),
('Tips for Healthy Living', 3, 'Healthy living tips and tricks...', 'healthy_living.jpg'),
('Championship Game Highlights', 4, 'Highlights of the championship game...', 'game_highlights.jpg'),
('Top Movies of the Year', 5, 'List of the best movies in 2025...', 'top_movies.jpg');

-- Insert comments
INSERT INTO comments (news_id, parent_id, user_id, content) VALUES
(1, NULL, 1, 'This is amazing news for tech enthusiasts!'),
(1, 1, 2, 'I agree! This is very exciting.'),
(2, NULL, 3, 'This is a significant step in cancer research.'),
(3, NULL, 1, 'I found these tips very helpful, thank you!'),
(3, 3, 2, 'Me too! I have started following some of them.');
