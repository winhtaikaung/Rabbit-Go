
  
  

  

# Rabbit-Go

  

  

![Logo](https://avatars3.githubusercontent.com/u/11961573?v=3&s=100)

  

[![Build Status](https://travis-ci.org/winhtaikaung/Rabbit-Go.svg?branch=master)](https://travis-ci.org/winhtaikaung/Rabbit-Go) 

[![GO Report](https://goreportcard.com/badge/github.com/winhtaikaung/Rabbit-Go)](https://goreportcard.com/report/github.com/winhtaikaung/Rabbit-Go)



  

**Zawgyi <=> Unicode Converter for Go lang (Go)**

  

## Installation

install using go get

`go get github.com/winhtaikaung/Rabbit-Go/rabbit`

  

## Usage

  

playground link

https://play.golang.org/p/8oUUuRKk0SD

  

```

import "rabbit"

  

func main() {

  

rabbit.Zg2uni("သီဟိုဠ္မွ ဉာဏ္ႀကီးရွင္သည္ အာယုဝၯနေဆးၫႊန္းစာကို ဇလြန္ေဈးေဘး ဗာဒံပင္ထက္ အဓိ႒ာန္လ်က္ ဂဃနဏဖတ္ခဲ့သည္။")

rabbit.Uni2zg("သီဟိုဠ်မှ ဉာဏ်ကြီးရှင်သည် အာယုဝဍ္ဎနဆေးညွှန်းစာကို ဇလွန်ဈေးဘေး ဗာဒံပင်ထက် အဓိဋ္ဌာန်လျက် ဂဃနဏဖတ်ခဲ့သည်။")

  

}

```

  

  

## Contributing

  

  

1. Fork it ( https://github.com/winhtaikaung/Rabbit-Go )

  

2. Create your feature branch (`git checkout -b my-new-feature`)

  

3. Commit your changes (`git commit -am 'Add some feature'`)

  

4. Push to the branch (`git push origin my-new-feature`)

  

5. Create a new Pull Request

  

  

## License

  

  

MIT