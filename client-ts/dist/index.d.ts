type CategorizeRequest = {
    prompt: string;
    categories: string[];
    allow_null_category?: boolean;
    allow_multiple_classes?: boolean;
};
type CategorizeResponse = {
    choices: string[];
};

export { CategorizeRequest, CategorizeResponse };
