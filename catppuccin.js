const fs = require("fs");
const {variants} = require('@catppuccin/palette');

function run() {
    for (let name in variants) {
        const variant = variants[name];
        const style = template
            .replaceAll("{{variant}}", name.charAt(0).toUpperCase() + name.substring(1))
            .replaceAll("{{variant-lower}}", name)
            .replaceAll("{{flamingo}}", variant.flamingo.hex)
            .replaceAll("{{pink}}", variant.pink.hex)
            .replaceAll("{{mauve}}", variant.mauve.hex)
            .replaceAll("{{red}}", variant.red.hex)
            .replaceAll("{{maroon}}", variant.maroon.hex)
            .replaceAll("{{peach}}", variant.peach.hex)
            .replaceAll("{{yellow}}", variant.yellow.hex)
            .replaceAll("{{green}}", variant.green.hex)
            .replaceAll("{{teal}}", variant.teal.hex)
            .replaceAll("{{sky}}", variant.sky.hex)
            .replaceAll("{{blue}}", variant.blue.hex)
            .replaceAll("{{lavender}}", variant.lavender.hex)
            .replaceAll("{{text}}", variant.text.hex)
            .replaceAll("{{overlay0}}", variant.overlay0.hex)
            .replaceAll("{{surface2}}", variant.surface2.hex)
            .replaceAll("{{surface0}}", variant.surface0.hex)
            .replaceAll("{{base}}", variant.base.hex);
        fs.writeFileSync(`catppuccin-${name}.go`, style);
    }
}

const template = `package catppuccin

import (
	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/styles"
)

var (
ctp{{variant}}Flamingo = "{{flamingo}}"
ctp{{variant}}Pink     = "{{pink}}"
ctp{{variant}}Mauve    = "{{mauve}}"
ctp{{variant}}Red      = "{{red}}"
ctp{{variant}}Maroon   = "{{maroon}}"
ctp{{variant}}Peach    = "{{peach}}"
ctp{{variant}}Yellow   = "{{yellow}}"
ctp{{variant}}Green    = "{{green}}"
ctp{{variant}}Teal     = "{{teal}}"
ctp{{variant}}Sky      = "{{sky}}"
ctp{{variant}}Blue     = "{{blue}}"
ctp{{variant}}Lavender = "{{lavender}}"
ctp{{variant}}Text     = "{{text}}"
ctp{{variant}}Overlay0 = "{{overlay0}}"
ctp{{variant}}Surface2 = "{{surface2}}"
ctp{{variant}}Surface0 = "{{surface0}}"
ctp{{variant}}Base     = "{{base}}"
)

// Catppuccin{{variant}} a soothing pastel theme for the high-spirited
var Catppuccin{{variant}} = styles.Register(chroma.MustNewStyle("catppuccin-{{variant-lower}}", chroma.StyleEntries{
chroma.TextWhitespace:        ctp{{variant}}Surface0,
chroma.Comment:               "italic " + ctp{{variant}}Surface2,
chroma.CommentPreproc:        ctp{{variant}}Blue,
chroma.Keyword:               ctp{{variant}}Mauve,
chroma.KeywordPseudo:         "bold " + ctp{{variant}}Mauve,
chroma.KeywordType:           ctp{{variant}}Flamingo,
chroma.KeywordConstant:       "italic " + ctp{{variant}}Mauve,
chroma.Operator:              ctp{{variant}}Sky,
chroma.OperatorWord:          "bold " + ctp{{variant}}Sky,
chroma.Name:                  ctp{{variant}}Lavender,
chroma.NameBuiltin:           "italic " + ctp{{variant}}Text,
chroma.NameFunction:          ctp{{variant}}Sky,
chroma.NameClass:             ctp{{variant}}Yellow,
chroma.NameNamespace:         ctp{{variant}}Yellow,
chroma.NameException:         ctp{{variant}}Maroon,
chroma.NameVariable:          ctp{{variant}}Peach,
chroma.NameConstant:          ctp{{variant}}Yellow,
chroma.NameLabel:             ctp{{variant}}Yellow,
chroma.NameEntity:            ctp{{variant}}Pink,
chroma.NameAttribute:         ctp{{variant}}Yellow,
chroma.NameTag:               ctp{{variant}}Mauve,
chroma.NameDecorator:         ctp{{variant}}Pink,
chroma.NameOther:             ctp{{variant}}Peach,
chroma.Punctuation:           ctp{{variant}}Text,
chroma.LiteralString:         ctp{{variant}}Green,
chroma.LiteralStringDoc:      ctp{{variant}}Green,
chroma.LiteralStringInterpol: ctp{{variant}}Green,
chroma.LiteralStringEscape:   ctp{{variant}}Blue,
chroma.LiteralStringRegex:    ctp{{variant}}Blue,
chroma.LiteralStringSymbol:   ctp{{variant}}Green,
chroma.LiteralStringOther:    ctp{{variant}}Green,
chroma.LiteralNumber:         ctp{{variant}}Teal,
chroma.GenericHeading:        "bold " + ctp{{variant}}Sky,
chroma.GenericSubheading:     "bold " + ctp{{variant}}Sky,
chroma.GenericDeleted:        ctp{{variant}}Maroon,
chroma.GenericInserted:       ctp{{variant}}Green,
chroma.GenericError:          ctp{{variant}}Maroon,
chroma.GenericEmph:           "italic",
chroma.GenericStrong:         "bold",
chroma.GenericPrompt:         "bold " + ctp{{variant}}Overlay0,
chroma.GenericOutput:         ctp{{variant}}Peach,
chroma.GenericTraceback:      ctp{{variant}}Maroon,
chroma.Error:                 ctp{{variant}}Red,
chroma.Background:            ctp{{variant}}Peach + " bg:" + ctp{{variant}}Base,
}))`

if (require.main === module) {
    run();
}
