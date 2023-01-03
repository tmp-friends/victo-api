CREATE TABLE IF NOT EXISTS vtuber (
  id INT(11) AUTO_INCREMENT NOT NULL COMMENT "id",
  hashtagId INT(11) NOT NULL COMMENT "hashtag id",
  familyName VARCHAR(191) COMMENT "family name",
  givenName VARCHAR(191) NOT NULL COMMENT "given name",
  twitterUid VARCHAR(191) NOT NULL COMMENT "twitter user id",
  profileImage VARCHAR(191) NOT NULL COMMENT "profile image",
  PRIMARY KEY (id),
  FOREIGN KEY(hashtagId) REFERENCES hashtag(id)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
