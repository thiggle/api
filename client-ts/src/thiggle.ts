const thiggleVersion = '0.0.5';

export type CategorizeRequest = {
    prompt: string;
    categories: string[];
    allow_null_category?: boolean;
    allow_multiple_classes?: boolean;
};

export type CategorizeResponse = {
    choices: string[];
};

export type ContextFreeCompletionRequest = {
    prompt: string;
    grammar: string;
    max_new_tokens?: number;
    stop_after_match?: boolean;
    temperature?: number;
    top_p?: number;
    top_k?: number;
    repetition_penalty?: number;
};

export type ContextFreeCompletionResponse = {
    completion: string;
    tokens_generated: number;
    stop_reason: string;
};

export type RegexCompletionRequest = {
    prompt: string;
    pattern: string[];
    max_new_tokens?: number;
    stop_after_match?: boolean;
    temperature?: number;
    top_p?: number;
    top_k?: number;
    repetition_penalty?: number;
};

export type RegexCompletionResponse = {
    completion: string;
    tokens_generated: number;
    stop_reason: string;
};

export type CompletionRequest = {
    prompt: string;
    model?: string | string[]; // Can be a single string or an array of strings
    max_tokens?: number;
    temperature?: number;
    top_p?: number;
    stop?: string[];
    presence_penalty?: number;
    frequency_penalty?: number;
};

export type CompletionResponse = {
    id: string;
    object: string;
    created: number;
    choices: CompletionChoice[];
    usage?: Usage;
};

export type CompletionChoice = {
    text: string;
    created: number;
    model: string;
    finish_reason: 'length' | 'stop' | 'timeout';
    usage: Usage;
};

export type Usage = {
    prompt_tokens: number;
    completion_tokens: number;
    prediction_time_ms: number;
    total_tokens: number;
};

class Thiggle {
    private baseUrl: string;
    private apiKey: string;

    constructor(apiKey: string, baseUrl: string = 'https://api.thiggle.com') {
        this.baseUrl = baseUrl;
        this.apiKey = apiKey;
    }

    private async makeRequest<T>(endpoint: string, request: any): Promise<T> {
        const response = await fetch(`${this.baseUrl}${endpoint}`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${this.apiKey}`,
                'User-Agent': `thiggle-client-ts/${thiggleVersion}`,
            },
            body: JSON.stringify(request)
        });

        if (!response.ok) {
            throw new Error(`API request failed with status ${response.status}`);
        }

        return response.json() as Promise<T>;
    }



    // Function for the completion request
    async complete(request: CompletionRequest): Promise<CompletionResponse> {
        return this.makeRequest<CompletionResponse>('/v1/completion', request);
    }

    async categorize(request: CategorizeRequest): Promise<CategorizeResponse> {
        return this.makeRequest<CategorizeResponse>('/v1/categorize', request);
    }

    async regexCompletion(request: RegexCompletionRequest): Promise<RegexCompletionResponse> {
        return this.makeRequest<RegexCompletionResponse>('/v1/completion/regex', request);
    }

    async contextFreeCompletion(request: ContextFreeCompletionRequest): Promise<ContextFreeCompletionResponse> {
        return this.makeRequest<ContextFreeCompletionResponse>('/v1/completion/cfg', request);
    }
}

export default Thiggle;