--
-- Add demo user
--

CREATE USER 'demo'@'localhost' IDENTIFIED BY 'welcome1';
GRANT SELECT, INSERT, UPDATE, DELETE, LOCK TABLES, EXECUTE ON *.* TO 'demo'@'localhost';
CREATE USER 'demo'@'%' IDENTIFIED BY 'welcome1';
GRANT SELECT, INSERT, UPDATE, DELETE, LOCK TABLES, EXECUTE ON *.* TO 'demo'@'%';
