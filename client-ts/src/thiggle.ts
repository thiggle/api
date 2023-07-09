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

class Thiggle {
    private baseUrl: string;
    private apiKey: string;

    constructor(apiKey: string, baseUrl: string = 'https://api.thiggle.com') {
        this.baseUrl = baseUrl;
        this.apiKey = apiKey;
    }

    async categorize(request: CategorizeRequest): Promise<CategorizeResponse> {
        const response = await fetch(`${this.baseUrl}/v1/categorize`, {
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

        return response.json() as Promise<CategorizeResponse>;
    }
}

export default Thiggle;