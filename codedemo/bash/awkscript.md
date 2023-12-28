```bash
#!/bin/awk -f
BEGIN {
    FS = ":"
    md5 = 0
    sha1 = 0
    sha256 = 0
}

function is_number(string) {
    return match(string, /^\[0-9]+$/)
}

function supported_target(virusname) {
    return match(virusname,/^Unix|^Win/)
}

function is_supported(filesize, virusname) {
    if (is_number(filesize) == 1) 
    {
        if (supported_target(virusname) == 1)
        {
            return 1
        }
    }
    return 0
}

{
    hash = $1
    size = $2
    name = $3
    if (is_supported(size,name) == 1) {
        if (length(hash) == 32) # md5
        {
            printf "%s:%s:%s\n",size,hash,name  >> "md5.db"
            md5++
            print size >> "size.db"
        }
        #if (length(hash) == 40) # sha1
        #{
        #    printf "%s:%s:%s\n",size,hash,name  >> "sha1.db"
        #    sha1++
        #}
        #if (length(hash) == 64) # sha256
        #{
        #    printf "%s:%s:%s\n",size,hash,name  >> "sha256.db"
        #    sha256++
        #}
    }
}

END {
    printf "md5: %s\nsha1: %s\nsha256: %s\n",md5,sha1,sha256
}
```