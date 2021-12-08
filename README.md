## Installation
```
go install github.com/provenance-io/extkey/cmd/extkey@latest
```

# Encoding

## Key generation interactive
```
# Using interactive mode
 ▷▷ extkey
Mnemonic: fly fly comfort
Passphrase: 
HRP: tp
HDPath: m/44'/1'/0'/420'


RootKey Private: xprv9s21ZrQH143K2gAYPH8TS4wyWvQ4F99jm7cHbckUwKwvk75QpVx6VmqqZxgUno53xUWNbhnUy83RnPqSCa7hfKQFzmp1Lcsv2pp2PikLtHU
RootKey Public : xpub661MyMwAqRbcFAF1VJfToCti4xEYebsb8LXtQ1A6VfUucuQZN3GM3aAKREjZt3ZFhQB2M5Le1FEfuhVQRQ8DgvmPkLjgKPMAFo5X923Ut1B
Address: tp148sw6szwkk3cj5fdudzyhamkj3juy0wx0pj2c0

Path: m/44'/1'/0'/420'
ChildKey Private: xprv9zqjpMDofQuSFaF8NsS2Ybq2Xndj9zB5PkKyS16JqWsvP8aQWELPkpBnTh6NUUFHmRqRxVpz3fT8S2ckHSRSQ8EDcS4ZifxwQsjWgJjn5GK
ChildKey Public : xpub6Dq6DrkhVnTjU4KbUty2ujmm5pUDZStvkyFaEPVvPrQuFvuZ3meeJcWGJwjLjb666HDPxVg2SDTMuh6JVfP897z5VJxRoSf82koiPucLPDm
Address: tp1ndh7g7xy48k52phkr3p37rnkazmc98zuv8fp38
```

## Key generation with env vars
```
# Using env vars
$ MNEMONIC="fly fly comfort" PASSPHRASE="" HDPATH="m/44'/1'/0'/0'" HRP="tp" extkey

RootKey Private: xprv9s21ZrQH143K2gAYPH8TS4wyWvQ4F99jm7cHbckUwKwvk75QpVx6VmqqZxgUno53xUWNbhnUy83RnPqSCa7hfKQFzmp1Lcsv2pp2PikLtHU
RootKey Public : xpub661MyMwAqRbcFAF1VJfToCti4xEYebsb8LXtQ1A6VfUucuQZN3GM3aAKREjZt3ZFhQB2M5Le1FEfuhVQRQ8DgvmPkLjgKPMAFo5X923Ut1B
Address: tp148sw6szwkk3cj5fdudzyhamkj3juy0wx0pj2c0

Path: m/44'/1'/0'/0'
ChildKey Private: xprv9zqjpMDofQu7oDsEZEoP8sFrrvqewMi4s8ntBLthjaGfnwe8sCXdjgqYH5HQZXDeHbtsS3mdqdkaFAGVNJ6Xde48hkUsfbtUsJTQKyvBVJv
ChildKey Public : xpub6Dq6DrkhVnTR1hwhfGLPW1CbQxg9LpRvEMiUyjJKHuoefjyHQjqtHVA28NCY3YqM35fd1LyG5jkAaYZbHciHhDdJPxSux97o1nvBgvcww7o
Address: tp1wjccte6zcr0d9d8l5mjj5ju6rmlcywlt02tlpn
```

# Decoding

## Decoding xprv keys
```
 ▷▷ extkey xprv9zqjpMDofQuSFaF8NsS2Ybq2Xndj9zB5PkKyS16JqWsvP8aQWELPkpBnTh6NUUFHmRqRxVpz3fT8S2ckHSRSQ8EDcS4ZifxwQsjWgJjn5GK
HRP: tp

Key Private: xprv9zqjpMDofQuSFaF8NsS2Ybq2Xndj9zB5PkKyS16JqWsvP8aQWELPkpBnTh6NUUFHmRqRxVpz3fT8S2ckHSRSQ8EDcS4ZifxwQsjWgJjn5GK
Key Public : xpub6Dq6DrkhVnTjU4KbUty2ujmm5pUDZStvkyFaEPVvPrQuFvuZ3meeJcWGJwjLjb666HDPxVg2SDTMuh6JVfP897z5VJxRoSf82koiPucLPDm
ChainCode  : F285FC31610476271F3EB344992EE7735D830235059FD301B1B5787A87A7B68F
Fingerprint: 28508294
Depth      : 4 (m/44/coin/account/*change/index)
Address    : tp1ndh7g7xy48k52phkr3p37rnkazmc98zuv8fp38
```

## Decoding xpub keys
```
 ▷▷ extkey xpub6Dq6DrkhVnTjU4KbUty2ujmm5pUDZStvkyFaEPVvPrQuFvuZ3meeJcWGJwjLjb666HDPxVg2SDTMuh6JVfP897z5VJxRoSf82koiPucLPDm
HRP: tp

Key Private: xpub6Dq6DrkhVnTjU4KbUty2ujmm5pUDZStvkyFaEPVvPrQuFvuZ3meeJcWGJwjLjb666HDPxVg2SDTMuh6JVfP897z5VJxRoSf82koiPucLPDm
Key Public : xpub6Dq6DrkhVnTjU4KbUty2ujmm5pUDZStvkyFaEPVvPrQuFvuZ3meeJcWGJwjLjb666HDPxVg2SDTMuh6JVfP897z5VJxRoSf82koiPucLPDm
ChainCode  : F285FC31610476271F3EB344992EE7735D830235059FD301B1B5787A87A7B68F
Fingerprint: 28508294
Depth      : 4 (m/44/coin/account/*change/index)
Address    : tp1ndh7g7xy48k52phkr3p37rnkazmc98zuv8fp38
```
