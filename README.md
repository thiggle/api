## [thiggle.com](https://thiggle.com) API

A simple structured API to run categorization tasks with an LLM. 

* **Zero Parsing**: Always returns structured JSON with only your categories
* **0,1,or N Labels**: Return exactly one class, or allow multiple classes, or allow uncategorized results
* **Deterministic**: Never returns unexpected or unparsable results.

### Examples

* [Building block for building higher-level AI agents](#ai-agents)
* [Answering multiple choice questions](#multiple-choice-questions)
* [Labeling training data](#labeling-training-data)
* [Sentiment analysis](#sentiment-analysis)

#### More documentation at [docs.thiggle.com](https://docs.thiggle.com)

### Quickstart

Get an API key at [thiggle.com/account](https://thiggle.com/account). Set it as an environment variable `THIGGLE_API_KEY`. Call the API directly over HTTPS or use one of the client libraries.

* [cURL](#curl)
* [Python](#python)
* [Go](#go)
* [TypeScript](#typescript)

#### cURL
```bash
curl -X POST "https://api.thiggle.com/v1/categorize" \
     -H "Content-Type: application/json" \
     -H "Authorization: Bearer $THIGGLE_API_KEY" \
     -d '{
         "prompt": "What animal barks?",
         "categories": ["Dog", "Cat", "Bird", "Fish"]
     }'
```

#### Python
```bash
pip install thiggle
```
```python
import thiggle as tg

api = tg.API(os.getenv("THIGGLE_API_KEY"))
response = api.categorize("What animal barks?", ["dog", "cat", "bird", "fish"])
print(response)
```

#### TypeScript
```bash
npm install @thiggle/client
```
```typescript
import Thiggle from '@thiggle/client';

const api = new Thiggle(process.env.THIGGLE_API_KEY);
const response = await api.categorize("What animal barks?", ["dog", "cat", "bird", "fish"]);
console.log(response);
```

#### Go
```bash
go get github.com/thiggle/api/client-go
```
```go
package main

import (
    "fmt"
    "os"

    "github.com/thiggle/api"
)

func main() {
    client := api.NewClient(os.Getenv("THIGGLE_API_KEY"))
    response, err := client.Categorize("What animal barks?", []string{"dog", "cat", "bird", "fish"})
    if err != nil {
        panic(err)
    }
    fmt.Println(response)
}
```

### API Keys
To get started, you'll need an API key. You can get one by signing up for an account at [https://thiggle.com](https://thiggle.com). Once you create an account, you generate API keys on your [account page](https://thiggle.com/account). Set the `THIGGLE_API_KEY` environment variable to your API key.

```bash
export THIGGLE_API_KEY=your-api-key
```

If you are using a client library, you can pass the API key as a parameter to the client constructor. If you are using the REST API directly, you can pass the API key in the `Authorization` header (be sure to include the `Bearer` prefix).

  ```bash copy
  curl -X POST "https://thiggle.com/api/v1/categorize" \
     -H "Content-Type: application/json" \
     -H "Authorization: Bearer $THIGGLE_API_KEY" \
     -d '{
         "prompt": "What animal barks?",
         "categories": ["Dog", "Cat", "Bird", "Fish"]
     }'
  ```


  ### Examples

#### AI Agents
Use the categorization API to choose the relevant tools to complete the task. Use this as a reliable building block for higher-order AI agents. Never worry about the API returning extraneous text or unknown categories that break your agent.
```json
{
  "prompt": "What tools do I need to complete the following task? Task: find the best restaurant in San Francisco. Tools:",
  "categories": ["google-maps-api", "python-repl", "calculator", "yelp-api", "ffmpeg"]
}
```
```json
{
  "choices": ["google-maps-api", "yelp-api"]
}
```
#### Multiple-Choice Questions
Answer multiple-choice questions. For questions with more than one correct answer, use the `allow_multiple_classes` flag. 
```json
{
  "prompt": "What animals have four legs?",
  "categories": ["cat", "dog", "bird", "fish", "elephant", "snake"],
  "allow_multiple_classes": true,
}
```
```json
{
  "choices": ["cat", "dog", "elephant"]
}
```

#### Labeling Training Data
You can use the categorization API to label text for training data. For example, you could use it to label text for a text classifier. This example bins text into different buckets: ['finance', 'sports', 'politics', 'science', 'technology', 'entertainment', 'health', 'other'].
```json
{
  "prompt": "What category does this text belong to? Text: The Dow Jones Industrial Average fell 200 points on Monday.",
  "categories": ["finance", "sports", "politics", "science", "technology", "entertainment", "health", "other"]
}
```
```json
{
  "choices": ["finance"]
}
```


#### Sentiment Analysis
Classify any text into sentiment classes. 

```json
{
  "prompt": "Is this a positive or negative review of Star Wars? The more one sees the main characters, the less appealing they become. Luke Skywalker is a whiner, Han Solo a sarcastic clod, Princess Leia a nag, and C-3PO just a drone",
  "categories": ["positive", "negative"]
}
```
```json
{
  "choices": ["negative"]
}
```

Use any sentiment categories you like. For example, you could use `["positive", "neutral", "negative"]` or `["positive", "negative", "very positive", "very negative"]`. Or even `["happy", "sad", "angry", "surprised", "disgusted", "fearful"]`.

### Rate Limits

The API is rate limited to 100 requests per minute. If you exceed this limit, you will receive a `429 Too Many Requests` response. If you need a higher rate limit, please contact us at [quota@thiggle.com](mailto:quota@thiggle.com).

Your current rate limit usage is returned in the `X-RateLimit-Remaining` header. If you are using a client library, you can use this to determine when you are approaching the rate limit.

### Quota

Quotas are determined by your current plan. You can view your current plan on your [account page](https://thiggle.com/account). The quota is reset at the beginning of each month. If you exceed your quota, you will receive a `402 Payment Required` response. Your current quota usage is returned in the `X-Quota-Remaining` header. If you are using a client library, you can use this to determine when you are approaching your quota.

