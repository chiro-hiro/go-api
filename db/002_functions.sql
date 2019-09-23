USE `GOAPI`;

DROP FUNCTION IF EXISTS `createNewUser`;
-- Dumping structure for function GOAPI.createNewUser
DELIMITER //
CREATE FUNCTION `createNewUser`(uname TEXT, uemail TEXT , upassword VARBINARY(32)) RETURNS BOOLEAN
BEGIN
   DECLARE usalt VARBINARY(16);
   IF (SELECT COUNT(`id`) FROM `GOAPI_users` WHERE `username`=uname OR `email`=uemail) > 0 THEN
   	RETURN FALSE;
   END IF;
   SET usalt = UNHEX(MD5(RAND()));
	INSERT INTO `GOAPI_users`
	(
		`id`,
		`username`,
		`email`,
		`password`,
		`salt`
	)
	VALUES
	(
		getNewUserId(),
		uname,
		uemail,
		UNHEX(SHA2(CONCAT(upassword, usalt), 256)),
		usalt
	);
	RETURN TRUE;
END//
DELIMITER ;

DROP FUNCTION IF EXISTS `getNewUserId`;
-- Dumping structure for function GOAPI.getNewUserId
DELIMITER //
CREATE FUNCTION `getNewUserId`() RETURNS VARBINARY(16)
BEGIN
	DECLARE new_id VARBINARY(16);
	findId: LOOP
		SET new_id = UNHEX(MD5(CONCAT(RAND(), NOW())));
		IF (SELECT COUNT(`id`) FROM `GOAPI_users` WHERE `id`=new_id) = 0 THEN
			LEAVE findId;
		END IF;
	END LOOP findId;
	RETURN new_id;
END//
DELIMITER ;

DROP FUNCTION IF EXISTS `updateProfileField`;
-- Dumping structure for function GOAPI.updateProfileField
DELIMITER //
CREATE FUNCTION `updateProfileField`(user_id VARBINARY(16), profile_field VARCHAR(64), profile_value VARCHAR(64)) RETURNS BOOL
BEGIN
	IF (SELECT COUNT(`id`) FROM `GOAPI_profiles` WHERE `id`=user_id AND `field`=profile_field) = 0 THEN
		INSERT INTO `GOAPI_profiles` (`id`, `field`, `value`) VALUES (user_id, profile_field, profile_value);
	ELSE
		UPDATE `GOAPI_profiles` SET `value`=profile_value WHERE `id`=user_id AND `field`=profile_field;
	END IF;
	RETURN TRUE;
END//
DELIMITER ;

DROP FUNCTION IF EXISTS `updateProfileFieldBySession`;
-- Dumping structure for function GOAPI.updateProfileFieldBySession
DELIMITER //
CREATE FUNCTION `updateProfileFieldBySession`(session_id VARBINARY(32), profile_field VARCHAR(64), profile_value VARCHAR(64)) RETURNS BOOL
BEGIN
    DECLARE user_id VARBINARY(32);
    SET user_id = (SELECT `uid` FROM `GOAPI_sessions` WHERE `id`=session_id);
	IF (SELECT COUNT(`id`) FROM `GOAPI_profiles` WHERE `id`=user_id AND `field`=profile_field) = 0 THEN
		INSERT INTO `GOAPI_profiles` (`id`, `field`, `value`) VALUES (user_id, profile_field, profile_value);
	ELSE
		UPDATE `GOAPI_profiles` SET `value`=profile_value WHERE `id`=user_id AND `field`=profile_field;
	END IF;
	RETURN TRUE;
END//
DELIMITER ;

DROP FUNCTION IF EXISTS `getNewSession`;
-- Dumping structure for function GOAPI.getSessionId
DELIMITER //
CREATE FUNCTION `getNewSession`() RETURNS VARBINARY(32)
BEGIN
	DECLARE new_id VARBINARY(32);
	findId: LOOP
		SET new_id = UNHEX(SHA2(CONCAT(RAND(), NOW()), 256));
		IF (SELECT COUNT(`id`) FROM `GOAPI_sessions` WHERE `id`=new_id) = 0 THEN
			LEAVE findId;
		END IF;
	END LOOP findId;
	RETURN new_id;
END//
DELIMITER ;

DROP FUNCTION IF EXISTS `loginUser`;
-- Dumping structure for function GOAPI.loginUser
DELIMITER //
CREATE FUNCTION `loginUser`(uname TEXT, upassword VARBINARY(32)) RETURNS VARBINARY(32)
BEGIN
	DECLARE new_id VARBINARY(32);
	DECLARE user_id VARBINARY(16);
	SET user_id = (SELECT `id` FROM `GOAPI_users` WHERE `username`=uname AND `password`=UNHEX(SHA2(CONCAT(upassword, salt), 256)));
	IF (LENGTH(user_id) > 0) THEN
	   SET new_id = getNewSession();
	   INSERT INTO `GOAPI_sessions` (`id`, `uid`, `last_active`) VALUES (new_id, user_id, NOW());
		RETURN new_id;
	ELSE
		RETURN new_id;
	END IF;
END//
DELIMITER ;

DROP FUNCTION IF EXISTS `logoutUser`;
-- Dumping structure for function GOAPI.logoutUser
DELIMITER //
CREATE FUNCTION `logoutUser` (session_id VARBINARY(32)) RETURNS bool
BEGIN
    DELETE FROM `GOAPI_sessions` WHERE `id`=session_id;
	RETURN TRUE;
END//
DELIMITER ;

DROP PROCEDURE IF EXISTS `updateLastActive`;
-- Dumping structure for procedure GOAPI.loginUser
DELIMITER //
CREATE PROCEDURE `updateLastActive`(sessionId VARBINARY(32))
BEGIN
	UPDATE `GOAPI_sessions` SET `last_active`=NOW() WHERE `id`=sessionId;
END//
DELIMITER ;

DROP FUNCTION IF EXISTS `updatePassword`;
-- Dumping structure for function GOAPI.updatePassword
DELIMITER //
CREATE FUNCTION `updatePassword`(session_id VARBINARY(32), upassword VARBINARY(32), npassword VARBINARY(32)) RETURNS BOOL
BEGIN
   DECLARE usalt VARBINARY(16);
   DECLARE user_id VARBINARY(16);
    SET user_id = (SELECT `id` FROM `GOAPI_users` WHERE `id`=(SELECT `uid` FROM `GOAPI_sessions` WHERE `id`=session_id) AND  `password`=UNHEX(SHA2(CONCAT(upassword, `salt`), 256)));
   IF LENGTH(user_id) <= 0 THEN
    RETURN FALSE;
   END IF;
   SET usalt = UNHEX(MD5(RAND()));
    UPDATE `GOAPI_users` SET `password` = UNHEX(SHA2(CONCAT(npassword, usalt), 256)), `salt`=usalt WHERE `id`=user_id;
    DELETE FROM `GOAPI_sessions` WHERE `uid`=user_id;
	RETURN TRUE;
END//
DELIMITER ;

DROP PROCEDURE IF EXISTS `getUser`;
-- Dumping structure for procedure GOAPI.getUser
DELIMITER //
CREATE PROCEDURE `getUser`(user_id VARBINARY(32))
BEGIN
	SELECT * FROM `GOAPI_users` WHERE `id`=user_id;
END//
DELIMITER ;

DROP PROCEDURE IF EXISTS `getUserBySession`;
-- Dumping structure for procedure GOAPI.getUserBySession
DELIMITER //
CREATE PROCEDURE `getUserBySession`(session_id VARBINARY(32))
BEGIN
	SELECT `id`, `username`, `email` FROM `GOAPI_users` WHERE `id`=(SELECT `uid` FROM `GOAPI_sessions` WHERE `id`=session_id);
END//
DELIMITER ;

DROP PROCEDURE IF EXISTS `getProfile`;
-- Dumping structure for procedure GOAPI.getProfile
DELIMITER //
CREATE PROCEDURE `getProfile`(user_id VARBINARY(32))
BEGIN
	SELECT * FROM `GOAPI_profiles` WHERE `id`=user_id;
END//
DELIMITER ;

DROP PROCEDURE IF EXISTS `getProfileBySession`;
-- Dumping structure for procedure GOAPI.getProfileBySession
DELIMITER //
CREATE PROCEDURE `getProfileBySession`(session_id VARBINARY(32))
BEGIN
	SELECT * FROM `GOAPI_profiles` WHERE `id`=(SELECT `uid` FROM `GOAPI_sessions` WHERE `id`=session_id);
END//
DELIMITER ;

DROP PROCEDURE IF EXISTS `getProfileField`;
-- Dumping structure for procedure GOAPI.getProfileField
DELIMITER //
CREATE PROCEDURE `getProfileField`(user_id VARBINARY(32), profile_field VARCHAR(64))
BEGIN
	SELECT * FROM `GOAPI_profiles` WHERE `id`=user_id AND `field`=profile_field;
END//
DELIMITER ;