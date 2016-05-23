package iespec

import (
	"encoding/base64"
)

func ExampleBytes() []byte {
	bytes, _ := base64.StdEncoding.DecodeString(ExPacket)
	return bytes
}

var ExPacket string = "AAoC1Fc+G4gAAABGAAAAAAACAsQBAAAxAAQAAQAIAAQABwACAAwABAALAAIAAQAIAAIACACYAAgAmQAIAD0AAQFz//+BBP//AACuS4ECAAQAAK5LgQMABAAArkuBBf//AACuS4EG//8AAK5LgQf//wAArkuBCP//AACuS4EJAAEAAK5LgQz//wAArkuBDf//AACuS4EO//8AAK5LgQ///wAArkuBEP//AACuS4ER//8AAK5LgRL//wAArkuBE///AACuS4EU//8AAK5LgRX//wAArkuBFv//AACuS4EX//8AAK5LgRj//wAArkuBGf//AACuS4EaAAEAAK5LgRv//wAArkuBHP//AACuS4Ed//8AAK5LgR7//wAArkuBH///AACuS4Eg//8AAK5LgSH//wAArkuBIv//AACuS4Ej//8AAK5LgST//wAArkuBJf//AACuS4EK//8AAK5LgQv//wAArkuBJv//AACuS4En//8AAK5LAQEAMQAEAAEAGwAQAAcAAgAcABAACwACAAEACAACAAgAmAAIAJkACAA9AAEBc///gQT//wAArkuBAgAEAACuS4EDAAQAAK5LgQX//wAArkuBBv//AACuS4EH//8AAK5LgQj//wAArkuBCQABAACuS4EM//8AAK5LgQ3//wAArkuBDv//AACuS4EP//8AAK5LgRD//wAArkuBEf//AACuS4ES//8AAK5LgRP//wAArkuBFP//AACuS4EV//8AAK5LgRb//wAArkuBF///AACuS4EY//8AAK5LgRn//wAArkuBGgABAACuS4Eb//8AAK5LgRz//wAArkuBHf//AACuS4Ee//8AAK5LgR///wAArkuBIP//AACuS4Eh//8AAK5LgSL//wAArkuBI///AACuS4Ek//8AAK5LgSX//wAArkuBCv//AACuS4EL//8AAK5LgSb//wAArkuBJ///AACuSw=="