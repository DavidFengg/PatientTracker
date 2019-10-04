USE rest_api

CREATE TABLE IF NOT EXISTS `patient` (
    id VARCHAR(10), 
    first_name VARCHAR(20), 
    last_name VARCHAR(20), 
    diagnosis VARCHAR(30), 
    physician VARCHAR(20), 
    dov DATE
)
