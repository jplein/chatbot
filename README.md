# Chatbot

A quick and dirty command-line client for OpenAI's `gpt-3.5-turbo` model, the same model used by [ChatGPT](https://openai.com/blog/chatgpt).

This tool is in a very early stage of development. Expect the configuration-file format and the CLI parameters to change quite a bit in the future.

This software is covered under the [MIT license](LICENSE.txt).

## Installation

Install [Go](https://go.dev/doc/install).

Run this command:

```
go install 'github.com/jplein/chatbot@v0.0.2'
```

## System requirements

This tool will work on macOS and Linux. I haven't tested it on Windows.

## Get your OpenAI token

You can get your OpenAI token from [the OpenAI API Keys page](https://platform.openai.com/account/api-keys). You'll need to make your OpenAI account a paid amount. 

EACH TIME YOU USE THE API, INCLUDING WHEN YOU USE THIS TOOL, OPENAI WILL CHARGE YOU MONEY. For details, see [OpenAI's pricing page](https://openai.com/pricing). The model used by this tool is `gpt-3.5-turbo`, currently hard-coded but with the ability to choose a model planned for a future version.

## Store your API key

Make a directory named `.chatbot` in your home directory. Edit the file `~/.chatbot/config.json` to look like this:

```
{"api_key":"your API key here"}
```

Replace "your API key here" with your API key.

You can also track the token usage of each message to the OpenAI API by setting the `log_token_usage` property to `true`, like this:

```
{"api_key":"your API key here", "log_token_usage": true}
```

This will add a message like this to the end of each response:

```
(tokens used: 1039)
```

You may also choose a model. The default model is `gpt-3.5-turbo`. To use `gpt-3.5-turbo-16k`:

```
{"api_key":"your API key here", "model": "gpt-3.5-turbo-16k"}
```

Or one of the other models. See [the OpenAI Models documentation](https://platform.openai.com/docs/models) and the [model endpoint documentation](https://platform.openai.com/docs/models/model-endpoint-compatibility).

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

There needs to be a better mechanism for storing the API key than hand-editing a JSON file.

The engine is hard-coded to `chatgpt-3.5-turbo`. In the future, you'll be able to choose the engine.

There is no mechanism to prevent storing the transcript, or resetting it for the current shell, other than finding the transcripts and deleting them.