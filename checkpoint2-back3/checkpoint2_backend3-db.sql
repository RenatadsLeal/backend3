CREATE SCHEMA `checkpoint2`;

CREATE TABLE `checkpoint2`.`dentist` (
	`id` INT NOT NULL AUTO_INCREMENT,
    `surname` VARCHAR(100) NOT NULL,
    `name` VARCHAR(100)  NOT NULL,
    `registration` VARCHAR(100)  NOT NULL,
    PRIMARY KEY (`id`)
);
    
CREATE TABLE `checkpoint2`.`patient` (
	`id` INT NOT NULL AUTO_INCREMENT,
    `surname` VARCHAR(100) NOT NULL,
    `name` VARCHAR(100)  NOT NULL,
    `rg` VARCHAR(100)  NOT NULL,
    `registration_date` VARCHAR(100)  NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE `checkpoint2`.`appointment` (
	`id` INT NOT NULL AUTO_INCREMENT,
    `patient_id` INT NOT NULL,
    `dentist_id` INT  NOT NULL,
    `date` VARCHAR(100)  NOT NULL,
    `description` VARCHAR(100)  NOT NULL,
    PRIMARY KEY (`id`),
		FOREIGN KEY (`patient_id`)
        REFERENCES `checkpoint2`.`patient` (`id`),
        FOREIGN KEY (`dentist_id`)
        REFERENCES `checkpoint2`.`dentist` (`id`)
);

INSERT INTO `checkpoint2`.`dentist` (`surname`, `name`, `registration`)
VALUES ('Renata', 'da Silva Leal', '001');

INSERT INTO `checkpoint2`.`patient` (`surname`, `name`, `rg`, `registration_date`)
VALUES ('Carolina', 'Haka', '36070666', '15/12/2022');

INSERT INTO `checkpoint2`.`appointment` (`patient_id`, `dentist_id`, `date`, `description`)
VALUES (1, 1, '13/12/2022', 'Limpeza bucal');