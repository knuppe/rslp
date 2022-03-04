package rslp

import (
	"fmt"
	"testing"
)

func TestPlural(t *testing.T) {
	step := steps["Plural"]

	tests := []struct {
		input string
		want  string
	}{
		{"vilã", "vilã"},
		{"bons", "bom"},
		{"bal\u00f5es", "bal\u00e3o"},
		{"capit\u00e3es", "capit\u00e3o"},
		{"normais", "normal"},
		{"am\u00e1veis", "am\u00e1vel"},
		{"len\u00e7\u00f3is", "len\u00e7ol"},
		{"barris", "barril"},
		{"males", "male"},
		{"mares", "mare"},
		{"casas", "casa"},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got, _ := applyStep(tt.input, step)

			if tt.want != got {
				t.Fatalf("invalid stem output, %q -> %q (got %q)", tt.input, tt.want, got)
			}
		})
	}
}

func TestFeminine(t *testing.T) {
	step := steps["Feminine"]

	tests := []struct {
		input string
		want  string
	}{
		{"chefona", "chef\u00e3o"},
		{"vilã", "vilã"},
		{"professora", "professor"},
		{"americana", "americano"},
		{"chilena", "chileno"},
		{"inglesa", "ingl\u00eas"},
		{"famosa", "famoso"},
		{"man\u00edaca", "man\u00edaco"},
		{"pr\u00e1tica", "pr\u00e1tico"},
		{"mantida", "mantido"},
		{"cansada", "cansado"},
		{"prima", "prima"},
		{"passiva", "passivo"},
		{"primeira", "primeiro"},
		{"sozinha", "sozinho"},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got, _ := applyStep(tt.input, step)

			if tt.want != got {
				t.Fatalf("invalid stem output, %q -> %q (got %q)", tt.input, tt.want, got)
			}
		})
	}
}

func TestAugmentative(t *testing.T) {
	step := steps["Augmentative"]

	tests := []struct {
		input string
		want  string
	}{
		{"cansad\u00edssimo", "cansa"},
		{"fort\u00edssimo", "fort"},
		{"chiqu\u00e9rrimo", "chiqu"},
		{"pezinho", "pe"},
		{"maluquinho", "maluc"},
		{"amiguinho", "amig"},
		{"cansadinho", "cans"},
		{"carrinho", "carr"},
		{"grandalh\u00e3o", "grand"},
		{"dentu\u00e7a", "dent"},
		{"mulhera\u00e7o", "mulher"},
		{"cansad\u00e3o", "cans"},
		{"corp\u00e1zio", "corp"},
		{"pratarraz", "prat"},
		{"bocarra", "boc"},
		{"calorz\u00e3o", "calor"},
		{"menin\u00e3o", "menin"},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got, _ := applyStep(tt.input, step)

			if tt.want != got {
				t.Fatalf("invalid stem output, %q -> %q (got %q)", tt.input, tt.want, got)
			}
		})
	}
}

func TestNoun(t *testing.T) {
	step := steps["Noun"]

	tests := []struct {
		input string
		want  string
	}{

		{"existencialista", "exist"},
		{"minimalista", "minim"},
		{"contagem", "cont"},
		{"gerenciamento", "gerenc"},
		{"monitoramento", "monitor"},
		{"nascimento", "nasc"},
		{"comercializado", "comerci"},
		{"traumatizado", "traum"},
		{"alfabetizado", "alfabet"},
		{"associativo", "associ"},
		{"contraceptivo", "contracep"},
		{"esportivo", "espor"},
		{"abalado", "abal"},
		{"impedido", "imped"},
		{"ralador", "ral"},
		{"entendido", "entend"},
		{"cumpridor", "cumpr"},
		{"obrigat\u00f3ria", "obrig"},
		{"produtor", "produt"},
		{"comparabilidade", "compar"},
		{"abolicionista", "abol"},
		{"intervencionista", "interven"},
		{"profissional", "profiss"},
		{"refer\u00eancia", "refer"},
		{"repugn\u00e2ncia", "repugn"},
		{"abatedouro", "abat"},
		{"fofoqueiro", "fofoc"},
		{"brasileiro", "brasil"},
		{"gostoso", "gost"},
		{"comercializa\u00e7", "comerci"},
		{"consumismo", "consum"},
		{"alega\u00e7", "aleg"},
		{"aboli\u00e7", "abol"},
		{"anedot\u00e1rio", "anedot"},
		{"minist\u00e9rio", "minist"},
		{"chin\u00eas", "chin"},
		{"beleza", "bel"},
		{"rigidez", "rigid"},
		{"parentesco", "parent"},
		{"ocupante", "ocup"},
		{"bomb\u00e1stico", "bomb"},
		{"problem\u00e1tico", "problem"},
		{"pol\u00eamico", "pol\u00eam"},
		{"produtividade", "produt"},
		{"profundidade", "profund"},
		{"aposentadoria", "aposentad"},
		{"anedot\u00e1rio", "anedot"},
		{"existencial", "exist"},
		{"artista", "artista"},
		{"maluquice", "maluc"},
		{"chatice", "chat"},
		{"demon\u00edaco", "demon"},
		{"decorrente", "decorr"},
		{"criminal", "crim"},
		{"americano", "americ"},
		{"am\u00e1vel", "am"},
		{"combust\u00edvel", "combust"},
		{"cobertura", "cobert"},
		{"consensual", "consens"},
		{"mundial", "mund"},
		{"experimental", "experiment"},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got, _ := applyStep(tt.input, step)

			if tt.want != got {
				t.Fatalf("invalid stem output, %q -> %q (got %q)", tt.input, tt.want, got)
			}
		})
	}
}

func TestVerb(t *testing.T) {
	step := steps["Verb"]

	tests := []struct {
		input string
		want  string
	}{

		{"cantar\u00edamo", "cant"},
		{"cant\u00e1ssemo", "cant"},
		{"beber\u00edamo", "beb"},
		{"beb\u00eassemo", "beb"},
		{"partir\u00edamo", "part"},
		{"part\u00edssemo", "part"},
		{"cant\u00e1ramo", "cant"},
		{"cant\u00e1rei", "cant"},
		{"cantaremo", "cant"},
		{"cantariam", "cant"},
		{"cantar\u00edei", "cant"},
		{"cant\u00e1ssei", "cant"},
		{"cantassem", "cant"},
		{"cant\u00e1vamo", "cant"},
		{"beb\u00earamo", "beb"},
		{"beberemo", "beb"},
		{"beberiam", "beb"},
		{"beber\u00edei", "beb"},
		{"beb\u00eassei", "beb"},
		{"bebessem", "beb"},
		{"partir\u00edamo", "part"},
		{"partiremo", "part"},
		{"partiriam", "part"},
		{"partir\u00edei", "part"},
		{"part\u00edssei", "part"},
		{"partissem", "part"},
		{"cantando", "cant"},
		{"bebendo", "beb"},
		{"partindo", "part"},
		{"propondo", "prop"},
		{"cantaram", "cant"},
		{"cantarde", "cant"},
		{"cantarei", "cant"},
		{"cantarem", "cant"},
		{"cantaria", "cant"},
		{"cantarmo", "cant"},
		{"cantasse", "cant"},
		{"cantaste", "cant"},
		{"cantavam", "cant"},
		{"cant\u00e1vei", "cant"},
		{"beberam", "beb"},
		{"beberde", "beb"},
		{"beberei", "beb"},
		{"beb\u00earei", "beb"},
		{"beberem", "beb"},
		{"beberia", "beb"},
		{"bebermo", "beb"},
		{"bebesse", "beb"},
		{"bebeste", "beb"},
		{"beb\u00edamo", "beb"},
		{"partiram", "part"},
		{"conclu\u00edram", "conclu"},
		{"partirde", "part"},
		{"partirei", "part"},
		{"partirem", "part"},
		{"partiria", "part"},
		{"partirmo", "part"},
		{"partisse", "part"},
		{"partiste", "part"},
		{"cantamo", "cant"},
		{"cantara", "cant"},
		{"cantar\u00e1", "cant"},
		{"cantare", "cant"},
		{"cantava", "cant"},
		{"cantemo", "cant"},
		{"bebera", "beb"},
		{"beber\u00e1", "beb"},
		{"bebere", "beb"},
		{"bebiam", "beb"},
		{"beb\u00edei", "beb"},
		{"partimo", "part"},
		{"partira", "part"},
		{"partir\u00e1", "part"},
		{"partire", "part"},
		{"compomo", "comp"},
		{"cantai", "cant"},
		{"cantam", "cant"},
		{"cheguei", "cheg"},
		{"cantei", "cant"},
		{"cantem", "cant"},
		{"beber", "beb"},
		{"bebeu", "beb"},
		{"bebia", "beb"},
		{"partir", "part"},
		{"partiu", "part"},
		{"chegou", "cheg"},
		{"bebi", "beb"},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got, _ := applyStep(tt.input, step)

			if tt.want != got {
				t.Fatalf("invalid stem output, %q -> %q (got %q)", tt.input, tt.want, got)
			}
		})
	}
}

func TestVowel(t *testing.T) {
	step := steps["Vowel"]

	tests := []struct {
		input string
		want  string
	}{
		{"menina", "menin"},
		{"grande", "grand"},
		{"menino", "menin"},
		{"coração", "coraçã"},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got, _ := applyStep(tt.input, step)

			if tt.want != got {
				t.Fatalf("invalid stem output, %q -> %q (got %q)", tt.input, tt.want, got)
			}
		})
	}
}

func TestStem(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{

		{"", ""},
		{"            ", ""},
		{"a1a", "a1a"},
		{"coração", "coraca"},
		{"coraçãozinho", "coraca"},
		{"funcionamento", "funcion"},
		{"nervosos", "nerv"},
		{"continuar", "continu"},
		{"continuando", "continu"},
		{"demonstração", "demonstr"},
		{"finalidades", "final"},
		{"utilizar-se", "utilizar-s"},
		{"infelizmente", "infeliz"},
		{"comentário", "coment"},
		{"comentários", "coment"},
		{"bons", "bom"},
		{"bal\u00f5es", "bal"},
		{"capit\u00e3es", "capita"},
		{"normais", "norm"},
		{"am\u00e1veis", "am"},
		{"len\u00e7\u00f3is", "lencol"},
		{"barris", "barril"},
		{"males", "mal"},
		{"mares", "mar"},
		{"casas", "cas"},
		{"chefona", "chef"},
		{"vilã", "vila"},
		{"professora", "profes"},
		{"americana", "americ"},
		{"chilena", "chilen"},
		{"inglesa", "ingl"},
		{"famosa", "fam"},
		{"man\u00edaca", "man"},
		{"pr\u00e1tica", "prat"},
		{"mantida", "mant"},
		{"cansada", "cans"},
		{"prima", "prim"},
		{"passiva", "passiv"},
		{"primeira", "prim"},
		{"sozinha", "so"},
		{"felizmente", "feliz"},
		{"cansad\u00edssimo", "cans"},
		{"amabil\u00edssimo", "amavel"},
		{"fort\u00edssimo", "fort"},
		{"chiqu\u00e9rrimo", "chiqu"},
		{"pezinho", "pe"},
		{"maluquinho", "maluc"},
		{"amiguinho", "amig"},
		{"cansadinho", "cans"},
		{"carrinho", "carr"},
		{"grandalh\u00e3o", "grand"},
		{"dentu\u00e7a", "dent"},
		{"mulhera\u00e7o", "mulh"},
		{"cansad\u00e3o", "cans"},
		{"corp\u00e1zio", "corp"},
		{"pratarraz", "prat"},
		{"bocarra", "boc"},
		{"calorz\u00e3o", "cal"},
		{"menin\u00e3o", "menin"},
		{"existencialista", "exist"},
		{"minimalista", "minim"},
		{"contagem", "cont"},
		{"gerenciamento", "gerenc"},
		{"monitoramento", "monitor"},
		{"nascimento", "nasc"},
		{"comercializado", "comerci"},
		{"traumatizado", "traum"},
		{"alfabetizado", "alfabet"},
		{"associativo", "associ"},
		{"contraceptivo", "contracep"},
		{"esportivo", "espor"}, // diferente do artigo. o artigo est\u00e1 errad},
		{"abalado", "abal"},
		{"impedido", "imped"},
		{"ralador", "ral"},
		{"entendido", "entend"},
		{"cumpridor", "cumpr"},
		{"obrigat\u00f3ria", "obrig"},
		{"produtor", "produt"},
		{"comparabilidade", "compar"},
		{"abolicionista", "abol"},
		{"intervencionista", "interven"},
		{"profissional", "profiss"},
		{"refer\u00eancia", "refer"},
		{"repugn\u00e2ncia", "repugn"},
		{"abatedouro", "abat"},
		{"fofoqueiro", "fofoc"},
		{"brasileiro", "brasil"},
		{"gostoso", "gost"},
		{"comercializa\u00e7", "comerci"},
		{"consumismo", "consum"},
		{"concretiza\u00e7\u00e3o", "concre"},
		{"alega\u00e7", "aleg"},
		{"aboli\u00e7", "abol"},
		{"anedot\u00e1rio", "anedot"},
		{"minist\u00e9rio", "minist"},
		{"chin\u00eas", "chin"},
		{"beleza", "bel"},
		{"rigidez", "rigid"},
		{"parentesco", "parent"},
		{"ocupante", "ocup"},
		{"bomb\u00e1stico", "bomb"},
		{"problem\u00e1tico", "problem"},
		{"pol\u00eamico", "polem"},
		{"produtividade", "produt"},
		{"profundidade", "profund"},
		{"aposentadoria", "aposentad"},
		{"anedot\u00e1rio", "anedot"},
		{"existencial", "exist"},
		{"artista", "artist"},
		{"maluquice", "maluc"},
		{"chatice", "chat"},
		{"demon\u00edaco", "demon"},
		{"decorrente", "decorr"},
		{"criminal", "crim"},
		{"americano", "americ"},
		{"am\u00e1vel", "am"},
		{"combust\u00edvel", "combust"},
		{"cobertura", "cobert"},
		{"consensual", "consens"},
		{"mundial", "mund"},
		{"experimental", "experiment"},
		{"cantar\u00edamo", "cant"},
		{"cant\u00e1ssemo", "cant"},
		{"beber\u00edamo", "beb"},
		{"beb\u00eassemo", "beb"},
		{"partir\u00edamo", "part"},
		{"part\u00edssemo", "part"},
		{"cant\u00e1ramo", "cant"},
		{"cant\u00e1rei", "cant"},
		{"cantaremo", "cant"},
		{"cantariam", "cant"},
		{"cantar\u00edei", "cant"},
		{"cant\u00e1ssei", "cant"},
		{"cantassem", "cant"},
		{"cant\u00e1vamo", "cant"},
		{"beb\u00earamo", "beb"},
		{"beberemo", "beb"},
		{"beberiam", "beb"},
		{"beber\u00edei", "beb"},
		{"beb\u00eassei", "beb"},
		{"bebessem", "beb"},
		{"partir\u00edamo", "part"},
		{"partiremo", "part"},
		{"partiriam", "part"},
		{"partir\u00edei", "part"},
		{"part\u00edssei", "part"},
		{"partissem", "part"},
		{"cantando", "cant"},
		{"bebendo", "beb"},
		{"partindo", "part"},
		{"propondo", "prop"},
		{"cantaram", "cant"},
		{"cantarde", "cant"},
		{"cantarei", "cant"},
		{"cantarem", "cant"},
		{"cantaria", "cant"},
		{"cantarmo", "cant"},
		{"cantasse", "cant"},
		{"cantaste", "cant"},
		{"cantavam", "cant"},
		{"cant\u00e1vei", "cant"},
		{"beberam", "beb"},
		{"beberde", "beb"},
		{"beberei", "beb"},
		{"beb\u00earei", "beb"},
		{"beberem", "beb"},
		{"beberia", "beb"},
		{"bebermo", "beb"},
		{"bebesse", "beb"},
		{"bebeste", "beb"},
		{"beb\u00edamo", "beb"},
		{"partiram", "part"},
		{"conclu\u00edram", "conclu"},
		{"partirde", "part"},
		{"partirei", "part"},
		{"partirem", "part"},
		{"partiria", "part"},
		{"partirmo", "part"},
		{"partisse", "part"},
		{"partiste", "part"},
		{"cantamo", "cant"},
		{"cantara", "cant"},
		{"cantar\u00e1", "cant"},
		{"cantare", "cant"},
		{"cantava", "cant"},
		{"cantemo", "cant"},
		{"bebera", "beb"},
		{"beber\u00e1", "beb"},
		{"bebere", "beb"},
		{"bebiam", "beb"},
		{"beb\u00edei", "beb"},
		{"partimo", "part"},
		{"partira", "part"},
		{"partir\u00e1", "part"},
		{"partire", "part"},
		{"compomo", "comp"},
		{"cantai", "cant"},
		{"cantam", "cant"},
		{"cheguei", "cheg"},
		{"cantei", "cant"},
		{"cantem", "cant"},
		{"beber", "beb"},
		{"bebeu", "beb"},
		{"bebia", "beb"},
		{"partir", "part"},
		{"partiu", "part"},
		{"chegou", "cheg"},
		{"bebi", "beb"},
		{"menina", "menin"},
		{"grande", "grand"},
		{"menino", "menin"},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := Stem(tt.input)

			if tt.want != got {
				t.Fatalf("invalid stem output, %q -> %q (got %q)", tt.input, tt.want, got)
			}
		})
	}
}

func TestSentence(t *testing.T) {
	tests := []struct {
		input         string
		want          string
		removeAccents bool
	}{
		{
			"Que você faça o bem e não o mal.",
			"que voc fac o bem e na o mal.",
			true,
		},
		{
			"Que você encontre perdão para si mesmo e perdoe os outros.",
			"que voc encontr perd par si mesm e perdo os outros.",
			true,
		},
		{
			"Que você compartilhe livremente, nunca recebendo mais do que você dá.",
			"que voc compartilh livremente, nunc receb mais do que voc dá.",
			false,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := StemSentence(tt.input, tt.removeAccents)

			if tt.want != got {
				t.Fatalf("invalid stem sentence output, %q -> %q (got %q)", tt.input, tt.want, got)
			}
		})
	}
}
