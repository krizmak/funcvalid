// Factory functions composed from Regexp factory function and regexp collection from
// https://github.com/go-playground/validator
package funcvalid

import (
	"errors"
	"net/url"
	"os"
	"strings"

	urn "github.com/leodido/go-urn"
)

var (
	Alpha                 = RegexpRE(alphaRegex)
	AlphaNumeric          = RegexpRE(alphaNumericRegex)
	AlphaUnicode          = RegexpRE(alphaUnicodeRegex)
	AlphaUnicodeNumeric   = RegexpRE(alphaUnicodeNumericRegex)
	Numeric               = RegexpRE(numericRegex)
	Number                = RegexpRE(numberRegex)
	Hexadecimal           = RegexpRE(hexadecimalRegex)
	HexColor              = RegexpRE(hexColorRegex)
	Rgb                   = RegexpRE(rgbRegex)
	Rgba                  = RegexpRE(rgbaRegex)
	Hsl                   = RegexpRE(hslRegex)
	Hsla                  = RegexpRE(hslaRegex)
	E164                  = RegexpRE(e164Regex)
	Email                 = RegexpRE(emailRegex)
	Base64                = RegexpRE(base64Regex)
	Base64URL             = RegexpRE(base64URLRegex)
	Base64RawURL          = RegexpRE(base64RawURLRegex)
	ISBN10                = RegexpRE(iSBN10Regex)
	ISBN13                = RegexpRE(iSBN13Regex)
	UUID3                 = RegexpRE(uUID3Regex)
	UUID4                 = RegexpRE(uUID4Regex)
	UUID5                 = RegexpRE(uUID5Regex)
	UUID                  = RegexpRE(uUIDRegex)
	UUID3RFC4122          = RegexpRE(uUID3RFC4122Regex)
	UUID4RFC4122          = RegexpRE(uUID4RFC4122Regex)
	UUID5RFC4122          = RegexpRE(uUID5RFC4122Regex)
	UUIDRFC4122           = RegexpRE(uUIDRFC4122Regex)
	ULID                  = RegexpRE(uLIDRegex)
	Md4                   = RegexpRE(md4Regex)
	Md5                   = RegexpRE(md5Regex)
	Sha256                = RegexpRE(sha256Regex)
	Sha384                = RegexpRE(sha384Regex)
	Sha512                = RegexpRE(sha512Regex)
	Ripemd128             = RegexpRE(ripemd128Regex)
	Ripemd160             = RegexpRE(ripemd160Regex)
	Tiger128              = RegexpRE(tiger128Regex)
	Tiger160              = RegexpRE(tiger160Regex)
	Tiger192              = RegexpRE(tiger192Regex)
	ASCII                 = RegexpRE(aSCIIRegex)
	PrintableASCII        = RegexpRE(printableASCIIRegex)
	Multibyte             = RegexpRE(multibyteRegex)
	DataURI               = RegexpRE(dataURIRegex)
	Latitude              = RegexpRE(latitudeRegex)
	Longitude             = RegexpRE(longitudeRegex)
	SSN                   = RegexpRE(sSNRegex)
	HostnameRFC952        = RegexpRE(hostnameRegexRFC952)
	HostnameRFC1123       = RegexpRE(hostnameRegexRFC1123)
	FqdnRFC1123           = RegexpRE(fqdnRegexRFC1123)
	BtcAddress            = RegexpRE(btcAddressRegex)
	BtcUpperAddressBech32 = RegexpRE(btcUpperAddressRegexBech32)
	BtcLowerAddressBech32 = RegexpRE(btcLowerAddressRegexBech32)
	EthAddress            = RegexpRE(ethAddressRegex)
	URLEncoded            = RegexpRE(uRLEncodedRegex)
	HTMLEncoded           = RegexpRE(hTMLEncodedRegex)
	HTML                  = RegexpRE(hTMLRegex)
	JWT                   = RegexpRE(jWTRegex)
	SplitParams           = RegexpRE(splitParamsRegex)
	Bic                   = RegexpRE(bicRegex)
	Semver                = RegexpRE(semverRegex)
	DnsRFC1035Label       = RegexpRE(dnsRegexRFC1035Label)
	Cve                   = RegexpRE(cveRegex)
	Mongodb               = RegexpRE(mongodbRegex)
	Cron                  = RegexpRE(cronRegex)
	SpicedbID             = RegexpRE(spicedbIDRegex)
	SpicedbPermission     = RegexpRE(spicedbPermissionRegex)
	SpicedbType           = RegexpRE(spicedbTypeRegex)
)

var (
	Iso3166Alpha2       = KeyIn(iso3166_1_alpha2)
	Iso3166Alpha3       = KeyIn(iso3166_1_alpha3)
	Iso3166AlphaNumeric = KeyIn(iso3166_1_alpha_numeric)
	Iso4217             = KeyIn(iso4217)
	Iso4217Numeric      = KeyIn(iso4217_numeric)
	PostCodeByIso3166   = func(country_code string) Validator[string] {
		if _, ok := postCodePatternDict[country_code]; !ok {
			return ErrorValidator[string]("invalid country code")
		}
		return Regexp(postCodePatternDict[country_code])
	}
)

func Url(input string) error {
	if len(input) != 0 {
		url, err := url.Parse(input)
		if (err == nil) &&
			(url.Scheme != "") &&
			!(url.Host == "" && url.Fragment == "" && url.Opaque == "") {
			return nil
		}
	}
	return errors.New("error: Url")
}

func HttpUrl(input string) error {
	if len(input) != 0 {
		url, err := url.Parse(input)
		if err == nil {
			scheme := strings.ToLower(url.Scheme)
			if ((scheme == "http") || (scheme == "https")) &&
				!(url.Host == "" && url.Fragment == "" && url.Opaque == "") {
				return nil
			}
		}
	}
	return errors.New("error: HttpUrl")
}

func URI(input string) error {
	// checks needed as of Go 1.6 because of change https://github.com/golang/go/commit/617c93ce740c3c3cc28cdd1a0d712be183d0b328#diff-6c2d018290e298803c0c9419d8739885L195
	// emulate browser and strip the '#' suffix prior to validation. see issue-#237
	if i := strings.Index(input, "#"); i > -1 {
		input = input[:i]
	}

	if len(input) != 0 {
		_, err := url.ParseRequestURI(input)
		if err == nil {
			return nil
		}
	}

	return errors.New("error: URI")
}

// UrnRFC2141 is the validation function for validating if the input is a valid URN as per RFC 2141.
func UrnRFC2141(input string) error {
	_, match := urn.Parse([]byte(input))
	if match {
		return nil
	}

	return errors.New("error: UrnRFC2141")
}

// File is the validation function for validating if the input is a valid existing file path.
func File(input string) error {
	fileInfo, err := os.Stat(input)
	if err == nil &&
		!fileInfo.IsDir() {
		return nil
	}

	return errors.New("error: File")
}
