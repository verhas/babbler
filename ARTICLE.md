# I built a brown-bag project: babbler

A confession first: I have spent the last few months building things that require actual thought.
**Turicum** is a programming language.
**perag** is a personal RAG system that reads your documents so your AI doesn't have to.
**mdship** manipulates markdown files with the seriousness of a tool that knows people will build CI pipelines on top of it.
All good, all real, all serious work with the hope that some will use them and make value.

Sometimes you just want a brown bag project.
The kind you can start, finish, and ship before your lunch sandwich gets soggy.
So I built **babbler**.

## The problem: humans are bad at remembering numbers with 9 digits and no personality

Quick, what's your tax file reference number?

Not the one you can look up.
The one you can *recall*, right now, from memory, while a clerk on the phone is waiting and your insurance is somehow involved.

Modern life hands you an endless parade of numbers exactly like this:

- Your health insurance claim is `CLM-88213047`.
- Your tax office reference is `2026/AT/0041938`.
- Your car insurance policy is `POL-773190-B`.
- The ticket you opened with customer support, which you will need to reference again in three weeks, is `#40921177`.

None of these are memorable.
None of them are *meant* to be memorable.
They're auto-increment counters or UUID fragments that some database handed out.
The human is expected to somehow carry that around in the head or more realistically, in a purse full of paper never to find.
Nobody has ever said "oh yes, claim eight-eight-two-one-three-oh-four-seven, of course" out loud without sounding like they're reading a hostage's proof-of-life number.

What humans *are* good at remembering are names.
So: what if instead of `88213047`, your claim was called **Talo Buno**?
Instead of `88213048`, it's **Patu Luta**.
Two little pronounceable words instead of eight forgettable digits.
You'd still write the number down somewhere official, but you could *talk* about "the Talo Buno claim" the way you'd talk about a pet, and actually remember which one it was.

That's it.
That's the whole idea.
It's not going to save the world.
It might save you one mildly awkward phone call.

## The algorithm (short version, I promise)

The core idea: every counter number gets converted into two pronounceable syllable-pairs, like `Talo Buno` or `Patu Luta`.
Always four syllables, always two capitalized words, always the same result for the same number.

Under the hood, it's mercifully simple — no cryptography, no hashing, no network calls, no "AI" (yes, I checked myself at the door on that one):

1. Take 36 consonant-vowel syllables (`ba`, `bo`, `bu`, `da`...
up to `zu`) — no `h`, no confusing look-alikes, just clean sounds that work across languages.
2. Combine them two at a time to build words, and quietly remove the ones you wouldn't want on an official document (more on that in a second).
3. Run the input number through a small, fixed bit of math — a *bijection*, meaning every number maps to exactly one name and every valid name maps back to exactly one number.
No two different claim numbers ever accidentally get the same nickname.
No retries, no collisions, no "sorry, try again" — just guaranteed-unique output, by construction, every time.
4. Format the result as two capitalized words and hand it back.

I went through a "hash it with MD5 and hope for the best" version first, because that's the obvious first instinct.
Turns out "hope for the best" produced actual name collisions for about a third of all numbers, which is a spectacular way to guarantee two different insurance claims share a name.
The fixed version replaced "hope" with actual math, which is generally a good trade.

> The md5 version was never shipped, you can only find my silliness in the git logs.

## The blacklist (the fun part)

Here's the part of this project that made me grin more than a numbering scheme has any right to: since the syllables get combined somewhat randomly, you occasionally get combinations you do *not* want stamped on someone's official tax document.
So there's a small blacklist of forbidden two-syllable words the generator is guaranteed to never produce — permanently filtered out before any number is ever assigned, not caught after the fact.

The list is a delightful little multilingual minefield: some are baby-talk for things nobody needs on a government form, one is a slang word for police (and also, unrelated, poop, in Spanish), one means "senile," a couple are just international slang for excrement in three different languages, and two are, let's say, anatomically specific in a way that would not go over well in a customer support ticket.
All tastefully avoided.

Meanwhile, — and this delighted me — `Mama`, `Papa`, and `Nana` all made it through completely clean.
Sometimes the universe hands you a nice coincidence.

## Try it yourself

```
10000 → "Talo Buno"
10001 → "Patu Luta"
```

Every number from 0 up to about 1.64 million gets its own guaranteed-unique, guaranteed-(major EU languages plus Hungarian)inoffensive two-word name.
One-way only — it makes names, it doesn't decode them back, because the actual point was never secrecy or reversibility, it was just "give me a name I can remember instead of a number I can't."

## MIT licensed, available in Java, Go, and Python

The code — encoder, tests, docs, the works — is MIT licensed, so grab it and use it however you like: https://github.com/verhas/babbler

It ships with matching, fully tested implementations in **Java**, **Go**, and **Python** — pick your poison, they all produce byte-identical names for the same number.

And by "Java, Go, or Python," I obviously also mean *any other language you like*, because with three clean, well-documented reference implementations sitting right there, any reasonably competent LLM today can knock out a **Rust**, **Kotlin**, or — sure, why not — **COBOL** port before you've finished reading this sentence twice.
Consider that a stretch goal I'm delegating to the machines.

Brown bag project, delivered.
Back to the serious stuff after the weekend.

## Disclosure: AI-assisted creation, human-verified content

This article and its visuals were created with LLM assistance, but every claim, argument, and technical detail was thoroughly reviewed, fact-checked, and approved by the author before publication.
No "slop" here.
To be clear about what this means: I used an AI as a tool for composition and iteration—much like using a typewriter instead of painstakingly handwriting each page with a goose feather, let alone carve into stone.
The tool speeds up the mechanical work; it doesn't replace the thinking.
The tone, the structure, the emphasis—all human judgment.
This is how AI tools should be used in knowledge work: as amplifiers of clarity and productivity, not as replacements for expertise or accountability.
A human person, specifically I stand behind every word.
