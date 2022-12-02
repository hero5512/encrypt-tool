# encrypt-tool

## Introduce
* A command-line applications to encrypt or decrypt your important data
* Very easy to use, support interactive prompt, and as simple as possible
* Currently support aes and ecies crypto

## Install
Require go version 1.17+.

### Download source codes
```
$ git clone https://github.com/hero5512/encrypt-tool
```

### Install
```
$ cd encrypt-tool
$ go mod tidy -compat=1.17
$ go build
```

### Test it
```
$ encrypt-tool e hello
Use the arrow keys to navigate: ↓ ↑ → ←
? Select Algorithm:
    aes
  ▸ ecies
```

## Examples
Help command:
```
$ encrypt-tool h
NAME:
   encrypt-tool - crypto for important data!

USAGE:
   encrypt-tool [global options] command [command options] [arguments...]

VERSION:
   1.0.1

AUTHOR:
   Allen <lvshuaino@gmail.com>

COMMANDS:
   encrypt, e, en  encrypt a message
   decrypt, d, de  decrypt a message
   gen, g, gen     generate key pair
   help, h         Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version

```
generate ecdsa key for encrypting or decrypting data:
```
$ ./encrypt-tool gen
generate key pair public ===> private 
 bbb268e7dca5a372cc122a057dc50ea43f759f6795a3b48ffe3c5b6eb7222b0affe38181e6b2a620d4f455fc19a4a147204b657f323b8d2fe34c2176e00bed07 ===> 65ef30244f73e2397ee91dabcd7d69c7a1fb9b27359a0c8888fb9e7ac47a8be8

```
ECIES encrypt use interactive prompt:
```
$ ./encrypt-tool e hello  
✔ ecies
Public Key(64 bytes): ********************************************************************************************************************************
ecies encrypt hello ===> 0x042a7f41d58ff4f838d521b376644333da76134af188eb0d1b8047b8925f9e4bbe5c10b376e4d6b9d77df087e115abc24fc4e92b6a72d6e08e6033363db6bf77971d068fdc151bec65f3884b8bfc372ee28c241a79c1a820d7782d904b3a242729b7aa3e2b40081ef33309a0accfd1c04158bb5439a1

```
ECIES decrypt use interactive prompt:
```
$ ./encrypt-tool d 0x042a7f41d58ff4f838d521b376644333da76134af188eb0d1b8047b8925f9e4bbe5c10b376e4d6b9d77df087e115abc24fc4e92b6a72d6e08e6033363db6bf77971d068fdc151bec65f3884b8bfc372ee28c241a79c1a820d7782d904b3a242729b7aa3e2b40081ef33309a0accfd1c04158bb5439a1
✔ ecies
Private Key(32 bytes): ****************************************************************
ecies decrypt 0x042a7f41d58ff4f838d521b376644333da76134af188eb0d1b8047b8925f9e4bbe5c10b376e4d6b9d77df087e115abc24fc4e92b6a72d6e08e6033363db6bf77971d068fdc151bec65f3884b8bfc372ee28c241a79c1a820d7782d904b3a242729b7aa3e2b40081ef33309a0accfd1c04158bb5439a1 ===> hello

```
AES encrypt use interactive prompt:
```
$ ./encrypt-tool e hello
✔ aes
✔ AES Key(16 bytes): ****************
aes encrypt hello ===> AMCHTDF-q0jKdiXuNvoKcXYWvrx9HnDDLRriC6QVuFk
```
AES decrypt use interactive prompt:
```
$ ./encrypt-tool d WGM7Dteil6I1WDByRV3a2-yeXvD01YmXgIQUHYmOIhc
✔ aes
✔ AES Key(16 bytes): ****************
aes decrypt WGM7Dteil6I1WDByRV3a2-yeXvD01YmXgIQUHYmOIhc ===> hello
```
AES encrypt use aeskey and alg flag:
```
$ ./encrypt-tool e hello -aeskey 1234567890123456 --alg aes
aes encrypt hello ===> u1S3RUI5MTi2Yz_HBgkCIXhE25jAFUON1TeJuHAmUWc
```
AES decrypt use key flag:
```
$ ./encrypt-tool d u1S3RUI5MTi2Yz_HBgkCIXhE25jAFUON1TeJuHAmUWc -aeskey 1234567890123456 --alg aes
aes decrypt u1S3RUI5MTi2Yz_HBgkCIXhE25jAFUON1TeJuHAmUWc ===> hello
```
