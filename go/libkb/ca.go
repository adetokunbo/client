// Copyright 2015 Keybase, Inc. All rights reserved. Use of
// this source code is governed by the included BSD license.

package libkb

var BundledCAs = map[string]string{
	"api.keybase.io":          apiCA,
	"gregord.kbfs.keybase.io": kbfsCA,
	"gregord.dev.keybase.io":  kbfsDevCA,
}

const apiCA = `-----BEGIN CERTIFICATE-----
MIIGmzCCBIOgAwIBAgIJAPzhpcIBaOeNMA0GCSqGSIb3DQEBBQUAMIGPMQswCQYD
VQQGEwJVUzELMAkGA1UECBMCTlkxETAPBgNVBAcTCE5ldyBZb3JrMRQwEgYDVQQK
EwtLZXliYXNlIExMQzEXMBUGA1UECxMOQ2VydCBBdXRob3JpdHkxEzARBgNVBAMT
CmtleWJhc2UuaW8xHDAaBgkqhkiG9w0BCQEWDWNhQGtleWJhc2UuaW8wHhcNMTQw
MTAyMTY0MjMzWhcNMjMxMjMxMTY0MjMzWjCBjzELMAkGA1UEBhMCVVMxCzAJBgNV
BAgTAk5ZMREwDwYDVQQHEwhOZXcgWW9yazEUMBIGA1UEChMLS2V5YmFzZSBMTEMx
FzAVBgNVBAsTDkNlcnQgQXV0aG9yaXR5MRMwEQYDVQQDEwprZXliYXNlLmlvMRww
GgYJKoZIhvcNAQkBFg1jYUBrZXliYXNlLmlvMIICIjANBgkqhkiG9w0BAQEFAAOC
Ag8AMIICCgKCAgEA3sLA6ZG8uOvmlFvFLVIOURmcQrZyMFKbVu9/TeDiemls3w3/
JzVTduD+7KiUi9R7QcCW/V1ZpReTfunm7rfACiJ1fpIkjSQrgsvKDLghIzxIS5FM
I8utet5p6QtuJhaAwmmXn8xX05FvqWNbrcXRdpL4goFdigPsFK2xhTUiWatLMste
oShI7+zmrgkx75LeLMD0bL2uOf87JjOzbY8x2sUIZLGwPoATyG8WS38ey6KkJxRj
AhG3p+OTYEjYSrsAtQA6ImbeDpfSHKOB8HF3nVp//Eb4HEiEsWwBRbQXvAWh3DYL
GukFW0wiO0HVCoWY+bHL/Mqa0NdRGOlLsbL4Z4pLrhqKgSDU8umX9YuNRRaB0P5n
TkzyU6axHqzq990Gep/I62bjsBdYYp+DjSPK43mXRrfWJl2NTcl8xKAyfsOW+9hQ
9vwK0tpSicNxfYuUZs0BhfjSZ/Tc6Z1ERdgUYRiXTtohl+SRA2IgZMloHCllVMNj
EjXhguvHgLAOrcuyhVBupiUQGUHQvkMsr1Uz8VPNDFOJedwucRU2AaR881bknnSb
ds9+zNLsvUFV+BK7Qdnt/WkFpYL78rGwY47msi9Ooddx6fPyeg3qkJGM6cwn/boy
w9lQeleYDq8kyJdixIAxtAskNzRPJ4nDu2izTfByQoM8epwAWboc/gNFObMCAwEA
AaOB9zCB9DAdBgNVHQ4EFgQURqpATOw1gVVrzlqqFKbkfaKXvwowgcQGA1UdIwSB
vDCBuYAURqpATOw1gVVrzlqqFKbkfaKXvwqhgZWkgZIwgY8xCzAJBgNVBAYTAlVT
MQswCQYDVQQIEwJOWTERMA8GA1UEBxMITmV3IFlvcmsxFDASBgNVBAoTC0tleWJh
c2UgTExDMRcwFQYDVQQLEw5DZXJ0IEF1dGhvcml0eTETMBEGA1UEAxMKa2V5YmFz
ZS5pbzEcMBoGCSqGSIb3DQEJARYNY2FAa2V5YmFzZS5pb4IJAPzhpcIBaOeNMAwG
A1UdEwQFMAMBAf8wDQYJKoZIhvcNAQEFBQADggIBAA3Z5FIhulYghMuHdcHYTYWc
7xT5WD4hXQ0WALZs4p5Y+b2Af54o6v1wUE1Au97FORq5CsFXX/kGl/JzzTimeucn
YJwGuXMpilrlHCBAL5/lSQjA7qbYIolQ3SB9ON+LYuF1jKB9k8SqNp7qzucxT3tO
b8ZMDEPNsseC7NE2uwNtcW3yrTh6WZnSqg/jwswiWjHYDdG7U8FjMYlRol3wPux2
PizGbSgiR+ztI2OthxtxNWMrT9XKxNQTpcxOXnLuhiSwqH8PoY17ecP8VPpaa0K6
zym0zSkbroqydazaxcXRk3eSlc02Ktk7HzRzuqQQXhRMkxVnHbFHgGsz03L533pm
mlIEgBMggZkHwNvs1LR7f3v2McdKulDH7Mv8yyfguuQ5Jxxt7RJhUuqSudbEhoaM
6jAJwBkMFxsV2YnyFEd3eZ/qBYPf7TYHhyzmHW6WkSypGqSnXd4gYpJ8o7LxSf4F
inLjxRD+H9Xn1UVXWLM0gaBB7zZcXd2zjMpRsWgezf5IR5vyakJsc7fxzgor3Qeq
Ri6LvdEkhhFVl5rHMQBwNOPngySrq8cs/ikTLTfQVTYXXA4Ba1YyiMOlfaR1LhKw
If1AkUV0tfCTNRZ01EotKSK77+o+k214n+BAu+7mO+9B5Kb7lMFQcuWCHXKYB2Md
cT7Yh09F0QpFUd0ymEfv
-----END CERTIFICATE-----`

const kbfsCA = `-----BEGIN CERTIFICATE-----
MIIDkDCCAnigAwIBAgIRAL1MQ3C37AuGO8gFqfhqb9EwDQYJKoZIhvcNAQELBQAw
UTESMBAGCgmSJomT8ixkARkWAmlvMRwwGgYKCZImiZPyLGQBGRYMa2Jmcy5rZXli
YXNlMR0wGwYDVQQDDBRLZXliYXNlIEtCRlMgQ0EgcHJvZDAeFw0xNTExMDgyMzM3
MDFaFw0xNzExMDcyMzM3MDFaMFExEjAQBgoJkiaJk/IsZAEZFgJpbzEcMBoGCgmS
JomT8ixkARkWDGtiZnMua2V5YmFzZTEdMBsGA1UEAwwUS2V5YmFzZSBLQkZTIENB
IHByb2QwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCr9ttTzL093jPt
WstzWR19qvLprd778ALqShZYZughuXPULgOck4AQW27vlp1nY8+7sBnWgstzL6Gv
dTQU61e34yOeAFYyKoWPFHyeo/g1y+LANgLdLbeOatOlWyM2sb/f0K3SKpusp/9J
0ylpDyko97MAI28spwX1d7L/qlDV6ryce4GrzElp3J8j3TZ3cju5rEldn8BSnLYw
i/2/Sc93GwhkjI03MZvuWaJQXQjTMALVzx5gFzshUymV4yrJfQbmBTwODf1yucsQ
NrWDiKWcFXe5dR8BWBZG7lslZeGYaHQ6lc3TgGwaPobpaZpzVEt3Crb9HAuTVl8/
Ynlw2XvzAgMBAAGjYzBhMA8GA1UdEwEB/wQFMAMBAf8wDgYDVR0PAQH/BAQDAgEG
MB0GA1UdDgQWBBSg9AYko8IqCwg2awZOZO6TW+ITsjAfBgNVHSMEGDAWgBSg9AYk
o8IqCwg2awZOZO6TW+ITsjANBgkqhkiG9w0BAQsFAAOCAQEAAJ2oOlY+DCDWr73m
TrR3Kfx+bDzvU1IZviKKooGGPjG+apcz5rWoKhjkO593ORCrygAvITnAI4v2Eaic
h2zYfWkOnCI2YYvVChR0TSJfa2+gxZFUqxRb68zMgcTxGZTZUonEX4nCJjkrSx3M
ATZkFWJDPPVci6o87VbpnKOc3mep3i1s3Cvw0GMHP+yVgw8Y0BpXII5hGbCODmoh
d2mdg2gjlOVBCfTEAe7cgUx9/lraQwUurUjDO3g54NZo/pcoc9koIW+Ai+saF5gA
UnFkqAOuEw0y4Fxzr9pw9naKF3KMlEJf6CiDJ4xspNzPZFupuepKitRrlrzofYuW
OXgZAw==
-----END CERTIFICATE-----
`
const kbfsDevCA = `
-----BEGIN CERTIFICATE-----
MIIDjDCCAnSgAwIBAgIRANRMeoRz3Xg5c8kT2f5k9wMwDQYJKoZIhvcNAQELBQAw
TzESMBAGCgmSJomT8ixkARkWAmlvMRswGQYKCZImiZPyLGQBGRYLZGV2LmtleWJh
c2UxHDAaBgNVBAMME0tleWJhc2UgS0JGUyBDQSBkZXYwHhcNMTUwOTIzMTk0NzM3
WhcNMTcwOTIyMTk0NzM3WjBPMRIwEAYKCZImiZPyLGQBGRYCaW8xGzAZBgoJkiaJ
k/IsZAEZFgtkZXYua2V5YmFzZTEcMBoGA1UEAwwTS2V5YmFzZSBLQkZTIENBIGRl
djCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBALWCTpo8NNeGM6nIkNW+
4qiM8ocuFjDw2Er6XJhncgj7xjTGf/9yZqnXeGHyHGT66AtKl5bc8son4+npWmvs
47OXORF7YGi89d9KBlIC4NCetZLBSVWiSG+XXSKrmIffi6D0UojpZc2blnzgejEO
ii1uCDSaj6TRLcC8z/eXKq+DtPcfNnPL0pu5CiUNrH1cA9PS+jO1OonCGPG5yVjW
bBw0nQfThhapm9IohtdbYzlQiSbE1+3ctNwCPLas3mmUWkcrrVbn1Fa54LnfNR2u
pnZRNZ7czfB/vtymUJ6/y8dLYTmnzMFFYy416FOmvr4NqLBkaMWg9xp+KeR30044
AicCAwEAAaNjMGEwDwYDVR0TAQH/BAUwAwEB/zAOBgNVHQ8BAf8EBAMCAQYwHQYD
VR0OBBYEFBdb6+h+Qq5vXUWo99QbatQTX6u+MB8GA1UdIwQYMBaAFBdb6+h+Qq5v
XUWo99QbatQTX6u+MA0GCSqGSIb3DQEBCwUAA4IBAQArhp0KXfJHEhVcUXqYYjdn
pZQjq3+0aKjMjgnVWekxwwBARh4ycy2e7066ru1eDZr6myGYK+/vjXituWtq7/c/
Fifezgje6o9lB1TPamgQeE8slqqAgc3OxTqbAAf+rxJelcI6aOm7tqX04k8Aiuhm
dr64cM/NsZTKUbrCHCVNHPNj8wWkrb9pbXH/q0+Gt/gw4MiL6p1YuSr4SIENqDpP
VFiOCcbOSiw5OHPe/VwLts/g3e3NSXqd53nQW1/CgpSBdT73oWw+SBfv21KuJN5K
745S8d9JfbLItWgM73o94MSLOpUEl2F7qqXj2eOBEYWIMbRjMMZ7Vzmuo5wo3M8i
-----END CERTIFICATE-----`
