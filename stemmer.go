package rslp

import (
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

type rule struct {
	suffix      string
	minLength   int
	replacement string
	exceptions  []string
}

func (r *rule) apply(word string) (string, bool) {
	if len(word) < r.minLength+len(r.suffix) {
		return word, false
	}

	for _, e := range r.exceptions {
		if word == e {
			return word, false
		}
	}
	if strings.HasSuffix(word, r.suffix) {
		return word[:len(word)-len(r.suffix)] + r.replacement, true
	}
	return word, false
}

type step struct {
	stepPass   string
	stepFail   string
	minLength  int
	entireWord bool
	endWords   []string
	rules      []rule
}

/**
  Default step´s application flow, as written in the original RSPL code.
  It differs from the Orengo's paper, shown below:

  { 0 "Plural"       , 1 "Feminine"     , 1 "Feminine"     }  ,
  { 1 "Feminine"     , 3 "Augmentative" , 3 "Augmentative" }  ,
  { 3 "Augmentative" , 2 "Adverb"       , 2 "Adverb"       }  ,
  { 2 "Adverb"       , 4 "Noun"         , 4 "Noun"         }  ,
  { 4 "Noun"         ,   NULL           , 5 "Verb"         }  ,
  { 5 "Verb"         ,   NULL           , 6 "Vowel"        }  ,
  { 6 "Vowel"        ,   NULL           ,   NULL           }  ,
*/
var steps = map[string]*step{

	// Step 1: Plural Reduction
	"Plural": {"Feminine", "Feminine", 3, false, []string{"s"}, []rule{
		{"ns", 1, "m", nil},
		{"\u00f5es", 3, "\u00e3o", nil},
		{"\u00e3es", 1, "\u00e3o", []string{"m\u00e3e"}},
		{"ais", 1, "al", []string{"cais", "mais"}},
		{"\u00e9is", 2, "el", nil},
		{"eis", 2, "el", nil},
		{"\u00f3is", 2, "ol", nil},
		{"is", 2, "il", []string{"l\u00e1pis", "cais", "mais", "cr\u00facis", "biqu\u00ednis", "pois", "depois", "dois", "leis"}},
		{"les", 3, "l", nil},
		{"res", 3, "r", nil},
		{"s", 2, "", []string{
			"ali\u00e1s", "pires", "l\u00e1pis", "cais", "mais", "mas", "menos", "f\u00e9rias", "fezes", "p\u00easames",
			"cr\u00facis", "g\u00e1s", "atr\u00e1s", "mois\u00e9s", "atrav\u00e9s", "conv\u00e9s", "\u00eas", "pa\u00eds", "ap\u00f3s", "ambas",
			"ambos", "messias",
		}},
	}},

	// Step 2: Feminine Reduction
	"Feminine": {"Augmentative", "Augmentative", 3, false, []string{"a"}, []rule{
		{"ona", 3, "\u00e3o", []string{
			"abandona", "lona", "iona", "cortisona", "mon\u00f3tona", "maratona", "acetona", "detona", "carona",
		}},
		{"ora", 3, "or", nil},
		{"na", 4, "no", []string{
			"carona", "abandona", "lona", "iona", "cortisona", "mon\u00f3tona", "maratona", "acetona", "detona",
			"guiana", "campana", "grana", "caravana", "banana", "paisana",
		}},
		{"inha", 3, "inho", []string{"rainha", "linha", "minha"}},
		{"esa", 3, "\u00eas", []string{
			"mesa", "obesa", "princesa", "turquesa", "ilesa", "pesa", "presa",
		}},
		{"osa", 3, "oso", []string{"mucosa", "prosa"}},
		{"\u00edaca", 3, "\u00edaco", nil},
		{"ica", 3, "ico", []string{"dica"}},
		{"ada", 2, "ado", []string{"pitada"}},
		{"ida", 3, "ido", []string{"vida"}},
		{"\u00edda", 3, "ido", []string{"reca\u00edda", "sa\u00edda", "d\u00favida"}},
		{"ima", 3, "imo", []string{"v\u00edtima"}},
		{"iva", 3, "ivo", []string{"saliva", "oliva"}},
		{"eira", 3, "eiro", []string{
			"beira", "cadeira", "frigideira", "bandeira", "feira", "capoeira", "barreira", "fronteira",
			"besteira", "poeira",
		}},
		{"\u00e3", 2, "\u00e3o", []string{"amanh\u00e3", "arapu\u00e3", "f\u00e3", "div\u00e3"}},
	}},

	// Step 3: Adverb Reduction
	"Adverb": {"Noun", "Noun", 0, false, nil, []rule{
		{"mente", 4, "", []string{"experimente"}},
	}},

	// Step 4: Augmentative/Diminutive Reduction
	"Augmentative": {"Adverb", "Adverb", 0, false, nil, []rule{
		{"d\u00edssimo", 5, "", nil},
		{"abil\u00edssimo", 5, "", nil},
		{"\u00edssimo", 3, "", nil},
		{"\u00e9simo", 3, "", nil},
		{"\u00e9rrimo", 4, "", nil},
		{"zinho", 2, "", nil},
		{"quinho", 4, "c", nil},
		{"uinho", 4, "", nil},
		{"adinho", 3, "", nil},
		{"inho", 3, "", []string{"caminho", "cominho"}},
		{"alh\u00e3o", 4, "", nil},
		{"u\u00e7a", 4, "", nil},
		{"a\u00e7o", 4, "", []string{"antebra\u00e7o"}},
		{"a\u00e7a", 4, "", nil},
		{"ad\u00e3o", 4, "", nil},
		{"id\u00e3o", 4, "", nil},
		{"\u00e1zio", 3, "", []string{"top\u00e1zio"}},
		{"arraz", 4, "", nil},
		{"zarr\u00e3o", 3, "", nil},
		{"arr\u00e3o", 4, "", nil},
		{"arra", 3, "", nil},
		{"z\u00e3o", 2, "", []string{"coaliz\u00e3o"}},
		{"\u00e3o", 3, "", []string{
			"camar\u00e3o", "chimarr\u00e3o", "can\u00e7\u00e3o", "cora\u00e7\u00e3o", "embri\u00e3o", "grot\u00e3o", "glut\u00e3o",
			"fic\u00e7\u00e3o", "fog\u00e3o", "fei\u00e7\u00e3o", "furac\u00e3o", "gam\u00e3o", "lampi\u00e3o", "le\u00e3o", "macac\u00e3o", "na\u00e7\u00e3o",
			"\u00f3rf\u00e3o", "org\u00e3o", "patr\u00e3o", "port\u00e3o", "quinh\u00e3o", "rinc\u00e3o", "tra\u00e7\u00e3o",
			"falc\u00e3o", "espi\u00e3o", "mam\u00e3o", "foli\u00e3o", "cord\u00e3o", "aptid\u00e3o", "campe\u00e3o",
			"colch\u00e3o", "lim\u00e3o", "leil\u00e3o", "mel\u00e3o", "bar\u00e3o", "milh\u00e3o", "bilh\u00e3o", "fus\u00e3o",
			"crist\u00e3o", "ilus\u00e3o", "capit\u00e3o", "esta\u00e7\u00e3o", "sen\u00e3o",
		}},
	}},

	// Step 5: Noun Suffix Reduction
	"Noun": {"", "Verb", 0, false, nil, []rule{

		{"encialista", 4, "", nil},
		{"alista", 5, "", nil},
		{"agem", 3, "", []string{"coragem", "chantagem", "vantagem", "carruagem"}},
		{"ático", 3, "", nil},
		{"iamento", 4, "", nil},
		{"amento", 3, "", []string{"firmamento", "fundamento", "departamento"}},
		{"imento", 3, "", nil},
		{"mento", 6, "", []string{"firmamento", "elemento", "complemento", "instrumento", "departamento"}},
		{"alizado", 4, "", nil},
		{"atizado", 4, "", nil},
		{"tizado", 4, "", []string{"alfabetizado"}},
		{"izado", 5, "", []string{"organizado", "pulverizado"}},
		{"ativo", 4, "", []string{"pejorativo", "relativo"}},
		{"tivo", 4, "", []string{"relativo"}},
		{"ivo", 4, "", []string{"passivo", "possessivo", "pejorativo", "positivo"}},
		{"ado", 2, "", []string{"grado"}},
		{"ido", 3, "", []string{"cândido", "consolido", "rápido", "decido", "tímido", "duvido", "marido"}},
		{"ador", 3, "", nil},
		{"edor", 3, "", nil},
		{"idor", 4, "", []string{"ouvidor"}},
		{"dor", 4, "", []string{"ouvidor"}},
		{"sor", 4, "", []string{"assessor"}},
		{"atória", 5, "", nil},
		{"tor", 3, "", []string{"benfeitor", "leitor", "editor", "pastor", "produtor", "promotor", "consultor"}},
		{"or", 2, "", []string{"motor", "melhor", "redor", "rigor", "sensor", "tambor", "tumor", "assessor", "benfeitor", "pastor", "terior", "favor", "autor"}},
		{"abilidade", 5, "", nil},
		{"icionista", 4, "", nil},
		{"cionista", 5, "", nil},
		{"ionista", 5, "", nil},
		{"ionar", 5, "", nil},
		{"ional", 4, "", nil},
		{"ência", 3, "", nil},
		{"ância", 4, "", []string{"ambulância"}},
		{"edouro", 3, "", nil},
		{"queiro", 3, "c", nil},
		{"adeiro", 4, "", []string{"desfiladeiro"}},
		{"eiro", 3, "", []string{"desfiladeiro", "pioneiro", "mosteiro"}},
		{"uoso", 3, "", nil},
		{"oso", 3, "", []string{"precioso"}},
		{"alizaç", 5, "", nil},
		{"atizaç", 5, "", nil},
		{"tizaç", 5, "", nil},
		{"izaç", 5, "", []string{"organizaç"}},
		{"aç", 3, "", []string{"equaç", "relaç"}},
		{"iç", 3, "", []string{"eleição"}},
		{"ário", 3, "", []string{"voluntário", "salário", "aniversário", "diário", "lionário", "armário"}},
		{"atório", 3, "", nil},
		{"ário", 5, "", []string{"voluntário", "salário", "aniversário", "diário", "compulsório", "lionário", "próprio", "stério", "armário"}},
		{"ério", 6, "", nil},
		{"ês", 4, "", nil},
		{"eza", 3, "", nil},
		{"ez", 4, "", nil},
		{"esco", 4, "", nil},
		{"ante", 2, "", []string{"gigante", "elefante", "adiante", "possante", "instante", "restaurante"}},
		{"ástico", 4, "", []string{"eclesiástico"}},
		{"alístico", 3, "", nil},
		{"áutico", 4, "", nil},
		{"êutico", 4, "", nil},
		{"tico", 3, "", []string{"político", "eclesiástico", "diagnostico", "prático", "doméstico", "diagnóstico", "idêntico", "alopático", "artístico", "autêntico", "eclético", "crítico", "critico"}},
		{"ico", 4, "", []string{"tico", "público", "explico"}},
		{"ividade", 5, "", nil},
		{"idade", 4, "", []string{"autoridade", "comunidade"}},
		{"oria", 4, "", []string{"categoria"}},
		{"encial", 5, "", nil},
		{"ista", 4, "", nil},
		{"auta", 5, "", nil},
		{"quice", 4, "c", nil},
		{"ice", 4, "", []string{"cúmplice"}},
		{"íaco", 3, "", nil},
		{"ente", 4, "", []string{"freqüente", "alimente", "acrescente", "permanente", "oriente", "aparente"}},
		{"ense", 5, "", nil},
		{"inal", 3, "", nil},
		{"ano", 4, "", nil},
		{"ável", 2, "", []string{"afável", "razoável", "potável", "vulnerável"}},
		{"ível", 3, "", []string{"possível"}},
		{"vel", 5, "", []string{"possível", "vulnerável", "solúvel"}},
		{"bil", 3, "vel", nil},
		{"ura", 4, "", []string{"imatura", "acupuntura", "costura"}},
		{"ural", 4, "", nil},
		{"ual", 3, "", []string{"bissexual", "virtual", "visual", "pontual"}},
		{"ial", 3, "", nil},
		{"al", 4, "", []string{"afinal", "animal", "estatal", "bissexual", "desleal", "fiscal", "formal", "pessoal", "liberal", "postal", "virtual", "visual", "pontual", "sideral", "sucursal"}},
		{"alismo", 4, "", nil},
		{"ivismo", 4, "", nil},
		{"ismo", 3, "", []string{"cinismo"}},
	}},

	// Step 6: Verb Suffix Reduction
	"Verb": {"", "Vowel", 0, false, nil, []rule{
		{"ar\u00edamo", 2, "", nil},
		{"\u00e1ssemo", 2, "", nil},
		{"er\u00edamo", 2, "", nil},
		{"\u00eassemo", 2, "", nil},
		{"ir\u00edamo", 3, "", nil},
		{"\u00edssemo", 3, "", nil},
		{"\u00e1ramo", 2, "", nil},
		{"\u00e1rei", 2, "", nil},
		{"aremo", 2, "", nil},
		{"ariam", 2, "", nil},
		{"ar\u00edei", 2, "", nil},
		{"\u00e1ssei", 2, "", nil},
		{"assem", 2, "", nil},
		{"\u00e1vamo", 2, "", nil},
		{"\u00earamo", 3, "", nil},
		{"eremo", 3, "", nil},
		{"eriam", 3, "", nil},
		{"er\u00edei", 3, "", nil},
		{"\u00eassei", 3, "", nil},
		{"essem", 3, "", nil},
		{"\u00edramo", 3, "", nil},
		{"iremo", 3, "", nil},
		{"iriam", 3, "", nil},
		{"ir\u00edei", 3, "", nil},
		{"\u00edssei", 3, "", nil},
		{"issem", 3, "", nil},
		{"ando", 2, "", nil},
		{"endo", 3, "", nil},
		{"indo", 3, "", nil},
		{"ondo", 3, "", nil},
		{"aram", 2, "", nil},
		{"ar\u00e3o", 2, "", nil},
		{"arde", 2, "", nil},
		{"arei", 2, "", nil},
		{"arem", 2, "", nil},
		{"aria", 2, "", nil},
		{"armo", 2, "", nil},
		{"asse", 2, "", nil},
		{"aste", 2, "", nil},
		{"avam", 2, "", []string{"agravam"}},
		{"\u00e1vei", 2, "", nil},
		{"eram", 3, "", nil},
		{"er\u00e3o", 3, "", nil},
		{"erde", 3, "", nil},
		{"erei", 3, "", nil},
		{"\u00earei", 3, "", nil},
		{"erem", 3, "", nil},
		{"eria", 3, "", nil},
		{"ermo", 3, "", nil},
		{"esse", 3, "", nil},
		{"este", 3, "", []string{"faroeste", "agreste"}},
		{"\u00edamo", 3, "", nil},
		{"iram", 3, "", nil},
		{"\u00edram", 3, "", nil},
		{"ir\u00e3o", 2, "", nil},
		{"irde", 2, "", nil},
		{"irei", 3, "", []string{"admirei"}},
		{"irem", 3, "", []string{"adquirem"}},
		{"iria", 3, "", nil},
		{"irmo", 3, "", nil},
		{"isse", 3, "", nil},
		{"iste", 4, "", nil},
		{"iava", 4, "", []string{"ampliava"}},
		{"amo", 2, "", nil},
		{"iona", 3, "", nil},
		{"ara", 2, "", []string{"arara", "prepara"}},
		{"ar\u00e1", 2, "", []string{"alvar\u00e1"}},
		{"are", 2, "", []string{"prepare"}},
		{"ava", 2, "", []string{"agrava"}},
		{"emo", 2, "", nil},
		{"era", 3, "", []string{"acelera", "espera"}},
		{"er\u00e1", 3, "", nil},
		{"ere", 3, "", []string{"espere"}},
		{"iam", 3, "", []string{"enfiam", "ampliam", "elogiam", "ensaiam"}},
		{"\u00edei", 3, "", nil},
		{"imo", 3, "", []string{"reprimo", "intimo", "\u00edntimo", "nimo", "queimo", "ximo"}},
		{"ira", 3, "", []string{"fronteira", "s\u00e1tira"}},
		{"\u00eddo", 3, "", nil},
		{"ir\u00e1", 3, "", nil},
		{"tizar", 4, "", []string{"alfabetizar"}},
		{"izar", 5, "", []string{"organizar"}},
		{"itar", 5, "", []string{"acreditar", "explicitar", "estreitar"}},
		{"ire", 3, "", []string{"adquire"}},
		{"omo", 3, "", nil},
		{"ai", 2, "", nil},
		{"am", 2, "", nil},
		{"ear", 4, "", []string{"alardear", "nuclear"}},
		{"ar", 2, "", []string{"azar", "bazaar", "patamar"}},
		{"uei", 3, "", nil},
		{"u\u00eda", 5, "u", nil},
		{"ei", 3, "", nil},
		{"guem", 3, "g", nil},
		{"em", 2, "", []string{"alem", "virgem"}},
		{"er", 2, "", []string{"\u00e9ter", "pier"}},
		{"eu", 3, "", []string{"chapeu"}},
		{"ia", 3, "", []string{
			"est\u00f3ria", "fatia", "acia", "praia", "elogia", "mania", "l\u00e1bia", "aprecia",
			"pol\u00edcia", "arredia", "cheia", "\u00e1sia",
		}},
		{"ir", 3, "", []string{"freir"}},
		{"iu", 3, "", nil},
		{"eou", 5, "", nil},
		{"ou", 3, "", nil},
		{"i", 3, "", nil},
	}},

	// Step 7: Vowel Removal
	"Vowel": {"", "", 0, false, nil, []rule{
		{"bil", 2, "vel", nil},
		{"gue", 2, "g", []string{"gangue", "jegue"}},
		{"á", 3, "", nil},
		{"ê", 3, "", []string{"bebê"}},
		{"a", 3, "", []string{"ásia"}},
		{"e", 3, "", nil},
		{"o", 3, "", []string{"ão"}},
	}},

	// Step 8: Accents Removal
}

var normalize = transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)

func StemSentence(sentence string, removeDiacritics ...bool) string {
	var buf strings.Builder
	for index, word := range strings.Fields(sentence) {
		if index > 0 {
			buf.WriteByte(' ')
		}
		buf.WriteString(Stem(word, removeDiacritics...))
	}
	return buf.String()
}

func Stem(word string, removeDiacritics ...bool) string {
	if len(word) <= 3 {
		return strings.TrimSpace(strings.ToLower(word))
	}

	word = strings.TrimSpace(strings.ToLower(word))

	var ok bool
	var cur *step
	cur = steps["Plural"]

	for cur != nil {
		if word, ok = applyStep(word, cur); !ok {
			cur = steps[cur.stepFail]
		} else {
			cur = steps[cur.stepPass]
		}
	}

	if len(removeDiacritics) == 0 || removeDiacritics[0] {
		if s, _, e := transform.String(normalize, word); e == nil {
			return s
		}
	}

	return word
}

func applyStep(word string, cur *step) (string, bool) {
	if cur.minLength > 0 && len(word) < cur.minLength {
		return word, false
	} else if !hasSuffix(word, cur.endWords...) {
		return word, false
	}

	var ok bool
	for _, r := range cur.rules {
		if word, ok = r.apply(word); ok {
			return word, true
		}
	}
	return word, false
}

func hasSuffix(word string, suffix ...string) bool {
	if len(suffix) == 0 {
		return true
	}
	for _, s := range suffix {
		if strings.HasSuffix(word, s) {
			return true
		}
	}
	return false
}
