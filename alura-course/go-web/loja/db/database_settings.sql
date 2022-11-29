CREATE TABLE produtos (
	id serial primary key,
	nome varchar,
	descricao varchar,
	preco decimal,
	quantidade integer
)

INSERT INTO produtos (nome, descricao, preco, quantidade) VALUES ('Camiseta', 'Azul, bem bonita', 39, 15);
INSERT INTO produtos (nome, descricao, preco, quantidade) VALUES ('Tenis', 'Confortável', 339, 35);
INSERT INTO produtos (nome, descricao, preco, quantidade) VALUES ('Fone', 'Muito Bom', 100.9, 1);
INSERT INTO produtos (nome, descricao, preco, quantidade) VALUES ('Produto novo	', 'Bem legal', 1.99, 25);
INSERT INTO produtos (nome, descricao, preco, quantidade) VALUES ('Notebook', 'Massa', 5998.58, 1);

SELECT * FROM PRODUTOS