CREATE TABLE `vtubers` (
    `id` INTEGER NOT NULL,
    `name` VARCHAR(191) NOT NULL,
    `belongs_to` VARCHAR(191),
    `profile_image_url` VARCHAR(191),
    `twitter_user_name` VARCHAR(191),
    `channel` VARCHAR(191),
    `created_at` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `updated_at` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),

    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

CREATE TABLE `hashtags` (
    `id` INTEGER NOT NULL,
    `name` VARCHAR(191) NOT NULL,
    `is_self` BOOLEAN NOT NULL DEFAULT true,
    `created_at` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `updated_at` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `vtuber_id` INTEGER NOT NULL,

    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

CREATE TABLE `tweet_objects` (
    `id` VARCHAR(30) NOT NULL UNIQUE,
    `text` VARCHAR(4096),
    `retweet_count` INTEGER NOT NULL,
    `like_count` INTEGER NOT NULL,
    `author_id` VARCHAR(30) NOT NULL,
    `url` VARCHAR(191) NOT NULL,
    `tweeted_at` DATETIME(3) NOT NULL,
    `created_at` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `updated_at` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `hashtag_id` INTEGER NOT NULL,

    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

CREATE TABLE `media_objects` (
    `id` INTEGER NOT NULL AUTO_INCREMENT,
    `url` VARCHAR(191) NOT NULL,
    `type` VARCHAR(191) NOT NULL,
    `created_at` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `updated_at` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `tweet_id` VARCHAR(191) NOT NULL,

    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

CREATE TABLE `authors` (
    `id` VARCHAR(30) NOT NULL UNIQUE,
    `name` VARCHAR(191) NOT NULL,
    `username` VARCHAR(191) NOT NULL,
    `url` VARCHAR(191) NOT NULL,
    `profile_image_url` VARCHAR(191),
    `banner_image_url` VARCHAR(191),
    `biography` VARCHAR(4096),
    `website` VARCHAR(191),
    `followers_count` INTEGER NOT NULL,
    `following_count` INTEGER NOT NULL,
    `tweets_count` INTEGER NOT NULL,
    `joined` DATETIME(3) NOT NULL,
    `created_at` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `updated_at` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),

    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
