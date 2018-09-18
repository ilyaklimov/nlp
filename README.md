Packages for NLP (Natural Language Processing).

# Tokenize
Tokenizer: string to n-terms, string to n-runes (in future).

## Terms

```
str := "I'm doing no harm--I'm not playing games, I'm mending the Primus,' said the cat with a hostile scowl, ' and I'd better warn you that a cat is an ancient and inviolable animal."

uniterms, err := tokenize.ToUniterms(str)
if err != nil {
	fmt.Println(err)
}

fmt.Printf("%s\n%#v\n", str, uniterms)
```

### TODO:
* english nouns plural (apple's, cat's, ...)
* n-runes ...


# Metrics
Calculation of text metrics.

## Tokens

```
tokens := []string{"I'M", "DOING", "NO", "HARM", "I'M", "NOT", "PLAYING", "GAMES", "I'M", "MENDING", "THE", "PRIMUS", "SAID", "THE", "CAT", "WITH", "A", "HOSTILE", "SCOWL", "AND", "I'D", "BETTER", "WARN", "YOU", "THAT", "A", "CAT", "IS", "AN", "ANCIENT", "AND", "INVIOLABLE", "ANIMAL"}

tf, err := metrics.TF("I'M", tokens)
if err != nil {
	fmt.Println(err)
}

fmt.Println(tf)

\\ 0.09090909090909091
```