CREATE TABLE IF NOT EXISTS tweet (
  id INT(11) NOT NULL AUTO_INCREMENT COMMENT "id",
  hashtagId INT(11) NOT NULL COMMENT "hashtag id",
  tweetId VARCHAR(191) NOT NULL COMMENT "tweet id",
  text VARCHAR(191) COMMENT "text",
  retweetCount VARCHAR(10) NOT NULL COMMENT "retweeet count",
  likeCount INT(11) NOT NULL COMMENT "like count",
  authorId VARCHAR(191) NOT NULL COMMENT "author id",
  tweetUrl VARCHAR(191) NOT NULL COMMENT "tweet url",
  tweetedAt DATETIME NOT NULL COMMENT "tweeted date",
  createdAt DATETIME NOT NULL COMMENT "created date",
  updatedAt DATETIME NOT NULL COMMENT "updated date",
  PRIMARY KEY (id),
  FOREIGN KEY(hashtagId) REFERENCES hashtag(id)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
