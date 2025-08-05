export interface User {
    username: string;
    password: string;
}

export interface Message {
    sender: string;
    receiver: string;
    text: string;
    created_at: string;
}