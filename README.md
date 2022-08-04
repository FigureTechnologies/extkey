## Installation
```shell
$ go install github.com/provenance-io/extkey/cmd/extkey@latest
```

# Encoding

## Key encoding interactive
```shell
# Using interactive mode
$ extkey encode --hrp tp --hd-path "m/44'/1'/0'/420'"
mnemonic: fly fly comfort
passphrase: 
seed: UeItRbah4gE-syrw1EaXVKyg3GKWqfZqOztWNAXfnUME-bMpjp4jT0YxzEBWBA33_QWDAmeFE1I_hNlt2xJJew==
hrp: tp
rootKey:
    address: tp148sw6szwkk3cj5fdudzyhamkj3juy0wx0pj2c0
    privateKey:
        base58: xprv9s21ZrQH143K2gAYPH8TS4wyWvQ4F99jm7cHbckUwKwvk75QpVx6VmqqZxgUno53xUWNbhnUy83RnPqSCa7hfKQFzmp1Lcsv2pp2PikLtHU
        bytes: 0488ADE40000000000000000003E02C49CF3397E45141E4C9BED309CE9630A2DC6C4C10892A8A71447C9AF79FE0088AAA01CC129E4C16BFE49AF8659B92CCA5D964F0A048BB1633F4514AB14850B4B6C702F
        bigInt: "61816016462973053634971943848394651064681471574639878248202317244753465607435"
    publicKey:
        base58: xpub661MyMwAqRbcFAF1VJfToCti4xEYebsb8LXtQ1A6VfUucuQZN3GM3aAKREjZt3ZFhQB2M5Le1FEfuhVQRQ8DgvmPkLjgKPMAFo5X923Ut1B
        bytes: 0488B21E0000000000000000003E02C49CF3397E45141E4C9BED309CE9630A2DC6C4C10892A8A71447C9AF79FE030F5D004B8355C738762E54CCD22ECBA7798A2327982BA412FDBFAC34F99B767354878D2E
        bigInt: "354325279253549169190772409736366026367206282379023929601832819184328435791475"
childKey:
    address: tp1ndh7g7xy48k52phkr3p37rnkazmc98zuv8fp38
    privateKey:
        base58: xprv9zqjpMDofQuSFaF8NsS2Ybq2Xndj9zB5PkKyS16JqWsvP8aQWELPkpBnTh6NUUFHmRqRxVpz3fT8S2ckHSRSQ8EDcS4ZifxwQsjWgJjn5GK
        bytes: 0488ADE40428508294800001A4F285FC31610476271F3EB344992EE7735D830235059FD301B1B5787A87A7B68F007531EE53DE80DCAA03BABDA2BA9FD00CC8CDDEAE481A61A996CC6E8EB417E1B327182390
        bigInt: "53008823667154757289335472780676933506146822681876901480085094723772554404275"
    publicKey:
        base58: xpub6Dq6DrkhVnTjU4KbUty2ujmm5pUDZStvkyFaEPVvPrQuFvuZ3meeJcWGJwjLjb666HDPxVg2SDTMuh6JVfP897z5VJxRoSf82koiPucLPDm
        bytes: 0488B21E0428508294800001A4F285FC31610476271F3EB344992EE7735D830235059FD301B1B5787A87A7B68F02416EF4140348D3BF8C0B2B81E1B1C75CF6E2B5335ECE3152106719A3728779095C9382A4
        bigInt: "261180551375323774831564216383503265774120811521638640313796615741160982804745"
```

## Key encoding with env vars
```shell
# Using env vars
$ MNEMONIC="fly fly comfort" PASSPHRASE="" extkey encode --hrp tp --hd-path "m/44'/1'/0'/0'" 
seed: UeItRbah4gE-syrw1EaXVKyg3GKWqfZqOztWNAXfnUME-bMpjp4jT0YxzEBWBA33_QWDAmeFE1I_hNlt2xJJew==
hrp: tp
rootKey:
    address: tp148sw6szwkk3cj5fdudzyhamkj3juy0wx0pj2c0
    privateKey:
        base58: xprv9s21ZrQH143K2gAYPH8TS4wyWvQ4F99jm7cHbckUwKwvk75QpVx6VmqqZxgUno53xUWNbhnUy83RnPqSCa7hfKQFzmp1Lcsv2pp2PikLtHU
        bytes: 0488ADE40000000000000000003E02C49CF3397E45141E4C9BED309CE9630A2DC6C4C10892A8A71447C9AF79FE0088AAA01CC129E4C16BFE49AF8659B92CCA5D964F0A048BB1633F4514AB14850B4B6C702F
        bigInt: "61816016462973053634971943848394651064681471574639878248202317244753465607435"
    publicKey:
        base58: xpub661MyMwAqRbcFAF1VJfToCti4xEYebsb8LXtQ1A6VfUucuQZN3GM3aAKREjZt3ZFhQB2M5Le1FEfuhVQRQ8DgvmPkLjgKPMAFo5X923Ut1B
        bytes: 0488B21E0000000000000000003E02C49CF3397E45141E4C9BED309CE9630A2DC6C4C10892A8A71447C9AF79FE030F5D004B8355C738762E54CCD22ECBA7798A2327982BA412FDBFAC34F99B767354878D2E
        bigInt: "354325279253549169190772409736366026367206282379023929601832819184328435791475"
childKey:
    address: tp1wjccte6zcr0d9d8l5mjj5ju6rmlcywlt02tlpn
    privateKey:
        base58: xprv9zqjpMDofQu7oDsEZEoP8sFrrvqewMi4s8ntBLthjaGfnwe8sCXdjgqYH5HQZXDeHbtsS3mdqdkaFAGVNJ6Xde48hkUsfbtUsJTQKyvBVJv
        bytes: 0488ADE4042850829480000000FEB0AB4988EAE6538E60190EBF17678B0F2E590ED1D88D5D7CD8E50B63AC23EF00436F8E356F0E13B8007D668625ED6B61FC65B8286778452DE55F124DD934CB9FCCE76EDF
        bigInt: "30502062367823121628620232764791700029549832820376726752590340403318586919839"
    publicKey:
        base58: xpub6Dq6DrkhVnTR1hwhfGLPW1CbQxg9LpRvEMiUyjJKHuoefjyHQjqtHVA28NCY3YqM35fd1LyG5jkAaYZbHciHhDdJPxSux97o1nvBgvcww7o
        bytes: 0488B21E042850829480000000FEB0AB4988EAE6538E60190EBF17678B0F2E590ED1D88D5D7CD8E50B63AC23EF033BC1E08A7DAF3D04F704B43DD53FF9E4C4FDC9855B416881859B3808EFEDD34543BDE2A2
        bigInt: "374405276986753977850763764723578643711636570020048640824572441823815521915717"
```

# Decoding

## Decoding xprv keys
```shell
$ extkey decode --hrp tp xprv9zqjpMDofQuSFaF8NsS2Ybq2Xndj9zB5PkKyS16JqWsvP8aQWELPkpBnTh6NUUFHmRqRxVpz3fT8S2ckHSRSQ8EDcS4ZifxwQsjWgJjn5GK
address: tp1ndh7g7xy48k52phkr3p37rnkazmc98zuv8fp38
xkey:
    private: xprv9zqjpMDofQuSFaF8NsS2Ybq2Xndj9zB5PkKyS16JqWsvP8aQWELPkpBnTh6NUUFHmRqRxVpz3fT8S2ckHSRSQ8EDcS4ZifxwQsjWgJjn5GK
    public: xpub6Dq6DrkhVnTjU4KbUty2ujmm5pUDZStvkyFaEPVvPrQuFvuZ3meeJcWGJwjLjb666HDPxVg2SDTMuh6JVfP897z5VJxRoSf82koiPucLPDm
    depth: 4
    depthLoc: m/44/coin/account/*change/index
    chaincode: F285FC31610476271F3EB344992EE7735D830235059FD301B1B5787A87A7B68F
    fingerprint: "28508294"
```

## Decoding xpub keys
```shell
$ extkey decode --hrp tp xpub6Dq6DrkhVnTjU4KbUty2ujmm5pUDZStvkyFaEPVvPrQuFvuZ3meeJcWGJwjLjb666HDPxVg2SDTMuh6JVfP897z5VJxRoSf82koiPucLPDm
address: tp1ndh7g7xy48k52phkr3p37rnkazmc98zuv8fp38
xkey:
    public: xpub6Dq6DrkhVnTjU4KbUty2ujmm5pUDZStvkyFaEPVvPrQuFvuZ3meeJcWGJwjLjb666HDPxVg2SDTMuh6JVfP897z5VJxRoSf82koiPucLPDm
    depth: 4
    depthLoc: m/44/coin/account/*change/index
    chaincode: F285FC31610476271F3EB344992EE7735D830235059FD301B1B5787A87A7B68F
    fingerprint: "28508294"
```

# Generating

## Key generation with random seed
```shell
$ extkey gen --hd-path "m/44'/505'/0'/0'/0" --hrp tp
seed: kwDcbPMhZOTRQpCcgtBKXxbtVcEe_4ne5tasLV-JAhlYAbaf13vRMaVANwDZTm4A75snNhI2LTH_T0NVGTYJuQ==
mnemonic: text gown action resource depth toe parrot limit dinosaur frog again march pool another pretty imitate spatial music lecture charge very fiber few ignore
hrp: tp
rootKey:
    address: tp1scl32wrzy043ldnz5qyaqn5kpjws2fktce4k7s
    privateKey:
        base58: xprv9s21ZrQH143K3hoYnZzA9WRAkGK7pJbjDDKf2MjstzMLyH9tUUgRtDWaPCULb6rsnijuqs3FDoXoGsZuzx8it8ZkKuyuqKZH9EghnUBBGtV
        bytes: 0488ADE4000000000000000000A54A051BCE830BF6C6A3656C5CA2A4E45D11B44E853FF5019D8DB8F57328B00F00712F6D5E492D9516A0D68E6F1A34068206D93C3E8193D2C2DCA6F9CC3636D3773700F716
        bigInt: "51195148534247021627172126139454315055974072071557156386597256238996714345335"
    publicKey:
        base58: xpub661MyMwAqRbcGBt1tbXAWeMuJJ9cDmKaaSFFpk9VTKtKr5V321zgS1q4EWQUMT7QN4xaXzkWMCzQAQ6Vp6DGtp5MyMQ7cDToY6tzNH91X57
        bytes: 0488B21E000000000000000000A54A051BCE830BF6C6A3656C5CA2A4E45D11B44E853FF5019D8DB8F57328B00F03EF7C0C7C458B18A3D5ADEF1EA181271F4DE1191554F0B4BFF95923703FA02DC43C4BC866
        bigInt: "455698213730695089965667949209897413589214143496016811921245211286232805420484"
childKey:
    address: tp1wy5vf0cem2ec88qvnj5qaeucyf3twrhhfzs7qe
    privateKey:
        base58: xprvA2vNQscVB7uwrpXTLYXw5ffoRxHrp75Ynwc3LVixemDticLn7BpGDnxXvc2jJEksJuef3bLP3QtKZLYRUPncbBoaBW5SMSjMENCLppnT2te
        bytes: 0488ADE40543478CF000000000E63E5C24103358AEB259B1E0926EDED31D6DCF316F16D6A74BB4AAF9ED5D1D610001628A7298303FB55D71174CCF7C8B20D556FCF97A815456B0301A14049BD9E64A4C9357
        bigInt: "626419391388641969500024507723796044901451795080364889081030583799756085734"
    publicKey:
        base58: xpub6FuipP9P1VUF5JbvSa4wSocXyz8MDZoQAAXe8t8aD6ksbQfvej8WmbH1mtypTMgVqxfkPpY94M122qqNhdWPLm1cUq4mDQW9VZeLwbZ2X6C
        bytes: 0488B21E0543478CF000000000E63E5C24103358AEB259B1E0926EDED31D6DCF316F16D6A74BB4AAF9ED5D1D6102FE281C136209674387FE584A4A18584D0BB82C2C7B5FF818AE25C713CB3E85790139E4CD
        bigInt: "346542509668834358586089346602711884922951328399776737952015035914050934310265"
```

## Key generation with known seed
```shell
$ extkey gen --hd-path "m/44'/505'/0'/0'/0" --hrp tp --seed kwDcbPMhZOTRQpCcgtBKXxbtVcEe_4ne5tasLV-JAhlYAbaf13vRMaVANwDZTm4A75snNhI2LTH_T0NVGTYJuQ==
seed: kwDcbPMhZOTRQpCcgtBKXxbtVcEe_4ne5tasLV-JAhlYAbaf13vRMaVANwDZTm4A75snNhI2LTH_T0NVGTYJuQ==
hrp: tp
rootKey:
    address: tp1scl32wrzy043ldnz5qyaqn5kpjws2fktce4k7s
    privateKey:
        base58: xprv9s21ZrQH143K3hoYnZzA9WRAkGK7pJbjDDKf2MjstzMLyH9tUUgRtDWaPCULb6rsnijuqs3FDoXoGsZuzx8it8ZkKuyuqKZH9EghnUBBGtV
        bytes: 0488ADE4000000000000000000A54A051BCE830BF6C6A3656C5CA2A4E45D11B44E853FF5019D8DB8F57328B00F00712F6D5E492D9516A0D68E6F1A34068206D93C3E8193D2C2DCA6F9CC3636D3773700F716
        bigInt: "51195148534247021627172126139454315055974072071557156386597256238996714345335"
    publicKey:
        base58: xpub661MyMwAqRbcGBt1tbXAWeMuJJ9cDmKaaSFFpk9VTKtKr5V321zgS1q4EWQUMT7QN4xaXzkWMCzQAQ6Vp6DGtp5MyMQ7cDToY6tzNH91X57
        bytes: 0488B21E000000000000000000A54A051BCE830BF6C6A3656C5CA2A4E45D11B44E853FF5019D8DB8F57328B00F03EF7C0C7C458B18A3D5ADEF1EA181271F4DE1191554F0B4BFF95923703FA02DC43C4BC866
        bigInt: "455698213730695089965667949209897413589214143496016811921245211286232805420484"
childKey:
    address: tp1wy5vf0cem2ec88qvnj5qaeucyf3twrhhfzs7qe
    privateKey:
        base58: xprvA2vNQscVB7uwrpXTLYXw5ffoRxHrp75Ynwc3LVixemDticLn7BpGDnxXvc2jJEksJuef3bLP3QtKZLYRUPncbBoaBW5SMSjMENCLppnT2te
        bytes: 0488ADE40543478CF000000000E63E5C24103358AEB259B1E0926EDED31D6DCF316F16D6A74BB4AAF9ED5D1D610001628A7298303FB55D71174CCF7C8B20D556FCF97A815456B0301A14049BD9E64A4C9357
        bigInt: "626419391388641969500024507723796044901451795080364889081030583799756085734"
    publicKey:
        base58: xpub6FuipP9P1VUF5JbvSa4wSocXyz8MDZoQAAXe8t8aD6ksbQfvej8WmbH1mtypTMgVqxfkPpY94M122qqNhdWPLm1cUq4mDQW9VZeLwbZ2X6C
        bytes: 0488B21E0543478CF000000000E63E5C24103358AEB259B1E0926EDED31D6DCF316F16D6A74BB4AAF9ED5D1D6102FE281C136209674387FE584A4A18584D0BB82C2C7B5FF818AE25C713CB3E85790139E4CD
        bigInt: "346542509668834358586089346602711884922951328399776737952015035914050934310265"
```

## Key generation with known mnemonic
```shell
TODO
```

# API key generation
```
$ extkey serve
...
[GIN-debug] GET    /generate                 --> github.com/provenance-io/extkey/cmd/extkey/commands.generateKeys (3 handlers)
...
```

```
$ curl "localhost:9000/generate?hrp=tp" | jq
{
  "seed": "UeItRbah4gE-syrw1EaXVKyg3GKWqfZqOztWNAXfnUME-bMpjp4jT0YxzEBWBA33_QWDAmeFE1I_hNlt2xJJew==",
  "hrp": "tp",
  "rootKey": {
    "address": "tp148sw6szwkk3cj5fdudzyhamkj3juy0wx0pj2c0",
    "privateKey": {
      "base58": "xprv9s21ZrQH143K2gAYPH8TS4wyWvQ4F99jm7cHbckUwKwvk75QpVx6VmqqZxgUno53xUWNbhnUy83RnPqSCa7hfKQFzmp1Lcsv2pp2PikLtHU",
      "bytes": "0488ADE40000000000000000003E02C49CF3397E45141E4C9BED309CE9630A2DC6C4C10892A8A71447C9AF79FE0088AAA01CC129E4C16BFE49AF8659B92CCA5D964F0A048BB1633F4514AB14850B4B6C702F",
      "bigInt": "61816016462973053634971943848394651064681471574639878248202317244753465607435"
    },
    "publicKey": {
      "base58": "xpub661MyMwAqRbcFAF1VJfToCti4xEYebsb8LXtQ1A6VfUucuQZN3GM3aAKREjZt3ZFhQB2M5Le1FEfuhVQRQ8DgvmPkLjgKPMAFo5X923Ut1B",
      "bytes": "0488B21E0000000000000000003E02C49CF3397E45141E4C9BED309CE9630A2DC6C4C10892A8A71447C9AF79FE030F5D004B8355C738762E54CCD22ECBA7798A2327982BA412FDBFAC34F99B767354878D2E",
      "bigInt": "354325279253549169190772409736366026367206282379023929601832819184328435791475"
    }
  }
}

$ curl "localhost:9000/generate?hrp=tp&seed=UeItRbah4gE-syrw1EaXVKyg3GKWqfZqOztWNAXfnUME-bMpjp4jT0YxzEBWBA33_QWDAmeFE1I_hNlt2xJJew==" | jq
{
  "seed": "UeItRbah4gE-syrw1EaXVKyg3GKWqfZqOztWNAXfnUME-bMpjp4jT0YxzEBWBA33_QWDAmeFE1I_hNlt2xJJew==",
  "hrp": "tp",
  "rootKey": {
    "address": "tp148sw6szwkk3cj5fdudzyhamkj3juy0wx0pj2c0",
    "privateKey": {
      "base58": "xprv9s21ZrQH143K2gAYPH8TS4wyWvQ4F99jm7cHbckUwKwvk75QpVx6VmqqZxgUno53xUWNbhnUy83RnPqSCa7hfKQFzmp1Lcsv2pp2PikLtHU",
      "bytes": "0488ADE40000000000000000003E02C49CF3397E45141E4C9BED309CE9630A2DC6C4C10892A8A71447C9AF79FE0088AAA01CC129E4C16BFE49AF8659B92CCA5D964F0A048BB1633F4514AB14850B4B6C702F",
      "bigInt": "61816016462973053634971943848394651064681471574639878248202317244753465607435"
    },
    "publicKey": {
      "base58": "xpub661MyMwAqRbcFAF1VJfToCti4xEYebsb8LXtQ1A6VfUucuQZN3GM3aAKREjZt3ZFhQB2M5Le1FEfuhVQRQ8DgvmPkLjgKPMAFo5X923Ut1B",
      "bytes": "0488B21E0000000000000000003E02C49CF3397E45141E4C9BED309CE9630A2DC6C4C10892A8A71447C9AF79FE030F5D004B8355C738762E54CCD22ECBA7798A2327982BA412FDBFAC34F99B767354878D2E",
      "bigInt": "354325279253549169190772409736366026367206282379023929601832819184328435791475"
    }
  }
}
```

# Running via docker
```
# Default listen port is 7000
$ docker run -p 7000:7000 -it provenanceio/extkey

# To specify listen port of 9000
$ docker run -p 9000:9000 -it provenanceio/extkey serve --laddr :9000

# Enable extra gin debug logs
$ docker run -p 7000:7000 -e GIN_MODE=debug -it provenanceio/extkey 
```