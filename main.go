package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// generateResourceName creates a resource name with a specified number of parts and separator
func generateResourceName(parts int, separator string, wordPools ...[]string) string {
	rand.Seed(time.Now().UnixNano()) // seed random generator

	var nameParts []string
	for i := 0; i < parts; i++ {
		// select a random pool and pick a word from it
		pool := wordPools[rand.Intn(len(wordPools))]
		nameParts = append(nameParts, pool[rand.Intn(len(pool))])
	}

	return strings.Join(nameParts, separator)
}

func main() {
	// define word pools
	prefixes := []string{"cosmic", "stellar", "orbit", "lunar", "solar", "galactic", "astro", "interstellar"}
	celestialTerms := []string{"nebula", "quasar", "pulsar", "galaxy", "comet", "asteroid", "meteor", "nova", "blackhole", "supernova"}
	suffixes := []string{"alpha", "beta", "gamma", "delta", "epsilon", "zeta", "theta", "iota", "kappa", "lambda"}

	// parse flags
	parts := flag.Int("p", 2, "number of name parts to include in the resource name")
	separator := flag.String("s", "-", "separator to use between name parts")
	count := flag.Int("n", 3, "number of names to generate (default: 3)")
	flag.Usage = func() {
		fmt.Println("Usage: planetary-resource-generator [options] <prefix>")
		fmt.Println("Options:")
		fmt.Println("  -p <number>   number of name parts (default: 2)")
		fmt.Println("  -s <string>   separator to use between name parts (default: -)")
		fmt.Println("  -n <number>   number of names to generate (default: 3)")
		fmt.Println("Arguments:")
		fmt.Println("  <prefix>      required string to prefix each generated name (e.g. example-resource)")
	}

	flag.Parse()

	// check if prefix is provided
	if flag.NArg() != 1 {
		flag.Usage()
		return
	}
	prefix := flag.Arg(0)

	// generate and print resource names
	fmt.Printf("Generating %d resource names with %d parts each:\n", *count, *parts)
	for i := 0; i < *count; i++ {
		name := generateResourceName(*parts, *separator, prefixes, celestialTerms, suffixes)
		fmt.Printf("%s%s%s\n", prefix, *separator, name)
	}
}
