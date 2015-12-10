deploy how-to
=============

deploy with https://chriseth.github.io/browser-solidity

compilation produces a number of strings, which are all copied here for reference.

to verify the validity of the resulting byte code, it should be possible to check your compilation
result with the byte code contained in this repository. it is important to use the same compiler
version if you want to check compilation results.

all files were compiled with optimizations enabled in the compiler options: this reduces byte code size by half, roughly.

`TweetAccount` compiled with: `Solidity version: 0.1.6-321b1ed2/.-Emscripten/clang/int linked to libethereum-1.1.0-998db95a/.-Emscripten/clang/int`

`TweetRegistry` compiled with: `Solidity version: 0.2.0-15a1468c/.-Emscripten/clang/int linked to libethereum-1.1.1-9f638c1c/.-Emscripten/clang/int`

if you want to deploy any of the contracts on your own, use the file `web3-deploy.txt` to deploy with geth.


deployment results
------------------

TweetRegistry: `0xe0f278b72097e563b09d7dc94c6f75aff5e83298`

TweetTestAccount: `0x9e82d1745c6c9c04a6cfcde102837cf0f25efc56`

