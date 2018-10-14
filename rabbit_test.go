package main

import (
	"testing"

	"github.com/winhtaikaung/Rabbit-Go/rabbit"
)

const UNI_STRING = string("သီဟိုဠ်မှ ဉာဏ်ကြီးရှင်သည် အာယုဝဍ္ဎနဆေးညွှန်းစာကို ဇလွန်ဈေးဘေး ဗာဒံပင်ထက် အဓိဋ္ဌာန်လျက် ဂဃနဏဖတ်ခဲ့သည်။")
const ZG_STRING = string("သီဟိုဠ္မွ ဉာဏ္ႀကီးရွင္သည္ အာယုဝၯနေဆးညႊန္းစာကို ဇလြန္ေဈးေဘး ဗာဒံပင္ထက္ အဓိ႒ာန္လ်က္ ဂဃနဏဖတ္ခဲ့သည္။")

func TestUni2zg(t *testing.T) {
	if rabbit.Uni2zg(UNI_STRING) != ZG_STRING {
		t.Log("Should match with Zawgyi string")
		t.Fail()
	}

}

func TestZg2uni(t *testing.T) {
	if rabbit.Zg2uni(ZG_STRING) != UNI_STRING {
		t.Log("Should match with Unicode string")
		t.Fail()
	}
}
