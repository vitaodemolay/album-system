-- Criação da base de dados
CREATE DATABASE SLQ_ALBUMSYSTEM_DB;
GO

-- Seleciona a base de dados
USE SLQ_ALBUMSYSTEM_DB;
GO

-- Criação da tabela TbAlbum
CREATE TABLE TbAlbum (
    ID INT IDENTITY(1,1) PRIMARY KEY,
    Title VARCHAR(50) NOT NULL,
    Artist VARCHAR(50) NOT NULL,
    Price DECIMAL(18, 2) NOT NULL,
    CreationDate DATETIME NOT NULL DEFAULT GETDATE()
);
GO
