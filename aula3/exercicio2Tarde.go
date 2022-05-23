package main

type Produto interface {
	CalcularCusto() float64
}

type produto struct {
	tipo  string
	nome  string
	preco float64
}

func (p produto) CalcularCusto() float64 {
	switch p.tipo {
	case "pequeno":
		return p.preco
	case "medio":
		return p.preco * 1.03
	case "grande":
		return p.preco*1.06 + 2_500.00
	}

	return 0.0
}

type Ecommerce interface {
	Total() float64
	Adicionar(produto Produto)
}

type loja struct {
	produtos []produto
}

func (l loja) Total() float64 {
	soma := 0.0
	for _, produto := range l.produtos {
		soma += produto.CalcularCusto()
	}

	return soma
}

func (l *loja) Adicionar(p produto) {
	l.produtos = append(l.produtos, p)
}

func novoProduto(tipo, nome string, preco float64) produto {
	novoProd := produto{
		tipo:  tipo,
		nome:  nome,
		preco: preco,
	}

	return novoProd
}

func novaLoja(produtos ...produto) loja {
	novaLo := loja{produtos: produtos}

	return novaLo
}
