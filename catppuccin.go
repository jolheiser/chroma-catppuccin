package catppuccin

import (
	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/styles"
)

var (
	// Morning
	flamingo = "#F2CDCD"
	mauve    = "#DDB6F2"
	pink     = "#F5C2E7"
	maroon   = "#E8A2AF"
	red      = "#F28FAD"
	peach    = "#F8BD96"
	yellow   = "#FAE3B0"
	green    = "#ABE9B3"
	teal     = "#B5E8E0"
	blue     = "#96CDFB"
	sky      = "#89DCEB"

	// Night
	black0    = "#161320"
	black1    = "#1A1826"
	black2    = "#1E1E2E"
	black3    = "#302D41"
	black4    = "#575268"
	gray0     = "#6E6C7E"
	gray1     = "#988BA2"
	gray2     = "#C3BAC6"
	white     = "#D9E0EE"
	lavender  = "#C9CBFF"
	rosewater = "#F5E0DC"
)

// Catppuccin a soothing pastel theme for the high-spirited
var Catppuccin = styles.Register(chroma.MustNewStyle("catppuccin", chroma.StyleEntries{
	chroma.TextWhitespace:        black3,
	chroma.Comment:               "italic " + gray0,
	chroma.CommentPreproc:        blue,
	chroma.Keyword:               mauve,
	chroma.KeywordPseudo:         "bold " + mauve,
	chroma.KeywordType:           flamingo,
	chroma.KeywordConstant:       "italic " + mauve,
	chroma.Operator:              sky,
	chroma.OperatorWord:          "bold " + sky,
	chroma.Name:                  lavender,
	chroma.NameBuiltin:           "italic " + white,
	chroma.NameFunction:          sky,
	chroma.NameClass:             yellow,
	chroma.NameNamespace:         yellow,
	chroma.NameException:         maroon,
	chroma.NameVariable:          peach,
	chroma.NameConstant:          yellow,
	chroma.NameLabel:             yellow,
	chroma.NameEntity:            pink,
	chroma.NameAttribute:         yellow,
	chroma.NameTag:               mauve,
	chroma.NameDecorator:         pink,
	chroma.NameOther:             peach,
	chroma.Punctuation:           white,
	chroma.LiteralString:         green,
	chroma.LiteralStringDoc:      green,
	chroma.LiteralStringInterpol: green,
	chroma.LiteralStringEscape:   blue,
	chroma.LiteralStringRegex:    blue,
	chroma.LiteralStringSymbol:   green,
	chroma.LiteralStringOther:    green,
	chroma.LiteralNumber:         teal,
	chroma.GenericHeading:        "bold " + sky,
	chroma.GenericSubheading:     "bold " + sky,
	chroma.GenericDeleted:        maroon,
	chroma.GenericInserted:       green,
	chroma.GenericError:          maroon,
	chroma.GenericEmph:           "italic",
	chroma.GenericStrong:         "bold",
	chroma.GenericPrompt:         "bold " + gray1,
	chroma.GenericOutput:         peach,
	chroma.GenericTraceback:      maroon,
	chroma.Error:                 red,
	chroma.Background:            peach + " bg:" + black2,
}))
