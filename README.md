# KONFIG

A simple utility to simplify the reading of configuraion info from multiple files or environment variables.

### Why?

Often when you build a micro service in golang, you want to provide some configurations that live with the executable, perhaps good for debugging or reasonable defaults that can be used in running the service. This can be baked into the docker image. 

Often some values have no good usable defaults (e.g. IP addresses) and these can be supplied from all manner of places like secrets and configmaps mounted as files. Some other are suppled as filename arguments at invocation.

Finally, some are supplied as environment variables. 

Usually, there is an order of preference so some way of passing the config settings overrides another.

## Solution

Konfig offers one call to read all the sources of configuraion info while letting you decide the order of preference. You just define and tag a struct to be filled with the config. You can even call it with different structs that are filled by different portions of the config sources.

It is able to read configuration from json files, yaml files, xml files and even the environment.

Try it out!

```golang
konfig.LoadCfg(cfg, <config file 1> <config file 2> ... "env" <more config fiels>)
```

Its a WIP. More to come.