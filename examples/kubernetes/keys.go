package main

// TODO: these should be generated on the fly and not hardcoded here.

const mkePubKeyPEM = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxiUit96Jn5NnHGYyIJ2y
K9n8/ZSipzva0CwTw9LYILy9V5jFPlzd4/5MEkB99CGqT2ZEKGuUxffSWu09SKCP
aSEEytOl55c/6n5AhxdvSqsAxVm3ujTeJGdlSuX8miB5U1p0TJk/79GvUpfsx/6o
ZmEI5OCSz1cZiBI+EkPA5dxVl4Dxdq8hwQr8lROhOucEu0MKgykcPzzznxmn6RgP
KIyt1yRt3RoPw2Ay6ximWqLBgKanXNYHbZi/kFCB7tiurhZarNqnvFco6ukkQHZM
2ramotu+6F1ax1JwvvMro8mIkv3L30GZYy/iiY/NGC1nNvs7ivmxA2YkHGVhKNpm
jwIDAQAB
-----END PUBLIC KEY-----`

const mkePrivKeyPEM = `-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDGJSK33omfk2cc
ZjIgnbIr2fz9lKKnO9rQLBPD0tggvL1XmMU+XN3j/kwSQH30IapPZkQoa5TF99Ja
7T1IoI9pIQTK06Xnlz/qfkCHF29KqwDFWbe6NN4kZ2VK5fyaIHlTWnRMmT/v0a9S
l+zH/qhmYQjk4JLPVxmIEj4SQ8Dl3FWXgPF2ryHBCvyVE6E65wS7QwqDKRw/PPOf
GafpGA8ojK3XJG3dGg/DYDLrGKZaosGApqdc1gdtmL+QUIHu2K6uFlqs2qe8Vyjq
6SRAdkzatqai277oXVrHUnC+8yujyYiS/cvfQZljL+KJj80YLWc2+zuK+bEDZiQc
ZWEo2maPAgMBAAECggEAbV5mEXOGJ0ynZ2eJV7Nzevzxxc5bEU/acXlC153zqe5D
3Kv4g3mjCIt+1bw/SZejd/wMtvravYMS2xmD8Cifv/l3ubSlKHW10+soFuj6hw5o
eIYTAXbY/uMLaBIWIcS28ylEdFtKjUcCLOaqzaFktk3tKNhEcKm08ASRumcnfDfu
/zAo/E/jg0/q7DMEU34KCw4FukLhcHN/Kx81NssOmXsA85vTxkl5W3vZDJVGu9os
MXp+PzdqiJeOv0oxV9TXnUqt36qmkn721mPBtpD3wzId29f/R361pd45LihhyG1/
3K8fU+EHY7GMaTBZiQ2H7ABXPhHnuYv1dViqKv1FAQKBgQDu0o5TohfqP8DtRcZx
jMOV85EKoCPDW7h3aduxYK6YKaI8kbG6AoFo0on1Pq/Etqy48+8NEi+MeDz7oxph
bjVyoIHSda3Bqh4XWwRpIXyOMdfCN1UD1AeIaGCCtVmeGg2J9uHAnh4LfD6qra+8
nGT5m46RXG/SL6LrzKbPkLa1QQKBgQDUZZXWSWG7FaOxuVQBMgyiXOsiCbaOGP+G
52S2ujtBxgxv1mCpKr2RWXQNWWxwWGpvHFTBIbhJnkFYR+cETtQ9LzUkLIcrh6yu
gQ5bSIMejPIwy1WfjcoH1wNCxg+Ymb9VM5dAuHBYCrAqok5qoiV36hbphGbQGNI+
hAq4PJUXzwKBgBQ8UQTzGhWmG0G/hMYASF4/Sg3dbcpSzjmIBa5s475O4MlDCw0Y
w0BPfpMCIcCIPfBZ8upnnRHI6lnkAws4XFz/DqD3iaZ8NJqEAsapqLUfsglpyNFP
OOgs9+h7V0GXMYh8G7rHawJMH780gkx37/JaZOUaMPtdP++84nF58JFBAoGAYUjO
YsJl001MFyFuCsYj51JsGMEeLuPgVqgyB0gx0CSomak1yQZcofC5KUwmUScOSSpO
389UG9Qy1f7JH74DlKrEynbiyttwCtXt/32tIRcp75OS+mv1d7XSNomFpz3010+0
Hd7dBlaO7lg9VrZCNzSvnfIZLYPQ+PV9+3k0SR0CgYEAmZYu8u5rvnjq0brcjXKH
pNvT3L0vaeYjDz4HtUeVo2URGUbYs426tMlxPiAUF3+8kdjpup+ZDfTBiZkLn3zs
mS5NGVjf85kCsOQ0gfZuVcUth3RRF6cTtrTDAMr7U8IIgXk3Rqhmn3k8UTjmLH9w
A9R9GFeLpUbT+M0IzlP4Iyc=
-----END PRIVATE KEY-----`

const k8sPubKeyPEM = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAtTy8tmVhDO1uDs1X0jXf
lErkUXYmWmS8HoZ/QnfZno6a5CgjH1Tx0Qe54sYampdEpyLj3FPETF/s2rBlj4nX
1xEWugjTfaNvQMhiGXsdJN/MQwfiYAevkz2GRjAKClVElUgDTW95PxNXyLry4+gC
Y2xmb0gecKwBcvOqpTKY99SvEFYsyKP/jLZlwoZZuMVl9xqN+w1zUAb7kfsCWLt3
dTZ5/qVYU62qCchkGP9Dcz8Qb/rgsZR1O+5rfUiePrwS1xmqOsw1HELGh9Ry9+8C
TY/OOGEYveh9wsEtgx4TiMkEyTM3qJ6xYn2Mmps2xYTPl1900lvnt6w3+0TVeL4C
8wIDAQAB
-----END PUBLIC KEY-----`

const k8sPrivKeyPEM = `-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQC1PLy2ZWEM7W4O
zVfSNd+USuRRdiZaZLwehn9Cd9mejprkKCMfVPHRB7nixhqal0SnIuPcU8RMX+za
sGWPidfXERa6CNN9o29AyGIZex0k38xDB+JgB6+TPYZGMAoKVUSVSANNb3k/E1fI
uvLj6AJjbGZvSB5wrAFy86qlMpj31K8QVizIo/+MtmXChlm4xWX3Go37DXNQBvuR
+wJYu3d1Nnn+pVhTraoJyGQY/0NzPxBv+uCxlHU77mt9SJ4+vBLXGao6zDUcQsaH
1HL37wJNj844YRi96H3CwS2DHhOIyQTJMzeonrFifYyamzbFhM+XX3TSW+e3rDf7
RNV4vgLzAgMBAAECggEAFCKQUK9Irff+3zQfCqKjmUUMwqQetLI2WCnXH5pTGhN7
z9dJt6RnTTLPiws5T2142hy9NfA0gcuyc8VfMyg8S4+Bd9+dJ3st1AOvKmeMIFmg
xt2sX8Da17/hGBEkPGumt/MNMuAslz5d5VYnv/w5r7QOWVSo1Rm2+39ouPu7FOVj
RHc3ePVldLXNdfhPXx7pK90eTkhNQLhw3HC/tgjjxVGnQHliYEuZeKdEfCfMCGUy
c6YWSPKqUh+f6oQ6NUxDgdvly+h5Mr7yLlLqRNfCAyPtnEcxc5vrpj/+v5BoAmnz
jjRFtb+DXtfjzTJb/xE4+UPWoFHIuAT5LxnWz/XNOQKBgQDZuSO1PzhGFpAGbvxv
SMtL8y77I0e+QxesgYYUG+dooRIfcteCOPkTawxlSdtjDWz5sVkITROPCZat6lAC
7Tau4NynDba3hAmJ8qKXdrMBrT68Y/6y3kOKYliwFdeA6LbVR9ozsjJsfxYu/Pbc
QX9CaCTQOJIt/WbHI7jpiuVT9QKBgQDVGYJsR4kaiNiOqOFPwN7Fo6RVXnvsNz44
/OvoIJvhCwMJf3DF6CVszlCloOUS/HkjsF+bLYPX29ndxpS6EKJegfwwgRJbjVGj
6JdF8VLTjgAa28ZeCjXEApQfoq2kSCwifZd2cV6fcqvwr+JneYlLawr4RqfCdEdM
D99A/PqSRwKBgGR5kAUAm9OcbfLKHSyuB0ORgkjbSyx+gdpWG64EApLCYj6pHNM3
v5o6eIn1v4zCkVvZgCDYkQIdhq/TxgDTv7yMgMeHCJ9AC7bhhi8n1AweCymda9jX
wYuHPy5jpgQTYOykMApTXfm73Fzq2HFkuHnI3dRJhww6OMgFsDv2oUZRAoGBAIpG
+UiFysN7FPrNPxbfUi7xFsuGeec3mZqlE/cWYc/Ps3LQTT8+ejp2TgKLutltFrY/
1mn9SNYjBOZR22Q5MwMcWaanul/J/bXcUXzDMTmxpQPIpJmVCnpcKf21M/OsGWdw
E3555iqU4FlX6BO424wis4WTY3xcs375taaYAYg1AoGBAJRqD8WNAXvN0tUsJfdA
v7A8PgdvWmfKYnoQg8TY3RdQ7pmr/NTfVP0Q8KcE/er+gsz/u4K1f7VPR027rhlQ
MNIYqyrXHWR+wM3zIGMnKta9pLito7RK72l/IUcut7hKNr2gESnzHZLw8S+4S6+E
atm8w9IC/uXWUoipYGgsLYMq
-----END PRIVATE KEY-----`

const edgeLBPubKeyPEM = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA4pc48iVWCUO6sUUFL95x
WnmHRE1krNMnJ5XCXKCn1crvxWViRvz1G5OH7gYJ7JPiTyemnBZRtBv0cx0J7OeY
Ny9XrajqIA+0VpltSWre9if+jNygT79nEVGHjqGtHbshQ8vAwvjYAaA1HBg7Dgy1
q6+vhUR5drXB7j0Ebt+EQxqqeWpyqLLG7JocY3vCkgtRI6broD4vSC9aAQU/t9rd
OalMndAbTh9//I5boU0BJF5kfW0Kz9dDyh+YCBixIxzD3AtRU9I+AZmrgmeKZOHG
H+LMoGWHoiYQWeygPKr3XhPV1RB9/FnX57nOiHU0O0yaFoQe9pnYugc6zmwUMOjM
QQIDAQAB
-----END PUBLIC KEY-----`

const edgeLBPrivKeyPEM = `-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDilzjyJVYJQ7qx
RQUv3nFaeYdETWSs0ycnlcJcoKfVyu/FZWJG/PUbk4fuBgnsk+JPJ6acFlG0G/Rz
HQns55g3L1etqOogD7RWmW1Jat72J/6M3KBPv2cRUYeOoa0duyFDy8DC+NgBoDUc
GDsODLWrr6+FRHl2tcHuPQRu34RDGqp5anKossbsmhxje8KSC1EjpuugPi9IL1oB
BT+32t05qUyd0BtOH3/8jluhTQEkXmR9bQrP10PKH5gIGLEjHMPcC1FT0j4BmauC
Z4pk4cYf4sygZYeiJhBZ7KA8qvdeE9XVEH38Wdfnuc6IdTQ7TJoWhB72mdi6BzrO
bBQw6MxBAgMBAAECggEAa1tw3LG56NvROapVAdAn7tEDuOXeXFATDhtTomVv/8iN
//3AKZgyXaVx0uHDk/ahVBWaR/v/LBHOqTfJ+xfjs8nXhe0xOI2YSSuabx2WlBRw
87aYshAiuj7htj998Uc3RJLmOrqCRYrtfukwGcEF1lmAFOlTj14gi0sXrZxzYwL7
9Dl8jzpp4Si6JpfpO7UwSkmOymerdVKY8J04LTJrc0O+bOExRG1V0jK6xcnaSr0R
WvYsGWqRBwHEJs+ePMkmfKZHI2TQqXecojyBTCxLqgbFSMtfNt3j3v+JWhq+cgRe
yw5nq/F8QJUUxM63/Fz1mDEZJUWiK4/t2ecmSmRmkQKBgQD7BmNT9Wr8GFVRCF4z
SZI3GH7eYxfdlwCK5CtyVNziNGJY1lMqhdlTFOSvhjKAfZXk3bcovFUiT14xCdj1
NHtWglOV+SHq9KSflP9wZma1A2wPOfv5A8pzyd5XZOhjjXYDm2p4sXo8b/HcucQa
OX60VX9HmS4aiqOcEPLkhZ4rFQKBgQDnFN0dsEeX8phh/y43C5vRwHqgxCqERa6t
2CofN9oad74BYsKH+zp5c2LT2SzjGjpc0LpZ44Dbgw80UVSog1Vcc20ODruJVpx/
9Bd/t6W3jGuqN/ttTw+3APX8MDsj9JryB2nDjEaaryHB27jrSOfCxXT9CgvKLeNI
Y3wR98t3fQKBgQCq1JYUNRgxp32oP8GbtZ3D0O/F+Dntmy3LV6wZipcnee7T9kdy
0NQtLjLTIMiNmOnBbwGOv1xQlSLMzJ7RgH3PSbIIhhsHAqZl08hifc23sjR/yD4q
IOJOGjstzoY3+bUujz3OFTnSl3xJckJ6dlY781NDLoOpnF/rfb1Ot4AEYQKBgC9E
ADpbXTmCQJMC3BQcRsHvieWqWjv1+NXMOklqZi01wuKLrdiclYhUBqnoaRsGuVtU
wlkyhJ/hvdFotVVGj8Y6Qds3PjrIQXiWl9vi132ktjz5+G6SVleLVcVApEgldy/8
PjDmV2a6XcBGThuqLOWU9+nuDR1Mp4md97nBIpgNAoGBAM+iSUJLvHWUR95ujG+a
rt4pDY2KMBGfCLhIUf1sXtvR30fFkmvi841h9WHD4lGnyFLf13Rh6nVEMHMJ2DPe
LC9JpKlZpMnuikBfCg3VWT7JX6MMidTBrTZ3uoCBLh4J3hqc0gocVTd50O5BaGv0
2KWwRsfk5dJUUxy5IwDLLoxv
-----END PRIVATE KEY-----`