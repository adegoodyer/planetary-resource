# planetary-resource-generator

[![Go Reference](https://pkg.go.dev/badge/github.com/adegoodyer/planetary-resource-generator.svg)](https://pkg.go.dev/github.com/adegoodyer/planetary-resource-generator)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Easily generate thousands of combinations of resource names with a celestial theme

## Install

```bash
go install github.com/adegoodyer/planetary-resource-generator@latest
```

## Usage
```bash
Usage: planetary-resource-generator [options] <prefix>
Options:
  -p <number>   number of name parts (default: 2)
  -s <string>   separator to use between name parts (default: -)
  -n <number>   number of names to generate (default: 3)
Arguments:
  <prefix>      required string to prefix each generated name (e.g. example-resource)
```

## Examples
```bash
# Generate default number of names with default options
planetary-resource-generator -n 2
# example-resource-orbit-alpha
# example-resource-nebula-solar

# Specify custom separator
planetary-resource-generator -n 2 -s _
# example_resource_theta_gamma
# example_resource_kappa_galaxy

# Generate names with 4 parts instead of 3
planetary-resource-generator -n 5 -p 4
# example-resource-galaxy-interstellar-nova-orbit
# example-resource-zeta-orbit-delta-interstellar

# Combine custom separator and parts
planetary-resource-generator -n 5 -p 4 -s .
# example.resource.stellar.stellar.astro.pulsar
# example.resource.zeta.lambda.galaxy.beta
# example.resource.cosmic.galactic.nebula.asteroid
# example.resource.asteroid.interstellar.alpha.supernova
# example.resource.galactic.stellar.zeta.delta
```

## Tags

- `latest`: Most recent stable build
- `x.y.z`: Specific version builds (e.g., `2.7.5`)
- `x.y`: Minor version builds (e.g., `2.7`)

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
