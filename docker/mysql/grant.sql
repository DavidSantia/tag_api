-- Add demo user
CREATE USER 'demo'@'%' IDENTIFIED BY 'welcome1';
GRANT SELECT, INSERT, UPDATE, DELETE, LOCK TABLES, EXECUTE ON *.* TO 'demo'@'%';

-- Add Wordpress user
CREATE USER 'wordpress'@'%' IDENTIFIED BY 'welcome1';
GRANT SELECT, INSERT, UPDATE, DELETE, CREATE, ALTER, LOCK TABLES, EXECUTE ON *.* TO 'wordpress'@'%';

-- Add APM user
CREATE USER 'newrelic'@'%' IDENTIFIED BY 'welcome1';
GRANT REPLICATION CLIENT ON *.* TO 'newrelic'@'%' WITH MAX_USER_CONNECTIONS 5;

-- Save changes
FLUSH PRIVILEGES;
