# command-line-chatbot

ChatGPT designed for use in the command line

## Usage

```bash
$ go install Entrecasteaux.de/x/command-line-chatbot@latest
$ echo $OPENAI_API_KEY # this should be the OpenAI API key
$ command-line-chatbot
```

## Example

- Use `<Ctrl+D>` to terminate user input. (`Ctrl+Z` + `Enter` in windows)
- Use `<Ctrl+C>` to terminate the chat session.

```
$ command-line-chatbot
> Hi, I'm a chatbot. How can I help you?
> User: How are you? <Ctrl+D>
> Assistant: As an AI language model, I don't have feelings, emotions, or physical experiences. But thank you for asking! How can I assist you today?
> User:
```

## License

MIT