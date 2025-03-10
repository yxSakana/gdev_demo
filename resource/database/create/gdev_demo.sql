DROP TABLE IF EXISTS `novel_tag_rel`;
DROP TABLE IF EXISTS `novel_tags`;
DROP TABLE IF EXISTS `novel_chapters`;

DROP TABLE IF EXISTS `image_tag_rel`;
DROP TABLE IF EXISTS `image_tags`;
DROP TABLE IF EXISTS `images`;
DROP TABLE IF EXISTS `image_collections`;

DROP TABLE IF EXISTS `novels`;
DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
  `id` BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '',
  `username` VARCHAR(50) NOT NULL UNIQUE COMMENT '',
  `nickname` VARCHAR(50) NOT NULL COMMENT '',
  `password` VARCHAR(255) NOT NULL COMMENT '',
  `email` VARCHAR(100) UNIQUE COMMENT '',
  `picture_url` VARCHAR(500) DEFAULT NULL COMMENT '',
  `role` TINYINT DEFAULT 1 COMMENT '角色: 0-管理, 1-普通, 2-vip',
--   'vip_expire'
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB CHARACTER SET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `novels` (
  `id` BIGINT PRIMARY KEY AUTO_INCREMENT,
  `user_id` BIGINT NOT NULL COMMENT '',
  `uploader` VARCHAR(50) NOT NULL,
  `title` VARCHAR(255) NOT NULL,
  `description` TEXT,
  `cover_url` VARCHAR(500),
  `status` TINYINT NOT NULL DEFAULT 0 COMMENT '0-连载中, 1-完结',
  `word_count` INT DEFAULT 0 COMMENT '总字数',
  `view` INT DEFAULT 0 COMMENT '阅读量',
  `like` INT DEFAULT 0,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=InnoDB CHARACTER SET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `novel_chapters` (
  `id` BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '',
  `novel_id` BIGINT NOT NULL COMMENT '',
  `title` VARCHAR(255) NOT NULL COMMENT '',
  `content` TEXT NOT NULL COMMENT '',
  `number` INT NOT NULL COMMENT '章节序号',
  `word_count` INT NOT NULL DEFAULT 0 COMMENT '总字数',
  `view` INT NOT NULL DEFAULT 0 COMMENT '阅读量',
  `like` INT NOT NULL DEFAULT 0,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (novel_id) REFERENCES novels(id) ON DELETE CASCADE
) ENGINE=InnoDB CHARACTER SET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT '';

CREATE TABLE `novel_tags` (
  `id` BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '',
  `name` VARCHAR(50) NOT NULL UNIQUE
) ENGINE=InnoDB CHARACTER SET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT '';

CREATE TABLE `novel_tag_rel` (
  `novel_id` BIGINT NOT NULL,
  `novel_tag_id` BIGINT NOT NULL,
  PRIMARY KEY (novel_id, novel_tag_id),
  FOREIGN KEY (novel_id) REFERENCES novels(id) ON DELETE CASCADE,
  FOREIGN KEY (novel_tag_id) REFERENCES novel_tags(id) ON DELETE CASCADE
) ENGINE=InnoDB CHARACTER SET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT 'novel与标签的映射表';

# SELECT name
# FROM novel_tags
# LEFT JOIN novel_tag_rel ON novel_tags.id = novel_tag_rel.novel_tag_id
# LEFT JOIN novels ON novel_tag_rel.novel_id = novels.id
# WHERE novel_tag_rel.novel_id = 1;  ????

# SELECT *
# FROM novels
# LEFT JOIN novel_tag_rel ON novels.id = novel_tag_rel.novel_id
# LEFT JOIN novel_tags ON novel_tag_rel.novel_tag_id = novel_tags.id
# WHERE novel_tags.name = 'tag_name';

DROP PROCEDURE IF EXISTS update_total_word_count;

DELIMITER $$
-- 触发器更新小说字数
CREATE PROCEDURE update_total_word_count(IN in_novel_id BIGINT)
BEGIN
    declare total_word_count INT;

    SELECT SUM(word_count) INTO total_word_count
    FROM novel_chapters
    WHERE novel_chapters.novel_id = in_novel_id;

    UPDATE novels
    SET word_count = total_word_count
    WHERE id = in_novel_id;
END $$
-- insert
CREATE TRIGGER update_novel_word_count_after_insert
AFTER INSERT ON novel_chapters
FOR EACH ROW
BEGIN
    call update_total_word_count(NEW.novel_id);
END $$
-- update
CREATE TRIGGER update_novel_word_count_after_update
AFTER UPDATE ON novel_chapters
FOR EACH ROW
BEGIN
    call update_total_word_count(NEW.novel_id);
END $$
-- delete
CREATE TRIGGER update_novel_word_count_after_delete
AFTER DELETE ON novel_chapters
FOR EACH ROW
BEGIN
    call update_total_word_count(OLD.novel_id);
END $$
DELIMITER ;
-- --------
-- Image
-- --------
CREATE TABLE `image_collections` (
  `id` BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '',
  `user_id` BIGINT NOT NULL,
  `uploader` VARCHAR(50) NOT NULL,
  `title` VARCHAR(255) NOT NULL,
  `description` TEXT,
  `cover_url` VARCHAR(255),
  `number` INT DEFAULT 0 NOT NULL COMMENT '该图包有x张',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB CHARACTER SET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT '';

CREATE TABLE `images` (
  `id` BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '',
  `collection_id` BIGINT NOT NULL,
  `image_url` VARCHAR(255) NOT NULL COMMENT '图片起始url, 结合图包中的图片数使用',
  FOREIGN KEY (collection_id) REFERENCES image_collections(id) ON DELETE CASCADE
) ENGINE=InnoDB CHARACTER SET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT '';

CREATE TABLE `image_tags` (
  `id` BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '',
  `name` VARCHAR(50) NOT NULL UNIQUE
) ENGINE=InnoDB CHARACTER SET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT '';

CREATE TABLE `image_tag_rel` (
  `collection_id` BIGINT NOT NULL,
  `image_tag_id` BIGINT NOT NULL,
  PRIMARY KEY (collection_id, image_tag_id),
  FOREIGN KEY (collection_id) REFERENCES image_collections(id) ON DELETE CASCADE,
  FOREIGN KEY (image_tag_id) REFERENCES image_tags(id) ON DELETE CASCADE
) ENGINE=InnoDB CHARACTER SET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT 'image collections与标签的映射表';


# UPDATE image_collections
# SET number = (
#     SELECT COUNT(*) FROM images WHERE collection_id = ?)
# WHERE image_collections.id = ?;

# SELECT name
# FROM image_tags
# LEFT JOIN image_tag_rel ON image_tags.id = image_tag_rel.image_tag_id
# LEFT JOIN images ON image_tag_rel.collection_id = images.id
# WHERE images.id = ?;
