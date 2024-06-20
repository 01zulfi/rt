# remember this (rt): cli utility to store key value pairs

i use it to store short texts, ids, memos for later use

## how to use

copy `rt` executable file to `/usr/local/bin` (or equivalent). this should make `rt` command available everywhere

```bash
curl https://raw.githubusercontent.com/01zulfi/rt/main/rt > ./rt
sudo mv ./rt /usr/local/bin/
chmod +x /usr/local/bin/rt
```

it stores key value pairs in a json format at `$HOME/.rt.json`

## commands

- rt a(dd) key ?value    
    if value isn't provided, key is the value

- rt v(iew) ?key

- rt x|clear-all

- rt r(emove) key
