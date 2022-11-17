DROP database IF EXISTS UTest;
create database UTest;

DROP table IF EXISTS `user`;
DROP table IF EXISTS `course`;
DROP table IF EXISTS `stu_course`;
DROP table IF EXISTS `teacher_course`;
DROP table IF EXISTS `exercise`;
DROP table IF EXISTS `stu_ex`;
DROP table IF EXISTS `answer`;
DROP table IF EXISTS `ex_ans`;
DROP table IF EXISTS `comment`;
DROP table IF EXISTS `notice`;
CREATE TABLE `user` (
 `id` varchar(255) PRIMARY KEY COMMENT 'schoolid',
 `email` varchar(255),
 `username` varchar(255),
 `password` varchar(255),
 `role` int COMMENT '0:管理员 1:老师 2:学生',
 `gender` varchar(255),
 `major` varchar(255)
);

CREATE TABLE `courses` (
   `id` int PRIMARY KEY,
   `name` varchar(255),
   `info` varchar(255),
   `createdAt` timestamp NOT NULL DEFAULT current_timestamp(),
   `updatedAt` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
   `deletedAt` timestamp NULL DEFAULT NULL
);

CREATE TABLE `stu_course` (
    `id` int PRIMARY KEY,
    `schoolid` int,
    `courseid` int,
    `score` int,
    `createdAt` timestamp NOT NULL DEFAULT current_timestamp(),
    `updatedAt` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
    `deletedAt` timestamp NULL DEFAULT NULL
);

CREATE TABLE `teacher_course` (
    `id` int PRIMARY KEY,
    `schoolid` int,
    `courseId` int,
    `createdAt` timestamp NOT NULL DEFAULT current_timestamp(),
    `updatedAt` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
    `deletedAt` timestamp NULL DEFAULT NULL
);

CREATE TABLE `exercises` (
     `id` int PRIMARY KEY,
     `name` varchar(255),
     `subjects` varchar(255),
     `contributorid` int,
     `exType` varchar(255) COMMENT '题目类型',
     `cnt` int COMMENT '总做题次数',
     `accuracy` float8,
     `content` text COMMENT '题目描述',
     `createdAt` timestamp NOT NULL DEFAULT current_timestamp(),
     `updatedAt` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
     `deletedAt` timestamp NULL DEFAULT NULL
);

CREATE TABLE `stu_ex` (
    `schoolid` int,
    `exerciseId` int,
    `pass` bool COMMENT '是否通过',
    `practiceCnt` int COMMENT '做题次数',
    `createdAt` timestamp NOT NULL DEFAULT current_timestamp(),
    `updatedAt` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
    `deletedAt` timestamp NULL DEFAULT NULL
);

CREATE TABLE `answers` (
   `id` int PRIMARY KEY,
   `schoolid` int COMMENT '发布人',
   `content` text COMMENT '答案内容',
   `createdAt` timestamp NOT NULL DEFAULT current_timestamp(),
   `updatedAt` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
   `deletedAt` timestamp NULL DEFAULT NULL
);

CREATE TABLE `ex_ans` (
    `id` int,
    `eid` int,
    `ansid` int,
    `createdAt` timestamp NOT NULL DEFAULT current_timestamp(),
    `updatedAt` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
    `deletedAt` timestamp NULL DEFAULT NULL
);

CREATE TABLE `comments` (
    `eaid` int,
    `schoolid` int,
    `content` text COMMENT '评论内容',
    `createdAt` timestamp NOT NULL DEFAULT current_timestamp(),
    `updatedAt` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
    `deletedAt` timestamp NULL DEFAULT NULL
);

CREATE TABLE `notice` (
    `scid` int,
    `tcid` int,
    `content` text,
    `createdAt` timestamp NOT NULL DEFAULT current_timestamp(),
    `updatedAt` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
    `deletedAt` timestamp NULL DEFAULT NULL
);

# ALTER TABLE `ex_ans` ADD FOREIGN KEY (`eid`) REFERENCES `exercises` (`id`);
#
# ALTER TABLE `ex_ans` ADD FOREIGN KEY (`ansid`) REFERENCES `answers` (`id`);
#
# ALTER TABLE `comments` ADD FOREIGN KEY (`eaid`) REFERENCES `ex_ans` (`id`);
#
# ALTER TABLE `teacher_course` ADD FOREIGN KEY (`schoolid`) REFERENCES `users` (`id`);
#
# ALTER TABLE `stu_ex` ADD FOREIGN KEY (`schoolid`) REFERENCES `users` (`id`);
#
# ALTER TABLE `stu_ex` ADD FOREIGN KEY (`exerciseId`) REFERENCES `exercises` (`id`);
#
# ALTER TABLE `stu_course` ADD FOREIGN KEY (`schoolid`) REFERENCES `users` (`id`);
#
# ALTER TABLE `teacher_course` ADD FOREIGN KEY (`courseId`) REFERENCES `courses` (`id`);
#
# ALTER TABLE `stu_course` ADD FOREIGN KEY (`courseid`) REFERENCES `courses` (`id`);
#
# ALTER TABLE `notice` ADD FOREIGN KEY (`tcid`) REFERENCES `teacher_course` (`id`);
#
# ALTER TABLE `notice` ADD FOREIGN KEY (`scid`) REFERENCES `stu_course` (`id`);
#
