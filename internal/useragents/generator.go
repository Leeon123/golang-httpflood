package useragents

import (
	"github.com/bonbonbon1414/golang-httpflood/internal/utils"
	"strconv"
)

// Generator handles user agent generation
type Generator struct{}

// NewGenerator creates a new user agent generator
func NewGenerator() *Generator {
	return &Generator{}
}

// Generate creates a random user agent string
func (g *Generator) Generate() string {
	platform := utils.String(platforms)
	
	var os string
	switch platform {
	case "Macintosh":
		os = utils.String(macOSVersions)
	case "Windows":
		os = utils.String(windowsVersions)
	case "X11":
		os = utils.String(linuxVersions)
	}
	
	browser := utils.String(browsers)
	
	switch browser {
	case "chrome":
		return g.generateChrome(os)
	case "ie":
		return g.generateIE(os)
	default: // spider
		return utils.String(spiderBots)
	}
}

// generateChrome generates a Chrome user agent
func (g *Generator) generateChrome(os string) string {
	webkit := strconv.Itoa(utils.IntRange(500, 599))
	version := strconv.Itoa(utils.IntRange(1, 99)) + ".0" +
		strconv.Itoa(utils.IntRange(1, 9999)) + "." +
		strconv.Itoa(utils.IntRange(1, 999))
	
	return "Mozilla/5.0 (" + os + ") AppleWebKit/" + webkit +
		".0 (KHTML, like Gecko) Chrome/" + version + " Safari/" + webkit
}

// generateIE generates an Internet Explorer user agent
func (g *Generator) generateIE(os string) string {
	version := strconv.Itoa(utils.IntRange(1, 99)) + ".0"
	engine := strconv.Itoa(utils.IntRange(1, 99)) + ".0"
	
	var token string
	if utils.Intn(2) == 1 {
		token = utils.String(ieOptions) + "; "
	}
	
	return "Mozilla/5.0 (compatible; MSIE " + version + "; " + os +
		"; " + token + "Trident/" + engine + ")"
}

// GetRandomSpider returns a random spider bot user agent
func (g *Generator) GetRandomSpider() string {
	return utils.String(spiderBots)
}

// GetRandomChrome returns a random Chrome user agent
func (g *Generator) GetRandomChrome() string {
	return g.generateChrome(utils.String(windowsVersions))
}

// GetRandomIE returns a random IE user agent
func (g *Generator) GetRandomIE() string {
	return g.generateIE(utils.String(windowsVersions))
}