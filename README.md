# tomlenv

tomlenv provieds a wrapper function of the `DecodeFile()` in [BurntSushi/toml](http://github.com/BurntSushi/toml) package.

It's inspired by [Configuring Rails Applications](http://guides.rubyonrails.org/configuring.html).

## Useage

Prepare toml dir. for example

```
test
├── common.toml
└── dev.toml
```

common.toml

```toml
top = "common_value"

[foo]
v1 = "common_value"
v2 = 1

[[bar]]
v = "common_value1"
[[bar]]
v = "common_value2"

[env.dev]
v = "common_value"
```

dev.toml

```toml
top = "dev_value"

[env.dev]
v = "dev_value"
```

parsed `common.toml` to tml at first, `dev.toml` is next (overwrite value).

```go
tml := Toml{}
DecodeEnv("dev", "./test", &tml)
// ---- result ------
//	Toml{
//		Top: "dev_value",
//		Foo: Foo{
//			V1: "common_value",
//			V2: 1,
//		},
//		Bar: []Bar{
//			{V: "common_value1"},
//			{V: "common_value2"},
//		},
//		Env: map[string]Env{
//			"dev": {V: "dev_value"},
//		},
//	}
```

see also [test code](tomlenv_test.go).

## Credit

This is an adaptation of [BurntSushi/toml](http://github.com/BurntSushi/toml). Thanks to the original authors for their work.

## Licence

tomlenv is licensed under the MIT