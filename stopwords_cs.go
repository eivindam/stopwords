// Copyright (c) 2005, Jacques Savoy.
// Use of this source code is governed by the BSD license
// license that can be found in the LICENSE file.

// stopwords package removes most frequent words from a text content
package stopwords

var czech = map[string]string{
	"ačkoli":     "",
	"ahoj":       "",
	"ale":        "",
	"anebo":      "",
	"ano":        "",
	"asi":        "",
	"aspoň":      "",
	"během":      "",
	"bez":        "",
	"beze":       "",
	"blízko":     "",
	"bohužel":    "",
	"brzo":       "",
	"bude":       "",
	"budeme":     "",
	"budeš":      "",
	"budete":     "",
	"budou":      "",
	"budu":       "",
	"byl":        "",
	"byla":       "",
	"byli":       "",
	"bylo":       "",
	"byly":       "",
	"bys":        "",
	"čau":        "",
	"chce":       "",
	"chceme":     "",
	"chceš":      "",
	"chcete":     "",
	"chci":       "",
	"chtějí":     "",
	"chtít":      "",
	"chut'":      "",
	"chuti":      "",
	"co":         "",
	"čtrnáct":    "",
	"čtyři":      "",
	"dál":        "",
	"dále":       "",
	"daleko":     "",
	"děkovat":    "",
	"děkujeme":   "",
	"děkuji":     "",
	"den":        "",
	"deset":      "",
	"devatenáct": "",
	"devět":      "",
	"do":         "",
	"dobrý":      "",
	"docela":     "",
	"dva":        "",
	"dvacet":     "",
	"dvanáct":    "",
	"dvě":        "",
	"hodně":      "",
	"já":         "",
	"jak":        "",
	"jde":        "",
	"je":         "",
	"jeden":      "",
	"jedenáct":   "",
	"jedna":      "",
	"jedno":      "",
	"jednou":     "",
	"jedou":      "",
	"jeho":       "",
	"její":       "",
	"jejich":     "",
	"jemu":       "",
	"jen":        "",
	"jenom":      "",
	"ještě":      "",
	"jestli":     "",
	"jestliže":   "",
	"jí":         "",
	"jich":       "",
	"jím":        "",
	"jimi":       "",
	"jinak":      "",
	"jsem":       "",
	"jsi":        "",
	"jsme":       "",
	"jsou":       "",
	"jste":       "",
	"kam":        "",
	"kde":        "",
	"kdo":        "",
	"kdy":        "",
	"když":       "",
	"ke":         "",
	"kolik":      "",
	"kromě":      "",
	"která":      "",
	"které":      "",
	"kteří":      "",
	"který":      "",
	"kvůli":      "",
	"má":         "",
	"mají":       "",
	"málo":       "",
	"mám":        "",
	"máme":       "",
	"máš":        "",
	"máte":       "",
	"mé":         "",
	"mě":         "",
	"mezi":       "",
	"mí":         "",
	"mít":        "",
	"mně":        "",
	"mnou":       "",
	"moc":        "",
	"mohl":       "",
	"mohou":      "",
	"moje":       "",
	"moji":       "",
	"možná":      "",
	"můj":        "",
	"musí":       "",
	"může":       "",
	"my":         "",
	"na":         "",
	"nad":        "",
	"nade":       "",
	"nám":        "",
	"námi":       "",
	"naproti":    "",
	"nás":        "",
	"náš":        "",
	"naše":       "",
	"naši":       "",
	"ne":         "",
	"ně":         "",
	"nebo":       "",
	"nebyl":      "",
	"nebyla":     "",
	"nebyli":     "",
	"nebyly":     "",
	"něco":       "",
	"nedělá":     "",
	"nedělají":   "",
	"nedělám":    "",
	"neděláme":   "",
	"neděláš":    "",
	"neděláte":   "",
	"nějak":      "",
	"nejsi":      "",
	"někde":      "",
	"někdo":      "",
	"nemají":     "",
	"nemáme":     "",
	"nemáte":     "",
	"neměl":      "",
	"němu":       "",
	"není":       "",
	"nestačí":    "",
	"nevadí":     "",
	"než":        "",
	"nic":        "",
	"nich":       "",
	"ním":        "",
	"nimi":       "",
	"nula":       "",
	"od":         "",
	"ode":        "",
	"on":         "",
	"ona":        "",
	"oni":        "",
	"ono":        "",
	"ony":        "",
	"osm":        "",
	"osmnáct":    "",
	"pak":        "",
	"patnáct":    "",
	"pět":        "",
	"po":         "",
	"pořád":      "",
	"potom":      "",
	"pozdě":      "",
	"před":       "",
	"přes":       "",
	"přese":      "",
	"pro":        "",
	"proč":       "",
	"prosím":     "",
	"prostě":     "",
	"proti":      "",
	"protože":    "",
	"rovně":      "",
	"se":         "",
	"sedm":       "",
	"sedmnáct":   "",
	"šest":       "",
	"šestnáct":   "",
	"skoro":      "",
	"smějí":      "",
	"smí":        "",
	"snad":       "",
	"spolu":      "",
	"sta":        "",
	"sté":        "",
	"sto":        "",
	"ta":         "",
	"tady":       "",
	"tak":        "",
	"takhle":     "",
	"taky":       "",
	"tam":        "",
	"tamhle":     "",
	"tamhleto":   "",
	"tamto":      "",
	"tě":         "",
	"tebe":       "",
	"tebou":      "",
	"ted'":       "",
	"tedy":       "",
	"ten":        "",
	"ti":         "",
	"tisíc":      "",
	"tisíce":     "",
	"to":         "",
	"tobě":       "",
	"tohle":      "",
	"toto":       "",
	"třeba":      "",
	"tři":        "",
	"třináct":    "",
	"trošku":     "",
	"tvá":        "",
	"tvé":        "",
	"tvoje":      "",
	"tvůj":       "",
	"ty":         "",
	"určitě":     "",
	"už":         "",
	"vám":        "",
	"vámi":       "",
	"vás":        "",
	"váš":        "",
	"vaše":       "",
	"vaši":       "",
	"ve":         "",
	"večer":      "",
	"vedle":      "",
	"vlastně":    "",
	"všechno":    "",
	"všichni":    "",
	"vůbec":      "",
	"vy":         "",
	"vždy":       "",
	"za":         "",
	"zač":        "",
	"zatímco":    "",
	"ze":         "",
	"že":         ""}
