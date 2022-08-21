package mangrep

import (
	"errors"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	ErrEmptyPath    = errors.New("путь не должен быть пустым")
	ErrEmptyPattern = errors.New("паттерн не должен быть пустым")
)

type GrepFlags struct {
	A     int
	B     int
	C     int
	Count bool
	I     bool
	V     bool
	F     bool
	N     bool
}

// Grep хранить в себе зависимости и методы для поиска совпадений.
// Path,Pattern, GrepFlags должны быть переданы, чтобы начать поиск по совпадениям.
type Grep struct {
	Path    string
	Pattern string
	GrepFlags
	data      []byte
	lines     []string
	grepLines map[int]bool
}

// Start знает как начать поиск совпадений.
// В случае невалидных данных либо ошибки поиска вовращает ошибку.
func (g *Grep) Start() error {
	if g.Path == "" {
		return ErrEmptyPath
	}

	if g.Pattern == "" {
		return ErrEmptyPattern
	}

	dat, err := os.ReadFile(g.Path)
	if err != nil {
		return err
	}

	g.data = dat

	return g.start()
}

// start знает какие переменные необходимо проинициализировать
// и какие методы запустить для поиска совпадений.
func (g *Grep) start() error {
	g.grepLines = make(map[int]bool)
	g.lines = strings.Split(string(g.data), "\n")

	if g.F {
		g.setFixedLines()
		return nil
	}

	return g.setPatternLines()
}

// copyOriginalGrepLines копирует g.grepLines, в которых значение полей true.
func (g *Grep) copyOriginalGrepLines() map[int]bool {
	originalGrepLines := make(map[int]bool)
	for key, val := range g.grepLines {
		if val {
			originalGrepLines[key] = val
		}
	}

	return originalGrepLines
}

// setFixedLines знает как найти совпавшие строки по фиксированному значению и записать
// их индекс в g.grepLines
// Если найденная строка является предметом поиска, значению индекса
// в g.grepLines будет присвоенно значение true.
func (g *Grep) setFixedLines() {
	for i, line := range g.lines {
		if strings.Contains(line, g.Pattern) {
			g.grepLines[i] = true
		}
	}
}

// setPatternLines знает как найти совпавшие строки по паттерну и записать
// их индекс в g.grepLines
// Если найденная строка является предметом поиска, значению индекса
// в g.grepLines будет присвоенно значение true.
func (g *Grep) setPatternLines() error {
	caseInsensitive := ""
	if g.I {
		caseInsensitive = `(?i)`
	}

	re, err := regexp.Compile(caseInsensitive + g.Pattern)
	if err != nil {
		return err
	}

	for i, line := range g.lines {
		if re.MatchString(line) {
			g.grepLines[i] = true
		}
	}

	return nil
}

// setPatternLinesAfter знает как добавить в g.grepLines необходимые индексы после найденных совпадений
func (g *Grep) setPatternLinesAfter(a int) {
	originalGrepLines := g.copyOriginalGrepLines()

	for key := range originalGrepLines {
		afterLastIndex := key + a

		if afterLastIndex > len(g.lines)-1 {
			afterLastIndex = len(g.lines) - 1
		}

		for i := key + 1; i <= afterLastIndex; i++ {
			if _, ok := g.grepLines[i]; !ok {
				g.grepLines[i] = false
			}
		}
	}
}

// setPatternLinesAfter знает как добавить в g.grepLines необходимые индексы до найденных совпадений
func (g *Grep) setPatternLinesBefore(b int) {
	originalGrepLines := g.copyOriginalGrepLines()

	for key := range originalGrepLines {
		beforeLastIndex := key - b
		if beforeLastIndex < 0 {
			beforeLastIndex = 0
		}

		for i := beforeLastIndex; i < key; i++ {
			if _, ok := g.grepLines[i]; !ok {
				g.grepLines[i] = false
			}
		}
	}
}

// invertLines убирает строки попавшие под совпадение
func (g *Grep) invertLines() {
	if !g.V {
		return
	}

	for i := range g.lines {
		if val := g.grepLines[i]; val {
			delete(g.grepLines, i)
			continue
		}
		g.grepLines[i] = false
	}
}

// setPatternLinesContext знает как добавить в g.grepLines необходимые индексы до и послей найденных совпадений.
func (g *Grep) setPatternLinesContext() {
	g.setPatternLinesBefore(g.C)
	g.setPatternLinesAfter(g.C)
}

// countOriginalLines знает как посчитать количество совпадений.
func (g *Grep) countOriginalLines() int {
	return len(g.grepLines)
}

// Result знает как получить совпавшие строки с необходимыми флагами.
func (g *Grep) Result() string {
	if g.Count {
		return strconv.Itoa(g.countOriginalLines())
	}

	if g.V {
		g.invertLines()
		return g.allGrepLinesToString()
	}

	if g.C > 0 {
		g.setPatternLinesContext()
	} else {
		if g.B > 0 {
			g.setPatternLinesBefore(g.B)
		}
		if g.A > 0 {
			g.setPatternLinesAfter(g.A)
		}
	}

	return g.allGrepLinesToString()
}

// allGrepLinesToString знает как преобразовать все данные в g.grepLines в строку
func (g *Grep) allGrepLinesToString() string {
	result := make([]string, 0, len(g.grepLines))

	for i, item := range g.lines {
		if _, ok := g.grepLines[i]; ok {
			numLine := ""
			if g.N {
				numLine = strconv.Itoa(i+1) + ". "
			}
			result = append(result, numLine+item)
		}
	}

	return strings.Join(result, "\n")
}
