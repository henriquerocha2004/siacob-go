CREATE TABLE clients(
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    full_name VARCHAR(255) NOT NULL,
    gender CHAR,
    type CHAR,
    birth_date DATE,
    place_of_birth VARCHAR(255),
    nationality VARCHAR(255),
    marital_status CHAR,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    deleted_at DATETIME, 
    PRIMARY KEY(id)
) ENGINE = innodb;

CREATE TABLE addresses(
    id BIGINT NOT NULL AUTO_INCREMENT,
    street VARCHAR(255),
    district VARCHAR(255),
    city VARCHAR(255),
    zip_code VARCHAR(255),
    state VARCHAR(255),
    type CHAR,
    client_id BIGINT UNSIGNED NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    deleted_at DATETIME,
    PRIMARY KEY(id)
) Engine = innodb;

CREATE TABLE contacts(
    id BIGINT NOT NULL AUTO_INCREMENT,
    phone VARCHAR(255),
    mobile_operator VARCHAR(255),
    site VARCHAR(255),
    type CHAR,
    client_id BIGINT UNSIGNED NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    deleted_at DATETIME,
    PRIMARY KEY(id)
) Engine = innodb;

CREATE TABLE documents(
    id BIGINT NOT NULL AUTO_INCREMENT,
    rg VARCHAR(255),
    cpf_or_cnpj VARCHAR(255),
    titulo_eleitor VARCHAR(255),
    ctps VARCHAR(255),
    pis VARCHAR(255),
    cnh VARCHAR(255),
    passport VARCHAR(255),
    reservista VARCHAR(255),
    ie VARCHAR(255),
    im VARCHAR(255),
    client_id BIGINT UNSIGNED NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    deleted_at DATETIME,
    PRIMARY KEY(id)
) Engine = innodb;

CREATE TABLE filiation (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    mother_name VARCHAR(255),
    father_name VARCHAR(255),
    client_id BIGINT UNSIGNED NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    deleted_at DATETIME,
    PRIMARY KEY(id)
) Engine = innodb;

CREATE TABLE bank_accounts(
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    type CHAR,
    name VARCHAR(255),
    agency VARCHAR(255),
    account VARCHAR(255),
    pix VARCHAR(255),
    client_id BIGINT UNSIGNED NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    deleted_at DATETIME,
    PRIMARY KEY(id)
) Engine = innodb;

ALTER TABLE addresses ADD CONSTRAINT fk_client_addresses FOREIGN KEY (client_id) REFERENCES clients (id);
ALTER TABLE contacts ADD CONSTRAINT fk_client_contacts FOREIGN KEY (client_id) REFERENCES clients (id);
ALTER TABLE documents ADD CONSTRAINT fk_client_documents FOREIGN KEY (client_id) REFERENCES clients (id);
ALTER TABLE filiation ADD CONSTRAINT fk_client_filiation FOREIGN KEY (client_id) REFERENCES clients (id);
ALTER TABLE bank_accounts ADD CONSTRAINT fk_client_bkAccount FOREIGN KEY (client_id) REFERENCES clients (id);