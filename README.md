# Chatbot

A quick and dirty UI to GPT 3.5.

## Installation

Install [Go](https://go.dev/doc/install).

Run this command:

```
go install 'github.com/jplein/chatbot@latest'
```

## System requirements

This tool will work on macOS and Linux. I haven't tested it on Windows.

## Get your OpenAI token

You can get your OpenAI token from [the OpenAI API Keys page](https://platform.openai.com/account/api-keys). You'll need to make your OpenAI account a paid amount. Each time you use the API, including when you use this tool, it will cost a small amount.

## Store your API key

Make a directory named `.chatbot` in your home directory. Edit the file `config.json` to look like this:

```
{"api_key":"your API key here"}
```

Replace "your API key here" with your API key.

## Ask a question

```
chatbot "What is a large language model?"
```

Each time you ask a question, the question and OpenAI's answer are stored, in the directory `~/.chatbot/transcripts`. This allows the context of your conversation to be saved and passed up to OpenAI, e.g.:

```
  $ chatbot "In Spanish, what is the yo conjugation of ser?"
The yo conjugation of ser is "soy".

  $ chatbot "What about el/ella/usted?"
The conjugation of "ser" for "él/ella/usted" is "es".

  $ chatbot "What about the preterite conjugations?"
The preterite conjugations of "ser" are:

- yo fui
- tú fuiste
- él/ella/usted fue
- nosotros/nosotras fuimos
- vosotros/vosotras fuisteis
- ellos/ellas/ustedes fueron
```

If you start a new shell, chatbot will have a new context.

## TODO

The engine is hard-coded to `chatgpt-3.5-turbo`. In the future, you'll be able to choose the engine.

There is no mechanism to prevent storing the transcript, or resetting it for the current shell, other than finding the transcripts and deleting them.