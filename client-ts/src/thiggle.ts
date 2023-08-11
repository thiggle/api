const thiggleVersion = '0.0.2';

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
    patterns: string[];
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