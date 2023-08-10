type CategorizeRequest = {
    prompt: string;
    categories: string[];
    allow_null_category?: boolean;
    allow_multiple_classes?: boolean;
};
type CategorizeResponse = {
    choices: string[];
};
type RegexCompletionRequest = {
    prompt: string;
    patterns: string[];
    max_new_tokens?: number;
    stop_after_match?: boolean;
    temperature?: number;
    top_p?: number;
    top_k?: number;
    repetition_penalty?: number;
};
type RegexCompletionResponse = {
    completion: string;
    tokens_generated: number;
};

export { CategorizeRequest, CategorizeResponse, RegexCompletionRequest, RegexCompletionResponse };
