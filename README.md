# crypto-cli
A command-line dependency built in Go Lang for generating sha1,sha512, md4 &amp; md5  
**Crypto CLI** is a test dependency built in Go for generating sha1,sha512, md4 & md5 hash.

The goal of this cli tool is to provide a generic usage for all backend infrastructure regardless of language, eliminating the need for language/framework dependencies

### Usage:
```console
$crypto-cli [command]
```

#### Available Commands:  
<pre>  help        Help about any command  
  md4         Generate md4 hash  
  md5         Generate md5 hash  
  sha1        Generate sha1 hash  
  sha512      Generate sha512 hash  

Flags:  
      --config string   config file (default is $HOME/.crypto-cli.yaml)  
  -h, --help            help for crypto-cli  
  -t, --toggle          Help message for toggle  

</pre>  


### Help:
```console
$crypto-cli [command] --help" for more information about a command. 
```
