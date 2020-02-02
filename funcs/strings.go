package funcs

import (
	"fmt"
	"strings"
)

func Strings() StringsFuncs {
	return StringsFuncs{}
}

type StringsFuncs struct{}

// From https://golang.org/pkg/strings/
func (StringsFuncs) Compare(a, b string) int            { return strings.Compare(a, b) }
func (StringsFuncs) Contains(s, substr string) bool     { return strings.Contains(s, substr) }
func (StringsFuncs) ContainsAny(s, chars string) bool   { return strings.ContainsAny(s, chars) }
func (StringsFuncs) ContainsRune(s string, r rune) bool { return strings.ContainsRune(s, r) }
func (StringsFuncs) Count(s, substr string) int         { return strings.Count(s, substr) }
func (StringsFuncs) EqualFold(s, t string) bool         { return strings.EqualFold(s, t) }
func (StringsFuncs) Fields(s string) []string           { return strings.Fields(s) }
func (StringsFuncs) HasPrefix(s, prefix string) bool    { return strings.HasPrefix(s, prefix) }
func (StringsFuncs) HasSuffix(s, suffix string) bool    { return strings.HasSuffix(s, suffix) }
func (StringsFuncs) Index(s, substr string) int         { return strings.Index(s, substr) }
func (StringsFuncs) IndexAny(s, chars string) int       { return strings.IndexAny(s, chars) }
func (StringsFuncs) IndexByte(s string, c byte) int     { return strings.IndexByte(s, c) }
func (StringsFuncs) IndexRune(s string, r rune) int     { return strings.IndexRune(s, r) }
func (StringsFuncs) Join(as []interface{}, sep string) string {
	ss := []string{}
	for _, a := range as {
		ss = append(ss, fmt.Sprintf("%s", a))
	}
	return strings.Join(ss, sep)
}
func (StringsFuncs) LastIndex(s, substr string) int            { return strings.LastIndex(s, substr) }
func (StringsFuncs) LastIndexAny(s, chars string) int          { return strings.LastIndexAny(s, chars) }
func (StringsFuncs) LastIndexByte(s string, c byte) int        { return strings.LastIndexByte(s, c) }
func (StringsFuncs) Repeat(s string, count int) string         { return strings.Repeat(s, count) }
func (StringsFuncs) Replace(s, old, new string, n int) string  { return strings.Replace(s, old, new, n) }
func (StringsFuncs) ReplaceAll(s, old, new string) string      { return strings.ReplaceAll(s, old, new) }
func (StringsFuncs) Split(s, sep string) []string              { return strings.Split(s, sep) }
func (StringsFuncs) SplitAfter(s, sep string) []string         { return strings.SplitAfter(s, sep) }
func (StringsFuncs) SplitAfterN(s, sep string, n int) []string { return strings.SplitAfterN(s, sep, n) }
func (StringsFuncs) SplitN(s, sep string, n int) []string      { return strings.SplitN(s, sep, n) }
func (StringsFuncs) Title(s string) string                     { return strings.Title(s) }
func (StringsFuncs) ToLower(s string) string                   { return strings.ToLower(s) }
func (StringsFuncs) ToTitle(s string) string                   { return strings.ToTitle(s) }
func (StringsFuncs) ToUpper(s string) string                   { return strings.ToUpper(s) }
func (StringsFuncs) ToValidUTF8(s, replacement string) string {
	return strings.ToValidUTF8(s, replacement)
}
func (StringsFuncs) Trim(s string, cutset string) string      { return strings.Trim(s, cutset) }
func (StringsFuncs) TrimLeft(s string, cutset string) string  { return strings.TrimLeft(s, cutset) }
func (StringsFuncs) TrimPrefix(s, prefix string) string       { return strings.TrimPrefix(s, prefix) }
func (StringsFuncs) TrimRight(s string, cutset string) string { return strings.TrimRight(s, cutset) }
func (StringsFuncs) TrimSpace(s string) string                { return strings.TrimSpace(s) }
func (StringsFuncs) TrimSuffix(s, suffix string) string       { return strings.TrimSuffix(s, suffix) }

// Not supported:
// func (StringsFuncs) FieldsFunc(s string, f func(rune) bool) []string { return strings.FieldsFunc(s , f )}
// func (StringsFuncs) IndexFunc(s string, f func(rune) bool) int { return strings.IndexFunc(s , f )}
// func (StringsFuncs) LastIndexFunc(s string, f func(rune) bool) int { return strings.LastIndexFunc(s , f )}
// func (StringsFuncs) Map(mapping func(rune) rune, s string) string { return strings.Map(mapping , s )}
// func (StringsFuncs) ToLowerSpecial(c unicode.SpecialCase, s string) string { return strings.ToLowerSpecial(c , s )}
// func (StringsFuncs) ToTitleSpecial(c unicode.SpecialCase, s string) string { return strings.ToTitleSpecial(c , s )}
// func (StringsFuncs) ToUpperSpecial(c unicode.SpecialCase, s string) string { return strings.ToUpperSpecial(c , s )}
// func (StringsFuncs) TrimFunc(s string, f func(rune) bool) string { return strings.TrimFunc(s , f )}
// func (StringsFuncs) TrimLeftFunc(s string, f func(rune) bool) string { return strings.TrimLeftFunc(s , f )}
// func (StringsFuncs) TrimRightFunc(s string, f func(rune) bool) string { return strings.TrimRightFunc(s , f )}
