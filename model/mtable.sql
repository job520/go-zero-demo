CREATE TABLE `mtable` (
      `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
      `name` varchar(255) CHARACTER SET latin1 DEFAULT '' COMMENT '名称',
      PRIMARY KEY (`id`),
      KEY `name` (`name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;