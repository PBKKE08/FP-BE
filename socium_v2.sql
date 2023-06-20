CREATE TABLE cities (
    id CHAR(36) NOT NULL,
    name VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE users (
  id CHAR(36) NOT NULL,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  telephone VARCHAR(20) NOT NULL,
  gender ENUM('f', 'm') NOT NULL,
  city_id CHAR(36) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  FOREIGN KEY (city_id) REFERENCES cities(id)
);

CREATE TABLE partners (
  id CHAR(36) NOT NULL,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  telephone VARCHAR(20) NOT NULL,
  gender ENUM('f', 'm') NOT NULL,
  price VARCHAR(100) NOT NULL,
  city_id CHAR(36) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  CONSTRAINT fk_city FOREIGN KEY (city_id) REFERENCES cities(id)
);

CREATE TABLE reviews (
  id CHAR(36) NOT NULL,
  user_id CHAR(36) NOT NULL,
  partner_id CHAR(36) NOT NULL,
  rating INT NOT NULL,
  comment VARCHAR(255),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (partner_id) REFERENCES partners(id)
);

INSERT INTO cities (id, name) VALUES ('4c2d8c23-573d-4f0e-9c1a-3d05f488f4e0', 'Bandung');
INSERT INTO cities (id, name) VALUES ('03b8552d-7905-4ce6-93c0-7d49538e5e0b', 'Jakarta');
INSERT INTO cities (id, name) VALUES ('f9e2a5cf-198a-4c88-9926-79ae5f44f7d2', 'Surabaya');
INSERT INTO cities (id, name) VALUES ('28b7b4c2-9f68-47e0-bd79-8a363fd98f13', 'Jogjakarta');
INSERT INTO cities (id, name) VALUES ('58dd3448-0bb2-42f2-8542-0e15eb2a59fe', 'Malang');


INSERT INTO users (id, name, email, telephone, gender, city_id)
VALUES
    ('4f150d84-2b5d-4cb6-82f1-73491df2407a', 'User 1', 'user1@example.com', '1234567890', 'm', '4c2d8c23-573d-4f0e-9c1a-3d05f488f4e0'),
    ('f1564f80-b3e1-4f8f-b2da-d739f80c7c62', 'User 2', 'user2@example.com', '0987654321', 'f', '03b8552d-7905-4ce6-93c0-7d49538e5e0b'),
    ('f218a2a6-b2f3-421a-88df-79a9e39adbed', 'User 3', 'user3@example.com', '5555555555', 'm', 'f9e2a5cf-198a-4c88-9926-79ae5f44f7d2'),
    ('2d8f9c77-ae79-439e-8a07-00d16db37e1a', 'User 4', 'user4@example.com', '9876543210', 'f', '28b7b4c2-9f68-47e0-bd79-8a363fd98f13'),
    ('c8ab0fd3-3d4a-4a39-b7e3-df84b8974a36', 'User 5', 'user5@example.com', '1112223333', 'm', '58dd3448-0bb2-42f2-8542-0e15eb2a59fe'),
    ('7d31483c-6990-4b15-bb3f-d81e72900ed0', 'User 6', 'user6@example.com', '4445556666', 'f', '4c2d8c23-573d-4f0e-9c1a-3d05f488f4e0'),
    ('1b1a6fb7-33e3-43d3-9678-5aebb65086be', 'User 7', 'user7@example.com', '7778889999', 'm', '03b8552d-7905-4ce6-93c0-7d49538e5e0b'),
    ('a1d9764a-0b71-4507-9f61-d3a6c11f7e79', 'User 8', 'user8@example.com', '2223334444', 'f', 'f9e2a5cf-198a-4c88-9926-79ae5f44f7d2'),
    ('b5d5f7e7-46a7-4b3b-bb76-78289e9e372a', 'User 9', 'user9@example.com', '6667778888', 'm', '28b7b4c2-9f68-47e0-bd79-8a363fd98f13'),
    ('d2d0378e-660d-4e33-9c5c-5dcdaf3e92e4', 'User 10', 'user10@example.com', '3334445555', 'f', '58dd3448-0bb2-42f2-8542-0e15eb2a59fe');

INSERT INTO partners (id, name, email, telephone, gender, price, city_id)
VALUES
    ('68a9d8a7-646a-4a95-8b10-ee1fd3d3ee5e', 'Partner 1', 'partner1@example.com', '1234567890', 'm', '100', '4c2d8c23-573d-4f0e-9c1a-3d05f488f4e0'),
    ('ef1371f9-c747-42f0-8e82-7a9d59ac135f', 'Partner 2', 'partner2@example.com', '0987654321', 'f', '200', '03b8552d-7905-4ce6-93c0-7d49538e5e0b'),
    ('bae2c874-b144-45e0-b2ae-5300b06c64e3', 'Partner 3', 'partner3@example.com', '5555555555', 'm', '150', 'f9e2a5cf-198a-4c88-9926-79ae5f44f7d2'),
    ('9a96b015-06ff-4e0f-8c25-cb775e36fe07', 'Partner 4', 'partner4@example.com', '9876543210', 'f', '300', '28b7b4c2-9f68-47e0-bd79-8a363fd98f13'),
    ('d1da7b0c-791e-4f7a-8a55-6e43ad0a7d4a', 'Partner 5', 'partner5@example.com', '1112223333', 'm', '250', '58dd3448-0bb2-42f2-8542-0e15eb2a59fe'),
    ('1e4ed79a-7c9e-4c7b-9b9f-9d86cc6c7c2d', 'Partner 6', 'partner6@example.com', '4445556666', 'f', '400', '4c2d8c23-573d-4f0e-9c1a-3d05f488f4e0'),
    ('dcaad68f-2c32-4d9c-81e6-1877f3ac8314', 'Partner 7', 'partner7@example.com', '7778889999', 'm', '350', '03b8552d-7905-4ce6-93c0-7d49538e5e0b'),
    ('42e68281-d8a0-4a39-968b-11e3ac3f99cd', 'Partner 8', 'partner8@example.com', '2223334444', 'f', '500', 'f9e2a5cf-198a-4c88-9926-79ae5f44f7d2'),
    ('6dd485be-cc1c-4d88-91e0-f960e7ef83f5', 'Partner 9', 'partner9@example.com', '6667778888', 'm', '450', '28b7b4c2-9f68-47e0-bd79-8a363fd98f13'),
    ('cd14509d-0325-4a8d-9e58-ff59db2cfd4c', 'Partner 10', 'partner10@example.com', '3334445555', 'f', '600', '58dd3448-0bb2-42f2-8542-0e15eb2a59fe');


INSERT INTO reviews (id, user_id, partner_id, rating, comment)
VALUES
    ('428881ad-8579-4380-8e85-690917976b9f', '4f150d84-2b5d-4cb6-82f1-73491df2407a', '68a9d8a7-646a-4a95-8b10-ee1fd3d3ee5e', 4, 'Good service'),
    ('0f441798-572b-4d37-95db-fc502e592f82', 'f1564f80-b3e1-4f8f-b2da-d739f80c7c62', 'ef1371f9-c747-42f0-8e82-7a9d59ac135f', 5, 'Excellent experience'),
    ('4ddcc8da-2f8a-434e-b358-f513a43c7c2c', 'f218a2a6-b2f3-421a-88df-79a9e39adbed', 'bae2c874-b144-45e0-b2ae-5300b06c64e3', 3, 'Average service'),
    ('aa4e5da1-00b1-4b03-b448-3ee47c18fbc2', '2d8f9c77-ae79-439e-8a07-00d16db37e1a', '9a96b015-06ff-4e0f-8c25-cb775e36fe07', 5, 'Highly recommended'),
    ('6a5117a9-d872-4b1c-93f2-684d0a09c52b', 'c8ab0fd3-3d4a-4a39-b7e3-df84b8974a36', 'd1da7b0c-791e-4f7a-8a55-6e43ad0a7d4a', 4, 'Good service'),
    ('7a44e847-59e1-4ebd-819a-2726e57de9ab', '4f150d84-2b5d-4cb6-82f1-73491df2407a', '1e4ed79a-7c9e-4c7b-9b9f-9d86cc6c7c2d', 5, 'Excellent experience'),
    ('6d2d6692-0e54-4477-9614-13ee26b836ed', 'f1564f80-b3e1-4f8f-b2da-d739f80c7c62', 'dcaad68f-2c32-4d9c-81e6-1877f3ac8314', 3, 'Average service'),
    ('0b96da1d-12a2-4c0b-a3a2-2b04f9f10824', '2d8f9c77-ae79-439e-8a07-00d16db37e1a', '42e68281-d8a0-4a39-968b-11e3ac3f99cd', 5, 'Highly recommended'),
    ('cbb9c4b9-6a37-4bcf-9f9f-9486e78e40fb', 'c8ab0fd3-3d4a-4a39-b7e3-df84b8974a36', '6dd485be-cc1c-4d88-91e0-f960e7ef83f5', 4, 'Good service'),
    ('d847be63-425f-447e-b53e-2039ab8731c0', '4f150d84-2b5d-4cb6-82f1-73491df2407a', 'cd14509d-0325-4a8d-9e58-ff59db2cfd4c', 5, 'Excellent experience'),
    ('ce0ac74e-6a67-4d4a-8b45-ecf5e7229967', 'f218a2a6-b2f3-421a-88df-79a9e39adbed', '68a9d8a7-646a-4a95-8b10-ee1fd3d3ee5e', 3, 'Average service'),
    ('0811a309-2c5d-46f7-8fc4-6d6944b58b8f', '2d8f9c77-ae79-439e-8a07-00d16db37e1a', 'ef1371f9-c747-42f0-8e82-7a9d59ac135f', 5, 'Highly recommended'),
    ('3703a5a9-3db3-456d-a2ab-ef5226b3a88c', 'c8ab0fd3-3d4a-4a39-b7e3-df84b8974a36', 'bae2c874-b144-45e0-b2ae-5300b06c64e3', 4, 'Good service'),
    ('d378b797-0dd0-4395-92d9-96f8fe85e2b2', '4f150d84-2b5d-4cb6-82f1-73491df2407a', '9a96b015-06ff-4e0f-8c25-cb775e36fe07', 5, 'Excellent experience'),
    ('9a42b6fb-4d8f-4e51-a47c-3450f05bfc17', 'f1564f80-b3e1-4f8f-b2da-d739f80c7c62', 'd1da7b0c-791e-4f7a-8a55-6e43ad0a7d4a', 3, 'Average service'),
    ('a0555d94-96b6-4f99-b5cc-37dd6ab8c5d4', '2d8f9c77-ae79-439e-8a07-00d16db37e1a', '1e4ed79a-7c9e-4c7b-9b9f-9d86cc6c7c2d', 5, 'Highly recommended'),
    ('db67d297-0385-4c68-96c9-9d7c8c4d5151', 'c8ab0fd3-3d4a-4a39-b7e3-df84b8974a36', '58dd3448-0bb2-42f2-8542-0e15eb2a59fe', 4, 'Good service'),
    ('6b6697a5-038e-4f18-9e94-bcc1f73d1ee8', '4f150d84-2b5d-4cb6-82f1-73491df2407a', 'f9e2a5cf-198a-4c88-9926-79ae5f44f7d2', 5, 'Excellent experience'),
    ('e07c2f8a-c660-4f3f-8b30-7f93b726eeb1', 'f1564f80-b3e1-4f8f-b2da-d739f80c7c62', '28b7b4c2-9f68-47e0-bd79-8a363fd98f13', 3, 'Average service'),
    ('51bebe59-6a21-4148-aa04-95a7e24f0e79', '2d8f9c77-ae79-439e-8a07-00d16db37e1a', '42e68281-d8a0-4a39-968b-11e3ac3f99cd', 5, 'Highly recommended');





