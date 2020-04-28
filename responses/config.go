package responses

var (
	Conf Config
)

type Config struct {
	AddressAPI     string `split_words:"true"`
	AddressWebsite string `split_words:"true"`
}
