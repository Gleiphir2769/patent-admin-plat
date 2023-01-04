package report

import (
	"bufio"
	"errors"
	"github.com/go-ego/gse"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type kv struct {
	Key   string
	Value float64
}

type Similarity struct {
	Count int
	Words []string
	Score float64
}

// TextSimilarity is a struct containing internal data to be re-used by the package.
type TextSimilarity struct {
	corpus            []string
	documents         []string
	documentFrequency map[string]int
}

// Option type describes functional options that
// allow modification of the internals of TextSimilarity
// before initialization. They are optional, and not using them
// allows you to use the defaults.
type Option func(TextSimilarity) TextSimilarity

func CutFirst(claims string) string {
	res := strings.Split(claims, "2.")
	res2 := strings.Split(res[0], "1.")
	if len(res2) > 0 {
		return res2[1]
	}
	return res2[0]
}

// Cosine returns the Cosine Similarity between two vectors.
func Cosine(a, b []float64) (float64, error) {
	count := 0
	lengthA := len(a)
	lengthB := len(b)
	if lengthA > lengthB {
		count = lengthA
	} else {
		count = lengthB
	}
	sumA := 0.0
	s1 := 0.0
	s2 := 0.0
	for k := 0; k < count; k++ {
		if k >= lengthA {
			s2 += math.Pow(b[k], 2)
			continue
		}
		if k >= lengthB {
			s1 += math.Pow(a[k], 2)
			continue
		}
		sumA += a[k] * b[k]
		s1 += math.Pow(a[k], 2)
		s2 += math.Pow(b[k], 2)
	}
	if s1 == 0 || s2 == 0 {
		return 0.0, errors.New("null vector")
	}
	return sumA / (math.Sqrt(s1) * math.Sqrt(s2)), nil
}

func count(key string, a []string) int {
	count := 0
	for _, s := range a {
		if key == s {
			count = count + 1
		}
	}
	return count
}

func getTime() string {
	t := time.Now()
	time := t.Format("2006年1月2日")
	return time
}

func tfidf(v string, tokens []string, n int, documentFrequency map[string]int) float64 {
	if documentFrequency[v] == 0 {
		return 0
	}
	tf := float64(count(v, tokens)) / float64(documentFrequency[v])
	idf := math.Log(float64(n) / (float64(documentFrequency[v])))

	return tf * idf
}

func union(a, b []string) []string {
	m := make(map[string]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; !ok {
			a = append(a, item)
		}
	}
	return a
}

func filter(vs []kv, f func(kv) bool) []kv {
	var vsf []kv
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// NewTextSimilarity accepts a slice of documents and
// creates the internal corpus and document frequency mapping.
func NewTextSimilarity(documents []string) *TextSimilarity {
	var (
		allTokens []string
	)

	ts := TextSimilarity{
		documents: documents,
	}

	ts.documentFrequency = map[string]int{}
	var seg gse.Segmenter
	seg.LoadDict()
	for _, doc := range documents {

		segments1 := seg.Segment([]byte(doc))
		resWords := RemoveStop(GetResult(segments1, 0))
		allTokens = append(allTokens, resWords...)
	}

	// Generate a corpus.
	for _, t := range allTokens {
		if ts.documentFrequency[t] == 0 {
			ts.documentFrequency[t] = 1
			ts.corpus = append(ts.corpus, t)
		} else {
			ts.documentFrequency[t] = ts.documentFrequency[t] + 1
		}
	}

	return &ts
}

// Similarity returns the cosine similarity between two documents using
// Tf-Idf vectorization using the corpus.
func (ts *TextSimilarity) Similarity(a, b []string) (float64, error) {
	combinedTokens := union(a, b)
	// Populate the vectors using frequency in the corpus.
	n := len(combinedTokens)
	vectorA := make([]float64, n)
	vectorB := make([]float64, n)
	for k, v := range combinedTokens {
		vectorA[k] = tfidf(v, a, n, ts.documentFrequency)
		vectorB[k] = tfidf(v, b, n, ts.documentFrequency)
	}

	similarity, err := Cosine(vectorA, vectorB)
	if err != nil {
		return 0.0, err
	}
	return similarity, nil
}

// Keywords accepts thresholds, which can be used to filter keyswords that
// are either they are too common or too Unique and returns a sorted list of
// keywords (index 0 being the lower tf-idf Score). Play with the thresholds
// according to your corpus.
func (ts *TextSimilarity) Keywords(threshLower, threshUpper float64, pattern int) []string {
	var (
		docKeywords = []kv{}
		result      = []string{}
	)
	var seg gse.Segmenter
	seg.LoadDict()
	for _, doc := range ts.documents {
		segments1 := seg.Segment([]byte(doc))
		tokens := RemoveStop(GetResult(segments1, 0))
		n := len(tokens)
		mapper := map[string]float64{}

		for _, v := range tokens {
			val := tfidf(v, tokens, n, ts.documentFrequency)
			mapper[v] = val
		}

		// Convert to a kv pair for convenience.
		i := 0
		vector := make([]kv, len(mapper))
		for k, v := range mapper {
			vector[i] = kv{
				Key:   k,
				Value: v,
			}
			i++
		}

		// Filter tf-idf, using threshold.
		vector = filter(vector, func(v kv) bool {
			return v.Value >= threshLower && v.Value <= threshUpper
		})

		// Select the most common Words relative to the corpus for this doc.

		if pattern != 0 {
			sort.Slice(vector, func(i, j int) bool {
				return vector[i].Value < vector[j].Value
			})
			docKeywords = append(docKeywords, vector...)
			break
		}
		docKeywords = append(docKeywords, vector...)
	}

	// Sort the vector based on tf-idf scores
	sort.Slice(docKeywords, func(i, j int) bool {
		return docKeywords[i].Value < docKeywords[j].Value
	})

	// Convert back to slice.
	for _, word := range docKeywords {
		result = append(result, word.Key)
	}
	return result
}

func Unique(resWords []string) []string {
	result := make([]string, len(resWords))
	result[0] = resWords[0]
	result_idx := 1
	for i := 0; i < len(resWords); i++ {
		is_repeat := false
		for j := 0; j < len(result); j++ {
			if resWords[i] == result[j] {
				is_repeat = true
				break
			}
		}
		if !is_repeat {
			result[result_idx] = resWords[i]
			result_idx++
		}
	}
	return result[:result_idx]
}

func RemoveStop(unstop []string) []string {
	file, err := os.Open("./app/user-agent/file1.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	stops := make([]string, 0)
	result := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		stop := strings.Split(line, "\n")
		stops = append(stops, stop[0])
	}
	for i := 0; i < len(unstop); i++ {
		same := 0
		for j := 0; j < len(stops); j++ {
			if stops[j] == unstop[i] {
				same = 1
				break
			}
		}
		if same == 0 {
			result = append(result, unstop[i])
		}
	}
	return result
}

func GetSimilar(word string) []string {
	var max = 4
	file, err := os.Open("./app/user-agent/cilin.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	similarwords := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		similarword := strings.Split(line, "\n")
		similarwords = append(similarwords, similarword[0])
	}
	for i := 0; i < len(similarwords); i++ {
		result := make([]string, 0)
		temp := strings.Split(similarwords[i], " ")
		var include = false
		var count = 0
		for j := 1; j < len(temp) && count < max; j++ {
			if temp[j] == word {
				include = true
			} else {
				result = append(result, temp[j])
				count++
			}
		}
		if include {
			return result
		}
	}
	return nil
}

func GetResult(segs []gse.Segment, pattern int, searchMode ...bool) []string {
	var mode bool
	var output []string
	if len(searchMode) > 0 {
		mode = searchMode[0]
	}

	if mode {
		for _, seg := range segs {
			output = append(output, seg.Token().Text())
		}
		return output
	}
	partOfSpeech := []string{"n", "v", "vn", "x", "an", "nz", "a", "l", "ns"}
	if pattern == 1 {
		partOfSpeech = []string{"n", "vn", "x", "an", "nz", "a", "l", "ns"}
	}
	if pattern == 2 {
		partOfSpeech = []string{"v"}
	}
	for _, seg := range segs {
		for i := 0; i < len(partOfSpeech); i++ {
			if seg.Token().Pos() == partOfSpeech[i] {
				output = append(output, seg.Token().Text())
				break
			}
		}
	}

	return output
}

func ToHtml(word string) string {
	result := strings.Replace(word, "\n", "<br> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;", -1)
	return result

}

func GenQueryExpression(word [][]string) string {
	result := ""
	num := len(word)
	for k := 0; k < 3; k++ {
		result += strconv.Itoa(k+1) + ".  "
		for i := k * num / 3; i < (k+1)*num/3; i++ {
			temp := ""
			if word[i] != nil {
				temp += " ( "
				for j := 0; j < len(word[i]); j++ {
					if j < len(word[i])-1 {
						temp += word[i][j] + " OR "
					} else {
						temp += word[i][j]
					}
				}
				temp += " ) "
				if i < (k+1)*num/3-1 {
					result += temp + " AND "
				} else {
					result += temp
				}
			}
		}
		result += "\n"
	}
	return result
}
