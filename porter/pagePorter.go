package porter

import "regexp"

var imports = regexp.MustCompile(`import \{.*\} from '((.*\/selenium\/.*[Ii]nteractions)|protractor)';`)
var class = regexp.MustCompile(`class (\w*)`)
var others = regexp.MustCompile(`(async|await)`)
var ef = regexp.MustCompile("ElementFinder")
var eaf = regexp.MustCompile("ElementArrayFinder")
var efLocator = regexp.MustCompile(`\$\((.*)\)`)
var eafLocator = regexp.MustCompile(`\$\$\((.*)\)`)
var pBool = regexp.MustCompile(`Promise<boolean>`)
var pVoid = regexp.MustCompile(`Promise<void>`)
var p = regexp.MustCompile(`Promise`)

// function patterns
var gtxtNan = regexp.MustCompile(`getNumberOrDefaultIfNaN(.|\n)*getText(.|\n)*elementOrLocator:\s*(\w.\w)(\,|.|\n)*\}\)\);`)
var gtxtNum = regexp.MustCompile(`getNumberFromText(.|\n)*getText(.|\n)*elementOrLocator:\s*(\w.\w)(\,|.|\n)*\}\)\);`)
var gtxt = regexp.MustCompile(`getText(.|\s)*elementOrLocator:\s*(\w*\.\w*)(\,)*(.|\s)*\}\);`)
var gtxtn = regexp.MustCompile(`getText(.|\s)*elementOrLocator:\s*(\w*\.\w*)(\,)*(.|\s)*\}\);\s*return getNumberFromText.*`)
var click = regexp.MustCompile(`click(.|\s)*elementOrLocator:\s*(\w*\.\w*)(.|\s)*\}\);`)
var stxt = regexp.MustCompile(`setText(.|\s)*elementOrLocator:\s*(\w*\.\w*)(.|\s)*\}\, value\);`)
var gattr = regexp.MustCompile(`getAttribute(.|\s)*elementOrLocator:\s*(\w*\.\w*)(.|\s)*\}\, 'value'\);`)
var gattrn = regexp.MustCompile(`getAttribute(.|\s)*elementOrLocator:\s*(\w*\.\w*)(.|\s)*\}\, 'value'\);\s*return getNumberFromText.*`)
var ddl = regexp.MustCompile(`selectDropdownByText(.|\s)*elementOrLocator:\s*(\w*\.\w*)(.|\s)*\}\, (.*)\);`)
var ied = regexp.MustCompile(`return isElementDisplayed(.|\n)*elementOrLocator:\s*(\w.\w)(\,|.|\n)*\}\);`)
var gtxtEls = regexp.MustCompile(`getTextFromElementArrayFinder(.|\n)*elementsOrLocator:\s*(\w.\w)(\,|.|\n)*\}\);`)

func PortPage(s string) string {
	className := class.FindStringSubmatch(s)[1]

	value := imports.ReplaceAllString(s, "")
	value = others.ReplaceAllString(value, "")
	value = ef.ReplaceAllString(value, "string")
	value = eaf.ReplaceAllString(value, "string")
	value = eafLocator.ReplaceAllString(value, "$1")
	value = efLocator.ReplaceAllString(value, "$1")
	value = pBool.ReplaceAllString(value, className)
	value = pVoid.ReplaceAllString(value, className)
	value = p.ReplaceAllString(value, "Chainable")

	value = gtxtNan.ReplaceAllString(value, `cy
      .get(${3})
      .getText<number|null>({ formatFn: getNumberOrDefaultIfNaN });`)

	value = gtxtNum.ReplaceAllString(value, `cy
      .get(${3})
      .getText<number>({ formatFn: getNumberFromText });`)

	value = gtxtn.ReplaceAllString(value, `cy
      .get(${2})
      .getText<number>({ formatFn: getNumberFromText });`)

	value = gtxt.ReplaceAllString(value, `cy
      .get(${2})
      .getText();`)

	value = gattrn.ReplaceAllString(value, `cy
      .get(${2})
      .getText();`)

	value = gattr.ReplaceAllString(value, `cy
      .get(${2})
      .getText();`)

	value = click.ReplaceAllString(value, `cy
      .get(${2})
      .click();
		
		return this;`)

	value = stxt.ReplaceAllString(value, `cy
      .get(${2})
      .enterText(value);
		
		return this;`)

	value = ddl.ReplaceAllString(value, `cy
      .get(${2})
      .enterText(value);
		
		return this;`)

	value = ied.ReplaceAllString(value, `cy
      .get(${2})
      .should('be.visible');
		
		return this;`)

	value = gtxtEls.ReplaceAllString(value, `cy
      .get(${2})
      .getTextFromElements();`)

	return value
}
